package tests

import "github.com/stretchr/testify/assert"

func (suite *easybillRestTestSuite) TestPositionsCreate() {
	client := suite.client
	t := suite.T()

	position, err := client.Positions().CreatePosition(map[string]string{
		"number":      "1234Test",
		"description": "Test Position",
		"sale_price":  "2000",
	})

	assert.Nil(t, err)
	assert.NotNil(t, position)
	assert.Equal(t, "1234Test", position["number"].(string))
	assert.Equal(t, "Test Position", position["description"].(string))
}

func (suite *easybillRestTestSuite) TestPositionsUpdate() {
	client := suite.client
	t := suite.T()

	position, err := client.Positions().UpdatePosition(12412, map[string]string{
		"number": "1234TestUpdated",
	})

	assert.Nil(t, err)
	assert.NotNil(t, position)
}

func (suite *easybillRestTestSuite) TestPositionFetch() {
	client := suite.client
	t := suite.T()

	position, err := client.Positions().GetPosition(124124)

	assert.Nil(t, err)
	assert.NotNil(t, position)
}

func (suite *easybillRestTestSuite) TestPositionsFetch() {
	client := suite.client
	t := suite.T()

	positions, err := client.Positions().GetPositions(map[string]string{
		"page":  "1",
		"limit": "1",
	})

	assert.Nil(t, err)
	assert.NotNil(t, positions)
	assert.Equal(t, 1, positions.Limit)
	assert.Equal(t, 1, len(positions.Items))
}

func (suite *easybillRestTestSuite) TestPositionsDelete() {
	client := suite.client
	t := suite.T()

	err := client.Positions().DeletePosition(1241)
	assert.Nil(t, err)
}

func (suite *easybillRestIntegrationTestSuite) TestPositionIntegration() {
	suiteAssert := suite.Assert()

	// Create Position
	position, err := suite.client.Positions().CreatePosition(map[string]string{
		"number":      "1234Test",
		"description": "Test Position",
		"sale_price":  "2000",
	})

	suiteAssert.Nil(err)
	suiteAssert.NotNil(position)
	suiteAssert.Equal("1234Test", position["number"].(string))

	positionID := int(position["id"].(float64))

	// Display Display
	position, err = suite.client.Positions().GetPosition(positionID)

	suiteAssert.Nil(err)
	suiteAssert.NotNil(position)
	suiteAssert.Equal("1234Test", position["number"].(string))

	// Update Position
	position, err = suite.client.Positions().UpdatePosition(positionID, map[string]interface{}{
		"description": "Updated Test Position",
	})

	suiteAssert.Nil(err)
	suiteAssert.NotNil(position)
	suiteAssert.Equal("Updated Test Position", position["description"].(string))

	// Delete Position
	err = suite.client.Positions().DeletePosition(positionID)
	suiteAssert.Nil(err)

	// Display List of Position
	positions, err := suite.client.Positions().GetPositions(map[string]string{
		"limit": "1",
	})

	suiteAssert.Nil(err)
	suiteAssert.NotNil(positions)
	suiteAssert.Equal(1, positions.Limit)
}
