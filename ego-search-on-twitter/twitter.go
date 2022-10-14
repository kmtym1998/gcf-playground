package f

import (
	"encoding/json"
	"net/http"
)

type tweets struct {
	Data []data `json:"data"`
	Meta meta   `json:"meta"`
}
type data struct {
	EditHistoryTweetIds []string `json:"edit_history_tweet_ids"`
	ID                  string   `json:"id"`
	Text                string   `json:"text"`
}
type meta struct {
	NewestID    string `json:"newest_id"`
	OldestID    string `json:"oldest_id"`
	ResultCount int    `json:"result_count"`
}

const TWITTER_API_ENDPOINT = "https://api.twitter.com/2/tweets/search/recent"

func searchRecentTweets(bearerToken string, startTime, query string) (*tweets, error) {
	resp, err := doRequest(
		http.MethodGet,
		TWITTER_API_ENDPOINT+"?query="+query+"&start_time="+startTime,
		nil,
		[]reqHeaders{
			{"Authorization", "Bearer " + bearerToken},
		},
	)
	if err != nil {
		return nil, err
	}

	var result tweets
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
