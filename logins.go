package ebrest

import (
	"net/http"
	"strconv"
)

type LoginsResource struct {
	client   *Client
	endpoint string
}

func (Logins *LoginsResource) GetLogins(params map[string]string) (PaginatedResponse, error) {
	var payload PaginatedResponse

	request, err := http.NewRequest(
		http.MethodGet,
		Logins.endpoint,
		nil,
	)

	if err != nil {
		return payload, err
	}

	Logins.client.createQueryUrl(request, params)

	if err = Logins.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (Logins *LoginsResource) GetLogin(id int) (map[string]interface{}, error) {
	var payload map[string]interface{}

	request, err := http.NewRequest(
		http.MethodGet,
		Logins.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return payload, err
	}

	if err = Logins.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}
