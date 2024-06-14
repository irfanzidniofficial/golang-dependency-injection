package main

import (
	"irfanzidniofficial/golang-restful-api/app"
	"irfanzidniofficial/golang-restful-api/controller"
	"irfanzidniofficial/golang-restful-api/helper"
	"irfanzidniofficial/golang-restful-api/middleware"
	"irfanzidniofficial/golang-restful-api/repository"
	"irfanzidniofficial/golang-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
