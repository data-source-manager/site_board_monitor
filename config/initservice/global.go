package initservice

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	SqlService *gorm.DB
	Rdb        *redis.Client
)
