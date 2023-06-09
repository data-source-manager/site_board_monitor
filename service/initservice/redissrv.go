package initservice

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"site_board_monitor/common"
	"site_board_monitor/config"
)

type (
	Xredis interface {
		Subscribe(key string)
		PushBoardsByPipeline()
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
		db:   NewBoard(),
		conf: Conf.Redis}
}

// PushBoardsByPipeline 板块推送
func (r *redisOp) PushBoardsByPipeline() {
	fmt.Println("开始推送板块数据")
	allData := r.db.QueryPushBoard()
	pipe := r.rdb.Pipeline()
	for _, v := range allData {
		boardMap := make(map[string]interface{})
		boardMap["meta"] = v
		boardMap["url"] = v.BoardUrl
		mapStr, _ := json.Marshal(boardMap)
		pipe.LPush(context.Background(), r.conf.NewsKey, string(mapStr))
	}
	_, err := pipe.Exec(context.Background())
	if err != nil {
		zap.L().Error(fmt.Sprintf("板块推送失败,%s", err.Error()))
		return
	}
	zap.L().Info(fmt.Sprintf("[%s] 一共推送了 %d 条数据", common.GetNowTimeStr(), len(allData)))
}

// Subscribe 更新板块的状态
func (r *redisOp) Subscribe(key string) {
	fmt.Println("订阅板块状态：" + key)
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
