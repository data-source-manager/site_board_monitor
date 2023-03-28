package initservice

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"site_board_monitor/config"
)

var (
	SqlService *gorm.DB
	Rdb        *redis.Client
	Conf       config.Config
)
