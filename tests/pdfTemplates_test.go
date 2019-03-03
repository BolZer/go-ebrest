package tests

import (
	"github.com/stretchr/testify/assert"
)

func (suite *easybillRestTestSuite) TestPdfTemplateFetch() {
	client := suite.client
	t := suite.T()

	templates, err := client.PdfTemplates().GetPdfTemplates()

	assert.Nil(t, err)
	assert.NotNil(t, templates)
}

func (suite *easybillRestIntegrationTestSuite) TestPDFTemplatesIntegration() {
	suiteAssert := suite.Assert()

	templates, err := suite.client.PdfTemplates().GetPdfTemplates()

	suiteAssert.Nil(err)
	suiteAssert.NotNil(templates)
	suiteAssert.NotNil(templates.Items)
	suiteAssert.NotZero(templates.Items)
}
