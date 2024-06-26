package repository

import (
	"backend-edu-experience/models"
	"context"
	"errors"

	"github.com/getsentry/sentry-go"
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
		sentry.CaptureMessage(err.Error())
		return nil, err
	}

	return &candidate, nil
}

func (c *CandidateRepository) UpdateCandidate(ctx context.Context, id int, candidate models.Candidate) error {
	data, err := c.GetById(id)

	if err != nil {
		return err
	}

	if data == nil {
		return errors.New("candidate not found")
	}

	err = c.DB.Where("id = ?", id).Updates(&candidate).Error
	if err != nil {
		sentry.CaptureMessage(err.Error())
		return err
	}

	return nil
}

func (c *CandidateRepository) DeleteCandidate(id int) error {
	data, err := c.GetById(id)
	if err != nil {
		sentry.CaptureMessage(err.Error())
		return err
	}

	if data == nil {
		return errors.New("candidate does not exist")
	}

	err = c.DB.Where("id = ?", data.ID).Delete(&data).Error
	if err != nil {
		sentry.CaptureMessage(err.Error())
		return err
	}

	return nil
}

func (c *CandidateRepository) GetOneCandidate(filters map[string]interface{}) (*models.Candidate, error) {
	var candidate models.Candidate

	query := c.DB.Model(&models.Candidate{})

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
		sentry.CaptureMessage(err.Error())
		return nil, err
	}

	return &candidate, nil
}
