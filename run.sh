

nohup go run . -web_addr=:80 >run.log 2>&1 &

curl -XGET localhost:80/ping
