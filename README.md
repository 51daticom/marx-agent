# marx-agent
一款日志文件监控报警（支持企业微信机器人、钉钉机器人）工具，体积小巧仅7M，部署简单仅需5分钟。

# 快速使用

### 编译

* 方法一：

```
git clone https://github.com/51daticom/marx-agent.git
cd marx-agent
export GO111MODULE=on
export GOPROXY=https://goproxy.io
go build -o marx-agent
```
* 方法二：

```
go get -u github.com/51daticom/marx-agent
ls {GORPATH}/bin/marx-agent
```
### 配置

```
mv config.nginx.example.in config.ini;
vim config.in
```

```
[pro]
buf = 1
whiteList = ""
blackList = "\ 500\ ","\ 502\ ","\ 501\ " #监控报警的状态码（正则匹配）
log = /var/log/nginx/access.log  #监控的日志文件路径
wxpush = https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key={{youkey}} #企业微信机器人webhook地址
dingpush =  #钉钉机器人webhook地址

#/var/log/nginx/access.log format data such as:
#127.0.0.1 - - [21/Jul/2020:05:57:48 +0800] "GET /thinkphp/html/public/index.php HTTP/1.1" 500 47 "-" "Mozilla/5.0 (Windows; U; Windows NT 6.0;en-US; rv:1.9.2) Gecko/20100115 Firefox/3.6)" "-" "0.001" "0.001

```

### 启动监控

```
./max-agent config.ini pro
```
