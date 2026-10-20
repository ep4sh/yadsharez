package yadsh

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/fatih/color"
)

// Download prints out download link for the particular file stored on the YandexDisk.
func Download(ctx context.Context, filePath, token string) error {
	downloadHref, err := getDownloadHref(ctx, filePath, token)
	if err != nil {
		return err
	}

	log.Printf("Copy the download link:")
	d := color.New(color.FgGreen, color.Bold)
	d.Printf("%s\n", downloadHref)

	return nil
}

// getDownloadHref return URL for direct file download.
func getDownloadHref(ctx context.Context, filePath, token string) (string, error) {
	url := getDownloadURL(filePath)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", AuthHeader(token))

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	downloadHref, err := getHref(res)
	if err != nil {
		return "", err
	}
	return downloadHref, nil
}

// getDownloadURL returns API URL endpoint for download request.
func getDownloadURL(file string) string {
	fp := filepath.Base(file)
	return fmt.Sprintf("https://cloud-api.yandex.net:443/v1/disk/resources/download/?path=app:/%s&overwrite=true", fp)
}
