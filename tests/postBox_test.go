package tests

func (suite *easybillRestIntegrationTestSuite) TestPostBoxIntegration() {
	suiteAssert := suite.Assert()

	postBoxes, err := suite.client.PostBoxes().GetPostBoxes(map[string]string{})

	suiteAssert.Nil(err)
	suiteAssert.NotNil(postBoxes)
	suiteAssert.NotEmpty(postBoxes.Items)
	suiteAssert.NotNil(postBoxes.Items[0]["id"].(float64))

	postBoxID := int(postBoxes.Items[0]["id"].(float64))

	postBox, err := suite.client.PostBoxes().GetPostBox(postBoxID)
	suiteAssert.Nil(err)
	suiteAssert.NotNil(postBox)
}
