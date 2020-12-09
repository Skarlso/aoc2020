package aoc

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// DownloadInput downloads and saves the requested input file.
func DownloadInput(year, day int) (string, error) {
	sessionID := os.Getenv("AOC_SESSION")
	if sessionID == "" {
		log.Fatal("No AOC_SESSION is set.")
	}
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	client := &http.Client{}

	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return "", err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: sessionID})

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", DownloadError{StatusCode: resp.StatusCode}
	}

	targetDir := fmt.Sprintf("days/day%d", day)
	targetFile := fmt.Sprintf("%s/input.txt", targetDir)

	ensurePath(targetDir)

	f, err := os.Create(targetFile)
	if err != nil {
		return "", err
	}

	defer f.Close()

	_, err = io.Copy(f, resp.Body)

	return targetFile, err
}
