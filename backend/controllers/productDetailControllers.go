package controllers

import (
	"net/http"
	"strconv"

	"github.com/iqbalpradipta/BalPay/tree/main/backend/helpers"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/model"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/services"
	"github.com/labstack/echo/v4"
)

type ProductDetailController struct {
	services.ProductDetailService
}

func NewProductDetailController(s services.ProductDetailService) *ProductDetailController {
	return &ProductDetailController{s}
}

func (pc *ProductDetailController) CreateProductDetail(c echo.Context) error {
	var data model.ProductDetail

	err := c.Bind(&data); if err != nil {
		return helpers.FailedResponse(c, http.StatusBadRequest, "failed to bind data", err)
	}

	err = pc.ProductDetailService.CreateProductDetail(&data); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "failed to create data", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success Create Product detail", data)
} 

func (pc *ProductDetailController) GetAllProductDetail(c echo.Context) error {
	productDetail, err := pc.ProductDetailService.GetAllProductDetail(); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to get data", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success Get All Product Detail", productDetail)
}

func (pc *ProductDetailController) GetByIdProductDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	productDetail, err := pc.ProductDetailService.GetByIdProductDetail(uint(id)); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to get data by id", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success Get data By Id", productDetail)
}

func (pc *ProductDetailController) UpdateProductDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	productDetail, err := pc.ProductDetailService.GetByIdProductDetail(uint(id)); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to find data", err)
	}

	err = c.Bind(&productDetail); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to bind data", err)
	}

	err = pc.ProductDetailService.UpdateProductDetail(uint(id), &productDetail); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to update data", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success Update data", productDetail)

}


func (pc *ProductDetailController) DeleteProductDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := pc.ProductDetailService.DeleteProductDetail(uint(id)); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to delete data", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success to delete data", model.ProductDetail{})
}