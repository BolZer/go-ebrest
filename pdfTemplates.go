package ebrest

import "net/http"

type PdfTemplatesResource struct {
	client   *Client
	endpoint string
}

func (PdfTemplates *PdfTemplatesResource) GetPdfTemplates() (SimpleResponse, error) {
	var payload SimpleResponse

	request, err := http.NewRequest(
		http.MethodGet,
		PdfTemplates.endpoint,
		nil,
	)

	if err != nil {
		return payload, err
	}

	if err = PdfTemplates.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}
