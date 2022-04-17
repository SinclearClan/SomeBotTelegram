package webserver

import (
	"strconv"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/SinclearClan/SomeBotTelegram/telegram"
	"github.com/SinclearClan/SomeBotTelegram/config"
	"github.com/gin-gonic/gin"
)

var (
	cfg = config.GetConfig()
)

func RegisterRoutes(engine *gin.Engine, tgbot *gotgbot.Bot) {
	
	engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	engine.GET("/notify", func(ctx *gin.Context) {
		
		if ctx.Query("id") == "" {
			ctx.JSON(400, gin.H{
				"message": "No id provided",
			})
			return
		} else if ctx.Query("key") == "" {
			ctx.JSON(400, gin.H{
				"message": "No key provided",
			})
			return
		}


		// if username is in whitelist, notify
		chatId, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
		if err != nil {
			panic(err)
		}

		if ctx.Query("key") == cfg.Webserver.Key {

			if telegram.IsOnWhitelist(chatId) {
				telegram.Notify(tgbot, chatId)
				ctx.JSON(200, gin.H{
					"message": "OK",
				})
			} else {
				ctx.JSON(406, gin.H{
					"message": "User not in whitelist",
				})
			}
		
		} else {
			ctx.JSON(401, gin.H{
				"message": "Wrong key",
			})
		}
	})

}
