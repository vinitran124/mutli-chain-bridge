package api

import (
	"bridge/app/utils"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"strings"
	"time"
)

func NewCommand() *cli.Command {
	return &cli.Command{
		Name:  "api",
		Usage: "Start the API server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  argsAddr,
				Value: "0.0.0.0:3030",
				Usage: "serve address",
			},
		},

		Action: func(c *cli.Context) error {
			return startAPIServer(c)
		},

		Before: func(c *cli.Context) error {
			return beforeStartApiServer(c)
		},
	}
}

func beforeStartApiServer(c *cli.Context) error {
	utils.SetContextSQL()
	utils.SetContextRedisClient()
	utils.SetChainClient()
	return nil
}

func startAPIServer(c *cli.Context) error {
	fmt.Println("start")
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	addr := strings.ToLower(c.String(argsAddr))
	if addr == "" {
		return fmt.Errorf("[API Server] start error: addr is empty")
	}
	fmt.Printf("ListenAndServe: %s\n", addr)

	v1 := router.Group("/api/v1")
	v1Router(v1)

	return router.Run(addr)
}
