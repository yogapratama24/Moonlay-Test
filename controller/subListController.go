package controller

import (
	"log"
	"moonlay_technical_test/helper"
	"moonlay_technical_test/model"
	"moonlay_technical_test/service"
	"strconv"

	"github.com/labstack/echo"
)

type SubListController struct {
	service service.SubListService
}

func NewSubListController(service service.SubListService) *SubListController {
	return &SubListController{service}
}

func (subListController *SubListController) GetSubListController(c echo.Context) error {
	var request model.SubListRequest
	if err := c.Bind(&request); err != nil {
		helper.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return nil
	}
	if err := helper.DoValidation(&request); err != nil {
		helper.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return nil
	}
	if request.Id != 0 {
		list, err := subListController.service.GetDetailSubList(request.Id)
		if err != nil {
			helper.NewHandlerResponse(err.Error(), nil).Failed(c)
			return nil
		}

		helper.NewHandlerResponse("Successfully get sub list detail", list).Success(c)
		return nil
	}
	lists, err := subListController.service.GetSubList(request.ListId)
	if err != nil {
		helper.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	helper.NewHandlerResponse("Successfully get sub list", lists).Success(c)
	return nil
}

func (subListController *SubListController) CreateSubListController(c echo.Context) error {
	var (
		request  model.SubListCreateRequest
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
		path := "upload/sublist"
		fullPath, err = helper.FileUploadHandler(file, path)
		if err != nil {
			helper.NewHandlerResponse("Failed to upload image file", nil).Failed(c)
			log.Printf("Error upload file with: %v", err)
			return nil
		}
	}
	request.FilePath = fullPath

	if err := subListController.service.CreateSubList(request); err != nil {
		helper.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}
	helper.NewHandlerResponse("Successfully create sub list", nil).SuccessCreate(c)
	return nil
}

func (subListController *SubListController) UpdateSubListController(c echo.Context) error {
	var (
		request  model.SubListUpdateRequest
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

	if err := subListController.service.UpdateSubList(request); err != nil {
		helper.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}
	helper.NewHandlerResponse("Successfully update sub list", nil).Success(c)
	return nil
}

func (subListController *SubListController) DeleteSubListController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.NewHandlerResponse("Error convert data", nil).BadRequest(c)
		return nil
	}

	if err := subListController.service.DeleteSubList(id); err != nil {
		helper.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	helper.NewHandlerResponse("Successfully delete sub list", nil).Success(c)
	return nil
}
