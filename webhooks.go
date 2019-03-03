package ebrest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type WebhooksResource struct {
	client   *Client
	endpoint string
}

func (Webhook *WebhooksResource) GetWebhooks(params map[string]string) (PaginatedResponse, error) {
	var payload PaginatedResponse

	request, err := http.NewRequest(
		http.MethodGet,
		Webhook.endpoint,
		nil,
	)

	if err != nil {
		return payload, err
	}

	Webhook.client.createQueryUrl(request, params)

	if err = Webhook.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Webhook *WebhooksResource) GetWebhook(id int) (map[string]interface{}, error) {
	var payload map[string]interface{}

	request, err := http.NewRequest(
		http.MethodGet,
		Webhook.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return payload, err
	}

	if err = Webhook.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Webhook *WebhooksResource) CreateWebhook(body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		Webhook.endpoint,
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = Webhook.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Webhook *WebhooksResource) UpdateWebhook(id int, body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPut,
		Webhook.endpoint+"/"+strconv.Itoa(id),
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = Webhook.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Webhook *WebhooksResource) DeleteWebhook(id int) error {

	request, err := http.NewRequest(
		http.MethodDelete,
		Webhook.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return err
	}

	err = Webhook.client.handle(request, nil)

	if err != nil {
		return err
	}

	return nil
}
