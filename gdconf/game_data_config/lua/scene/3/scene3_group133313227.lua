-- 基础信息
local base_info = {
	group_id = 133313227
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
	{ config_id = 227001, gadget_id = 70290687, pos = { x = -429.060, y = -72.894, z = 5523.045 }, rot = { x = 0.000, y = 121.973, z = 0.000 }, level = 32, area_id = 32 },
	{ config_id = 227002, gadget_id = 70290683, pos = { x = -428.556, y = -72.777, z = 5524.670 }, rot = { x = 0.000, y = 154.567, z = 0.000 }, level = 32, area_id = 32 }
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
		gadgets = { 227001, 227002 },
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