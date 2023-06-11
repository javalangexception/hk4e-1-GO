package main

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"hk4e/common/config"
	"hk4e/common/mq"
	"hk4e/common/rpc"
	"hk4e/node/api"
	"hk4e/pkg/endec"
	"hk4e/pkg/logger"
	"hk4e/protocol/cmd"
	"hk4e/protocol/proto"
	"hk4e/robot/client"
	"hk4e/robot/login"
	"hk4e/robot/net"

	pb "google.golang.org/protobuf/proto"
)

var APPID string

func main() {
	config.InitConfig("application.toml")
	logger.InitLogger("robot")
	defer func() {
		logger.CloseLogger()
	}()

	// // DPDK模式需开启
	// err := engine.InitEngine("00:0C:29:3E:3E:DF", "192.168.199.199", "255.255.255.0", "192.168.199.1")
	// if err != nil {
	// 	panic(err)
	// }
	// engine.RunEngine([]int{0, 1, 2, 3}, 4, 1, "0.0.0.0")
	// time.Sleep(time.Second * 30)

	if config.GetConfig().Hk4e.ForwardModeEnable {
		// natsrpc client
		discoveryClient, err := rpc.NewDiscoveryClient()
		if err != nil {
			logger.Error("find discovery service error: %v", err)
			return
		}

		// 注册到节点服务器
		rsp, err := discoveryClient.RegisterServer(context.TODO(), &api.RegisterServerReq{
			ServerType: api.ROBOT,
		})
		if err != nil {
			logger.Error("register to node server error: %v", err)
			return
		}
		APPID = rsp.GetAppId()
		go func() {
			ticker := time.NewTicker(time.Second * 15)
			for {
				<-ticker.C
				_, err := discoveryClient.KeepaliveServer(context.TODO(), &api.KeepaliveServerReq{
					ServerType: api.ROBOT,
					AppId:      APPID,
				})
				if err != nil {
					logger.Error("keepalive error: %v", err)
				}
			}
		}()
		defer func() {
			_, _ = discoveryClient.CancelServer(context.TODO(), &api.CancelServerReq{
				ServerType: api.ROBOT,
				AppId:      APPID,
			})
		}()

		messageQueue := mq.NewMessageQueue(api.ROBOT, APPID, discoveryClient)
		defer messageQueue.Close()

		go runForward(messageQueue)
	} else {
		go runRobot()
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			// // DPDK模式需开启
			// engine.StopEngine()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func runForward(messageQueue *mq.MessageQueue) {
	dispatchInfo, err := login.GetDispatchInfo(config.GetConfig().Hk4eRobot.RegionListUrl,
		config.GetConfig().Hk4eRobot.RegionListParam,
		config.GetConfig().Hk4eRobot.CurRegionUrl,
		config.GetConfig().Hk4eRobot.CurRegionParam,
		config.GetConfig().Hk4eRobot.KeyId)
	if err != nil {
		logger.Error("get dispatch info error: %v", err)
		return
	}
	for {
		netMsg := <-messageQueue.GetNetMsg()
		if netMsg.OriginServerType != api.GATE {
			continue
		}
		if netMsg.MsgType != mq.MsgTypeServer || netMsg.EventId != mq.ServerForwardModeClientConnNotify {
			continue
		}
		gateAppId := netMsg.OriginServerAppId
		logger.Info("client connect, gateAppId: %v", gateAppId)
		for {
			netMsg = <-messageQueue.GetNetMsg()
			if netMsg.OriginServerType != api.GATE || netMsg.OriginServerAppId != gateAppId {
				continue
			}
			if netMsg.MsgType != mq.MsgTypeGame || netMsg.EventId != mq.NormalMsg {
				continue
			}
			gameMsg := netMsg.GameMsg
			if gameMsg.CmdId != cmd.GetPlayerTokenReq {
				continue
			}
			req := gameMsg.PayloadMessage.(*proto.GetPlayerTokenReq)
			session, err, rsp := login.GateLogin(dispatchInfo, nil, config.GetConfig().Hk4eRobot.KeyId, req, gameMsg.ClientSeq)
			if err != nil {
				logger.Error("remote gate login error: %v", err)
				continue
			}
			logger.Info("remote gate login ok, uid: %v", session.Uid)
			messageQueue.SendToGate(gateAppId, &mq.NetMsg{
				MsgType: mq.MsgTypeGame,
				EventId: mq.NormalMsg,
				GameMsg: &mq.GameMsg{
					UserId:         rsp.Uid,
					CmdId:          cmd.GetPlayerTokenRsp,
					ClientSeq:      gameMsg.ClientSeq,
					PayloadMessage: rsp,
				},
			})
			forwardLoop(session, messageQueue, gateAppId)
			break
		}
	}
}

func forwardLoop(session *net.Session, messageQueue *mq.MessageQueue, gateAppId string) {
	for {
		select {
		case netMsg := <-messageQueue.GetNetMsg():
			if netMsg.OriginServerType != api.GATE || netMsg.OriginServerAppId != gateAppId {
				continue
			}
			switch netMsg.MsgType {
			case mq.MsgTypeGame:
				if netMsg.EventId != mq.NormalMsg {
					continue
				}
				gameMsg := netMsg.GameMsg
				session.SendMsgFwd(gameMsg.CmdId, gameMsg.ClientSeq, gameMsg.PayloadMessage)
				if gameMsg.CmdId == cmd.PlayerLoginReq {
					data, _ := json.Marshal(gameMsg.PayloadMessage)
					logger.Debug("PlayerLoginReq: %v", string(data))
				}
			case mq.MsgTypeServer:
				switch netMsg.EventId {
				case mq.ServerForwardModeClientCloseNotify:
					logger.Info("client conn close, uid: %v", session.Uid)
					session.Close()
				}
			}
		case protoMsg := <-session.RecvChan:
			gameMsg := new(mq.GameMsg)
			gameMsg.UserId = session.Uid
			gameMsg.CmdId = protoMsg.CmdId
			if protoMsg.HeadMessage != nil {
				gameMsg.ClientSeq = protoMsg.HeadMessage.ClientSequenceId
			}
			// 在这里直接序列化成二进制数据 防止发送的消息内包含各种游戏数据指针 而造成并发读写的问题
			payloadMessageData, err := pb.Marshal(protoMsg.PayloadMessage)
			if err != nil {
				logger.Error("parse payload msg to bin error: %v", err)
				continue
			}
			gameMsg.PayloadMessageData = payloadMessageData
			messageQueue.SendToGate(gateAppId, &mq.NetMsg{
				MsgType: mq.MsgTypeGame,
				EventId: mq.NormalMsg,
				GameMsg: gameMsg,
			})
			if gameMsg.CmdId == cmd.PlayerLoginRsp {
				data, _ := json.Marshal(protoMsg.PayloadMessage)
				logger.Debug("PlayerLoginRsp: %v", string(data))
			}
		case <-session.DeadEvent:
			logger.Info("remote gate conn close, uid: %v", session.Uid)
			close(session.SendChan)
			messageQueue.SendToGate(gateAppId, &mq.NetMsg{
				MsgType: mq.MsgTypeServer,
				EventId: mq.ServerForwardModeServerCloseNotify,
			})
			return
		}
	}
}

func runRobot() {
	dispatchInfo, err := login.GetDispatchInfo(config.GetConfig().Hk4eRobot.RegionListUrl,
		config.GetConfig().Hk4eRobot.RegionListParam,
		config.GetConfig().Hk4eRobot.CurRegionUrl,
		config.GetConfig().Hk4eRobot.CurRegionParam,
		config.GetConfig().Hk4eRobot.KeyId)
	if err != nil {
		logger.Error("get dispatch info error: %v", err)
		return
	}
	if config.GetConfig().Hk4eRobot.DosEnable {
		dosBatchNum := int(config.GetConfig().Hk4eRobot.DosBatchNum)
		for i := 0; i < int(config.GetConfig().Hk4eRobot.DosTotalNum); i += dosBatchNum {
			wg := new(sync.WaitGroup)
			wg.Add(dosBatchNum)
			for j := 0; j < dosBatchNum; j++ {
				go httpLogin(config.GetConfig().Hk4eRobot.Account+"_"+strconv.Itoa(i+j), dispatchInfo, wg)
			}
			wg.Wait()
		}
	} else {
		httpLogin(config.GetConfig().Hk4eRobot.Account, dispatchInfo, nil)
	}
}

func httpLogin(account string, dispatchInfo *login.DispatchInfo, wg *sync.WaitGroup) {
	defer func() {
		if config.GetConfig().Hk4eRobot.DosEnable {
			wg.Done()
		}
	}()
	accountInfo, err := login.AccountLogin(config.GetConfig().Hk4eRobot.LoginSdkUrl, account, config.GetConfig().Hk4eRobot.Password)
	if err != nil {
		logger.Error("account login error: %v", err)
		return
	}
	logger.Info("robot http login ok, account: %v", account)
	go func() {
		for {
			gateLogin(account, dispatchInfo, accountInfo)
			if !config.GetConfig().Hk4eRobot.DosLoopLogin {
				break
			}
			time.Sleep(time.Second)
			continue
		}
	}()
}

func gateLogin(account string, dispatchInfo *login.DispatchInfo, accountInfo *login.AccountInfo) {
	session, err, _ := login.GateLogin(dispatchInfo, accountInfo, config.GetConfig().Hk4eRobot.KeyId, nil, 1)
	if err != nil {
		logger.Error("gate login error: %v", err)
		return
	}
	logger.Info("robot gate login ok, account: %v", account)
	clientVersionHashData, err := hex.DecodeString(
		endec.Sha1Str(config.GetConfig().Hk4eRobot.ClientVersion + session.ClientVersionRandomKey + "mhy2020"),
	)
	if err != nil {
		logger.Error("gen clientVersionHashData error: %v", err)
		return
	}
	checksumClientVersion := strings.Split(config.GetConfig().Hk4eRobot.ClientVersion, "_")[0]
	session.SendMsg(cmd.PlayerLoginReq, &proto.PlayerLoginReq{
		AccountType:           1,
		SubChannelId:          1,
		LanguageType:          2,
		PlatformType:          3,
		Checksum:              "$008094416f86a051270e64eb0b405a38825",
		ChecksumClientVersion: checksumClientVersion,
		ClientDataVersion:     11793813,
		ClientVerisonHash:     base64.StdEncoding.EncodeToString(clientVersionHashData),
		ClientVersion:         config.GetConfig().Hk4eRobot.ClientVersion,
		SecurityCmdReply:      session.SecurityCmdBuffer,
		SecurityLibraryMd5:    "574a507ffee2eb6f997d11f71c8ae1fa",
		Token:                 accountInfo.ComboToken,
	})
	client.Logic(account, session)
}
