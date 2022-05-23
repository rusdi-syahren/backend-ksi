package main

import (
	"fmt"
	"os"
	"sync"

	"gitlab.com/k1476/scaffolding/config"
	"gitlab.com/k1476/scaffolding/config/database"
)

func main() {

	// call config.Load() before start up
	err := config.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	writeDB, err := database.GetGormConn(config.ReadDBHost, config.ReadDBUser, config.ReadDBName, config.ReadDBPassword, 5432)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	readDB, err := database.GetGormConn(config.WriteDBHost, config.WriteDBUser, config.WriteDBName, config.WriteDBPassword, 5432)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	redis, err := database.GetRedis(config.RedisHost, config.RedisTLS, config.RedisPassword, config.RedisPort)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	echoServer, err := NewEchoServer(writeDB, readDB, redis)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		echoServer.Run()
		wg.Done()
	}()

	wg.Wait()

}
