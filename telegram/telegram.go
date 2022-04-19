package telegram

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/SinclearClan/SomeBotTelegram/config"
	"github.com/SinclearClan/SomeBotTelegram/dhl"
)

var (
	cfg = config.GetConfig()
)

func Start(bot *gotgbot.Bot, ctx *ext.Context) error {

	_, err := ctx.EffectiveMessage.Reply(
		bot, 
		fmt.Sprintf("Hallo, Ich bin <b>%s</b>. Ich kann ein paar praktische Dinge. Marc hat dich sicherlich schon mit Details genervt.", bot.User.FirstName), 
		&gotgbot.SendMessageOpts{ParseMode: "html",},
	)
	if err != nil {
		return err
	}

	return nil

}

// @Command
func Where(bot *gotgbot.Bot, ctx *ext.Context) error {

	if IsOnWhitelist(ctx.EffectiveChat.Id) {
		resp, err := http.Get(cfg.Macrodroid.RestUrl + "/GetLocation?chatId=" + fmt.Sprintf("%d", ctx.EffectiveChat.Id))
		if err != nil {
			return err
		}
		defer resp.Body.Close()
	}

	return nil

}

func SendLocation(chatId int64, lat float64, lon float64, acc float64, bot *gotgbot.Bot) error {
	
	_, err := bot.SendLocation(chatId, lat, lon, &gotgbot.SendLocationOpts{HorizontalAccuracy: acc})
	if err != nil {
		return err
	}

	return nil

}

func Notify(chatId int64, bot *gotgbot.Bot) error {

	_, err := bot.SendMessage(chatId, "Marc nähert sich deinem Standort!", nil)
	if err != nil {
		return err
	}

	return nil
	
}

func DHL(bot *gotgbot.Bot, ctx *ext.Context) error {

	var args = strings.Split(ctx.EffectiveMessage.Text, " ")
	if len(args) < 2 {
		_, err := ctx.EffectiveMessage.Reply(bot, "Bitte gib eine Tracking-Nummer an.", nil)
		if err != nil {
			return err
		}
	} else {
		packageInfo, err := dhl.GetDHLPackageInfo(args[1])
		if err != nil {
			panic(err)
		}

		if packageInfo.Shipments == nil {
			_, err := ctx.EffectiveMessage.Reply(bot, "Es konnten keine Informationen über dieses Paket gefunden werden. Entweder ist die Tracking-Nummer falsch oder es wurde noch nicht abgeschickt.", nil)
			if err != nil {
				return err
			}
			return nil
		}

		for _, shipment := range *packageInfo.Shipments {
			_, err := ctx.EffectiveMessage.Reply(
				bot, 
				fmt.Sprintf("<b>DHL Tracking Status</b>\nTracking-Nummer: %s\nStatus: %s\nAddressiert an: %s\nAnkunft: %s", shipment.Id, shipment.Status.Status, shipment.Destination.Address.AddressLocality, shipment.EstimatedTimeOfDelivery), 
				&gotgbot.SendMessageOpts{
					ParseMode: "html",
				},
			)
			if err != nil {
				return err
			}
		}
	}

	return nil

}

func IsOnWhitelist(chatId int64) bool {
	for _, v := range cfg.Telegram.Whitelist {
		if v == chatId {
			return true
		}
	}
	return false
}
