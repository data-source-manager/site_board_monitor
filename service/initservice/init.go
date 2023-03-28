package initservice

import (
	"site_board_monitor/config"
	"site_board_monitor/service/logservice"
	"site_board_monitor/service/mysqlservice"
	"site_board_monitor/service/redisservice"
)

func InitService(c config.Config) {
	SqlService = mysqlservice.InitMysql(c.Mysql)
	logservice.InitLogger(c.Log)
	Rdb = redisservice.InitRedis(c.Redis)
}
