package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io/ioutil"
	"net/http"
	"todo/main/config"
	"todo/main/model/web"
	"todo/main/service"
)

func HandleToDo() {
	config.MyRouter.HandleFunc("/todos", GetAll).Methods("GET")
	config.MyRouter.HandleFunc("/todos", Insert).Methods("POST")
}

func Insert(w http.ResponseWriter, r *http.Request) {

	fmt.Println("API Insert ToDo is hit")

	var insertToDoRequest web.InsertToDoRequest

	err := GetRequestBodyAndValidate(&w, r, &insertToDoRequest)

	if err != nil {
		return
	}

	result := service.CreateToDo(insertToDoRequest.Title, insertToDoRequest.Description)

	json.NewEncoder(w).Encode(web.Response{
		StatusCode: 200,
		Error:      nil,
		Data:       result,
	})
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API GetAll ToDo is hit")
	fmt.Println("API GetAll ToDo is hit")
	todos := service.GetAllToDo()
	fmt.Println("API GetAll ToDo is hit")
	json.NewEncoder(w).Encode(web.Response{
		StatusCode: 200,
		Error:      nil,
		Data:       todos,
	})
}

func GetRequestBodyAndValidate(w *http.ResponseWriter, r *http.Request, mod interface{}) error {
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, mod)
	val := validator.New()
	err := val.Struct(mod)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors).Error()
		json.NewEncoder(*w).Encode(web.Response{
			StatusCode: 400,
			Error:      &validationErrors,
			Data:       nil,
		})
	}
	
	return err
}
