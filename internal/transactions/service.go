package transactions

type Repository interface {
	GetRoleByName(name string) (*Role, error)
}

type Service struct {
	Repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		Repository: repository,
	}
}
