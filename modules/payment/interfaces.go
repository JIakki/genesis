package payment

type IProductMapper interface {
	Find([]string, ...interface{}) (*Product, error)
}

type IProductRepository interface {
	FindById(int64) (*Product, error)
}
