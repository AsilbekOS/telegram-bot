package bot

import (
	downloader "bot/service/download"
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	API *tgbotapi.BotAPI
}

func NewBot(token string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	log.Printf("Authorized on account %s", api.Self.UserName)
	return &Bot{API: api}, nil
}

func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.API.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Foydalanuvchiga yuklash jarayoni haqida ma'lumot beramiz
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ma'lumot yuklanmoqda...")
		msg.ReplyToMessageID = update.Message.MessageID // Bu xabarni reply qilib yuboradi

		sendMessage, err := b.API.Send(msg)
		if err != nil {
			log.Println("Ma'lumotni yuborishda xatolik...", err)
		}

		url := update.Message.Text
		chatID := update.Message.Chat.ID
		// chatIDStr := strconv.Itoa(int(chatID))

		// Yuklab olish va yuborish jarayoni
		filePath, err := downloader.DownloadMedia(url, "360")
		if err != nil {
			msg = tgbotapi.NewMessage(chatID, fmt.Sprintf("Download failed: %v", err))
			msg.ReplyToMessageID = update.Message.MessageID // Xatolikni ham reply qiladi
			b.API.Send(msg)
			continue
		}

		deleteMsg := tgbotapi.DeleteMessageConfig{
			ChatID:    chatID,
			MessageID: sendMessage.MessageID,
		}
		if _, err := b.API.DeleteMessage(deleteMsg); err != nil {
			log.Println("Xabarni o'chirishda xatolik:", err)
		}

		// Fayl yuklanganidan so'ng foydalanuvchiga video yuborish
		video := tgbotapi.NewVideoUpload(chatID, filePath)
		video.ReplyToMessageID = update.Message.MessageID
		if _, err := b.API.Send(video); err != nil {
			msg = tgbotapi.NewMessage(chatID, "Video yuborishda xatolik: "+err.Error())
			msg.ReplyToMessageID = update.Message.MessageID
			b.API.Send(msg)
			log.Println("Video yuborishda xatolik:", err)
		}

		// Faylni o'chirib tashlaymiz, chunki u endi kerak emas
		if err := os.Remove(filePath); err != nil {
			log.Println("Faylni o'chirishda xatolik:", err)
		}
	}
}
