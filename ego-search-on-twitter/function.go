package f

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("Run", Run)
}

func Run(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now().Add(-time.Hour * 12).Format(time.RFC3339)
	query := "zenn.dev/kmtym1998"
	bearerToken := os.Getenv("BEARER_TOKEN")

	tweets, err := searchRecentTweets(bearerToken, startTime, query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	webhookURL := os.Getenv("SLACK_WEBHOOK_URL")
	for _, t := range tweets.Data {
		if err := sendSlack(
			fmt.Sprintf("%s ~ %s の検索結果:\n%s", startTime, time.Now().Format(time.RFC3339), t.Text),
			webhookURL,
		); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	w.WriteHeader(http.StatusOK)
}
