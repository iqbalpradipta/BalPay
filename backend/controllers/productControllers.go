package controllers

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/iqbalpradipta/BalPay/tree/main/backend/helpers"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/model"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/services"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/utils"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	services.ProductService
}

func NewProductController(s services.ProductService) *ProductController {
	return &ProductController{s}
}

func (pc *ProductController) CreateProduct(c echo.Context) error {
	var product model.Product

	product.Name = c.FormValue("name")
	product.Description = c.FormValue("description")
	file, err := c.FormFile("image"); if err != nil {
		return helpers.FailedResponse(c, http.StatusBadRequest, "File not found", err)
	}

	src, _ := file.Open()
	defer src.Close()

	tempPath := "uploads/" + file.Filename
	dst, err := os.Create(tempPath); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to create temp file", err)
	}
	defer dst.Close()
	io.Copy(dst, src)

	product.Image, err = utils.UploadToCloudinary(tempPath); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed upload to cloudinary", err)
	}
	
	os.Remove(tempPath)

	err = pc.ProductService.CreateProduct(&product); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to create product", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success Create Data", product)
}

func (pc *ProductController) GetAllProduct(c echo.Context) error {
	product, err := pc.ProductService.GetAllProduct()
	if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to get Product", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success Get Data", product)
}

func (pc *ProductController) GetByIdProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id")) 
	product, err := pc.ProductService.GetByIdProduct(uint(id)); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to get product by Id", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success Get Product By id", product)
}

func (pc *ProductController) UpdateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id")) 
	product, err := pc.ProductService.GetByIdProduct(uint(id)); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to get product by Id", err)
	}

	if err := c.Bind(&product); err != nil {
		return helpers.FailedResponse(c, http.StatusBadRequest, "Bind data Failed", err)
	}

	if err := pc.ProductService.UpdateProduct(uint(id), &product); err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to Update product", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success Update Product", product)
}

func (pc *ProductController) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id")) 

	if err := pc.ProductService.DeleteProduct(uint(id)); err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to delete product", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success Delete Product", model.Product{})
}