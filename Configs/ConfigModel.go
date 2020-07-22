package Configs

import (
	"regexp"
	"strings"
	"time"
)

type Config struct {
	Buf       int
	WiteList  []string
	BlackList []string
	Log       string
	WxPush    string
	DingPush  string
}

func (c *Config) GetLog() string {
	reg, _ := regexp.CompilePOSIX(`\{([YMD]{0,4})\}`)
	mapt := reg.FindAllString(c.Log, -1)
	now := time.Now()
	if mapt == nil || len(mapt) <= 0 {
		return c.Log
	}
	newLog := c.Log
	for _, temp := range mapt {
		nt := strings.Trim(temp, "{")
		nt = strings.Trim(nt, "}")
		if nt[0] == 'Y' {
			year := now.Format("2006")
			start := 4 - len(nt)
			Y := year[start:]
			newLog = strings.Replace(newLog, temp, Y, 3)
		}
		if nt[0] == 'M' {
			month := now.Format("01")
			start := 2 - len(nt)
			m := month[start:]
			newLog = strings.Replace(newLog, temp, m, 3)
		}
		if nt[0] == 'D' {
			day := now.Format("02")
			start := 2 - len(nt)
			d := day[start:]
			newLog = strings.Replace(newLog, temp, d, 3)
		}
	}
	return newLog
}
