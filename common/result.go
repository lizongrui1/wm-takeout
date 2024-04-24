package common

type PageResult struct {
	Total   int64       `json:"total"`
	Records interface{} `json:"records"`
}
