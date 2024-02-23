package cron_interest

import (
	"app/pkg"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func (c *CronInterest) DeleteInterestAfterOneDay() {
	schedule, err := gocron.NewScheduler(gocron.WithLocation(time.Local))
	if err != nil {
		c.logger.Error(err)
		return
	}
	j, err := schedule.NewJob(
		gocron.DurationJob(
			5*time.Second,
		),
		gocron.NewTask(
			c.repoInterest.DeletePartnerExpired,
		),
	)
	if err != nil {
		c.logger.Error(err)
		return
	}
	pkg.Prt(j.Name())
	schedule.Start()
}
