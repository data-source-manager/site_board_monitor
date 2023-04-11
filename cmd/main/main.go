package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"site_board_monitor/service/initservice"
)

var configFile = flag.String("f", "config.yaml", "the config file")

func main() {
	flag.Parse()
	conf.MustLoad(*configFile, &initservice.Conf, conf.UseEnv())

	initservice.InitService(initservice.Conf)

	board := initservice.NewReid()
	//initservice.Rdb.LPush(context.Background(), "BasicSpider:start_urls", `{"meta":{"appid":5024,"country":"uk","site_name":"未来论坛","site_uuid":"","board_uuid":"4a8f89a4-fe6f-11ec-a30b-d4619d029786","board_name":"新闻稿","post_url":"","post_data":"","header":"","board_url":"https://www.forumforthefuture.org/press-releases"},"url":"https://www.forumforthefuture.org/press-releases"}`)
	board.PushBoardsByPipeline()
	//board.Subscribe("news:board:error")
}
