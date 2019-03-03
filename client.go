// Package easybillRest provides a interface to interact with the easybill.de
// REST API. The Package exposes a series of Resources which exposes direct
// methods to call the several endpoints of the previously mentioned api.
package ebrest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const baseUri = "https://api.easybill.de/rest/v1"

// Endpoints is a enum containing the exposed endpoints of the
// easybill API. It is used internally for building the request urls.
type Endpoint string

const (
	endpointDocument              Endpoint = "/documents"
	endpointCustomer              Endpoint = "/customers"
	endpointCustomerGroup         Endpoint = "/customer-groups"
	endpointDiscountPosition      Endpoint = "/discounts/position"
	endpointDiscountPositionGroup Endpoint = "/discounts/position-group"
	endpointDocumentPayment       Endpoint = "/document-payments"
	endpointPosition              Endpoint = "/positions"
	endpointPositionGroup         Endpoint = "/position-groups"
	endpointProject               Endpoint = "/projects"
	endpointTask                  Endpoint = "/tasks"
	endpointLogin                 Endpoint = "/logins"
	endpointSerialNumber          Endpoint = "/serial-numbers"
	endpointPDFTemplate           Endpoint = "/pdf-templates"
	endpointPostBox               Endpoint = "/post-boxes"
	endpointSepaPayment           Endpoint = "/sepa-payments"
	endpointWebhook               Endpoint = "/webhooks"
	endpointTextTemplate          Endpoint = "/text-templates"
	endpointTimeTracking          Endpoint = "/time-trackings"
	endpointContact               Endpoint = "/customers/{id}/contacts"
	endpointStock                 Endpoint = "/stocks"
)

// String returns the Endpoint Type as a string representation.
func (Endpoint Endpoint) String() string {
	return string(Endpoint)
}

// Client is the entry type which needs to to be created for further
// interaction with the easybill API.
type Client struct {
	HttpClient HttpClient
	AuthKey    string
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// PaginatedResponse will be returned on methods like GetDocuments.
// I. E. for methods which return a paginated responses.
type PaginatedResponse struct {
	Page  int
	Pages int
	Limit int
	Total int
	Items []map[string]interface{}
}

// SimpleResponse will be used for response which are akin to the PaginatedResponse,
// however does not contain any information regarding the pagination but a "items" construct.
type SimpleResponse struct {
	Items []map[string]interface{}
}

// New takes the API Key as arguments and returns
// a prepared Client for further consumption.
func New(authKey string) Client {
	return Client{
		HttpClient: &http.Client{
			Timeout: 15 * time.Second,
		},
		AuthKey: authKey,
	}
}

// Documents exposes different methods to interact with the document endpoints.
func (Client Client) Documents() *DocumentsResource {
	return &DocumentsResource{
		client:   &Client,
		endpoint: baseUri + endpointDocument.String(),
	}
}

// Customers exposes different methods to interact with the customer endpoints.
func (Client Client) Customers() *CustomersResource {
	return &CustomersResource{
		client:   &Client,
		endpoint: baseUri + endpointCustomer.String(),
	}
}

// DocumentPayments exposes different methods to interact with the document payment endpoints.
func (Client Client) DocumentPayments() *DocumentPaymentsResource {
	return &DocumentPaymentsResource{
		client:   &Client,
		endpoint: baseUri + endpointDocumentPayment.String(),
	}
}

// Positions exposes different methods to interact with the position endpoints.
func (Client Client) Positions() *PositionsResource {
	return &PositionsResource{
		client:   &Client,
		endpoint: baseUri + endpointPosition.String(),
	}
}

// Logins exposes different methods to interact with the login endpoints.
func (Client Client) Logins() *LoginsResource {
	return &LoginsResource{
		client:   &Client,
		endpoint: baseUri + endpointLogin.String(),
	}
}

// SerialNumbers exposes different methods to interact with the serial number endpoints.
func (Client Client) SerialNumbers() *SerialNumbersResource {
	return &SerialNumbersResource{
		client:   &Client,
		endpoint: baseUri + endpointSerialNumber.String(),
	}
}

// Projects exposes different methods to interact with the project endpoints.
func (Client Client) Projects() *ProjectsResource {
	return &ProjectsResource{
		client:   &Client,
		endpoint: baseUri + endpointProject.String(),
	}
}

// Tasks exposes different methods to interact with the task endpoints.
func (Client Client) Tasks() *TasksResource {
	return &TasksResource{
		client:   &Client,
		endpoint: baseUri + endpointTask.String(),
	}
}

// PositionGroups exposes different methods to interact with the position group endpoints.
func (Client Client) PositionGroups() *PositionGroupsResource {
	return &PositionGroupsResource{
		client:   &Client,
		endpoint: baseUri + endpointPositionGroup.String(),
	}
}

// PdfTemplates exposes different methods to interact with the pdf template endpoints.
func (Client Client) PdfTemplates() *PdfTemplatesResource {
	return &PdfTemplatesResource{
		client:   &Client,
		endpoint: baseUri + endpointPDFTemplate.String(),
	}
}

// PostBoxes exposes different methods to interact with the post box endpoints.
func (Client Client) PostBoxes() *PostBoxesResource {
	return &PostBoxesResource{
		client:   &Client,
		endpoint: baseUri + endpointPostBox.String(),
	}
}

// CustomerGroups exposes different methods to interact with the customer group endpoints.
func (Client Client) CustomerGroups() *CustomerGroupsResource {
	return &CustomerGroupsResource{
		client:   &Client,
		endpoint: baseUri + endpointCustomerGroup.String(),
	}
}

// PositionDiscounts exposes different methods to interact with the position discount endpoints.
func (Client Client) PositionDiscounts() *DiscountPositionsResource {
	return &DiscountPositionsResource{
		client:   &Client,
		endpoint: baseUri + endpointDiscountPosition.String(),
	}
}

// PositionGroupDiscounts exposes different methods to interact with the position group discount endpoints.
func (Client Client) PositionGroupDiscounts() *DiscountPositionGroupsResource {
	return &DiscountPositionGroupsResource{
		client:   &Client,
		endpoint: baseUri + endpointDiscountPositionGroup.String(),
	}
}

// SepaPayments exposes different methods to interact with the sepa payment endpoints.
func (Client Client) SepaPayments() *SepaPaymentsResource {
	return &SepaPaymentsResource{
		client:   &Client,
		endpoint: baseUri + endpointSepaPayment.String(),
	}
}

// Webhooks exposes different methods to interact with the webhook endpoints.
func (Client Client) Webhooks() *WebhooksResource {
	return &WebhooksResource{
		client:   &Client,
		endpoint: baseUri + endpointWebhook.String(),
	}
}

// TextTemplates exposes different methods to interact with the text template endpoints.
func (Client Client) TextTemplates() *TextTemplatesResource {
	return &TextTemplatesResource{
		client:   &Client,
		endpoint: baseUri + endpointTextTemplate.String(),
	}
}

// TimeTrackings exposes different methods to interact with the time tracking endpoints.
func (Client Client) TimeTrackings() *TimeTrackingsResource {
	return &TimeTrackingsResource{
		client:   &Client,
		endpoint: baseUri + endpointTimeTracking.String(),
	}
}

// Contacts exposes different methods to interact with the contacts endpoints.
// The function takes the customer id as parameters due to the fact this resource
// is more of a subresource of the customer resource.
func (Client Client) Contacts(customerID int) *ContactsResource {
	return &ContactsResource{
		client:   &Client,
		endpoint: baseUri + strings.Replace(endpointContact.String(), "{id}", string(customerID), 1),
	}
}

// Stocks exposes different methods to interact with the stock endpoints.
func (Client Client) Stocks() *StocksResource {
	return &StocksResource{
		client:   &Client,
		endpoint: baseUri + endpointStock.String(),
	}
}

func (Client Client) call(request *http.Request, additionalHeaders map[string]string) (*http.Response, error) {
	var response *http.Response
	var errPayload map[string]interface{}

	request.Header.Set("Authorization", "Bearer "+Client.AuthKey)

	for key, value := range additionalHeaders {
		request.Header.Set(key, value)
	}

	response, err := Client.HttpClient.Do(request)

	if err != nil {
		return response, err
	}

	if response.StatusCode/100 == 2 {
		return response, nil
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(body, &errPayload); err != nil {
		return response, err
	}

	return response, newHttpError("{call}", errPayload["code"].(float64), errPayload["message"].(string), errPayload["arguments"])
}

func (Client *Client) createQueryUrl(request *http.Request, params map[string]string) {
	query := request.URL.Query()

	for key, value := range params {
		query.Add(key, value)
	}

	request.URL.RawQuery = query.Encode()
}

func (Client Client) handle(request *http.Request, payload interface{}) error {
	response, err := Client.call(request, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	})

	defer Client.closeResponse(response)

	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	if payload == nil {
		return nil
	}

	if payloadAsPaginatedResponse, ok := payload.(*PaginatedResponse); ok {
		if err := json.Unmarshal(body, &payloadAsPaginatedResponse); err != nil {
			return err
		}
		return nil
	}

	if payloadAsSimpleResponse, ok := payload.(*SimpleResponse); ok {
		if err := json.Unmarshal(body, &payloadAsSimpleResponse); err != nil {
			return err
		}
		return nil
	}

	if payloadAsMapInterface, ok := payload.(*map[string]interface{}); ok {
		if err := json.Unmarshal(body, &payloadAsMapInterface); err != nil {
			return err
		}

		return nil
	}

	if err := json.Unmarshal(body, &payload); err != nil {
		return err
	}

	return nil
}

func (Client Client) handleDownload(request *http.Request, contentType string) ([]byte, error) {
	response, err := Client.call(request, map[string]string{
		"Content-Type": contentType,
	})

	defer Client.closeResponse(response)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (Client Client) closeResponse(response *http.Response) {
	if response != nil {
		_ = response.Body.Close()
	}
}
