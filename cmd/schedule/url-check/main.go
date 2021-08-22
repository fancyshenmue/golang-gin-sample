package main

import (
	customApp "sre-app/cmd/schedule/url-check/app"

	cron "github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc(customApp.V.GetStringMapString("schedule")["check_app_endpoint_schedule_time"], customApp.CheckURLStatus)

	c.Start()

	select {}
}
