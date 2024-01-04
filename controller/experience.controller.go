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

type ExperienceController struct {
	ExperienceRepository *repository.ExperienceRepository
}

func NewExperienceController(er *repository.ExperienceRepository) *ExperienceController {
	return &ExperienceController{
		ExperienceRepository: er,
	}
}

func (ec *ExperienceController) CreateExperience(c *gin.Context) {

	ctx := context.Background()

	var (
		newExperienceReq models.ExperienceCreateRequest
		resp             template.Response
		experienceModel  []models.Experience
		endYear          *string
		e                time.Time
		s                time.Time
	)

	if err := c.ShouldBindJSON(&newExperienceReq); err != nil {
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

	for _, v := range newExperienceReq.Data {

		if v.UntilNow && v.EndYear != nil {
			resp = template.Response{
				Data:    nil,
				Error:   err,
				Message: "Bad Request - End Year should be empty",
			}

			c.JSON(500, resp)
			return

		}

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

		expModel := models.Experience{
			CandidateID:    uint(id),
			CompanyName:    v.CompanyName,
			CompanyAddress: v.CompanyAddress,
			Position:       v.Position,
			JobDesc:        v.JobDesc,
			Flag:           &v.Flag,
			StartYear:      &s,
			EndYear:        &e,
			UntilNow:       v.UntilNow,
		}
		experienceModel = append(experienceModel, expModel)
	}

	err = ec.ExperienceRepository.BatchInsertExperience(ctx, experienceModel)
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

func (ec *ExperienceController) UpdateExperience(c *gin.Context) {
	var (
		newExperienceReq models.ExperienceRequest
		resp             template.Response
		expModel         models.Experience
		endyear          time.Time
		startyear        time.Time
	)

	experienceId := c.Param("id")
	ctx := context.Background()

	if err := c.ShouldBindJSON(&newExperienceReq); err != nil {
		resp = template.Response{
			Data:    nil,
			Error:   err,
			Message: "Bad Request",
		}

		c.JSON(400, resp)
		return
	}

	expId, err := strconv.Atoi(experienceId)
	if err != nil {
		resp = template.Response{
			Data:    nil,
			Error:   err,
			Message: "Internal Server Error",
		}

		c.JSON(500, resp)
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
	filterCheck["id"] = expId

	data, err := ec.ExperienceRepository.GetOneExperience(filterCheck)
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

	if newExperienceReq.StartYear != nil {
		startYear := newExperienceReq.StartYear
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

	if newExperienceReq.EndYear != nil {
		endYear := newExperienceReq.EndYear
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

	if newExperienceReq.UntilNow && newExperienceReq.StartYear != nil {

		resp = template.Response{
			Data:    nil,
			Error:   err,
			Message: "Bad Request",
		}
		c.JSON(400, resp)
		return
	}

	expModel = models.Experience{
		CompanyName:    newExperienceReq.CompanyName,
		CompanyAddress: newExperienceReq.CompanyAddress,
		StartYear:      &startyear,
		EndYear:        &endyear,
		Position:       newExperienceReq.Position,
		JobDesc:        newExperienceReq.Position,
		UntilNow:       newExperienceReq.UntilNow,
		Flag:           &newExperienceReq.Flag,
	}

	err = ec.ExperienceRepository.UpdateExperience(ctx, expId, expModel)
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

func (ec *ExperienceController) DeleteExperience(c *gin.Context) {

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

	data, err := ec.ExperienceRepository.GetOneExperience(filters)
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
			Message: "Bad Request - Experience Does Not Exist",
		}

		c.JSON(400, resp)
		return
	}

	err = ec.ExperienceRepository.DeleteExperience(int(data.ID))
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

func (ec *ExperienceController) GetExperience(c *gin.Context) {
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

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	companyName := c.Query("company_name")
	companyAddress := c.Query("company_address")
	position := c.Query("position")
	jobDesc := c.Query("job_desc")

	filters := make(map[string]interface{})
	filters["candidate_id"] = id

	if companyAddress != "" {
		filters["company_address"] = companyAddress
	}
	if companyName != "" {
		filters["company_name"] = companyName
	}
	if jobDesc != "" {
		filters["job_desc"] = jobDesc
	}
	if position != "" {
		filters["position"] = position
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

	data, count, err := ec.ExperienceRepository.GetExperience(limit, offset, filters)
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
