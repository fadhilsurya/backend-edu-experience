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
}

func initializeCandidateRoutes(r *gin.Engine, db *gorm.DB) {
	candidateRepo := repository.NewCandidateRepository(db)
	candidateController := controller.NewCandidateController(candidateRepo)

	r.POST("/login", candidateController.Login)
	r.POST("/signup", candidateController.CreateCandidate)
	candidateRoutes := r.Group("/candidate")
	candidateRoutes.Use(middleware.TokenAuthMiddleware())
	{
		candidateRoutes.PUT("/", candidateController.UpdateCandidate)
		// userRoutes.GET("/:id", userController.GetById)
		// userRoutes.POST("/", userController.CreateUser)

	}
}
