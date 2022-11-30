-- 基础信息
local base_info = {
	group_id = 246202003
}

-- Trigger变量
local defs = {
	center_point = 3007,
	rogue_exit = 3015
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
	{ config_id = 3015, gadget_id = 70800208, pos = { x = 17.536, y = -41.090, z = -0.008 }, rot = { x = 0.000, y = 270.000, z = 0.000 }, level = 1 }
}

-- 区域
regions = {
}

-- 触发器
triggers = {
}

-- 点位
points = {
	{ config_id = 3001, pos = { x = 0.109, y = -43.990, z = 0.060 }, rot = { x = 0.000, y = 270.000, z = 0.000 }, tag = 2 },
	{ config_id = 3002, pos = { x = 0.109, y = -43.990, z = -6.473 }, rot = { x = 0.000, y = 270.000, z = 0.000 }, tag = 2 },
	{ config_id = 3003, pos = { x = 0.109, y = -43.990, z = 6.197 }, rot = { x = 0.000, y = 270.000, z = 0.000 }, tag = 2 },
	{ config_id = 3004, pos = { x = 5.140, y = -43.990, z = 0.060 }, rot = { x = 0.000, y = 270.000, z = 0.000 }, tag = 2 },
	{ config_id = 3005, pos = { x = 5.140, y = -43.990, z = -6.473 }, rot = { x = 0.000, y = 270.000, z = 0.000 }, tag = 2 },
	{ config_id = 3006, pos = { x = 5.140, y = -43.990, z = 6.197 }, rot = { x = 0.000, y = 270.000, z = 0.000 }, tag = 2 },
	{ config_id = 3007, pos = { x = -7.392, y = -43.990, z = 0.010 }, rot = { x = 0.000, y = 0.000, z = 0.000 } },
	{ config_id = 3008, pos = { x = -12.523, y = -43.990, z = -0.011 }, rot = { x = 0.000, y = 90.000, z = 0.000 } },
	{ config_id = 3009, pos = { x = -14.635, y = -43.990, z = 0.060 }, rot = { x = 0.000, y = 90.000, z = 0.000 }, tag = 2 },
	{ config_id = 3010, pos = { x = -14.635, y = -43.990, z = -6.473 }, rot = { x = 0.000, y = 90.000, z = 0.000 }, tag = 2 },
	{ config_id = 3011, pos = { x = -14.633, y = -43.990, z = 6.197 }, rot = { x = 0.000, y = 90.000, z = 0.000 }, tag = 2 },
	{ config_id = 3012, pos = { x = -9.604, y = -43.990, z = 0.060 }, rot = { x = 0.000, y = 90.000, z = 0.000 }, tag = 2 },
	{ config_id = 3013, pos = { x = -9.604, y = -43.990, z = -6.473 }, rot = { x = 0.000, y = 90.000, z = 0.000 }, tag = 2 },
	{ config_id = 3014, pos = { x = -9.604, y = -43.990, z = 6.197 }, rot = { x = 0.000, y = 90.000, z = 0.000 }, tag = 2 }
}

-- 变量
variables = {
}

--================================================================
-- 
-- 初始化配置
-- 
--================================================================

-- 初始化时创建
init_config = {
	suite = 1,
	end_suite = 0,
	rand_suite = false
}

--================================================================
-- 
-- 小组配置
-- 
--================================================================

suites = {
	{
		-- suite_id = 1,
		-- description = ,
		monsters = { },
		gadgets = { },
		regions = { },
		triggers = { },
		rand_weight = 100
	}
}

--================================================================
-- 
-- 触发器
-- 
--================================================================

require "V2_7/Rogue_Cell"
require "V2_7/Rogue_Terrain_2"