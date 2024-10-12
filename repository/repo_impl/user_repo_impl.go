package repo_impl

import (
	"cinema/banana"
	"cinema/model"
	"cinema/model/req"
	"cinema/repository"
	"context"
	"errors"
	"log"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repository.UserRepo {
	return &UserRepoImpl{db: db}
}

// Lưu User vào CSDL
func (u *UserRepoImpl) SaveUser(ctx context.Context, user model.User) (model.User, error) {
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	// Sử dụng GORM để lưu user vào CSDL
	err := u.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		// Kiểm tra xem lỗi có phải là unique constraint từ PostgreSQL
		var pqErr *pq.Error
		if ok := errors.As(err, &pqErr); ok && pqErr.Code.Name() == "unique_violation" {
			return user, banana.UserConfilict
		}

		log.Println("Error saving user:", err)
		return user, banana.SignUpFail
	}

	return user, nil
}

// Kiểm tra thông tin đăng nhập
func (u *UserRepoImpl) CheckLogin(ctx context.Context, loginReq req.ReqSignIn) (model.User, error) {
	var user model.User

	// Tìm user bằng email
	if err := u.db.WithContext(ctx).Where("email = ?", loginReq.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, banana.UserNotFound
		}
		log.Println("Error finding user:", err)
		return user, err
	}

	return user, nil
}
