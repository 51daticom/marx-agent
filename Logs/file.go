package Logs

import (
	"bytes"
	"github.com/howeyc/fsnotify"
	"log"
	"github.com/51daticom/marx-agent/Configs"
	"github.com/51daticom/marx-agent/Message"

	"os"
	"regexp"
	"strings"
)

var filebuf = make(chan byte, 2>>30)
var slip = []byte("\n")

var lines = make(chan string, 1000)

func getLog(config *Configs.Config) {
	f, err := os.Open(config.GetLog())

	defer f.Close()
	if err != nil {
		panic(err.Error())
	}
	watcher, _ := fsnotify.NewWatcher()
	fileInfo, _ := f.Stat()
	offset := fileInfo.Size()
	for {
		watchError := watcher.Watch(config.GetLog())
		if watchError != nil {
			log.Println(watchError.Error())
		} else {
			select {
			case <-watcher.Event:
				f2, err := os.Open(config.GetLog())
				if err != nil {
					log.Println("open file error:" + err.Error())
				} else {
					for {
						var tmp [10240]byte
						n, _ := f2.ReadAt(tmp[:], offset)
						if n == 0 {
							f2.Close()
							break
						}
						offset = offset + (int64(n))
						for _, v := range tmp[0:n] {
							filebuf <- v
						}
					}
				}
			case err := <-watcher.Error:
				log.Println("watcher error:", err)
			}
		}
	}
	log.Println("end")

}

func readFileBuf() {
	var d []byte
	s := len(slip)
	for {
		select {
		case b := <-filebuf:
			d = append(d, b)
			td := len(d)
			if td >= s {
				if bytes.Equal(d[td-s:td], slip) {
					lines <- string(d)
					d = []byte{}
				}
			}
		}
	}
}

func DoLine(config *Configs.Config) {
	go func() {
		for {
			select {
			case l := <-lines:
				if len(config.BlackList) <= 0 {
					break
				}
				for _, v := range config.BlackList {
					v = strings.Trim(v, `"`)
					reg, _ := regexp.CompilePOSIX(v)
					find := reg.FindAllString(l, -1)
					if len(find) > 0 {
						if config.WxPush != "" {
							Message.WxSend(find[0]+": "+l, config)
						}
						if config.DingPush != "" {
							Message.DingSend(find[0]+": "+l, config)
						}
					}
				}
			}
		}
	}()
	go readFileBuf()
	getLog(config)
}
