package controller

import (
	"backend-edu-experience/helper"
	"backend-edu-experience/middleware"
	"backend-edu-experience/models"
	"backend-edu-experience/repository"
	"backend-edu-experience/template"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

type CandidateController struct {
	CandidateRepository *repository.CandidateRepository
}

func NewCandidateController(cr *repository.CandidateRepository) *CandidateController {
	return &CandidateController{
		CandidateRepository: cr,
	}
}

func (cc *CandidateController) CreateCandidate(c *gin.Context) {
	var newCandidateReq models.CandidateCreateRequest
	ctx := context.Background()
	resp := template.Response{}

	if err := c.ShouldBindJSON(&newCandidateReq); err != nil {
		resp.Data = nil
		resp.Error = err
		resp.Message = "bad request"

		c.JSON(400, resp)
		return
	}

	dob := newCandidateReq.DOB
	layout := "2006-01-02"

	dd, err := time.Parse(layout, dob)
	if err != nil {
		resp.Data = nil
		resp.Error = err
		resp.Message = "bad request"

		c.JSON(400, resp)
		return
	}

	hashPass, err := helper.HashPassword(newCandidateReq.Password)
	if err != nil {
		resp.Data = nil
		resp.Error = err
		resp.Message = "bad request"

		c.JSON(400, resp)
		return
	}

	if newCandidateReq.Gender != "male" && newCandidateReq.Gender != "female" {
		resp.Data = nil
		resp.Error = nil
		resp.Message = "bad request gender is not male or female"

		c.JSON(400, resp)
		return

	}

	candidateModel := models.Candidate{
		Fullname:       newCandidateReq.Fullname,
		DOB:            dd,
		Latitude:       newCandidateReq.Latitude,
		Longitude:      newCandidateReq.Longitude,
		Email:          newCandidateReq.Email,
		MobilePhone:    newCandidateReq.MobilePhone,
		Password:       hashPass,
		Gender:         newCandidateReq.Gender,
		CityID:         newCandidateReq.CityID,
		ProvinceID:     newCandidateReq.ProvinceID,
		LastEducation:  newCandidateReq.LastEducation,
		LastExperience: &newCandidateReq.LastExperience,
	}

	if err := cc.CandidateRepository.CreateCandidate(ctx, &candidateModel); err != nil {
		resp.Data = nil
		resp.Error = err
		resp.Message = "internal server error"

		c.JSON(500, resp)
		return
	}

	resp.Data = nil
	resp.Error = nil
	resp.Message = "Success"

	c.JSON(200, resp)
}

func (cc *CandidateController) Login(c *gin.Context) {
	var loginReq models.LoginReq
	resp := template.Response{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		resp.Data = nil
		resp.Error = err
		resp.Message = "bad request"

		c.JSON(400, resp)
		return
	}

	filter := make(map[string]interface{})

	if loginReq.Email == nil && loginReq.MobilePhone == nil {
		resp.Data = nil
		resp.Error = nil
		resp.Message = "bad request"

		c.JSON(400, resp)
		return
	}

	if loginReq.Email != nil {
		filter["email"] = loginReq.Email
	}

	if loginReq.MobilePhone != nil {
		filter["mobile_phone"] = loginReq.MobilePhone
	}

	data, err := cc.CandidateRepository.GetOneCandidate(filter)
	if err != nil {
		resp.Data = nil
		resp.Error = err
		resp.Message = "Internal Server Error"

		c.JSON(500, resp)
		return
	}

	err = helper.ComparePassword(data.Password, loginReq.Password)
	if err != nil {
		resp.Data = nil
		resp.Error = nil
		resp.Message = "invalid password"

		c.JSON(400, resp)
		return
	}

	token, err := middleware.GenerateJWT(data.ID)
	if err != nil {
		resp.Data = nil
		resp.Error = err
		resp.Message = "Internal Server Errror"

		c.JSON(500, resp)
		return
	}

	resp.Data = token
	resp.Error = nil
	resp.Message = "Success"

	c.JSON(200, resp)
}

func (cc *CandidateController) UpdateCandidate(c *gin.Context) {
	var newCandidateReq models.CandidateCreateRequest
	ctx := context.Background()
	resp := template.Response{}

	if err := c.ShouldBindJSON(&newCandidateReq); err != nil {
		resp.Data = nil
		resp.Error = err
		resp.Message = "bad request"

		c.JSON(400, resp)
		return
	}

	getToken := middleware.GetToken(c)
	if getToken == nil {
		resp.Data = nil
		resp.Error = nil
		resp.Message = "Bad Request - token is not exist"

		c.JSON(400, resp)
		return
	}

	id, err := middleware.GetUserIDFromToken(*getToken)
	if err != nil {
		resp.Data = nil
		resp.Error = err
		resp.Message = "Internal Server Error"

		c.JSON(500, resp)
		return
	}

	if newCandidateReq.Gender != "" {
		if newCandidateReq.Gender != "male" && newCandidateReq.Gender != "female" {
			resp.Data = nil
			resp.Error = nil
			resp.Message = "bad request gender is not male or female"

			c.JSON(400, resp)
			return
		}
	}

	candidateModel := models.Candidate{
		Fullname:       newCandidateReq.Fullname,
		Latitude:       newCandidateReq.Latitude,
		Longitude:      newCandidateReq.Longitude,
		Email:          newCandidateReq.Email,
		MobilePhone:    newCandidateReq.MobilePhone,
		Gender:         newCandidateReq.Gender,
		CityID:         newCandidateReq.CityID,
		ProvinceID:     newCandidateReq.ProvinceID,
		LastEducation:  newCandidateReq.LastEducation,
		LastExperience: &newCandidateReq.LastExperience,
	}

	if newCandidateReq.Password != "" {
		hashPass, err := helper.HashPassword(newCandidateReq.Password)
		if err != nil {
			resp.Data = nil
			resp.Error = err
			resp.Message = "bad request"

			c.JSON(400, resp)
			return
		}

		candidateModel.Password = hashPass
	}

	err = cc.CandidateRepository.UpdateCandidate(ctx, id, candidateModel)
	if err != nil {
		resp.Data = nil
		resp.Error = err
		resp.Message = "Internal Server Error"

		c.JSON(500, resp)
		return
	}

	resp.Data = nil
	resp.Error = nil
	resp.Message = "Success"

	c.JSON(200, resp)

}

func (cc *CandidateController) DeleteCandidate(c *gin.Context) {
	resp := template.Response{}

	getToken := middleware.GetToken(c)
	if getToken == nil {
		resp.Data = nil
		resp.Error = nil
		resp.Message = "Bad Request - token is not exist"

		c.JSON(400, resp)
		return
	}

	id, err := middleware.GetUserIDFromToken(*getToken)
	if err != nil {
		resp.Data = nil
		resp.Error = err
		resp.Message = "Internal Server Error"

		c.JSON(500, resp)
		return
	}

	err = cc.CandidateRepository.DeleteCandidate(id)
	if err != nil {
		resp.Data = nil
		resp.Error = err
		resp.Message = "Internal Server Error"

		c.JSON(500, resp)
		return
	}

	resp.Data = nil
	resp.Error = nil
	resp.Message = "Success"

	c.JSON(200, resp)

}
