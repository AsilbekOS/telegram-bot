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

// DownloadMedia downloads media from the given URL based on its type and returns the file path
func DownloadMedia(mediaURL, format string) (string, error) {
	outputFile := "output.mp4" // Output fayl nomini beramiz, bu so'ngra o'chiriladi
	if IsYouTubeURL(mediaURL) {
		err := DownloadingFromYoutube(mediaURL, format, outputFile)
		if err != nil {
			return "", err
		}
	} else {
		err := DownloadingFromOtherOlatform(outputFile, mediaURL)
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
		return fmt.Errorf("error downloading media: %v", err)
	}
	return nil
}

// Boshqa platformalardan media yuklashni o'zgarishlar kiritish uchun beyerda ishlang
func DownloadingFromOtherOlatform(outputFile, mediaURL string) error {
	fmt.Println("Downloading from other platform")
	// cmd := exec.Command("yt-dlp", "-f", "best", outputFile, mediaURL)
	cmd := exec.Command("yt-dlp", "-f", "b", "--merge-output-format", "mp4", "-o", outputFile, mediaURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error downloading media: %v", err)
	}
	return nil
}
