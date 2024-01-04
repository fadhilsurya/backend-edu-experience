package controller

import (
	"backend-edu-experience/middleware"
	"backend-edu-experience/models"
	"backend-edu-experience/repository"
	"backend-edu-experience/template"
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type EducationController struct {
	EducationRepository *repository.EducationRepository
}

func NewEducationController(er *repository.EducationRepository) *EducationController {
	return &EducationController{
		EducationRepository: er,
	}
}

func (ec *EducationController) CreateEducation(c *gin.Context) {

	ctx := context.Background()

	var (
		newEducationReq models.EducationCreateRequest
		resp            template.Response
		eduModels       []models.Education
		endYear         *string
		e               time.Time
		s               time.Time
	)

	if err := c.ShouldBindJSON(&newEducationReq); err != nil {
		resp = template.Response{
			Data:    nil,
			Error:   err,
			Message: "Bad Request",
		}

		c.JSON(400, resp)
		return
	}

	id, err := middleware.GetUserIDFromToken(c.GetHeader("Authorization"))
	if err != nil {
		resp = template.Response{
			Data:    nil,
			Error:   nil,
			Message: "Internal Server Error",
		}

		c.JSON(500, resp)
		return
	}

	for _, v := range newEducationReq.Data {
		layout := "2006-01-02"
		startYear := v.StartYear
		sy, err := time.Parse(layout, *startYear)
		if err != nil {
			resp = template.Response{
				Data:    nil,
				Error:   err,
				Message: "Bad Request",
			}

			c.JSON(400, resp)
			return
		}
		s = sy

		if v.EndYear != nil {
			endYear = v.EndYear
			ey, err := time.Parse(layout, *endYear)
			if err != nil {
				resp = template.Response{
					Data:    nil,
					Error:   err,
					Message: "Bad Request",
				}

				c.JSON(400, resp)
				return
			}

			e = ey

		}

		eduModel := models.Education{
			CandidateID:     uint(id),
			InstitutionName: v.InstitutionName,
			Major:           v.Major,
			GPA:             &v.GPA,
			Role:            v.Role,
			Flag:            &v.Flag,
			StartYear:       &s,
			EndYear:         &e,
		}
		eduModels = append(eduModels, eduModel)
	}

	err = ec.EducationRepository.BatchInsertEducation(ctx, eduModels)
	if err != nil {
		resp = template.Response{
			Data:    nil,
			Error:   err,
			Message: "Internal Server Error",
		}

		c.JSON(500, resp)
		return
	}

	resp = template.Response{
		Data:    nil,
		Error:   nil,
		Message: "Success",
	}
	c.JSON(200, resp)

}

func (ec *EducationController) UpdateEducation(c *gin.Context) {
	var (
		newEducationReq models.EducationRequest
		resp            template.Response
		eduModels       models.Education
		endyear         time.Time
		startyear       time.Time
	)

	educationId := c.Param("id")
	ctx := context.Background()

	if err := c.ShouldBindJSON(&newEducationReq); err != nil {
		resp = template.Response{
			Data:    nil,
			Error:   err,
			Message: "Bad Request",
		}

		c.JSON(400, resp)
		return
	}

	id, err := middleware.GetUserIDFromToken(c.GetHeader("Authorization"))
	if err != nil {
		resp = template.Response{
			Data:    nil,
			Error:   err,
			Message: "Internal Server Error",
		}

		c.JSON(500, resp)
		return
	}

	filterCheck := make(map[string]interface{})
	filterCheck["candidate_id"] = id
	filterCheck["id"] = educationId

	data, err := ec.EducationRepository.GetOneEducation(filterCheck)
	if err != nil {
		resp = template.Response{
			Data:    nil,
			Error:   err,
			Message: "Internal Server Error",
		}

		c.JSON(500, resp)
		return
	}

	if data == nil {
		resp = template.Response{
			Data:    nil,
			Error:   nil,
			Message: "Education Does Not Exist",
		}

		c.JSON(500, resp)
		return
	}
	layout := "2006-01-02"

	if newEducationReq.StartYear != nil {
		startYear := newEducationReq.StartYear
		sy, err := time.Parse(layout, *startYear)
		if err != nil {
			resp = template.Response{
				Data:    nil,
				Error:   err,
				Message: "Bad Request",
			}

			c.JSON(400, resp)
			return
		}
		startyear = sy
	}

	if newEducationReq.EndYear != nil {
		endYear := newEducationReq.EndYear
		ey, err := time.Parse(layout, *endYear)
		if err != nil {
			resp = template.Response{
				Data:    nil,
				Error:   err,
				Message: "Bad Request",
			}

			c.JSON(400, resp)
			return
		}

		endyear = ey
	}

	if endyear.After(startyear) {
		resp = template.Response{
			Data:    nil,
			Error:   err,
			Message: "Bad Request",
		}

		c.JSON(400, resp)
		return
	}

	if newEducationReq.UntilNow && newEducationReq.StartYear != nil {

		resp = template.Response{
			Data:    nil,
			Error:   err,
			Message: "Bad Request",
		}
		c.JSON(400, resp)
		return
	}

	eduModels = models.Education{
		InstitutionName: newEducationReq.InstitutionName,
		StartYear:       &startyear,
		EndYear:         &endyear,
		Major:           newEducationReq.Major,
		UntilNow:        newEducationReq.UntilNow,
		GPA:             &newEducationReq.GPA,
		Flag:            &newEducationReq.Flag,
		Role:            newEducationReq.Role,
	}

	err = ec.EducationRepository.UpdateCandidate(ctx, id, eduModels)
	if err != nil {
		resp = template.Response{
			Data:    nil,
			Error:   err,
			Message: "Internal Server Error",
		}
		c.JSON(400, resp)
		return
	}

	resp = template.Response{
		Data:    nil,
		Error:   nil,
		Message: "Success",
	}
	c.JSON(200, resp)
}

func (ec *EducationController) DeleteEducation(c *gin.Context) {

	var (
		resp template.Response
	)

	educationId := c.Param("id")

	id, err := middleware.GetUserIDFromToken(c.GetHeader("Authorization"))
	if err != nil {
		resp = template.Response{
			Data:    nil,
			Error:   err,
			Message: "Internal Server Error",
		}

		c.JSON(500, resp)
		return
	}

	filters := make(map[string]interface{})
	filters["id"] = educationId
	filters["candidate_id"] = id

	data, err := ec.EducationRepository.GetOneEducation(filters)
	if err != nil {
		resp = template.Response{
			Data:    nil,
			Error:   err,
			Message: "Internal Server Error",
		}

		c.JSON(500, resp)
		return
	}

	if data == nil {
		resp = template.Response{
			Data:    nil,
			Error:   err,
			Message: "Bad Request - Education Does Not Exist",
		}

		c.JSON(400, resp)
		return
	}

	err = ec.EducationRepository.DeleteEducation(int(data.ID))
	if err != nil {
		resp = template.Response{
			Data:    nil,
			Error:   err,
			Message: "Internal Server Error",
		}

		c.JSON(500, resp)
		return
	}

	resp = template.Response{
		Data:    nil,
		Error:   nil,
		Message: "Success",
	}

	c.JSON(200, resp)
}

func (ec *EducationController) GetEducation(c *gin.Context) {
	var (
		responsePagination template.ResponsePagination
		response           template.Response
	)

	id, err := middleware.GetUserIDFromToken(c.GetHeader("Authorization"))
	if err != nil {
		response = template.Response{
			Data:    nil,
			Error:   err,
			Message: "Internal Server Error",
		}

		c.JSON(500, response)
		return
	}

	filters := make(map[string]interface{})
	filters["candidate_id"] = id

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	institutionName := c.Query("institution_name")
	major := c.Query("major")
	role := c.Query("role")

	if institutionName != "" {
		filters["institution_name"] = institutionName
	}
	if major != "" {
		filters["major"] = major
	}
	if role != "" {
		filters["role"] = role
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		response = template.Response{
			Data:    nil,
			Message: "Internal Server Error",
			Error:   err,
		}
		c.JSON(500, response)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		response = template.Response{
			Data:    nil,
			Message: "Internal Server Error",
			Error:   err,
		}
		c.JSON(500, response)
		return
	}

	offset := (page - 1) * limit

	data, count, err := ec.EducationRepository.GetEducation(limit, offset, filters)
	if err != nil {
		response = template.Response{
			Data:    nil,
			Message: "Internal Server Error",
			Error:   err,
		}
		c.JSON(500, response)
		return
	}

	responsePagination = template.ResponsePagination{
		Page:    offset,
		Perpage: limit,
		Data:    data,
		Error:   nil,
		Total:   *count,
	}

	c.JSON(200, responsePagination)
}
