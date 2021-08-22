package main

import (
	customApp "golang-gin-sample/cmd/schedule/cdn/app"

	cron "github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds())

	c.AddFunc(customApp.ScheduleSyncCloudflareZoneListCron, customApp.SyncCloudflareZoneList)
	c.AddFunc(customApp.ScheduleSyncCloudflareZoneRecord, customApp.SyncCloudflareZoneRecord)
	c.AddFunc(customApp.ScheduleSyncVerizonHttpLargeList, customApp.SyncVerizonHTTPLargeList)

	c.Start()

	defer c.Stop()

	select {}
}
