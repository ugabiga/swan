package core

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/spf13/cobra"
)
import "github.com/adhocore/gronx"

var (
	ErrCronExpressionInvalid = fmt.Errorf("invalid cron expression")
)

type CronJob struct {
	cronExpression string
	cronFunc       func() error
}

func NewCronJob(
	cronExpression string,
	cronFunc func() error,
) *CronJob {
	return &CronJob{
		cronExpression: cronExpression,
		cronFunc:       cronFunc,
	}
}

type CronTab struct {
	cronJobs []*CronJob
	logger   *slog.Logger
	gx       *gronx.Gronx
}

func NewCronTab(
	logger *slog.Logger,
) *CronTab {
	return &CronTab{
		logger: logger,
		gx:     gronx.New(),
	}
}

func (c *CronTab) RegisterCronJob(expression string, cronFunc func() error) {
	c.cronJobs = append(c.cronJobs, NewCronJob(expression, cronFunc))
}

func (c *CronTab) Start() {
	c.logger.Info("CronTab started")
	c.start()
}

func (c *CronTab) start() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := c.startJobs(); err != nil {
				c.logger.Error("Error", slog.Any("error", err))
				continue
			}
		}
	}
}

func (c *CronTab) startJobs() error {
	for _, cronJob := range c.cronJobs {
		cronExpression := "* " + cronJob.cronExpression

		if !c.gx.IsValid(cronExpression) {
			c.logger.Error("Error", slog.Any("error", ErrCronExpressionInvalid))
			continue
		}

		isDue, err := c.gx.IsDue(cronExpression, time.Now())
		if err != nil {
			c.logger.Error("Error", slog.Any("error", err))
			continue
		}

		if !isDue {
			continue
		}

		// Run asynchronously
		go func(cronJob *CronJob) {
			if err := cronJob.cronFunc(); err != nil {
				c.logger.Error("Error", slog.Any("error", err))
			}
		}(cronJob)

	}

	return nil
}

func InvokeSetCronCommand(
	crontab *CronTab,
	command *Command,
) {
	command.RegisterCommand(
		&cobra.Command{
			Use:   "cron",
			Short: "",
			Run: func(cmd *cobra.Command, args []string) {
				crontab.Start()
			},
		},
	)
}
