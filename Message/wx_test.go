package Message

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestSend(t *testing.T) {
	//WxSend()
}

func TestReg(t *testing.T)  {
	v := `47.115.174.5 - - [21/Jul/2020:05:57:48 +0800] "GET /thinkphp/html/public/index.php?404 HTTP/1.1" 404 47 "-" "Mozilla/5.0 (Windows; U; Windows NT 6.0;en-US; rv:1.9.2) Gecko/20100115 Firefox/3.6)" "-" "0.001" "0.001"
47.115.174.5 - - [21/Jul/2020:05:57:48 +0800] "GET /thinkphp/html/public/index.php HTTP/1.1" 500 47 "-" "Mozilla/5.0 (Windows; U; Windows NT 6.0;en-US; rv:1.9.2) Gecko/20100115 Firefox/3.6)" "-" "0.001" "0.001"`
	reg, _ := regexp.CompilePOSIX("\\ 404\\ |\\ 500\\ ")
	find := reg.FindAllString(v, -1)
	fmt.Println(find)
}

func TestTrimSend(t *testing.T) {
	a := `"\\ 404 \\ "`
	b := strings.Trim(a,`"`)
	fmt.Println(b)
}

func TestGetLog(t *testing.T){
	l := "{YYYY}-{MM}-{DD}"
	reg,_ := regexp.CompilePOSIX(`\{([YMD]{0,4})\}`)
	mapt := reg.FindAllString(l,-1)
	fmt.Println(mapt)
}