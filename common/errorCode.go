package common

// 错误码
const (
	ERR_NO_GATEWAY           = 1
	ERR_NO_GATEWAY_PORT      = 2
	ERR_ETCD_GET_PORT        = 3
	ERR_NO_CONNECT_EXIST     = 3
	ERR_NO_ROOMID_EXIST      = 4
	ERR_IS_NOT_ROOM_MASATER  = 5
	ERR_IS_ALREADY_IN_ROOM   = 6
	ERR_IS_ALREADY_IN_GAME   = 7
	ERR_IS_NOT_IN_ROOM       = 8
	ERR_IS_NOT_IN_GAMEROOM   = 9
	ERR_IS_NO_PREPARE_EXIST  = 10
	ERR_ROOM_PLAYERNUM_FULL  = 11
	ERR_ROOM_POSTION_USING   = 12
	ERR_PLAYERNUM_NOT_ENOUGH = 13
	ERR_GAME_POSTION_PARAM   = 14

	ERR_REDIS_WRITE_ERROR = 10000
	ERR_REDIS_QUERY       = 10001
	ERR_SDK_GET_USERINFO  = 10002
)
