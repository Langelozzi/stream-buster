package api

type SearchPage struct {
	Page         int
	TotalPages   int
	TotalResults int
	Results      []interface{}
}
