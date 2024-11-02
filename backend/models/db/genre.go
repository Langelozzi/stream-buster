package db

type Genre struct {
	ID   int `gorm:"primaryKey"`
	Name string
}
