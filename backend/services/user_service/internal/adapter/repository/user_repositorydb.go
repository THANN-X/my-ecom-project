package repository

import (
	"context"
	"errors"
	"user_service/internal/core/domain"
	"user_service/internal/core/port"

	"gorm.io/gorm"
)

// 1. ประกาศ Struct: เอาไว้เก็บ Connection ของ Database (เช่น *sql.DB หรือ *gorm.DB)
type userRepositoryDB struct {
	db *gorm.DB
}

// 2. NewRepo (Constructor): ฟังก์ชันสำหรับสร้าง Instance
func NewUserRepositoryDB(db *gorm.DB) port.UserRepository {
	return userRepositoryDB{db: db}
}

// 3. Implement Method ตาม Interface ที่สัญญาไว้ใน port
func (r userRepositoryDB) Save(ctx context.Context, user *domain.User) error {
	if err := r.db.Create(user); err != nil {
		return err.Error
	}

	return nil
}

func (r userRepositoryDB) Update(ctx context.Context, user *domain.User) error {
	if err := r.db.Save(user); err != nil {
		return err.Error
	}
	return nil
}

func (r userRepositoryDB) FindById(ctx context.Context, id int) (*domain.User, error) {
	user := &domain.User{}
	result := r.db.WithContext(ctx).First(user, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return user, nil
}

func (r userRepositoryDB) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}

	result := r.db.WithContext(ctx).First(user, email)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}

	return user, nil
}
