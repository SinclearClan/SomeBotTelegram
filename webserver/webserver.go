package webserver

import (
	"strconv"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/gin-gonic/gin"
	"github.com/SinclearClan/SomeBotTelegram/telegram"
	"github.com/SinclearClan/SomeBotTelegram/config"
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
		}
		if ctx.Query("key") == "" {
			ctx.JSON(400, gin.H{
				"message": "No key provided",
			})
			return
		}
		if ctx.Query("status") == "" {
			ctx.JSON(400, gin.H{
				"message": "No status provided",
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
				if ctx.Query("status") == "arrived" {
					telegram.Notify(ctx.Query("status"), chatId, tgbot)
					ctx.JSON(200, gin.H{
						"message": "OK",
					})
				} else if ctx.Query("status") == "b10b14" {
					if ctx.Query("loc_lat") == "" {
						ctx.JSON(400, gin.H{
							"message": "No latitude provided",
						})
						return
					}
					if ctx.Query("loc_lon") == "" {
						ctx.JSON(400, gin.H{
							"message": "No longitude provided",
						})
						return
					}
	
					lat, err := strconv.ParseFloat(ctx.Query("loc_lat"), 64)
					if err != nil {
						panic(err)
					}
					lon, err := strconv.ParseFloat(ctx.Query("loc_lon"), 64)
					if err != nil {
						panic(err)
					}
	
					err = telegram.NotifyWithLocation(ctx.Query("status"), lat, lon, chatId, tgbot)
					if err != nil {
						panic(err)
					}
	
					ctx.JSON(200, gin.H{
						"message": "OK",
					})
				}
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

	engine.GET("/sendLocation", func(ctx *gin.Context) {
		
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
		} else if ctx.Query("loc_lat") == "" {
			ctx.JSON(400, gin.H{
				"message": "Insuficcient location provided, missing latitude",
			})
			return
		} else if ctx.Query("loc_lon") == "" {
			ctx.JSON(400, gin.H{
				"message": "Insuficcient location provided, missing longitude",
			})
			return
		} else if ctx.Query("loc_acc") == "" {
			ctx.JSON(400, gin.H{
				"message": "Insuficcient location provided, missing accuracy",
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
				lat, err := strconv.ParseFloat(ctx.Query("loc_lat"), 64)
				if err != nil {
					panic(err)
				}
				lon, err := strconv.ParseFloat(ctx.Query("loc_lon"), 64)
				if err != nil {
					panic(err)
				}
				acc, err := strconv.ParseFloat(ctx.Query("loc_acc"), 64)
				if err != nil {
					panic(err)
				}
				telegram.SendLocation(chatId, lat, lon, acc, tgbot)
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
