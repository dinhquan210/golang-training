package main

import (
	"golang-training/db"
	"golang-training/handler"
	"golang-training/repository/repo_impl"
	"os"
	"os/signal"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func main() {
	sql := &db.Sql{
		Host:     "127.0.0.1",
		Port:     5432,
		UserName: "postgres",
		PassWord: "postgres",
		DbName:   "golang-training",
	}

	sql.Connect()
	defer sql.Close()

	imageHandler := handler.ImageHandler{
		ImageRepo: repo_impl.NewImageRepo(sql),
	}
	cron := cron.New()
	cron.AddFunc("* * * * *", func() {
		err := imageHandler.CronJobRandomImage()
		if err != nil {
			logrus.Error(err)
		}

	})
	cron.Start()
	logrus.Info("Cronjob start successfully")
	defer cron.Stop()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
