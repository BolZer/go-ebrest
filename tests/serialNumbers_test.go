package tests

import (
	"github.com/stretchr/testify/assert"
)

func (suite *easybillRestTestSuite) TestSerialNumbersCreate() {
	client := suite.client
	t := suite.T()

	serialNumber, err := client.SerialNumbers().CreateSerialNumber(map[string]interface{}{
		"serial_number": "DHEZ-DHSNR-2344D-FFW",
		"position_id":   12414,
	})

	assert.Nil(t, err)
	assert.NotNil(t, serialNumber)
	assert.Equal(t, "DHEZ-DHSNR-2344D-FFW", serialNumber["serial_number"].(string))
}

func (suite *easybillRestTestSuite) TestSerialNumberFetch() {
	client := suite.client
	t := suite.T()

	serialNumber, err := client.SerialNumbers().GetSerialNumber(2442)
	assert.Nil(t, err)
	assert.NotNil(t, serialNumber)
}

func (suite *easybillRestTestSuite) TestSerialNumbersFetch() {
	client := suite.client
	t := suite.T()

	serialNumbers, err := client.SerialNumbers().GetSerialNumbers(map[string]string{
		"page":  "1",
		"limit": "1",
	})

	assert.Nil(t, err)
	assert.NotNil(t, serialNumbers)
	assert.Equal(t, 1, serialNumbers.Limit)
	assert.Equal(t, 1, len(serialNumbers.Items))
}

func (suite *easybillRestTestSuite) TestSerialNumberDelete() {
	client := suite.client
	t := suite.T()

	err := client.SerialNumbers().DeleteSerialNumber(14124)
	assert.Nil(t, err)
}

func (suite *easybillRestIntegrationTestSuite) TestSerialNumbersIntegration() {
	suiteAssert := suite.Assert()

	// Create Position
	position, err := suite.client.Positions().CreatePosition(map[string]string{
		"number":      "1234Test",
		"description": "Test Position",
		"sale_price":  "2000",
	})

	suiteAssert.Nil(err)
	suiteAssert.NotNil(position)

	positionID := int(position["id"].(float64))

	// Create Serial Number
	serialNumber, err := suite.client.SerialNumbers().CreateSerialNumber(map[string]interface{}{
		"serial_number": "DHEZ-DHSNR-2344D-FFW",
		"position_id":   positionID,
	})

	suiteAssert.Nil(err)
	suiteAssert.NotNil(position)
	suiteAssert.Equal("DHEZ-DHSNR-2344D-FFW", serialNumber["serial_number"].(string))

	serialNumberID := int(serialNumber["id"].(float64))

	// Display Serial Number
	serialNumber, err = suite.client.SerialNumbers().GetSerialNumber(serialNumberID)

	suiteAssert.Nil(err)
	suiteAssert.NotNil(serialNumber)

	// Delete Serial Number
	err = suite.client.SerialNumbers().DeleteSerialNumber(serialNumberID)
	suiteAssert.Nil(err)

	// Display List of Serial Numbers
	serialNumbers, err := suite.client.SerialNumbers().GetSerialNumbers(map[string]string{
		"limit": "1",
	})

	suiteAssert.Nil(err)
	suiteAssert.NotNil(serialNumbers)
	suiteAssert.Equal(1, serialNumbers.Limit)

	err = suite.client.Positions().DeletePosition(positionID)
	suiteAssert.Nil(err)
}
