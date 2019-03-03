package ebrest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type ProjectsResource struct {
	client   *Client
	endpoint string
}

func (Projects *ProjectsResource) GetProjects(params map[string]string) (PaginatedResponse, error) {
	var payload PaginatedResponse

	request, err := http.NewRequest(
		http.MethodGet,
		Projects.endpoint,
		nil,
	)

	if err != nil {
		return payload, err
	}

	Projects.client.createQueryUrl(request, params)

	if err = Projects.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Projects *ProjectsResource) GetProject(id int) (map[string]interface{}, error) {
	var payload map[string]interface{}

	request, err := http.NewRequest(
		http.MethodGet,
		Projects.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return payload, err
	}

	if err = Projects.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Projects *ProjectsResource) CreateProject(body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		Projects.endpoint,
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = Projects.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Projects *ProjectsResource) UpdateProject(id int, body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPut,
		Projects.endpoint+"/"+strconv.Itoa(id),
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = Projects.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Projects *ProjectsResource) DeleteProject(id int) error {

	request, err := http.NewRequest(
		http.MethodDelete,
		Projects.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return err
	}

	err = Projects.client.handle(request, nil)

	if err != nil {
		return err
	}

	return nil
}
