package models

import (
	"time"

	"gorm.io/gorm"
)

type Education struct {
	gorm.Model
	CandidateID     uint       `json:"candidate_id,omitempty"`
	InstitutionName string     `json:"institution_name,omitempty"`
	Major           string     `json:"major,omitempty"`
	StartYear       *time.Time `json:"start_year,omitempty"`
	EndYear         *time.Time `json:"end_year,omitempty"`
	UntilNow        bool       `json:"until_now,omitempty"`
	GPA             *float64   `json:"gpa,omitempty"`
	Flag            *bool      `json:"flag,omitempty"`
	Role            string     `json:"role,omitempty"`
}

type EducationRequest struct {
	CandidateID     int     `json:"candidate_id,omitempty"`
	InstitutionName string  `json:"institution_name,omitempty"`
	Major           string  `json:"major,omitempty"`
	StartYear       *string `json:"start_year,omitempty"`
	EndYear         *string `json:"end_year,omitempty"`
	UntilNow        bool    `json:"until_now,omitempty"`
	GPA             float64 `json:"gpa,omitempty"`
	Flag            bool    `json:"flag,omitempty"`
	Role            string  `json:"role,omitempty"`
}

type EducationCreateRequest struct {
	Data []EducationRequest `json:"data"`
}

func (Education) TableName() string {
	return "education"
}
