package repository

import (
	"backend-edu-experience/models"
	"context"
	"errors"

	"github.com/getsentry/sentry-go"
	"gorm.io/gorm"
)

type ExperienceRepository struct {
	DB *gorm.DB
}

func NewExperienceRepository(db *gorm.DB) *ExperienceRepository {
	return &ExperienceRepository{
		DB: db,
	}
}

func (exp *ExperienceRepository) GetExperience(pageSize, Offset int, filters map[string]interface{}) (*[]models.Experience, *int64, error) {

	var ex []models.Experience

	query := exp.DB.Model(&models.Experience{})
	query = query.Offset(pageSize).Limit(Offset)

	// Filtering
	if len(filters) > 0 {
		for key, value := range filters {
			query = query.Where(key, value)
		}
	}

	err := query.Find(&ex).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, nil
		}
		sentry.CaptureMessage(err.Error())
		return nil, nil, err
	}

	total, err := exp.CountExperience(filters)
	if err != nil {
		sentry.CaptureMessage(err.Error())
		return nil, nil, err
	}

	return &ex, total, nil
}

func (exp *ExperienceRepository) CountExperience(filters map[string]interface{}) (*int64, error) {
	var (
		num int64
	)
	query := exp.DB.Model(&[]models.Experience{})

	// Filtering
	if len(filters) > 0 {
		for key, value := range filters {
			query = query.Where(key, value)
		}
	}

	err := query.Count(&num).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		sentry.CaptureMessage(err.Error())
		return nil, err
	}

	return &num, nil
}

func (exp *ExperienceRepository) GetOneExperience(filters map[string]interface{}) (*models.Experience, error) {
	var ex models.Experience

	query := exp.DB.Model(&models.Experience{})

	if len(filters) > 0 {
		for key, value := range filters {
			query = query.Where(key, value)
		}
	}

	err := query.First(&ex).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		sentry.CaptureMessage(err.Error())
		return nil, err
	}

	return &ex, nil
}

func (exp *ExperienceRepository) UpdateExperience(ctx context.Context, id int, ex models.Experience) error {
	filters := make(map[string]interface{})
	filters["id"] = id

	data, err := exp.GetOneExperience(filters)
	if err != nil {
		sentry.CaptureMessage(err.Error())
		return err
	}

	if data == nil {
		return errors.New("experience not found")
	}

	err = exp.DB.Where("id = ?", id).Updates(&ex).Error
	if err != nil {
		sentry.CaptureMessage(err.Error())
		return err
	}

	return nil
}

func (exp *ExperienceRepository) DeleteExperience(id int) error {

	filters := make(map[string]interface{})
	filters["id"] = id

	data, err := exp.GetOneExperience(filters)
	if err != nil {
		sentry.CaptureMessage(err.Error())
		return err
	}

	if data == nil {
		return errors.New("experience does not exist")
	}

	err = exp.DB.Where("id = ?", data.ID).Delete(&data).Error
	if err != nil {
		sentry.CaptureMessage(err.Error())
		return err
	}

	return nil
}

func (exp *ExperienceRepository) BatchInsertExperience(ctx context.Context, ex []models.Experience) error {
	err := exp.DB.WithContext(ctx).Create(ex).Error
	if err != nil {
		sentry.CaptureMessage(err.Error())
		return err
	}

	return nil
}
