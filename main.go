package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/shuyi-tangerine/wechat_official_account/gin"
	"os"
)

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

func main() {
	flag.Usage = Usage
	webAddr := flag.String("web_addr", "localhost:8080", "Address to listen to web server")
	flag.Parse()

	ginServer := gin.NewServer(*webAddr)
	err := ginServer.Start(context.Background())
	if err != nil {
		fmt.Println("start gin server", err)
		panic(err)
	}
	fmt.Println("main exit")
}
