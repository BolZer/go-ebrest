package ebrest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type StocksResource struct {
	client   *Client
	endpoint string
}

// GetStocks takes a map for arguments (filters) and returns a PaginatedResponse containing
// stock entries and infos regarding the pagination.
//
// Example:
//
// stocks, err := client.Stocks().GetStocks(map[string]string{
// 	 "page": "2",
// })
//
func (Stocks *StocksResource) GetStocks(params map[string]string) (PaginatedResponse, error) {
	var payload PaginatedResponse

	request, err := http.NewRequest(
		http.MethodGet,
		Stocks.endpoint,
		nil,
	)

	if err != nil {
		return payload, err
	}

	Stocks.client.createQueryUrl(request, params)

	if err = Stocks.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

// GetStock needs an id (int) as parameters and returns the requested stock entry if successful.
// If the stock entry is not found the api will return an 404 and GetStock err (HttpError)
//
// Example:
//
// stock, err := client.Stocks().GetStock(2)
//
func (Stocks *StocksResource) GetStock(id int) (map[string]interface{}, error) {
	var payload map[string]interface{}

	request, err := http.NewRequest(
		http.MethodGet,
		Stocks.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return payload, err
	}

	if err = Stocks.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

// CreateStock takes a map with the stock parameters as arguments and returns on success
// the created object as map[string]interface{}. If not enough arguments are provided, CreateStock
// will return an error (HttpError) containing the error and may contain a list of the missing parameters.
//
// Example:
//
// stocks, err := client.Stocks().CreateStock(map[string]string{
// 	 "stock_count": "666",
//   "position_id": "22",
// })
//
func (Stocks *StocksResource) CreateStock(body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		Stocks.endpoint,
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = Stocks.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}
