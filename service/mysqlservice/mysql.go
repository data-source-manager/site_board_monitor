package mysqlservice

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"site_board_monitor/config"
)

func InitMysql(c config.MysqlConfig) (MysqlDB *gorm.DB) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.Db)
	var err error
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error(fmt.Sprintf("conn mysql failed,err:%v", err.Error()))
		panic(err.Error())
	}
	return MysqlDB

}
