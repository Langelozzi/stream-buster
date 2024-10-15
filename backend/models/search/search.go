package search

type Search struct {
	Params          []Param         `json:"Params"`
	LogicalOperator LogicalOperator `json:"LogicalOperator"`
}
