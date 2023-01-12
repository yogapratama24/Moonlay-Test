package routes

import (
	"moonlay_technical_test/config"
	"moonlay_technical_test/controller"
	"moonlay_technical_test/repository"
	"moonlay_technical_test/service"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	db := config.Connect()

	// LIST
	listRepository := repository.NewListRepository(db)
	listService := service.NewListService(listRepository)
	listController := controller.NewListController(listService)

	// SUB LIST
	subListRepository := repository.NewSubListRepository(db)
	subListService := service.NewSubListService(subListRepository)
	subListController := controller.NewSubListController(subListService)

	e := echo.New()
	e.HideBanner = true
	v1 := e.Group("/api/v1")
	v1.Static("/public", "public/")
	list := v1.Group("/list")
	list.POST("", listController.CreateListController)
	list.GET("", listController.GetListController)
	list.GET("/relational", listController.GetListRelationalController)
	list.PUT("", listController.UpdateListController)
	list.DELETE("/:id", listController.DeleteListController)

	subList := v1.Group("/sub-list")
	subList.POST("", subListController.CreateSubListController)
	subList.GET("", subListController.GetSubListController)
	subList.PUT("", subListController.UpdateSubListController)
	subList.DELETE("/:id", subListController.DeleteSubListController)
	return e
}
