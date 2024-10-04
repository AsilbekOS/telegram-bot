package atrash

// package bot

// import (
// 	downloader "bot/service/download"
// 	"log"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
// )

// type Bot struct {
// 	API *tgbotapi.BotAPI
// }

// func NewBot(token string) (*Bot, error) {
// 	api, err := tgbotapi.NewBotAPI(token)
// 	if err != nil {
// 		return nil, err
// 	}
// 	log.Printf("Authorized on account %s", api.Self.UserName)
// 	return &Bot{API: api}, nil
// }

// func (b *Bot) Start() {
// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 60

// 	updates, err := b.API.GetUpdatesChan(u)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for update := range updates {
// 		if update.Message == nil {
// 			continue
// 		}

// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ma'lumot yuklanmoqda")
// 		if _, err := b.API.Send(msg); err != nil {
// 			log.Println("Ma'lumotni yuboishda xatolik...", err)
// 		}

// 		url := update.Message.Text
// 		// chatid := update.Message.Chat.ID
// 		// chatidStr := strconv.Itoa(int(chatid))

// 		message := ""
// 		// err := downloader.DownloadAndSendToTelegram(url, "1080", b.API.Token, chatidStr)
// 		err = downloader.DownloadMedia(url, "2160")
// 		if err != nil {
// 			message = "Download failed:"
// 		} else {
// 			message = "Download successful!"
// 		}

// 		// videoConfig := tgbotapi.NewVideoUpload(chatid, video)
// 		msg = tgbotapi.NewMessage(update.Message.Chat.ID, message)
// 		if _, err := b.API.Send(msg); err != nil {
// 			log.Println("Video yuborishda xatolik:", err)
// 		}

// 	}
// }
