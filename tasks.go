package ebrest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type TasksResource struct {
	client   *Client
	endpoint string
}

func (Tasks *TasksResource) GetTasks(params map[string]string) (PaginatedResponse, error) {
	var payload PaginatedResponse

	request, err := http.NewRequest(
		http.MethodGet,
		Tasks.endpoint,
		nil,
	)

	if err != nil {
		return payload, err
	}

	Tasks.client.createQueryUrl(request, params)

	if err = Tasks.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Tasks *TasksResource) GetTask(id int) (map[string]interface{}, error) {
	var payload map[string]interface{}

	request, err := http.NewRequest(
		http.MethodGet,
		Tasks.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return payload, err
	}

	if err = Tasks.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Tasks *TasksResource) CreateTask(body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		Tasks.endpoint,
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = Tasks.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Tasks *TasksResource) UpdateTask(id int, body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPut,
		Tasks.endpoint+"/"+strconv.Itoa(id),
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = Tasks.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Tasks *TasksResource) DeleteTask(id int) error {

	request, err := http.NewRequest(
		http.MethodDelete,
		Tasks.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return err
	}

	err = Tasks.client.handle(request, nil)

	if err != nil {
		return err
	}

	return nil
}
