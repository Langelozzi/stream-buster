package daos

type TMDBDao struct{}

func NewTMDBDao() *TMDBDao {
	return &TMDBDao{}
}

func (dao TMDBDao) SearchMultiMedia(query string) ([]interface{}, error) {
	return []interface{}{}, nil
}
