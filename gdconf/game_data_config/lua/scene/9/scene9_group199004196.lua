-- 基础信息
local base_info = {
	group_id = 199004196
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
}

-- 区域
regions = {
}

-- 触发器
triggers = {
}

-- 变量
variables = {
}

-- 废弃数据
garbages = {
	monsters = {
		{ config_id = 196001, monster_id = 20060201, pos = { x = -300.829, y = 120.000, z = -614.474 }, rot = { x = 0.000, y = 327.353, z = 0.000 }, level = 20, drop_tag = "飘浮灵", pose_id = 101, area_id = 400 },
		{ config_id = 196003, monster_id = 20011501, pos = { x = -306.148, y = 120.000, z = -613.398 }, rot = { x = 0.000, y = 336.772, z = 0.000 }, level = 20, drop_tag = "大史莱姆", area_id = 400 },
		{ config_id = 196004, monster_id = 20010401, pos = { x = -312.944, y = 120.000, z = -600.755 }, rot = { x = 0.000, y = 152.433, z = 0.000 }, level = 20, drop_tag = "大史莱姆", area_id = 400 },
		{ config_id = 196005, monster_id = 20060301, pos = { x = -317.384, y = 120.000, z = -598.796 }, rot = { x = 0.000, y = 143.852, z = 0.000 }, level = 20, drop_tag = "飘浮灵", pose_id = 101, area_id = 400 }
	}
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