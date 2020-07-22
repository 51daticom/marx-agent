package Configs

import (
	"fmt"
	"testing"
)

func TestConfig_GetLog(t *testing.T) {
	c := new(Config)
	c.Log = "name-{YYYY}-{MM}-{DD}.log"
	fmt.Println(c.GetLog())
}
