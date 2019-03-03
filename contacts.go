package ebrest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type ContactsResource struct {
	client   *Client
	endpoint string
}

func (Contact *ContactsResource) GetContacts(params map[string]string) (PaginatedResponse, error) {
	var payload PaginatedResponse

	request, err := http.NewRequest(
		http.MethodGet,
		Contact.endpoint,
		nil,
	)

	if err != nil {
		return payload, err
	}

	Contact.client.createQueryUrl(request, params)

	if err = Contact.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Contact *ContactsResource) GetContact(id int) (map[string]interface{}, error) {
	var payload map[string]interface{}

	request, err := http.NewRequest(
		http.MethodGet,
		Contact.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return payload, err
	}

	if err = Contact.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Contact *ContactsResource) CreateContact(body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		Contact.endpoint,
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = Contact.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Contact *ContactsResource) UpdateContact(id int, body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPut,
		Contact.endpoint+"/"+strconv.Itoa(id),
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = Contact.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Contact *ContactsResource) DeleteContact(id int) error {

	request, err := http.NewRequest(
		http.MethodDelete,
		Contact.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return err
	}

	err = Contact.client.handle(request, nil)

	if err != nil {
		return err
	}

	return nil
}
