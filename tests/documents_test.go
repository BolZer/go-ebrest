package tests

import (
	"github.com/BolZer/go-ebrest"
	"github.com/stretchr/testify/assert"
)

func (suite *easybillRestTestSuite) TestDocumentResource() {
	client := suite.client
	t := suite.T()

	document, err := client.Documents().CreateDocument(map[string]string{
		"title": "Test Document",
	})

	assert.Nil(t, err)
	assert.NotNil(t, document)
	assert.Equal(t, "Test Document", document["title"].(string))
}

func (suite *easybillRestTestSuite) TestDocumentUpdate() {
	client := suite.client
	t := suite.T()

	document, err := client.Documents().UpdateDocument(2525, map[string]interface{}{
		"title": "Title Updated",
	})

	assert.Nil(t, err)
	assert.NotNil(t, document)
}

func (suite *easybillRestTestSuite) TestDocumentFetch() {
	client := suite.client
	t := suite.T()

	document, err := client.Documents().GetDocument(1241)

	assert.Nil(t, err)
	assert.NotNil(t, document)
}

func (suite *easybillRestTestSuite) TestDocumentsFetch() {
	client := suite.client
	t := suite.T()

	documents, err := client.Documents().GetDocuments(map[string]string{
		"page":  "1",
		"limit": "1",
	})

	assert.Nil(t, err)
	assert.NotNil(t, documents)
	assert.Equal(t, 1, documents.Limit)
	assert.Equal(t, 1, len(documents.Items))
}

func (suite *easybillRestTestSuite) TestDocumentFinalize() {
	client := suite.client
	t := suite.T()

	err := client.Documents().FinalizeDocument(24124)
	assert.Nil(t, err)
}

func (suite *easybillRestTestSuite) TestDocumentDelete() {
	client := suite.client
	t := suite.T()

	err := client.Documents().DeleteDocument(24124)
	assert.Nil(t, err)
}

func (suite *easybillRestTestSuite) TestDocumentCancel() {
	client := suite.client
	t := suite.T()

	err := client.Documents().CancelDocument(24124)
	assert.Nil(t, err)
}

func (suite *easybillRestTestSuite) TestDocumentSend() {
	client := suite.client
	t := suite.T()

	err := client.Documents().SendDocument(12412, ebrest.EMAIL, map[string]string{
		"to":      "test@test.de",
		"from":    "noehles@easybill.de",
		"subject": "Test Email",
		"message": "Test",
	})

	assert.Nil(t, err)
}

func (suite *easybillRestTestSuite) TestDocumentDownload() {
	client := suite.client
	t := suite.T()

	bytes, err := client.Documents().DownloadDocument(12412)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)
}

func (suite *easybillRestIntegrationTestSuite) TestDocumentIntegration() {
	suiteAssert := suite.Assert()

	// Create Document
	document, err := suite.client.Documents().CreateDocument(map[string]string{
		"title": "test",
	})

	suiteAssert.Nil(err)
	suiteAssert.NotNil(document)
	suiteAssert.Equal("test", document["title"].(string))

	documentID := int(document["id"].(float64))

	// Display Document
	document, err = suite.client.Documents().GetDocument(documentID)

	suiteAssert.Nil(err)
	suiteAssert.NotNil(document)
	suiteAssert.Equal("test", document["title"].(string))

	// Create Customer for Update
	customer, err := suite.client.Customers().CreateCustomer(map[string]string{
		"last_name":    "Test Customer",
		"company_name": "Test Company",
	})

	suiteAssert.Nil(err)
	suiteAssert.NotNil(customer)

	customerID := int(customer["id"].(float64))

	// Update Document with Customer
	document, err = suite.client.Documents().UpdateDocument(documentID, map[string]interface{}{
		"title":       "testUpdate",
		"customer_id": customerID,
	})

	suiteAssert.Nil(err)
	suiteAssert.NotNil(document)
	suiteAssert.Equal("testUpdate", document["title"].(string))
	suiteAssert.Equal(customerID, int(document["customer_id"].(float64)))

	// Finalize Document
	err = suite.client.Documents().FinalizeDocument(documentID)
	suiteAssert.Nil(err)

	// Download the Document
	bytes, err := suite.client.Documents().DownloadDocument(documentID)
	suiteAssert.Nil(err)
	suiteAssert.NotNil(bytes)

	// Delete Document
	err = suite.client.Documents().DeleteDocument(documentID)
	suiteAssert.Nil(err)

	// Display List of Documents
	documents, err := suite.client.Documents().GetDocuments(map[string]string{
		"limit": "1",
	})

	suiteAssert.Nil(err)
	suiteAssert.NotNil(documents)
	suiteAssert.Equal(1, documents.Limit)
}
