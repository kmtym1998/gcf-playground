package f

import (
	"encoding/json"
	"log"
	"net/http"
)

func sendSlack(text, webhookURL string) error {
	body := map[string]string{"text": text}

	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	log.Println(string(b))

	resp, err := doRequest(
		http.MethodPost,
		webhookURL,
		&b,
		nil,
	)
	if err != nil {
		return err
	}

	log.Printf(`{"severity": "INFO", "textPayload": %s}`, string(resp))

	return nil
}
