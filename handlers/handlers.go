package handlers

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"shop/contracts"
	"shop/model"
	"strconv"
)

type ProductHandler struct {
	service contracts.IProduct
}

func NewProductHandler(service contracts.IProduct) ProductHandler {
	return ProductHandler{
		service: service,
	}
}

func (handler *ProductHandler) Index(w http.ResponseWriter, request *http.Request) {
	products, err := handler.service.GetAll()

	if err != nil {
		panic(err.Error())
	}

	var tmpl = template.Must(template.ParseGlob("templates/*"))
	tmpl.ExecuteTemplate(w, "Index.gohtml", products)
}

func (handler *ProductHandler) Show(w http.ResponseWriter, request *http.Request) {
	queryId := request.URL.Query().Get("id")
	productId, _ := strconv.Atoi(queryId)
	product, err := handler.service.GetOneById(productId)

	if err != nil {
		panic(err.Error())
	}
	var tmpl = template.Must(template.ParseGlob("templates/*"))
	tmpl.ExecuteTemplate(w, "Show", product)
}

func (handler *ProductHandler) Add(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseGlob("templates/*"))
	tmpl.ExecuteTemplate(w, "New", nil)
}

func (handler *ProductHandler) Edit(w http.ResponseWriter, r *http.Request) {
	queryId := r.URL.Query().Get("id")
	productId, _ := strconv.Atoi(queryId)
	product, err := handler.service.GetOneById(productId)

	if err != nil {
		panic(err.Error())
	}
	var tmpl = template.Must(template.ParseGlob("templates/*"))
	tmpl.ExecuteTemplate(w, "Edit", product)
}

func (handler *ProductHandler) Post(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()
	if r.Method == "POST" {
		title := r.FormValue("title")
		description := r.FormValue("description")

		// Validation
		product := model.Product{
			Title:       title,
			Description: description,
		}

		errs := validate.Struct(product)
		if errs != nil {
			var errors []string

			for _, err := range errs.(validator.ValidationErrors) {
				errors = append(errors, err.StructField())

				fmt.Println(err.StructField())
				fmt.Println(err.ActualTag())
				fmt.Println(err.Kind())
				fmt.Println(err.Value())
				fmt.Println(err.Param())
			}
			fmt.Print(errors)
			var tmpl = template.Must(template.ParseGlob("templates/*"))
			err := tmpl.ExecuteTemplate(w, "New", errors)
			if err != nil {
				panic(err)
			}
			return
		}

		err := handler.service.Create(&product)

		if err != nil {
			panic(err)
		}

	}
	http.Redirect(w, r, "/", 303)
}

func (handler *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		title := r.FormValue("title")
		description := r.FormValue("description")
		id := r.FormValue("id")

		idInt, _ := strconv.Atoi(id)

		emp := model.Product{
			Id:          idInt,
			Title:       title,
			Description: description,
		}

		err := handler.service.UpdateOne(&emp)
		if err != nil {
			panic(err)
		}

	}
	http.Redirect(w, r, "/", 301)
}

func (handler *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	queryId := r.URL.Query().Get("id")
	productId, _ := strconv.Atoi(queryId)
	err := handler.service.Delete(productId)
	if err != nil {
		panic(err.Error())
	}
	http.Redirect(w, r, "/", 301)
}

func (handler *ProductHandler) Router() *mux.Router {

	router := mux.NewRouter()

	//GET
	router.HandleFunc("/", handler.Index).Methods(http.MethodGet)
	router.HandleFunc("/show", handler.Show).Methods(http.MethodGet)
	router.HandleFunc("/add", handler.Add).Methods(http.MethodGet)
	router.HandleFunc("/edit", handler.Edit).Methods(http.MethodGet)
	router.HandleFunc("/delete", handler.Delete).Methods(http.MethodGet)

	//Post
	router.HandleFunc("/post", handler.Post).Methods(http.MethodPost)
	router.HandleFunc("/update", handler.Update).Methods(http.MethodPost)

	return router
}
