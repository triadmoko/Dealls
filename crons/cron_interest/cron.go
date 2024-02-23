package cron_interest

import (
	"app/domain"

	"github.com/sirupsen/logrus"
)

type CronInterest struct {
	logger       *logrus.Logger
	repoInterest domain.RepositoryInterest
}

func NewCronInterest(logger *logrus.Logger, repoInterest domain.RepositoryInterest) *CronInterest {
	return &CronInterest{
		logger:       logger,
		repoInterest: repoInterest,
	}
}
