package initservice

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"site_board_monitor/config"
)

type (
	Xredis interface {
		Subscribe(key string)
		PushBoardsByPipeline([]string, error)
	}

	redisOp struct {
		rdb  *redis.Client
		db   SiteBoard
		conf config.RedisConfig
	}

	boardMsg struct {
		SiteBoardUUID string `json:"site_board_uuid"`
		Msg           string `json:"msg"`
	}
)

func NewReid() Xredis {
	return &redisOp{rdb: Rdb,
		db: NewBoard()}
}

func (r *redisOp) PushBoardsByPipeline(alldata []string, err error) {
	_, err = r.rdb.Pipelined(context.Background(), func(pipe redis.Pipeliner) error {
		for _, v := range alldata {
			pipe.LPush(context.Background(), r.conf.NewsKey, v)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func (r *redisOp) Subscribe(key string) {
	fmt.Println(key)
	sub := r.rdb.Subscribe(context.Background(), key)
	_, err := sub.Receive(context.Background())
	if err != nil {
		zap.L().Error(fmt.Sprintf("redis 订阅 error:%s", err.Error()))
		panic(err)
	}

	ch := sub.Channel()
	for msg := range ch {
		var m boardMsg
		err := json.Unmarshal([]byte(msg.Payload), &m)
		if err != nil {
			zap.L().Error(fmt.Sprintf("反序列化msg error:%s", msg.Payload))
			continue
		}
		fmt.Println("更新板块状态：", m.SiteBoardUUID)
		err = r.db.UpdateBoardByUUID(m.SiteBoardUUID, m.Msg)
		if err != nil {
			zap.L().Error("")
			return
		}

	}

}
