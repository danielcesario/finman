package transactions

import "gorm.io/gorm"

type MySQLRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *MySQLRepository {
	return &MySQLRepository{
		db: db,
	}
}

func (r *MySQLRepository) GetRoleByName(name string) (*Role, error) {
	var role Role
	record := r.db.Model(&User{}).Where("role = ?", name).First(&role)
	if record.Error != nil {
		return nil, record.Error
	}

	return &role, nil
}
