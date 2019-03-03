package ebrest

import (
	"net/http"
	"strconv"
)

type PostBoxesResource struct {
	client   *Client
	endpoint string
}

func (PostBox *PostBoxesResource) GetPostBoxes(params map[string]string) (PaginatedResponse, error) {
	var payload PaginatedResponse

	request, err := http.NewRequest(
		http.MethodGet,
		PostBox.endpoint,
		nil,
	)

	if err != nil {
		return payload, err
	}

	PostBox.client.createQueryUrl(request, params)

	if err = PostBox.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (PostBox *PostBoxesResource) GetPostBox(id int) (map[string]interface{}, error) {
	var payload map[string]interface{}

	request, err := http.NewRequest(
		http.MethodGet,
		PostBox.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return payload, err
	}

	if err = PostBox.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (PostBox *PostBoxesResource) DeletePostBox(id int) error {

	request, err := http.NewRequest(
		http.MethodDelete,
		PostBox.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return err
	}

	err = PostBox.client.handle(request, nil)

	if err != nil {
		return err
	}

	return nil
}
