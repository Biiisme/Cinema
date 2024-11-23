package repo_impl

import (
	"cinema/banana"
	"cinema/model"
	"context"
	"log"

	"gorm.io/gorm"
)

type ScheduleRepoImpl struct {
	db *gorm.DB
}

func NewScheduleRepoImpl(db *gorm.DB) *ScheduleRepoImpl {
	return &ScheduleRepoImpl{db: db}
}

// SaveSchedule lưu lịch chiếu vào cơ sở dữ liệu
func (s *ScheduleRepoImpl) SaveSchedule(ctx context.Context, schedule model.Schedule) (model.Schedule, error) {
	// Sử dụng GORM để lưu lịch chiếu
	if err := s.db.WithContext(ctx).Create(&schedule).Error; err != nil {
		// Kiểm tra lỗi nếu có
		log.Println("Error saving schedule:", err)
		return schedule, banana.SaveScheduleFail
	}

	return schedule, nil
}

func (s *ScheduleRepoImpl) GetSchedulesByMovieID(ctx context.Context, movieID int) ([]model.Schedule, error) {
	var schedules []model.Schedule
	if err := s.db.WithContext(ctx).Where("film_id = ?", movieID).Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}
