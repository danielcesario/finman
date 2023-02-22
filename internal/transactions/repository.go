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

func (r *MySQLRepository) SaveUser(user User) (*User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *MySQLRepository) GetUserByEmail(email string) (*User, error) {
	var user User
	record := r.db.Model(&User{}).Preload("Roles").Where("email = ?", email).First(&user)
	if record.Error != nil {
		return nil, record.Error
	}

	return &user, nil
}
