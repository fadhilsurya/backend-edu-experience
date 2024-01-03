package models

import (
	"time"

	"gorm.io/gorm"
)

type Candidate struct {
	gorm.Model
	Fullname       string     `json:"fullname,omitempty"`
	DOB            time.Time  `json:"dob,omitempty"`
	Latitude       float64    `json:"latitude,omitempty"`
	Longitude      float64    `json:"longitude,omitempty"`
	Email          string     `json:"email,omitempty"`
	MobilePhone    string     `json:"mobile_phone,omitempty"`
	Password       string     `json:"password,omitempty"`
	Gender         string     `json:"gender,omitempty"`
	CityID         int        `json:"city_id,omitempty"`
	ProvinceID     int        `json:"province_id,omitempty"`
	LastEducation  string     `json:"last_education,omitempty"`
	LastExperience *int       `json:"last_experience,omitempty"`
	LoginDate      *time.Time `json:"login_date,omitempty"`
}

type CandidateCreateRequest struct {
	Fullname       string  `json:"fullname"`
	DOB            string  `json:"dob"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	Email          string  `json:"email"`
	MobilePhone    string  `json:"mobile_phone"`
	Password       string  `json:"password"`
	Gender         string  `json:"gender"`
	CityID         int     `json:"city_id"`
	ProvinceID     int     `json:"province_id"`
	LastEducation  string  `json:"last_education"`
	LastExperience int     `json:"last_experience"`
}

type LoginReq struct {
	Email       *string `json:"email,omitempty"`
	MobilePhone *string `json:"mobile_phone,omitempty"`
	Password    string  `json:"password"`
}
