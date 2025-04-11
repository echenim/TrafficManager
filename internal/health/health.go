package health

import (
	"net/http"
	"time"
)

func CheckEndpoint(url string) bool {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		// logger.Log.WithFields(logrus.Fields{"url": url, "error": err}).Error("Failed to check endpoint health")
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == 200
}
