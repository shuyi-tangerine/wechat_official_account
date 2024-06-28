

* 微信公众号相关接口
* 服务器URL
* 启动
```shell
go run . -web_addr=:80
curl -XGET localhost:80/ping
curl -XPOST localhost/wechat/official/account/callback -d '{}'
```
* 查看是否在公网启动 netstat -ntpl


