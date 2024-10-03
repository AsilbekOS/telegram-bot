package downloader

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
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

// DownloadMedia downloads media from the given URL based on its type
func DownloadMedia(url, format string) error {
	if IsYouTubeURL(url) {
		fmt.Println("Downloading from YouTube")
		cmd := exec.Command("yt-dlp", "-f", fmt.Sprintf("bestvideo[height<=%v]+bestaudio", format), "--merge-output-format", "mp4", url)
		// cmd := exec.Command("yt-dlp", "-f", "bestvideo+bestaudio/best", "--merge-output-format", "mp4", url)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("error downloading media: %v", err)
		}
		return nil
	} else {
		fmt.Println("Downloading from other platform")
		// cmd := exec.Command("yt-dlp", "--format", "best", url)
		cmd := exec.Command("yt-dlp", "-f", "b", "--merge-output-format", "mp4", url)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("error downloading media: %v", err)
		}
		return nil
	}
}
