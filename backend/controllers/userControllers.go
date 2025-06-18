package controllers

import (
	"net/http"
	"strconv"

	"github.com/iqbalpradipta/BalPay/tree/main/backend/helpers"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/model"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/services"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	services.UserService
}

func NewUserController(u services.UserService) *UserController {
	return &UserController{u}
}

func (u *UserController) CreateUser(c echo.Context) error {
	var data model.User

	err := c.Bind(&data); if err != nil {
		return helpers.FailedResponse(c, http.StatusBadRequest, "Failed to bind data", err)
	}

	err = u.UserService.CreateUser(&data); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to Create User", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success Create User", data)

}

func (u *UserController) LoginUser(c echo.Context) error {
	var data model.User

	err := c.Bind(&data); if err != nil {
		return helpers.FailedResponse(c, http.StatusBadRequest, "Failed to bind data", err)
	}

	user, err := u.UserService.LoginUser(&data); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to Login User", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success Login User", user)

}

func (u *UserController) GetEmailUser(c echo.Context) error {
	var data model.User

	err := c.Bind(&data); if err != nil {
		return helpers.FailedResponse(c, http.StatusBadRequest, "Failed to bind data", err)
	}

	user, err := u.UserService.GetEmailUser(string(data.Email)); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to Create User", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success Find User", user)

}

func (u *UserController) GetAllUser(c echo.Context) error {
	data, err := u.UserService.GetAllUser(); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to get data", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success Get All Users", data)
}

func (u *UserController) GetByIdUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := u.UserService.GetByIdUser(uint(id)); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to find user by id", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success Get Data By Id", data)
}

func (u *UserController) UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := u.UserService.GetByIdUser(uint(id)); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to find data", err)
	}

	err = c.Bind(&data); if err != nil {
		return helpers.FailedResponse(c, http.StatusBadRequest, "Failed to bind data", err)
	}

	err = u.UserService.UpdateUser(uint(id), &data); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to update user", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success Update Users", data)
}

func (u *UserController) DeleteUser(c echo.Context) error {
	email := c.Param("email")
	err := u.UserService.DeleteUser(email); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to delete user", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success Delete User", model.User{})
}