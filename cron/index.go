package cron

import (
	"fmt"
	"gin/models/test"
	"gin/tool/logger"
	"time"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func InitCron() {

	go func() {
		c := cron.New(cron.WithSeconds())

		logger.ZapLog.Info(
			"定时任务启动",
			zap.String("cron", "ok"),
		)

		c.AddFunc("1/30 * * * * *", func() {
			fmt.Println("Run models.CleanAllTag...")
			test.UpdateTest0301Cron()
		})

		c.Start()

		defer c.Stop()

		t1 := time.NewTimer(time.Second * 10)
		for {
			select {
			case <-t1.C:
				t1.Reset(time.Second * 10)
			}
		}

	}()

}
