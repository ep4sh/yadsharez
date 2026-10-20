package yadsh

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const (
	tokenEnv         = "YANDEX_OAUTH_TOKEN"
	tokenNotFoundErr = "YANDEX_OAUTH_TOKEN env missing"
)

// OAuthResp represents correct response on OAuth request.
type OAuthResp struct {
	OperationID string `json:"operation_id,omitempty"`
	Href        string `json:"href"`
	Method      string `json:"method"`
	Templated   bool   `json:"templated"`
}

// Upload request temporary URL for uploads
// and upload the file.
func Upload(ctx context.Context, filePath, token string) error {
	uploadHref, err := getUploadHref(ctx, filePath, token)
	if err != nil {
		return err
	}

	err = uploadYandexDisk(ctx, filePath, uploadHref)
	if err != nil {
		return err
	}

	return nil
}

// getUploadHref return URL for direct file download.
func getUploadHref(ctx context.Context, filePath, token string) (string, error) {
	url := getUploadURL(filePath)
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

	uploadHref, err := getHref(res)
	if err != nil {
		return "", err
	}
	return uploadHref, nil
}

// uploadYandexDisk performs direct file uploading to the Yandex Disk storage.
func uploadYandexDisk(ctx context.Context, filePath, href string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fieldName := "file"
	part, err := writer.CreateFormFile(fieldName, filepath.Base(filePath))
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, href, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("got %v status", res.Status)
	}

	return nil
}

// getUploadURL forms correct URL for upload request.
func getUploadURL(file string) string {
	fp := filepath.Base(file)
	return fmt.Sprintf("https://cloud-api.yandex.net:443/v1/disk/resources/upload/?path=app:/%s&overwrite=true", fp)
}

// GetOAuthToken looks up into env var and returns error if OAuth token not set.
func GetOAuthToken() (string, error) {
	val, ok := os.LookupEnv(tokenEnv)
	if !ok {
		return "", fmt.Errorf(tokenNotFoundErr)
	}

	return val, nil
}

// AuthHeader forms string for Authorization header.
func AuthHeader(token string) string {
	return fmt.Sprintf("OAuth %s", token)
}

// getHref decodes JSON response into OAuthResp struct and
// returns href field from it.
func getHref(res *http.Response) (string, error) {
	var oauth OAuthResp
	err := json.NewDecoder(res.Body).Decode(&oauth)
	if err != nil {
		return "", err
	}

	if oauth.Href == "" {
		return "", fmt.Errorf("href is empty, yandex api is broken: %v", res.Body)
	}

	return oauth.Href, nil
}

// Help prints out the app usage.
func Help() {
	log.Println("Usage: yadsharez [file]")
}
