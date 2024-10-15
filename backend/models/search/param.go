package search

type Param struct {
	Field                 string      `json:"Field"`
	Operator              Operator    `json:"Operator"`
	Value                 interface{} `json:"Value"`
	AssociationForeignKey string      `json:"AssociationForeignKey"`
}
