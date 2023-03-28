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

}
