package server

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/threading"
	"site_board_monitor/service/initservice"
	"time"
)

func StartServer() {
	srv := initservice.NewReid()
	threading.GoSafe(pushSrv)
	srv.Subscribe(initservice.Conf.Redis.BoardErrorKey)
}

func pushSrv() {
	srv := initservice.NewReid()

	ch := make(chan string)

	threading.GoSafe(func() {
		for {
			ticker := time.Tick(time.Minute * 5)
			for {
				<-ticker
				ch <- "开始推送"
			}
		}
	})

	for {
		select {
		case <-ch:
			fmt.Println("开始推送")
			srv.PushBoardsByPipeline()
		}
	}

}
