package ebrest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type CustomersResource struct {
	client   *Client
	endpoint string
}

func (Customers *CustomersResource) GetCustomers(params map[string]string) (PaginatedResponse, error) {
	var payload PaginatedResponse

	request, err := http.NewRequest(
		http.MethodGet,
		Customers.endpoint,
		nil,
	)

	if err != nil {
		return payload, err
	}

	Customers.client.createQueryUrl(request, params)

	if err = Customers.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Customers *CustomersResource) GetCustomer(id int) (map[string]interface{}, error) {
	var payload map[string]interface{}

	request, err := http.NewRequest(
		http.MethodGet,
		Customers.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return payload, err
	}

	if err = Customers.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Customers *CustomersResource) CreateCustomer(body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		Customers.endpoint,
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = Customers.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Customers *CustomersResource) UpdateCustomer(id int, body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPut,
		Customers.endpoint+"/"+strconv.Itoa(id),
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = Customers.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Customers *CustomersResource) DeleteCustomer(id int) error {

	request, err := http.NewRequest(
		http.MethodDelete,
		Customers.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return err
	}

	err = Customers.client.handle(request, nil)

	if err != nil {
		return err
	}

	return nil
}
