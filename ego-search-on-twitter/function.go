package f

import (
	"context"
	"fmt"
	"os"
	"time"
)

func init() {
}

type PubSubMessage struct{}

func Run(ctx context.Context, _ PubSubMessage) error {
	startTime := time.Now().Add(-time.Hour * 12).Format(time.RFC3339)
	query := "zenn.dev/kmtym1998"
	bearerToken := os.Getenv("BEARER_TOKEN")

	tweets, err := searchRecentTweets(bearerToken, startTime, query)
	if err != nil {
		return err
	}

	webhookURL := os.Getenv("SLACK_WEBHOOK_URL")
	for _, t := range tweets.Data {
		if err := sendSlack(
			fmt.Sprintf("%s ~ %s の検索結果:\n%s", startTime, time.Now().Format(time.RFC3339), t.Text),
			webhookURL,
		); err != nil {
			return err
		}
	}

	return nil
}
