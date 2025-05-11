package repo_impl

import (
	"cinema/banana"
	"cinema/model"
	"cinema/model/req"
	"cinema/repository"
	"context"
	"errors"
	"fmt"
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

// Same User to CSDL
func (u *UserRepoImpl) SaveUser(ctx context.Context, user model.User) (model.User, error) {
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	// Using GORM same user to CSDL
	err := u.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		// Check if the error is a unique constraint from PostgreSQL
		var pqErr *pq.Error
		if ok := errors.As(err, &pqErr); ok && pqErr.Code.Name() == "unique_violation" {
			return user, banana.UserConfilict
		}

		log.Println("Error saving user:", err)
		return user, banana.SignUpFail
	}

	return user, nil
}

// Check login information
func (u *UserRepoImpl) CheckLogin(ctx context.Context, loginReq req.ReqSignIn) (model.User, error) {
	var user model.User

	// check user by email
	if err := u.db.WithContext(ctx).Where("email = ?", loginReq.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, banana.UserNotFound
		}

		return user, err
	}

	return user, nil
}

func (u *UserRepoImpl) CountUser(ctx context.Context, UserID string) (int64, error) {
	var count int64
	// err := u.db.WithContext(ctx).Table("tokens").Select("user_id").Where("user_id = ?", UserID).Group("ip").Count(&count)
	err := u.db.WithContext(ctx).Table("tokens").Select("user_id").Where("user_id = ?", UserID).Count(&count).Error

	if err != nil {

		return 0, nil
	}
	return count, nil
}
func (u *UserRepoImpl) RemoveUser(UserID string) error {
	var oldestToken model.Token

	err := u.db.Where("user_id=?", UserID).Order("updated_at ASC").
		First(&oldestToken).Error
	if err != nil {

		return err
	}

	return u.db.Unscoped().Delete(&oldestToken).Error
}
func (u *UserRepoImpl) SaveToken(ctx context.Context, userID string, token string, ip string) error {

	ipCount, err := u.CountUser(ctx, userID)
	fmt.Println(ipCount)
	if err != nil {
		return err
	}
	if ipCount >= 3 {

		err := u.RemoveUser(userID)
		if err != nil {
			// log.Println(err)

			return err
		}

	}

	newToken := model.Token{
		UserID:    userID,
		IP:        ip,
		Token:     token,
		UpdatedAt: time.Now(),
	}

	return u.db.WithContext(ctx).Create(&newToken).Error
}

func (u *UserRepoImpl) GetUser(ctx context.Context, userID string) (model.User, error) {
	var user model.User

	if err := u.db.WithContext(ctx).Where("user_id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, banana.UserNotFound
		}
		return user, err
	}

	return user, nil
}

func (f *UserRepoImpl) UpdateUser(userReq req.ReqUpdateProfile, id string) (model.User, error) {
	// Using GORM to same film
	var user model.User
	if err := f.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, banana.UserConfilict
		}
		log.Println("Error retrieving user by ID:", err)
		return user, err
	}

	// Cập nhật các trường từ FilmReq vào Film
	updateData := map[string]interface{}{
		"full_name": userReq.FullName,
		"email":     userReq.Email,
	}
	if userReq.BirthDate != "" {
		parsedTime, _ := time.Parse("2006-01-02", userReq.BirthDate)
		updateData["birth_date"] = parsedTime
	}
	if userReq.PhoneNumber != "" {
		updateData["phone_number"] = userReq.PhoneNumber
	}
	if err := f.db.Model(&user).Updates(updateData).Error; err != nil {
		// Check error is unique constraint for PostgreSQL
		var pqErr *pq.Error
		if ok := errors.As(err, &pqErr); ok && pqErr.Code.Name() == "unique_violation" {
			return user, banana.UserConfilict
		}
		log.Println("Error saving user:", err)
		return user, banana.SaveUserFail
	}

	return user, nil
}
