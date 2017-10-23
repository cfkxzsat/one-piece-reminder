package main

import (
	"github.com/cfkxzsat/one-piece-reminder/reminder"
	"github.com/cfkxzsat/one-piece-reminder/wechat"
)

func main() {
	go wechat.RunService()
	go reminder.Cron()
}
