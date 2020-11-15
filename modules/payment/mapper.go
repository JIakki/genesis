package payment

import (
	"github.com/JIakki/genesis/db"
)

type ProductMapper struct {
	DB *db.DB
}

func NewMapper(DB *db.DB) *ProductMapper {
	return &ProductMapper{DB}
}

func (mapper *ProductMapper) Find(fields []string, args ...interface{}) (*Product, error) {
	return &Product{ID: 1, Price: 10, Name: "Awesome plan"}, nil
}
