package search

type Operator string

const (
	Equal              Operator = "="
	NotEqual           Operator = "!="
	GreaterThan        Operator = ">"
	GreaterThanOrEqual Operator = ">="
	LessThan           Operator = "<"
	LessThanOrEqual    Operator = "<="
	Like               Operator = "LIKE"
	NotLike            Operator = "NOT LIKE"
	In                 Operator = "IN"
	NotIn              Operator = "NOT IN"
	IsNull             Operator = "IS NULL"
	IsNotNull          Operator = "IS NOT NULL"
)
