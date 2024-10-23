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

		// Har bir yangilanishni alohida goroutine’da qayta ishlash
		go b.handleUpdate(update)
	}
}

func (b *Bot) handleUpdate(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	messageID := update.Message.MessageID

	// /start buyruqini tekshirish
	if update.Message.Text == "/start" {
		welcomeMsg := tgbotapi.NewMessage(chatID, "DOWNLOADER BOTIGA XUSH KELIBSIZ...")
		if _, err := b.API.Send(welcomeMsg); err != nil {
			log.Println("Xabar yuborishda xatolik:", err)
		}
		return
	}

	// Foydalanuvchiga "Yuklanmoqda..." xabarini yuborish
	msg := tgbotapi.NewMessage(chatID, "⏳")
	msg.ReplyToMessageID = messageID

	sentMessage, err := b.API.Send(msg)
	if err != nil {
		log.Println("Ma'lumotni yuborishda xatolik...", err)
		return
	}

	url := update.Message.Text

	// Media yuklab olish
	filePath, err := downloader.DownloadMedia(url, "720")
	if err != nil {
		// "Ma'lumot yuklanmoqda..." xabarini o'chirish
		deleteMsg := tgbotapi.DeleteMessageConfig{
			ChatID:    chatID,
			MessageID: sentMessage.MessageID,
		}
		if _, err := b.API.DeleteMessage(deleteMsg); err != nil {
			log.Println("Xabarni o'chirishda xatolik:", err)
		}
		// Xatolik haqida xabar yuborish
		errorMsg := tgbotapi.NewMessage(chatID, fmt.Sprintf("%v", err))
		errorMsg.ReplyToMessageID = messageID
		b.API.Send(errorMsg)
		return
	}

	// "Ma'lumot yuklanmoqda..." xabarini o'chirish
	deleteMsg := tgbotapi.DeleteMessageConfig{
		ChatID:    chatID,
		MessageID: sentMessage.MessageID,
	}
	if _, err := b.API.DeleteMessage(deleteMsg); err != nil {
		log.Println("Xabarni o'chirishda xatolik:", err)
	}

	// Yuklangan faylni foydalanuvchiga yuborish
	video := tgbotapi.NewVideoUpload(chatID, filePath)
	video.ReplyToMessageID = messageID
	if _, err := b.API.Send(video); err != nil {
		errorMsg := tgbotapi.NewMessage(chatID, "Video yuborishda xatolik: "+err.Error())
		errorMsg.ReplyToMessageID = messageID
		b.API.Send(errorMsg)
		log.Println("Video yuborishda xatolik:", err)
	}

	// Faylni o'chirib tashlash
	if err := os.Remove(filePath); err != nil {
		log.Println("Faylni o'chirishda xatolik:", err)
		log.Println("Ma'lumot muvaffaqiyatli yuborildi...")
	}
}
