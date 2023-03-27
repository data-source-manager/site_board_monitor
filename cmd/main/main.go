package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"site_board_monitor/config"
	"site_board_monitor/config/initservice"
)

var configFile = flag.String("f", "config.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	initservice.InitService(c)

	//var board mysqlservice.SiteBoard
	//initservice2.SqlService.Find(&board)
	//fmt.Println(board)
	xredis := initservice.NewReid()
	xredis.Subscribe(c.Redis.BoardErrorKey)

}
