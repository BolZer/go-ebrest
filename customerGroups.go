package ebrest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type CustomerGroupsResource struct {
	client   *Client
	endpoint string
}

func (CustomerGroup *CustomerGroupsResource) GetCustomerGroups(params map[string]string) (PaginatedResponse, error) {
	var payload PaginatedResponse

	request, err := http.NewRequest(
		http.MethodGet,
		CustomerGroup.endpoint,
		nil,
	)

	if err != nil {
		return payload, err
	}

	CustomerGroup.client.createQueryUrl(request, params)

	if err = CustomerGroup.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (CustomerGroup *CustomerGroupsResource) GetCustomerGroup(id int) (map[string]interface{}, error) {
	var payload map[string]interface{}

	request, err := http.NewRequest(
		http.MethodGet,
		CustomerGroup.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return payload, err
	}

	if err = CustomerGroup.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (CustomerGroup *CustomerGroupsResource) CreateCustomerGroup(body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		CustomerGroup.endpoint,
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = CustomerGroup.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (CustomerGroup *CustomerGroupsResource) UpdateCustomerGroup(id int, body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPut,
		CustomerGroup.endpoint+"/"+strconv.Itoa(id),
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = CustomerGroup.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (CustomerGroup *CustomerGroupsResource) DeleteCustomerGroup(id int) error {

	request, err := http.NewRequest(
		http.MethodDelete,
		CustomerGroup.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return err
	}

	err = CustomerGroup.client.handle(request, nil)

	if err != nil {
		return err
	}

	return nil
}
