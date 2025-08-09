package model

type URLCleanerRequest struct {
	URL       string `json:"url"`
	Operation string `json:"operation" example:"canonical,redirection,all"`
}

type URLCleanerResponse struct {
	ProcessedURL string `json:"processed_url"`
}
