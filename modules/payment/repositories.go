package payment

type ProductRepository struct {
	mapper IProductMapper
}

func (r *ProductRepository) FindById(id int64) (*Product, error) {
	return r.mapper.Find([]string{"ID"}, id)
}

func NewRepository(mapper IProductMapper) *ProductRepository {
	return &ProductRepository{
		mapper,
	}
}
