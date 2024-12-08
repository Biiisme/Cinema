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
	if err := s.db.WithContext(ctx).Where("film_id = ?", filmID).Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}
