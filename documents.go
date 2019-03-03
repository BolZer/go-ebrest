package ebrest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type DocumentsResource struct {
	client   *Client
	endpoint string
}

type SendMethod string

const (
	EMAIL SendMethod = "email"
	FAX   SendMethod = "fax"
	POST  SendMethod = "post"
)

func (SendMethod SendMethod) String() string {
	return string(SendMethod)
}

func (Documents *DocumentsResource) GetDocuments(params map[string]string) (PaginatedResponse, error) {
	var payload PaginatedResponse

	request, err := http.NewRequest(
		http.MethodGet,
		Documents.endpoint,
		nil,
	)

	if err != nil {
		return payload, err
	}

	Documents.client.createQueryUrl(request, params)

	if err = Documents.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Documents *DocumentsResource) GetDocument(id int) (map[string]interface{}, error) {
	var payload map[string]interface{}

	request, err := http.NewRequest(
		http.MethodGet,
		Documents.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return payload, err
	}

	if err = Documents.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Documents *DocumentsResource) UpdateDocument(id int, body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPut,
		Documents.endpoint+"/"+strconv.Itoa(id),
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = Documents.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Documents *DocumentsResource) CreateDocument(body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		Documents.endpoint,
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = Documents.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Documents *DocumentsResource) DeleteDocument(id int) error {

	request, err := http.NewRequest(
		http.MethodDelete,
		Documents.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return err
	}

	err = Documents.client.handle(request, nil)

	if err != nil {
		return err
	}

	return nil
}

func (Documents *DocumentsResource) FinalizeDocument(id int) error {

	request, err := http.NewRequest(
		http.MethodPut,
		Documents.endpoint+"/"+strconv.Itoa(id)+"/done",
		nil,
	)

	if err != nil {
		return err
	}

	err = Documents.client.handle(request, nil)

	if err != nil {
		return err
	}

	return nil
}

func (Documents *DocumentsResource) CancelDocument(id int) error {

	request, err := http.NewRequest(
		http.MethodPost,
		Documents.endpoint+"/"+strconv.Itoa(id)+"/cancel",
		nil,
	)

	if err != nil {
		return err
	}

	err = Documents.client.handle(request, nil)

	if err != nil {
		return err
	}

	return nil
}

func (Documents *DocumentsResource) SendDocument(id int, sendMethod SendMethod, body interface{}) error {
	jsonBody, err := json.Marshal(body)

	if err != nil {
		return err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		Documents.endpoint+"/"+strconv.Itoa(id)+"/send/"+sendMethod.String(),
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return err
	}

	err = Documents.client.handle(request, nil)

	if err != nil {
		return err
	}

	return nil
}

func (Documents *DocumentsResource) DownloadDocument(id int) ([]byte, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		Documents.endpoint+"/"+strconv.Itoa(id)+"/pdf",
		nil,
	)

	if err != nil {
		return nil, err
	}

	payload, err := Documents.client.handleDownload(request, "application/pdf; charset=utf-8")

	if err != nil {
		return nil, err
	}

	return payload, nil
}
