package Message

import (
	"bytes"
	"encoding/json"
	"github.com/51daticom/marx-agent/Configs"

	"net/http"
)


func DingSend(message string, config *Configs.Config) {
	body := make(map[string]interface{})
	body["msgtype"] = "text"
	body["text"] = map[string]string{
		"content": message,
	}
	b, _ := json.Marshal(body)
	http.DefaultClient.Post(config.DingPush, "application/json", bytes.NewBuffer(b))
}
