package repository

import (
	"backend-edu-experience/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

type CandidateRepository struct {
	DB *gorm.DB
}

func NewCandidateRepository(db *gorm.DB) *CandidateRepository {
	return &CandidateRepository{
		DB: db,
	}
}

func (c *CandidateRepository) CreateCandidate(ctx context.Context, user *models.Candidate) error {
	return c.DB.WithContext(ctx).Create(user).Error
}

func (c *CandidateRepository) GetById(id int) (*models.Candidate, error) {
	var (
		candidate models.Candidate
	)

	err := c.DB.Where("id = ?", id).Find(&candidate).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &candidate, nil
}

// func (c *CandidateRepository) GetCandidateById(id int) (*models.Candidate, error) {
// 	var (
// 		candidate models.Candidate
// 	)

// 	err := c.DB.Where("id = ?", id).Find(&candidate).Error
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}

// 	return &candidate, nil
// }

func (c *CandidateRepository) UpdateCandidate(ctx context.Context, id int, candidate models.Candidate) error {
	data, err := c.GetById(id)
	if err != nil {
		return err
	}

	if data == nil {
		return errors.New("user not found")
	}

	err = c.DB.Updates(&candidate).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *CandidateRepository) DeleteCandidate(id int) error {

	data, err := c.GetById(id)
	if err != nil {
		return err
	}

	err = c.DB.Delete(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *CandidateRepository) GetOneCandidate(filters map[string]interface{}) (*models.Candidate, error) {
	var candidate models.Candidate

	query := c.DB.Model(&models.Candidate{})
	// query = query.Offset(pageSize).Limit(offset)

	// Filtering
	if len(filters) > 0 {
		for key, value := range filters {
			query = query.Where(key, value)
		}
	}

	err := query.First(&candidate).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &candidate, nil
}
