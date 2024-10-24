package downloader

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// IsYouTubeURL checks if the provided URL is a YouTube URL
func IsYouTubeURL(u string) bool {
	parsedURL, err := url.Parse(u)
	if err != nil {
		return false
	}
	return strings.Contains(parsedURL.Hostname(), "youtube.com") || strings.Contains(parsedURL.Hostname(), "youtu.be")
}

// IsInstagramURL checks if the provided URL is an Instagram URL
func IsInstagramURL(u string) bool {
	parsedURL, err := url.Parse(u)
	if err != nil {
		return false
	}
	return strings.Contains(parsedURL.Hostname(), "instagram.com")
}

// DownloadMedia downloads media from the given URL based on its type and returns the file path
func DownloadMedia(mediaURL, format string) (string, error) {
	// -----------------------------------------------------------------------------------------------
	parsedURL, err := url.Parse(mediaURL) // mediaURL ni tahlil qiling
	if err != nil {
		return "", fmt.Errorf("error parsing media URL: %v", err)
	}
	outputFile := fmt.Sprintf("%s.mp4", strings.ReplaceAll(parsedURL.Path, "/", "_"))
	// Output fayl nomini beramiz, bu so'ngra o'chiriladi

	count := 1
	// Fayl mavjud bo'lsa, yangi nom yaratish
	for {
		if _, err := os.Stat(outputFile); os.IsNotExist(err) {
			// Agar fayl mavjud bo'lmasa, tsiklni to'xtatish
			break
		}
		// Agar fayl mavjud bo'lsa, oxiriga raqam qo'shish
		outputFile = fmt.Sprintf("%s_%d.mp4", strings.TrimSuffix(outputFile, ".mp4"), count)
		count++
	}
	// -----------------------------------------------------------------------------------------------

	if IsYouTubeURL(mediaURL) {
		err := DownloadingFromYoutube(mediaURL, format, outputFile)
		if err != nil {
			return "", err
		}
	} else if IsInstagramURL(mediaURL) {
		err := DownloadingFromInstagram(mediaURL, outputFile)
		if err != nil {
			return "", err
		}
	} else {
		err := DownloadingFromOtherPlatform(outputFile, mediaURL)
		if err != nil {
			return "", err
		}
	}

	// Yuklab olingan faylning to'liq yo'lini qaytaramiz
	absPath, err := filepath.Abs(outputFile)
	if err != nil {
		return "", fmt.Errorf("error getting absolute file path: %v", err)
	}
	return absPath, nil
}

// Youtubedan media yuklashni o'zgarishlar kiritish uchun beyerda ishlang
func DownloadingFromYoutube(mediaURL, format, outputFile string) error {
	fmt.Println("Downloading from YouTube")
	cmd := exec.Command("yt-dlp", "-f", fmt.Sprintf("bestvideo[height<=%v]+bestaudio", format), "--merge-output-format", "mp4", "-o", outputFile, mediaURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("YouTubedan ma'lumotni yuklab bo'lmadi")
	}
	return nil
}

// Instagramdan media yuklash uchun
func DownloadingFromInstagram(mediaURL, outputFile string) error {
	fmt.Println("Downloading from Instagram")
	cmd := exec.Command("yt-dlp", "--cookies", "./service/download/cookies.txt", "--merge-output-format", "mp4", "-o", outputFile, mediaURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("instagramdan ma'lumotni yuklab bo'lmadi")
	}
	return nil
}

// Boshqa platformalardan media yuklashni o'zgarishlar kiritish uchun beyerda ishlang
func DownloadingFromOtherPlatform(outputFile, mediaURL string) error {
	fmt.Println("Downloading from other platform")
	// cmd := exec.Command("yt-dlp", "-f", "best", outputFile, mediaURL)
	cmd := exec.Command("yt-dlp", "-f", "b", "--merge-output-format", "mp4", "-o", outputFile, mediaURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ma'lumotni yuklab bo'lmadi")
	}
	return nil
}
