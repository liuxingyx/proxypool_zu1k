package cron

import (
	"runtime"

	"github.com/jasonlvhit/gocron"
	"github.com/liuxingyx/proxypool_zu1k/internal/app"
)

func Cron() {
	_ = gocron.Every(10).Minutes().Do(crawlTask)
	<-gocron.Start()
}

func crawlTask() {
	_ = app.InitConfigAndGetters("")
	app.CrawlGo()
	app.Getters = nil
	runtime.GC()
}
