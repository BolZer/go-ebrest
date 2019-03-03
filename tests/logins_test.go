package tests

import "github.com/stretchr/testify/assert"

func (suite *easybillRestTestSuite) TestLoginsFetch() {
	client := suite.client
	t := suite.T()

	logins, err := client.Logins().GetLogins(map[string]string{})

	assert.Nil(t, err)
	assert.NotNil(t, logins)
}

func (suite *easybillRestTestSuite) TestLoginFetch() {
	client := suite.client
	t := suite.T()

	login, err := client.Logins().GetLogin(124124)

	assert.Nil(t, err)
	assert.NotNil(t, login)
}

func (suite *easybillRestIntegrationTestSuite) TestLoginIntegration() {
	suiteAssert := suite.Assert()

	logins, err := suite.client.Logins().GetLogins(map[string]string{})

	suiteAssert.Nil(err)
	suiteAssert.NotNil(logins)
	suiteAssert.NotNil(logins.Items)
	suiteAssert.NotZero(logins.Items)
	suiteAssert.NotNil(logins.Items[0]["id"].(float64))

	login, err := suite.client.Logins().GetLogin(int(logins.Items[0]["id"].(float64)))
	suiteAssert.Nil(err)
	suiteAssert.NotNil(login)
}
