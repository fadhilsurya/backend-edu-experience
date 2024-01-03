package repository

import (
	"backend-edu-experience/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

type EducationRepository struct {
	DB *gorm.DB
}

func NewEducationRepository(db *gorm.DB) *EducationRepository {
	return &EducationRepository{
		DB: db,
	}
}

func (ed *EducationRepository) GetEducation(pageSize, Offset int, filters map[string]interface{}) (*[]models.Education, *int64, error) {

	var edu []models.Education

	query := ed.DB.Model(&models.Education{})
	query = query.Offset(pageSize).Limit(Offset)

	// Filtering
	if len(filters) > 0 {
		for key, value := range filters {
			query = query.Where(key, value)
		}
	}

	err := query.Find(&edu).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, nil
		}
		return nil, nil, err
	}

	total, err := ed.CountEducation(filters)
	if err != nil {
		return nil, nil, err
	}

	return &edu, total, nil
}

func (ed *EducationRepository) CountEducation(filters map[string]interface{}) (*int64, error) {
	var (
		num int64
	)
	query := ed.DB.Model(&[]models.Education{})

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
		return nil, err
	}

	return &num, nil
}

func (ed *EducationRepository) GetOneEducation(filters map[string]interface{}) (*models.Education, error) {
	var education models.Education

	query := ed.DB.Model(&models.Education{})

	if len(filters) > 0 {
		for key, value := range filters {
			query = query.Where(key, value)
		}
	}

	err := query.First(&education).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &education, nil
}

func (ed *EducationRepository) UpdateCandidate(ctx context.Context, id int, education models.Education) error {
	filters := make(map[string]interface{})
	filters["id"] = id

	data, err := ed.GetOneEducation(filters)
	if err != nil {
		return err
	}

	if data == nil {
		return errors.New("education not found")
	}

	err = ed.DB.Where("id = ?", id).Updates(&education).Error
	if err != nil {
		return err
	}

	return nil
}

func (ed *EducationRepository) DeleteEducation(id int) error {

	filters := make(map[string]interface{})
	filters["id"] = id

	data, err := ed.GetOneEducation(filters)
	if err != nil {
		return err
	}

	if data == nil {
		return errors.New("education does not exist")
	}

	err = ed.DB.Where("id = ?", data.ID).Delete(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (ed *EducationRepository) BatchInsertEducation(ctx context.Context, edu []models.Education) error {
	err := ed.DB.WithContext(ctx).Create(edu).Error
	if err != nil {
		return err
	}

	return nil
}
