package downloader

// import (
// 	"bytes"
// 	"fmt"
// 	"mime/multipart"
// 	"net/http"
// 	"os/exec"
// )

// func DownloadAndSendToTelegram(url, format, telegramBotToken, chatID string) error {
// 	// yt-dlp buyruqni stdout ga stream sifatida yo'naltirish
// 	cmd := exec.Command("yt-dlp", "-f", fmt.Sprintf("bestvideo[height<=%v]+bestaudio", format), "-o", "-", url)
// 	var out bytes.Buffer
// 	cmd.Stdout = &out
// 	cmd.Stderr = &out

// 	// Buyruqni ishga tushiramiz
// 	if err := cmd.Run(); err != nil {
// 		return fmt.Errorf("error downloading video: %v", err)
// 	}

// 	// Telegram API orqali videoni yuborish
// 	urlTelegram := fmt.Sprintf("https://api.telegram.org/bot%s/sendVideo", telegramBotToken)

// 	// POST so'rovi uchun form-data yaratish
// 	body := &bytes.Buffer{}
// 	writer := multipart.NewWriter(body)
// 	part, err := writer.CreateFormFile(	"video", "video.mp4")
// 	if err != nil {
// 		return fmt.Errorf("error creating form file: %v", err)
// 	}

// 	// Video stream ni form-data ga qo'shish
// 	part.Write(out.Bytes())
// 	writer.WriteField("chat_id", chatID)
// 	writer.Close()

// 	// HTTP POST so'rovi yuborish
// 	req, err := http.NewRequest("POST", urlTelegram, body)
// 	if err != nil {
// 		return fmt.Errorf("error creating request: %v", err)
// 	}
// 	req.Header.Set("Content-Type", writer.FormDataContentType())

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return fmt.Errorf("error sending video: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return fmt.Errorf("failed to send video, status code: %v", resp.StatusCode)
// 	}

// 	return nil
// }
