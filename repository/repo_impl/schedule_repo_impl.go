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

// SaveSchedule saves the show schedule to the database
func (s *ScheduleRepoImpl) SaveSchedule(ctx context.Context, schedule model.Schedule) (model.Schedule, error) {
	// Use GORM to save showtimes
	if err := s.db.WithContext(ctx).Create(&schedule).Error; err != nil {
		// Check for errors if any
		log.Println("Error saving schedule:", err)
		return schedule, banana.SaveScheduleFail
	}

	return schedule, nil
}

func (s *ScheduleRepoImpl) GetSchedulesByFilmID(ctx context.Context, filmID int) ([]model.Schedule, error) {
	var schedules []model.Schedule
	//Check schedule by film
	if err := s.db.WithContext(ctx).Preload("Cinema").Where("film_id = ?", filmID).Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

func (s *ScheduleRepoImpl) ReadSchedule(ctx context.Context, scheduleId int) (model.Schedule, error) {
	var schedules model.Schedule
	//Check schedule by film
	if err := s.db.WithContext(ctx).Preload("Cinema").Preload("Room").Preload("Film").First(&schedules, "id = ?", scheduleId).Error; err != nil {
		return schedules, err
	}
	return schedules, nil
}

func (s *ScheduleRepoImpl) GetAllSchedule(ctx context.Context, offset int, length int) ([]model.Schedule, error) {
	var schedules []model.Schedule

	if err := s.db.WithContext(ctx).Preload("Cinema").Preload("Room").Preload("Film").Limit(length).Offset(offset).Find(&schedules).Error; err != nil {
		log.Println("Error retrieving all schedule:", err)
		return nil, err
	}

	return schedules, nil
}

func (s *ScheduleRepoImpl) Delete(schedule model.Schedule) error {
	err := s.db.Unscoped().Delete(&schedule).Error
	if err != nil {
		log.Println("Error delete schedule")
		return err
	}
	return nil
}

func (s *ScheduleRepoImpl) TotalPage(schedule model.Schedule, length int) int {
	var total int64

	s.db.Model(&schedule).Count((&total))

	totalPage := int((total + int64(length) - 1) / int64(length))
	return totalPage
}
