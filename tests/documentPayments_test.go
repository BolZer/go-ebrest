package tests

import "github.com/stretchr/testify/assert"

func (suite *easybillRestTestSuite) TestDocumentPaymentsCreate() {
	client := suite.client
	t := suite.T()

	payment, err := client.DocumentPayments().CreatePayment(map[string]interface{}{
		"document_id": 2424,
		"amount":      1000,
		"type":        "VISA",
		"provider":    "Stripe",
		"reference":   "111111-VISA-222222-6666",
	})

	assert.Nil(t, err)
	assert.NotNil(t, payment)
	assert.Equal(t, "Stripe", payment["provider"].(string))
}

func (suite *easybillRestTestSuite) TestDocumentPaymentFetch() {
	client := suite.client
	t := suite.T()

	payment, err := client.DocumentPayments().GetPayment(241)

	assert.Nil(t, err)
	assert.NotNil(t, payment)
	assert.Equal(t, "Stripe", payment["provider"].(string))
}

func (suite *easybillRestTestSuite) TestDocumentPaymentsFetch() {
	client := suite.client
	t := suite.T()

	payments, err := client.DocumentPayments().GetPayments(map[string]string{
		"page":  "1",
		"limit": "1",
	})

	assert.Nil(t, err)
	assert.NotNil(t, payments)
	assert.Equal(t, 1, payments.Limit)
	assert.Equal(t, 1, len(payments.Items))
}

func (suite *easybillRestTestSuite) TestDocumentPaymentsDelete() {
	client := suite.client
	t := suite.T()

	err := client.DocumentPayments().DeletePayments(124)
	assert.Nil(t, err)
}
