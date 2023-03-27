package initservice

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type (
	Xredis interface {
		Subscribe(key string)
	}

	redisOp struct {
		rdb *redis.Client
	}

	boardMsg struct {
		SiteBoardUUID string `json:"site_board_uuid"`
		Msg           string `json:"msg"`
	}
)

func NewReid() Xredis {
	return &redisOp{rdb: Rdb}
}

func (r *redisOp) Subscribe(key string) {
	fmt.Println(key)
	sub := r.rdb.Subscribe(context.Background(), key)
	_, err := sub.Receive(context.Background())
	if err != nil {
		zap.L().Error(fmt.Sprintf("redis 订阅 error:%s", err.Error()))
		panic(err)
	}
	board := NewBoard()

	ch := sub.Channel()
	for msg := range ch {
		var m boardMsg
		err := json.Unmarshal([]byte(msg.Payload), &m)
		if err != nil {
			zap.L().Error(fmt.Sprintf("反序列化msg error:%s", msg.Payload))
			continue
		}
		fmt.Println("更新板块状态：", m.SiteBoardUUID)
		err = board.UpdateBoardByUUID(m.SiteBoardUUID, m.Msg)
		if err != nil {
			zap.L().Error("")
			return
		}

	}

}
