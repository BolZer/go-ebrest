package ebrest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type SerialNumbersResource struct {
	client   *Client
	endpoint string
}

func (SerialNumbers *SerialNumbersResource) GetSerialNumbers(params map[string]string) (PaginatedResponse, error) {
	var payload PaginatedResponse

	request, err := http.NewRequest(
		http.MethodGet,
		SerialNumbers.endpoint,
		nil,
	)

	if err != nil {
		return payload, err
	}

	SerialNumbers.client.createQueryUrl(request, params)

	if err = SerialNumbers.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (SerialNumbers *SerialNumbersResource) GetSerialNumber(id int) (map[string]interface{}, error) {
	var payload map[string]interface{}

	request, err := http.NewRequest(
		http.MethodGet,
		SerialNumbers.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return payload, err
	}

	if err = SerialNumbers.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (SerialNumbers *SerialNumbersResource) CreateSerialNumber(body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		SerialNumbers.endpoint,
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = SerialNumbers.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (SerialNumbers *SerialNumbersResource) DeleteSerialNumber(id int) error {

	request, err := http.NewRequest(
		http.MethodDelete,
		SerialNumbers.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return err
	}

	err = SerialNumbers.client.handle(request, nil)

	if err != nil {
		return err
	}

	return nil
}
