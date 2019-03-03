package ebrest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type PositionsResource struct {
	client   *Client
	endpoint string
}

func (Positions *PositionsResource) GetPositions(params map[string]string) (PaginatedResponse, error) {
	var payload PaginatedResponse

	request, err := http.NewRequest(
		http.MethodGet,
		Positions.endpoint,
		nil,
	)

	if err != nil {
		return payload, err
	}

	Positions.client.createQueryUrl(request, params)

	if err = Positions.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Positions *PositionsResource) GetPosition(id int) (map[string]interface{}, error) {
	var payload map[string]interface{}

	request, err := http.NewRequest(
		http.MethodGet,
		Positions.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return payload, err
	}

	if err = Positions.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Positions *PositionsResource) CreatePosition(body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		Positions.endpoint,
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = Positions.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Positions *PositionsResource) UpdatePosition(id int, body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPut,
		Positions.endpoint+"/"+strconv.Itoa(id),
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = Positions.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Positions *PositionsResource) DeletePosition(id int) error {

	request, err := http.NewRequest(
		http.MethodDelete,
		Positions.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return err
	}

	err = Positions.client.handle(request, nil)

	if err != nil {
		return err
	}

	return nil
}
