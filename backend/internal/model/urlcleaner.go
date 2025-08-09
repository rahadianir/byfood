package model

type URLCleanerRequest struct {
	URL       string `json:"url"`
	Operation string `json:"operation"`
}

type URLCleanerResponse struct {
	ProcessedURL string `json:"processed_url"`
}
