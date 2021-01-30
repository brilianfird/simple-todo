package web

type Response struct {
	StatusCode int `json:"statusCode"`
	Error *string `json:"error"`
	Data interface{} `json:"data"`
}