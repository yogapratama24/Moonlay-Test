package controller

import (
	"log"
	"moonlay_technical_test/helper"
	"moonlay_technical_test/model"
	"moonlay_technical_test/service"
	"strconv"

	"github.com/labstack/echo"
)

type ListController struct {
	service service.ListService
}

func NewListController(service service.ListService) *ListController {
	return &ListController{service}
}

func (listController *ListController) GetListController(c echo.Context) error {
	// var request model.ListRequest
	var request helper.Pagination
	if err := c.Bind(&request); err != nil {
		helper.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return nil
	}
	if request.Id != 0 {
		list, err := listController.service.GetDetailList(request.Id)
		if err != nil {
			helper.NewHandlerResponse(err.Error(), nil).Failed(c)
			return nil
		}

		helper.NewHandlerResponse("Successfully get list detail", list).Success(c)
		return nil
	}
	lists, err := listController.service.GetList(request)
	if err != nil {
		helper.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	helper.NewHandlerResponse("Successfully get list", lists).Success(c)
	return nil
}

func (listController *ListController) GetListRelationalController(c echo.Context) error {
	lists, err := listController.service.GetListRelational()
	if err != nil {
		helper.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	helper.NewHandlerResponse("Successfully get list with relational", lists).Success(c)
	return nil
}

func (listController *ListController) CreateListController(c echo.Context) error {
	var (
		request  model.ListCreateRequest
		fullPath string
		err      error
	)
	if err := c.Bind(&request); err != nil {
		helper.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return nil
	}

	if err := helper.DoValidation(&request); err != nil {
		helper.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return nil
	}

	// UPLOAD IMAGE TO LOCAL
	file, _ := c.FormFile("file")
	if file != nil {
		path := "upload/list"
		fullPath, err = helper.FileUploadHandler(file, path)
		if err != nil {
			helper.NewHandlerResponse("Failed to upload image file", nil).Failed(c)
			log.Printf("Error upload file with: %v", err)
			return nil
		}
	}
	request.FilePath = fullPath

	if err := listController.service.CreateList(request); err != nil {
		helper.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}
	helper.NewHandlerResponse("Successfully create list", nil).SuccessCreate(c)
	return nil
}

func (listController *ListController) UpdateListController(c echo.Context) error {
	var (
		request  model.ListUpdateRequest
		fullPath string
		err      error
	)
	if err := c.Bind(&request); err != nil {
		helper.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return nil
	}

	if err := helper.DoValidation(&request); err != nil {
		helper.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return nil
	}

	// UPLOAD IMAGE TO LOCAL
	file, _ := c.FormFile("file")
	if file != nil {
		path := "upload/list"
		fullPath, err = helper.FileUploadHandler(file, path)
		if err != nil {
			helper.NewHandlerResponse("Failed to upload image file", nil).Failed(c)
			log.Printf("Error upload file with: %v", err)
			return nil
		}
	}
	request.FilePath = fullPath

	if err := listController.service.UpdateList(request); err != nil {
		helper.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}
	helper.NewHandlerResponse("Successfully update list", nil).Success(c)
	return nil
}

func (listController *ListController) DeleteListController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.NewHandlerResponse("Error convert data", nil).BadRequest(c)
		return nil
	}

	if err := listController.service.DeleteList(id); err != nil {
		helper.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	helper.NewHandlerResponse("Successfully delete list", nil).Success(c)
	return nil
}
