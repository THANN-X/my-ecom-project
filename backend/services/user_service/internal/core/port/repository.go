package port

import "user_service/internal/core/domain"

type UserRepository interface {
	// Write methods (บันทึกลง DB)
	Save(user *domain.User) error
	Update(user *domain.User) error

	// Read methods (ดึงจาก DB)
}
