package tests

import (
	"github.com/stretchr/testify/assert"
)

func (suite *easybillRestTestSuite) TestCustomerCreate() {
	client := suite.client
	t := suite.T()

	customer, err := client.Customers().CreateCustomer(map[string]string{
		"last_name":    "Test Customer",
		"company_name": "Test Company",
	})

	assert.Nil(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "Test Customer", customer["last_name"].(string))
	assert.Equal(t, "Test Company", customer["company_name"].(string))
}

func (suite *easybillRestTestSuite) TestCustomerFetch() {
	client := suite.client
	t := suite.T()

	customer, err := client.Customers().GetCustomer(24142)

	assert.Nil(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "Test Customer", customer["last_name"].(string))

}

func (suite *easybillRestTestSuite) TestCustomersFetch() {
	client := suite.client
	t := suite.T()

	customers, err := client.Customers().GetCustomers(map[string]string{
		"page":  "1",
		"limit": "1",
	})

	assert.Nil(t, err)
	assert.NotNil(t, customers)
	assert.Equal(t, 1, customers.Limit)
	assert.Equal(t, 1, len(customers.Items))
}

func (suite *easybillRestTestSuite) TestCustomerUpdate() {
	client := suite.client
	t := suite.T()

	customer, err := client.Customers().UpdateCustomer(4124124, map[string]string{
		"last_name":    "Test Customer",
		"company_name": "Test Company",
	})

	assert.Nil(t, err)
	assert.NotNil(t, customer)
}

func (suite *easybillRestTestSuite) TestCustomerDelete() {
	client := suite.client
	t := suite.T()

	err := client.Customers().DeleteCustomer(4124124)
	assert.Nil(t, err)

}

func (suite *easybillRestIntegrationTestSuite) TestCustomerIntegration() {
	suiteAssert := suite.Assert()

	// Create Customer
	customer, err := suite.client.Customers().CreateCustomer(map[string]string{
		"last_name":    "Test Customer",
		"company_name": "Test Company",
	})

	suiteAssert.Nil(err)
	suiteAssert.NotNil(customer)
	suiteAssert.Equal("Test Customer", customer["last_name"].(string))

	customerID := int(customer["id"].(float64))

	// Display Customer
	customer, err = suite.client.Customers().GetCustomer(customerID)

	suiteAssert.Nil(err)
	suiteAssert.NotNil(customer)
	suiteAssert.Equal("Test Customer", customer["last_name"].(string))

	// Update Customer
	customer, err = suite.client.Customers().UpdateCustomer(customerID, map[string]interface{}{
		"last_name": "Test Customer2",
	})

	suiteAssert.Nil(err)
	suiteAssert.NotNil(customer)
	suiteAssert.Equal("Test Customer2", customer["last_name"].(string))
	suiteAssert.Equal(customerID, int(customer["id"].(float64)))

	// Delete Customer
	err = suite.client.Customers().DeleteCustomer(customerID)
	suiteAssert.Nil(err)

	// Display List of Customers
	customers, err := suite.client.Customers().GetCustomers(map[string]string{
		"limit": "1",
	})

	suiteAssert.Nil(err)
	suiteAssert.NotNil(customers)
	suiteAssert.Equal(1, customers.Limit)
}
