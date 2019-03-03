package ebrest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type DiscountPositionsResource struct {
	client   *Client
	endpoint string
}

type DiscountPositionGroupsResource struct {
	client   *Client
	endpoint string
}

func (DiscountPosition *DiscountPositionsResource) GetPositionDiscounts(params map[string]string) (PaginatedResponse, error) {
	var payload PaginatedResponse

	request, err := http.NewRequest(
		http.MethodGet,
		DiscountPosition.endpoint,
		nil,
	)

	if err != nil {
		return payload, err
	}

	DiscountPosition.client.createQueryUrl(request, params)

	if err = DiscountPosition.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (DiscountPosition *DiscountPositionsResource) GetPositionDiscount(id int) (map[string]interface{}, error) {
	var payload map[string]interface{}

	request, err := http.NewRequest(
		http.MethodGet,
		DiscountPosition.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return payload, err
	}

	if err = DiscountPosition.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (DiscountPosition *DiscountPositionsResource) CreatePositionDiscount(body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		DiscountPosition.endpoint,
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = DiscountPosition.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (DiscountPosition *DiscountPositionsResource) UpdatePositionDiscount(id int, body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPut,
		DiscountPosition.endpoint+"/"+strconv.Itoa(id),
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = DiscountPosition.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (DiscountPosition *DiscountPositionsResource) DeletePositionDiscount(id int) error {

	request, err := http.NewRequest(
		http.MethodDelete,
		DiscountPosition.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return err
	}

	err = DiscountPosition.client.handle(request, nil)

	if err != nil {
		return err
	}

	return nil
}

func (DiscountPositionGroup *DiscountPositionGroupsResource) GetPositionGroupDiscounts(params map[string]string) (PaginatedResponse, error) {
	var payload PaginatedResponse

	request, err := http.NewRequest(
		http.MethodGet,
		DiscountPositionGroup.endpoint,
		nil,
	)

	if err != nil {
		return payload, err
	}

	DiscountPositionGroup.client.createQueryUrl(request, params)

	if err = DiscountPositionGroup.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (DiscountPositionGroup *DiscountPositionGroupsResource) GetPositionGroupDiscount(id int) (map[string]interface{}, error) {
	var payload map[string]interface{}

	request, err := http.NewRequest(
		http.MethodGet,
		DiscountPositionGroup.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return payload, err
	}

	if err = DiscountPositionGroup.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (DiscountPositionGroup *DiscountPositionGroupsResource) CreatePositionGroupDiscount(body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		DiscountPositionGroup.endpoint,
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = DiscountPositionGroup.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (DiscountPositionGroup *DiscountPositionGroupsResource) UpdatePositionGroupDiscount(id int, body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPut,
		DiscountPositionGroup.endpoint+"/"+strconv.Itoa(id),
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = DiscountPositionGroup.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (DiscountPositionGroup *DiscountPositionGroupsResource) DeletePositionGroupDiscount(id int) error {

	request, err := http.NewRequest(
		http.MethodDelete,
		DiscountPositionGroup.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return err
	}

	err = DiscountPositionGroup.client.handle(request, nil)

	if err != nil {
		return err
	}

	return nil
}
