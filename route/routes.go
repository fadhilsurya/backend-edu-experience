package route

import (
	"backend-edu-experience/controller"
	"backend-edu-experience/middleware"
	"backend-edu-experience/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeRoutes(r *gin.Engine, db *gorm.DB) {
	initializeCandidateRoutes(r, db)
	initializeEducationRoutes(r, db)
	initializeExperienceRoutes(r, db)
}

func initializeCandidateRoutes(r *gin.Engine, db *gorm.DB) {
	candidateRepo := repository.NewCandidateRepository(db)
	candidateController := controller.NewCandidateController(candidateRepo)

	r.POST("/signin", candidateController.Login)
	r.POST("/signup", candidateController.CreateCandidate)
	candidateRoutes := r.Group("/candidate")
	candidateRoutes.Use(middleware.TokenAuthMiddleware())
	{
		candidateRoutes.POST("/signout", candidateController.SignOut)
		candidateRoutes.PUT("/", candidateController.UpdateCandidate)
		candidateRoutes.DELETE("/", candidateController.DeleteCandidate)
	}
}

func initializeEducationRoutes(r *gin.Engine, db *gorm.DB) {
	educationRepo := repository.NewEducationRepository(db)
	educationController := controller.NewEducationController(educationRepo)

	educationRoutes := r.Group("/education")
	educationRoutes.Use(middleware.TokenAuthMiddleware())
	{
		educationRoutes.POST("/", educationController.CreateEducation)
		educationRoutes.GET("/", educationController.GetEducation)
		educationRoutes.PUT("/:id", educationController.UpdateEducation)
		educationRoutes.DELETE("/:id", educationController.DeleteEducation)
	}
}

func initializeExperienceRoutes(r *gin.Engine, db *gorm.DB) {
	experienceRepo := repository.NewExperienceRepository(db)
	experienceController := controller.NewExperienceController(experienceRepo)

	exprienceRoutes := r.Group("/experience")
	exprienceRoutes.Use(middleware.TokenAuthMiddleware())
	{
		exprienceRoutes.POST("/", experienceController.CreateExperience)
		exprienceRoutes.GET("/", experienceController.GetExperience)
		exprienceRoutes.PUT("/:id", experienceController.UpdateExperience)
		exprienceRoutes.DELETE("/:id", experienceController.DeleteExperience)
	}
}
