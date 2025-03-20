package main

import (
	"fmt"

	"github.com/lmbaila/gopportunities/config"
	"github.com/lmbaila/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		fmt.Println(err)
		return

	}
	router.Initialize()
}
