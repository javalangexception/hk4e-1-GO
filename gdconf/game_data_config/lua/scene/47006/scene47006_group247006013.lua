-- 基础信息
local base_info = {
	group_id = 247006013
}

-- DEFS_MISCS
local GroupId = 247006013
local point_id_list = {}
local MainGroupID = 247006001
local RegionID = 13001

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
	{ config_id = 13001, shape = RegionShape.CUBIC, size = { x = 40.000, y = 25.000, z = 40.000 }, pos = { x = 0.021, y = 0.050, z = 240.087 } }
}

-- 触发器
triggers = {
	{ config_id = 1013001, name = "ENTER_REGION_13001", event = EventType.EVENT_ENTER_REGION, source = "", condition = "", action = "" }
}

-- 点位
points = {
	{ config_id = 13002, pos = { x = -0.029, y = 0.039, z = 253.464 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 2 },
	{ config_id = 13003, pos = { x = -3.025, y = 0.039, z = 252.762 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 2 },
	{ config_id = 13004, pos = { x = 3.197, y = 0.039, z = 252.878 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 2 },
	{ config_id = 13005, pos = { x = -6.937, y = 0.039, z = 251.554 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 2 },
	{ config_id = 13006, pos = { x = 7.136, y = 0.039, z = 251.106 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 2 },
	{ config_id = 13007, pos = { x = -9.384, y = 0.039, z = 249.055 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 2 },
	{ config_id = 13008, pos = { x = 10.371, y = 0.039, z = 249.072 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 2 },
	{ config_id = 13009, pos = { x = -10.196, y = 0.039, z = 246.520 }, rot = { x = 0.000, y = 90.000, z = 0.000 }, tag = 2 },
	{ config_id = 13010, pos = { x = 10.484, y = 0.039, z = 246.171 }, rot = { x = 0.000, y = 270.000, z = 0.000 }, tag = 2 },
	{ config_id = 13011, pos = { x = -10.246, y = 0.039, z = 243.425 }, rot = { x = 0.000, y = 90.000, z = 0.000 }, tag = 2 },
	{ config_id = 13012, pos = { x = 0.040, y = 0.039, z = 252.151 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 4 },
	{ config_id = 13013, pos = { x = -4.856, y = 0.039, z = 250.185 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 4 },
	{ config_id = 13014, pos = { x = 5.094, y = 0.039, z = 250.359 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 4 },
	{ config_id = 13015, pos = { x = -8.256, y = 0.039, z = 247.088 }, rot = { x = 0.000, y = 120.000, z = 0.000 }, tag = 4 },
	{ config_id = 13016, pos = { x = 9.280, y = 0.039, z = 246.067 }, rot = { x = 0.000, y = 270.000, z = 0.000 }, tag = 4 },
	{ config_id = 13017, pos = { x = -7.807, y = 0.039, z = 241.431 }, rot = { x = 0.000, y = 80.000, z = 0.000 }, tag = 4 },
	{ config_id = 13018, pos = { x = 8.691, y = 0.039, z = 240.151 }, rot = { x = 0.000, y = 270.000, z = 0.000 }, tag = 4 },
	{ config_id = 13019, pos = { x = -4.800, y = 0.039, z = 237.620 }, rot = { x = 0.000, y = 60.000, z = 0.000 }, tag = 4 },
	{ config_id = 13020, pos = { x = 5.272, y = 0.039, z = 236.810 }, rot = { x = 0.000, y = 300.000, z = 0.000 }, tag = 4 },
	{ config_id = 13021, pos = { x = 0.032, y = 0.039, z = 234.756 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, tag = 4 },
	{ config_id = 13022, pos = { x = -9.975, y = 0.039, z = 252.493 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 8 },
	{ config_id = 13023, pos = { x = 10.091, y = 0.039, z = 252.493 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 8 },
	{ config_id = 13024, pos = { x = -9.975, y = 0.039, z = 233.467 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, tag = 8 },
	{ config_id = 13025, pos = { x = 10.091, y = 0.039, z = 233.467 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, tag = 8 },
	{ config_id = 13026, pos = { x = -8.241, y = 0.039, z = 252.493 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 8 },
	{ config_id = 13027, pos = { x = 11.825, y = 0.039, z = 252.493 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 8 },
	{ config_id = 13028, pos = { x = -8.241, y = 0.039, z = 233.467 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, tag = 8 },
	{ config_id = 13029, pos = { x = 11.825, y = 0.039, z = 233.467 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, tag = 8 },
	{ config_id = 13030, pos = { x = -12.657, y = 0.039, z = 252.493 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 8 },
	{ config_id = 13031, pos = { x = 7.409, y = 0.039, z = 252.493 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 8 },
	{ config_id = 13032, pos = { x = -12.657, y = 0.039, z = 233.467 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, tag = 8 },
	{ config_id = 13033, pos = { x = 7.409, y = 0.039, z = 233.467 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, tag = 8 },
	{ config_id = 13034, pos = { x = -1.155, y = 0.039, z = 246.503 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 16 },
	{ config_id = 13035, pos = { x = 1.077, y = 0.039, z = 246.503 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 16 },
	{ config_id = 13036, pos = { x = -3.068, y = 0.039, z = 246.503 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 16 },
	{ config_id = 13037, pos = { x = 3.123, y = 0.039, z = 246.503 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 16 },
	{ config_id = 13038, pos = { x = -4.940, y = 0.039, z = 246.503 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 16 },
	{ config_id = 13039, pos = { x = 4.961, y = 0.039, z = 246.503 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 16 },
	{ config_id = 13040, pos = { x = -3.185, y = 0.039, z = 244.131 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 16 },
	{ config_id = 13041, pos = { x = 2.861, y = 0.039, z = 244.131 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 16 },
	{ config_id = 13042, pos = { x = 1.077, y = 0.039, z = 244.131 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 16 },
	{ config_id = 13043, pos = { x = -1.155, y = 0.039, z = 244.131 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 16 },
	{ config_id = 13044, pos = { x = -0.065, y = 0.054, z = 247.238 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, tag = 32 }
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
		regions = { 13001 },
		triggers = { "ENTER_REGION_13001" },
		rand_weight = 100
	}
}

--================================================================
-- 
-- 触发器
-- 
--================================================================

require "V2_2/RogueDungeon_Require"