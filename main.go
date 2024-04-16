package main

import (
	"github.com/aifuxi/banking/app"
	"github.com/aifuxi/banking/logger"
)

func main() {

	logger.Info("Starting banking app...")
	app.Start()
}
