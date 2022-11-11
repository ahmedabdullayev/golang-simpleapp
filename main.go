package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"shop/database"
	"shop/handlers"
	"shop/repositories"
	"shop/services"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
func Run() error {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	db, err := database.Init()
	if err != nil {
		return err
	}
	empRep := repositories.NewProductRepository(db)
	empSer := services.NewProductService(empRep)
	hand := handlers.NewProductHandler(empSer)
	r := hand.Router()
	http.ListenAndServe(":8083", r)
	return err
}
