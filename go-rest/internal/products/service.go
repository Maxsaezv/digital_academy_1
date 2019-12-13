package products

type Service interface {
	Create(products Products) error
	FindById(id int32) (*Products, error)
	FindAll() *[]Products
	Delete(id int32) error
	Update(id int32, products Products) (*Products, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(products Products) error {
	return nil
}

func (s *service) FindById(id int32) (*Products, error) {
	return nil, nil
}

func (s *service) FindAll() *[]Products {
	return nil
}

func (s *service) Delete(id int32) error {
	return nil
}

func (s *service) Update(id int32, products Products) (*Products, error) {
	return nil, nil
}
