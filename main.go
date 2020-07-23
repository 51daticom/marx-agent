package main

import (
	"flag"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/51daticom/marx-agent/Configs"
	"github.com/51daticom/marx-agent/Logs"
	"strings"
)

var config = new(Configs.Config)

func initConfig(f string, t string) {
	data, _ := ini.Load(f)
	section, _ := data.GetSection(t)
	log, _ := section.GetKey("log")
	config.Log = log.String()
	buf, _ := section.GetKey("buf")
	config.Buf, _ = buf.Int()
	wxpush, _ := section.GetKey("wxpush")
	config.WxPush = wxpush.String()
	dingpush, _ := section.GetKey("dingpush")
	config.DingPush = dingpush.String()
	whiteList, _ := section.GetKey("whiteList")
	if whiteList.String() == "" {
		config.WiteList = []string{}
	} else {
		config.WiteList = strings.Split(whiteList.String(), ",")
	}
	blackList, _ := section.GetKey("blackList")
	if blackList.String() == "" {
		config.WiteList = []string{}
	} else {
		config.BlackList = strings.Split(blackList.String(), ",")
	}
}

func main() {

	flag.Parse()
	f := flag.Arg(0)
	t := flag.Arg(1)
	initConfig(f, t)
	fmt.Println(f)
	Logs.DoLine(config)
}
