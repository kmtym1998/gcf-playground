package f

import (
	"encoding/json"
	"log"
	"net/http"
)

func sendSlack(text, webhookURL string) error {
	reqBody := map[string]string{"text": text}

	reqBodyB, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	log.Printf(`{"severity": "INFO", "message": %s}`, string(reqBodyB))

	respBodyB, err := doRequest(
		http.MethodPost,
		webhookURL,
		&reqBodyB,
		nil,
	)
	if err != nil {
		return err
	}

	log.Printf(`{"severity": "INFO", "message": %s}`, string(respBodyB))

	return nil
}
