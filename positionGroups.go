package ebrest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type PositionGroupsResource struct {
	client   *Client
	endpoint string
}

func (PositionGroup *PositionGroupsResource) GetPositionGroups(params map[string]string) (PaginatedResponse, error) {
	var payload PaginatedResponse

	request, err := http.NewRequest(
		http.MethodGet,
		PositionGroup.endpoint,
		nil,
	)

	if err != nil {
		return payload, err
	}

	PositionGroup.client.createQueryUrl(request, params)

	if err = PositionGroup.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (PositionGroup *PositionGroupsResource) GetPositionGroup(id int) (map[string]interface{}, error) {
	var payload map[string]interface{}

	request, err := http.NewRequest(
		http.MethodGet,
		PositionGroup.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return payload, err
	}

	if err = PositionGroup.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (PositionGroup *PositionGroupsResource) CreatePositionGroup(body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		PositionGroup.endpoint,
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = PositionGroup.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (PositionGroup *PositionGroupsResource) UpdatePositionGroup(id int, body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPut,
		PositionGroup.endpoint+"/"+strconv.Itoa(id),
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = PositionGroup.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (PositionGroup *PositionGroupsResource) DeletePositionGroup(id int) error {

	request, err := http.NewRequest(
		http.MethodDelete,
		PositionGroup.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return err
	}

	err = PositionGroup.client.handle(request, nil)

	if err != nil {
		return err
	}

	return nil
}
