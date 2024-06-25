

* 微信公众号相关接口
* 服务器URL
* 启动
```shell
go run . -web_addr=localhost:23456
curl -XGET localhost:23456/ping
curl -XPOST localhost:23456/wechat/official/account/callback -d '{}'
```



