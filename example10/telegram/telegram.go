package telegram

import (
	"fmt"

	"strings"

	"gopkg.in/telegram-bot-api.v4"
	"gobot.io/x/gobot/drivers/gpio"
	"github.com/spf13/viper"
	"os"
	"strconv"
)

var bot *tgbotapi.BotAPI
var rgbled *gpio.RgbLedDriver

var token string
var myChatID int64

func Start(led *gpio.RgbLedDriver) {

	rgbled = led
	readConfig()
	go func() {
		var err error
		bot, err = tgbotapi.NewBotAPI(token)
		if err != nil {
			fmt.Println("Error starting Telegram bot:", err)
		}
		bot.Debug = false

		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		updates, err := bot.GetUpdatesChan(u)

		for update := range updates {

			parse(update)

		}
	}()

}

func parse(update tgbotapi.Update) {

	if update.CallbackQuery != nil {
		chatID := int64(update.CallbackQuery.From.ID)
		if chatID != myChatID {
			return
		}

		command := strings.Split(update.CallbackQuery.Data, " ")
		bot.DeleteMessage(tgbotapi.DeleteMessageConfig{
			ChatID:    chatID,
			MessageID: update.CallbackQuery.Message.MessageID,
		})


		switch(command[0]) {
		case "red":
			setRGB(255,0,0)
			break
		case "yellow":
			setRGB(0,255,255)
			break
		case "green":
			setRGB(0,255,0)
			break
		case "blue":
			setRGB(0,0,255)
			break
		case "purple":
			setRGB(255,0,255)
			break
		case "black":
			setRGB(0,0,0)
			break
		default:
			break
		}

		help(chatID)


	}

	if update.Message != nil {

		chatID := update.Message.Chat.ID
		if chatID != myChatID {
			return
		}

		help(chatID)
	}
}

func setRGB(red, green, blue byte) {
	rgbled.SetRGB(255-red, 255-green, 255-blue)
}

func help(chatID int64) {
	text := "*Opciones*\nElige una de las opciones del menú"
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = tgbotapi.ModeMarkdown

	redCB := "red"
	greenCB := "green"
	yellowCB := "yellow"
	blueCB := "blue"
	purpleCB := "purple"
	blackCB := "black"

	markup := tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			[]tgbotapi.InlineKeyboardButton{
				tgbotapi.InlineKeyboardButton{
					Text:         "\u2764\ufe0f Rojo",
					CallbackData: &redCB,
				},
			},
			[]tgbotapi.InlineKeyboardButton{
				tgbotapi.InlineKeyboardButton{
					Text:         "\U0001f49b Amarillo",
					CallbackData: &yellowCB,
				},
			},
			[]tgbotapi.InlineKeyboardButton{
				tgbotapi.InlineKeyboardButton{
					Text:         "\U0001f49a Verde",
					CallbackData: &greenCB,
				},
			},
			[]tgbotapi.InlineKeyboardButton{
				tgbotapi.InlineKeyboardButton{
					Text:         "\U0001f499 Azul",
					CallbackData: &blueCB,
				},
			},
			[]tgbotapi.InlineKeyboardButton{
				tgbotapi.InlineKeyboardButton{
					Text:         "\U0001f49c Lila",
					CallbackData: &purpleCB,
				},
			},
			[]tgbotapi.InlineKeyboardButton{
				tgbotapi.InlineKeyboardButton{
					Text:         "\U0001f5a4 Negro",
					CallbackData: &blackCB,
				},
			},
		},
	}
	msg.ReplyMarkup = markup

	bot.Send(msg)
}

func readConfig() {
	if _, err := os.Stat("./example10/telegram/config.yml"); err != nil {
		fmt.Println("Error: config.yml file does not exist")
	}

	viper.SetConfigName("example10/telegram/config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	i, err := strconv.Atoi(fmt.Sprint(viper.Get("tg_chat")))
	if err != nil {
		fmt.Println("Telegram Chat not integer")
	}
	myChatID = int64(i)
	token = fmt.Sprint(viper.Get("tg_token"))

}
