package tests

import (
	"bytes"
	"encoding/json"
	"github.com/BolZer/go-ebrest"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type ClientMock struct {
	mock.Mock
}

func (client *ClientMock) Do(req *http.Request) (*http.Response, error) {
	args := client.Called(req)
	return args.Get(0).(*http.Response), nil
}

type easybillRestTestSuite struct {
	suite.Suite
	client ebrest.Client
}

type easybillRestIntegrationTestSuite struct {
	suite.Suite
	client ebrest.Client
}

func (suite *easybillRestTestSuite) SetupTest() {
	clientMock := new(ClientMock)

	var items []map[string]interface{}

	items = append(items, map[string]interface{}{
		"test": "test",
	})

	buf, _ := json.Marshal(map[string]interface{}{
		"number":        "1234Test",
		"description":   "Test Position",
		"sale_price":    "2000",
		"serial_number": "DHEZ-DHSNR-2344D-FFW",
		"last_name":     "Test Customer",
		"company_name":  "Test Company",
		"document_id":   2424,
		"amount":        1000,
		"type":          "VISA",
		"provider":      "Stripe",
		"reference":     "111111-VISA-222222-6666",
		"id":            666,
		"items":         items,
		"Page":          1,
		"Pages":         1,
		"Limit":         1,
		"Total":         1,
		"title":         "Test Document",
	})

	clientMock.On("Do", mock.Anything).Return(&http.Response{
		StatusCode: 200,
		Status:     "200 OK",
		Body:       ioutil.NopCloser(bytes.NewBuffer(buf)),
		Header:     make(http.Header),
	}, nil)

	suite.client = ebrest.Client{
		HttpClient: clientMock,
		AuthKey:    "testString",
	}
}

func (suite *easybillRestIntegrationTestSuite) SetupTest() {
	suite.client = ebrest.New("wmB4I5mu7cIZDibNde2VggVj2Cj9uAah1G1mTYfjd3EA3vgx5MHOoW59FbLKeiTR")
}

func (suite *easybillRestIntegrationTestSuite) AfterTest(suiteName, testName string) {
	time.Sleep(10 * time.Second)
}

func TestEasybillRestTestSuite(t *testing.T) {
	suite.Run(t, new(easybillRestTestSuite))
}

func TestEasybillRestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(easybillRestIntegrationTestSuite))
}
