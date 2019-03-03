package ebrest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type DocumentPaymentsResource struct {
	client   *Client
	endpoint string
}

func (Payments *DocumentPaymentsResource) GetPayments(params map[string]string) (PaginatedResponse, error) {
	var payload PaginatedResponse

	request, err := http.NewRequest(
		http.MethodGet,
		Payments.endpoint,
		nil,
	)

	if err != nil {
		return payload, err
	}

	Payments.client.createQueryUrl(request, params)

	if err = Payments.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Payments *DocumentPaymentsResource) GetPayment(id int) (map[string]interface{}, error) {
	var payload map[string]interface{}

	request, err := http.NewRequest(
		http.MethodGet,
		Payments.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return payload, err
	}

	if err = Payments.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Payments *DocumentPaymentsResource) CreatePayment(body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		Payments.endpoint,
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = Payments.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Payments *DocumentPaymentsResource) DeletePayments(id int) error {

	request, err := http.NewRequest(
		http.MethodDelete,
		Payments.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return err
	}

	err = Payments.client.handle(request, nil)

	if err != nil {
		return err
	}

	return nil
}
