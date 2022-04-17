package main

import (
	"fmt"
	"net/http"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/SinclearClan/SomeBotTelegram/telegram"
	"github.com/SinclearClan/SomeBotTelegram/config"
	"github.com/SinclearClan/SomeBotTelegram/webserver"
	"github.com/gin-gonic/gin"
)

var (
	cfg = config.GetConfig()
)

func main() {
	
	tgbot, err := gotgbot.NewBot(cfg.Telegram.Key, &gotgbot.BotOpts{
		Client: http.Client{},
		GetTimeout: gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	})
	if err != nil {
		panic(err)
	}

	updater := ext.NewUpdater(&ext.UpdaterOpts{
		ErrorLog: nil,
		DispatcherOpts: ext.DispatcherOpts{
			Error: func(tgbot *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
				fmt.Println("an error occurred while handling update:", err.Error())
				return ext.DispatcherActionNoop
			},
			Panic:       nil,
			ErrorLog:    nil,
			MaxRoutines: 0,
		},
	})
	
	dispatcher := updater.Dispatcher

	dispatcher.AddHandler(handlers.NewCommand("start", telegram.Start))
	dispatcher.AddHandler(handlers.NewCommand("dhl", telegram.DHL))

	err = updater.StartPolling(tgbot, &ext.PollingOpts{DropPendingUpdates: true})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s has been started...\n", tgbot.User.Username)

	// init gin
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.SetTrustedProxies(nil)
	webserver.RegisterRoutes(engine, tgbot)

	// start gin (this also does the same as "updater.Idle()" at the same time, so the Telegram tgbot will still run)
	engine.Run(":" + cfg.Webserver.Port)
	

}
