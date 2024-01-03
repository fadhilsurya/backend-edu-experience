package models

import (
	"time"

	"gorm.io/gorm"
)

type Experience struct {
	gorm.Model
	CandidateID    uint      `json:"candidate_id,omitempty"`
	CompanyName    string    `json:"company_name,omitempty"`
	CompanyAddress string    `json:"company_address,omitempty"`
	Position       string    `json:"position,omitempty"`
	JobDesc        string    `json:"job_desc,omitempty"`
	StartYear      time.Time `json:"start_year,omitempty"`
	EndYear        time.Time `json:"end_year,omitempty"`
	UntilNow       bool      `json:"until_now,omitempty"`
	Flag           *bool     `json:"flag,omitempty"`
}

type ExperienceCreateRequest struct {
	CandidateID    int    `json:"candidate_id"`
	CompanyName    string `json:"company_name"`
	CompanyAddress string `json:"company_address"`
	Position       string `json:"position"`
	JobDesc        string `json:"job_desc"`
	StartYear      string `json:"start_year"`
	EndYear        string `json:"end_year"`
	UntilNow       bool   `json:"until_now"`
	Flag           bool   `json:"flag"`
}

func (Experience) TableName() string {
	return "experience"
}
