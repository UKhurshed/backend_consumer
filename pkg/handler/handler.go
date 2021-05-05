package handler

import (
	_ "backend_consumer/docs"
	"backend_consumer/pkg/service"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler{
	return &Handler{services: services}
}

func (h *Handler) InitRoute() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		api.GET("/", h.GetAllBuildings)
		api.POST("/",  h.CreateBuildingItem)
		//api.GET("/:id")
		api.PUT("/:id", h.UpdateBuildingItem)
		api.DELETE("/:id", h.DeleteBuildingItem)

	}

	subject := router.Group("/subject")
	{
		subject.GET("/", h.GetAllSubject)
		subject.POST("/", h.CreateSubject)
		subject.PUT("/:id", h.UpdateSubject)
		subject.DELETE("/:id", h.DeleteSubject)
	}

	//config := cors.DefaultConfig()
	//config.AllowOrigins = []string{"http://localhost:5000/"}

	//router.Use(cors.New(cors.Config{
	//	AllowOrigins: []string{"*"},
	//	AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	//
	//}))

	router.Use(CORSMiddleware())

	return router
}
