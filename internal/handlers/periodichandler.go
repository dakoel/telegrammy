package handlers

import (
	"log/slog"
	"telegrammy/internal/config"
	"telegrammy/internal/domain"
	"telegrammy/internal/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
)

func SetUpPeriodicJobs() {
	periodicJobs := config.GetPeriodicJobs()
	for _, periodicJob := range periodicJobs {
		setUpPeriodicJob(periodicJob)
	}
}

func setUpPeriodicJob(periodicJob domain.PeriodicJob) {
	cron := cron.New()
	cron.AddFunc(periodicJob.Schedule, func() {
		message := tgbotapi.NewMessage(config.GetTelegramChatId(), *periodicJob.GetMessage(config.GetShellPath()))
		message.ParseMode = periodicJob.ParseMode
		slog.Info("Sending scheduled message.", "content", message.Text)
		telegram.SendMessage(&message)
	})
	cron.Start()
}
