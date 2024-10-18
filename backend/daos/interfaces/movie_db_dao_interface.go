package interfaces

type MovieDatabaseDaoInterface interface {
	SearchMultiMedia(query string) ([]interface{}, error)
}
