package gin

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	addr string // 启动ip:port，空则默认是 localhost:8080
	c    chan error
}

func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
		c:    make(chan error),
	}
}

func (m *Server) IsBlock(ctx context.Context) (isBlock bool) {
	return true
}

func (m *Server) Start(ctx context.Context) (err error) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/wechat/official/account/callback", func(c *gin.Context) {
		data, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data":    string(data),
			"message": "ok",
		})
	})
	return r.Run(m.addr)
}

func (m *Server) AsyncStart(ctx context.Context) {
	go func() {
		err := m.Start(ctx)
		if err != nil {
			fmt.Println("[AsyncStart] http Start panic", err)
			m.c <- err
		}
	}()
}

func (m *Server) ErrorC() (c chan error) {
	return m.c
}
