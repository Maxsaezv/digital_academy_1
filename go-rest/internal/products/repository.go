package products

type Repository interface {
	Create(products Products) error
	FindById(id int32) (*Products, error)
	FindAll() *[]Products
	Delete(id int32) error
	Update(id int32, products Products) (*Products, error)
}

type repository struct {
	productsMap map[int32]*Products
}

func NewRepository(productsMap map[int32]*Products) Repository {
	return &repository{
		productsMap: productsMap,
	}
}

func (r *repository) Create(products Products) error {
	return nil
}

func (r *repository) FindById(id int32) (*Products, error) {
	return nil, nil
}

func (r *repository) FindAll() *[]Products {
	return nil
}

func (r *repository) Delete(id int32) error {
	return nil
}

func (r *repository) Update(id int32, products Products) (*Products, error) {
	return nil, nil
}
