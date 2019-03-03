package tests

func (suite *easybillRestIntegrationTestSuite) TestStockIntegration() {
	suiteAssert := suite.Assert()

	position, err := suite.client.Positions().CreatePosition(map[string]string{
		"number":      "1234Test",
		"description": "Test Position",
		"sale_price":  "2000",
	})

	suiteAssert.Nil(err)
	suiteAssert.NotNil(position)

	positionID := int(position["id"].(float64))

	// Create Stock
	stock, err := suite.client.Stocks().CreateStock(map[string]interface{}{
		"position_id": positionID,
		"stock_count": 666,
	})

	suiteAssert.Nil(err)
	suiteAssert.NotNil(stock)
	suiteAssert.Equal(positionID, int(stock["position_id"].(float64)))

	stockID := int(stock["id"].(float64))

	// Display Stock
	stock, err = suite.client.Stocks().GetStock(stockID)

	suiteAssert.Nil(err)
	suiteAssert.NotNil(stock)
	suiteAssert.Equal(666, int(stock["stock_count"].(float64)))

	// Display List of Documents
	stocks, err := suite.client.Stocks().GetStocks(map[string]string{
		"limit": "1",
	})

	suiteAssert.Nil(err)
	suiteAssert.NotNil(stocks)
	suiteAssert.Equal(1, stocks.Limit)

	err = suite.client.Positions().DeletePosition(positionID)
	suiteAssert.Nil(err)
}
