package main

import (
	customJob "github.com/fancyshenmue/golang-gin-sample/cmd/schedule/job"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds())

	c.AddFunc("0 0 3 * * ?", customJob.SyncCloudflareZoneList)
	c.AddFunc("0 10 3 * * ?", customJob.SyncCloudflareZoneRecord)
	c.AddFunc("0 0 3 * * ?", customJob.SyncVerizonHTTPLargeList)

	c.Start()

	defer c.Stop()

	select {}
}
