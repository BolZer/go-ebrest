package ebrest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type SepaPaymentsResource struct {
	client   *Client
	endpoint string
}

func (SepaPayment *SepaPaymentsResource) GetSepaPayments(params map[string]string) (PaginatedResponse, error) {
	var payload PaginatedResponse

	request, err := http.NewRequest(
		http.MethodGet,
		SepaPayment.endpoint,
		nil,
	)

	if err != nil {
		return payload, err
	}

	SepaPayment.client.createQueryUrl(request, params)

	if err = SepaPayment.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (SepaPayment *SepaPaymentsResource) GetSepaPayment(id int) (map[string]interface{}, error) {
	var payload map[string]interface{}

	request, err := http.NewRequest(
		http.MethodGet,
		SepaPayment.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return payload, err
	}

	if err = SepaPayment.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (SepaPayment *SepaPaymentsResource) CreateSepaPayment(body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		SepaPayment.endpoint,
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = SepaPayment.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (SepaPayment *SepaPaymentsResource) UpdateSepaPayment(id int, body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPut,
		SepaPayment.endpoint+"/"+strconv.Itoa(id),
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = SepaPayment.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (SepaPayment *SepaPaymentsResource) DeleteSepaPayment(id int) error {

	request, err := http.NewRequest(
		http.MethodDelete,
		SepaPayment.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return err
	}

	err = SepaPayment.client.handle(request, nil)

	if err != nil {
		return err
	}

	return nil
}
