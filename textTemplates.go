package ebrest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type TextTemplatesResource struct {
	client   *Client
	endpoint string
}

func (TextTemplate *TextTemplatesResource) GetTextTemplates(params map[string]string) (PaginatedResponse, error) {
	var payload PaginatedResponse

	request, err := http.NewRequest(
		http.MethodGet,
		TextTemplate.endpoint,
		nil,
	)

	if err != nil {
		return payload, err
	}

	TextTemplate.client.createQueryUrl(request, params)

	if err = TextTemplate.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (TextTemplate *TextTemplatesResource) GetTextTemplate(id int) (map[string]interface{}, error) {
	var payload map[string]interface{}

	request, err := http.NewRequest(
		http.MethodGet,
		TextTemplate.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return payload, err
	}

	if err = TextTemplate.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (TextTemplate *TextTemplatesResource) CreateTextTemplate(body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		TextTemplate.endpoint,
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = TextTemplate.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (TextTemplate *TextTemplatesResource) UpdateTextTemplate(id int, body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPut,
		TextTemplate.endpoint+"/"+strconv.Itoa(id),
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = TextTemplate.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (TextTemplate *TextTemplatesResource) DeleteTextTemplate(id int) error {

	request, err := http.NewRequest(
		http.MethodDelete,
		TextTemplate.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return err
	}

	err = TextTemplate.client.handle(request, nil)

	if err != nil {
		return err
	}

	return nil
}
