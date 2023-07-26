package main

import (
	"bridge/app/content/bob"
	"bridge/app/utils"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"strings"
	"time"
)

const (
	envPath = ".env,.env.local"
)

func init() {
	if err := godotenv.Overload(strings.Split(envPath, ",")...); err != nil {
		fmt.Println("Load env error", err.Error())
	}
}

func main() {
	utils.SetContextSQL()
	for {
		expiredRequest, err := bob.BridgeRequests(context.Background(),
			SQLRepository(),
			bob.SelectWhere.BridgeRequests.CreatedAt.LTE(time.Now().Add(-10*time.Minute)),
		).All()
		if err != nil {
			log.Println(err)
			continue
		}

		_, err = expiredRequest.DeleteAll(context.Background(), SQLRepository())
		if err != nil {
			log.Println(err)
			continue
		}
		time.Sleep(10 * time.Minute)
	}
}
