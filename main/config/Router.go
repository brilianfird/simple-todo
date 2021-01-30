package config

import "github.com/gorilla/mux"

var MyRouter *mux.Router

func InitRouter() {
	MyRouter = mux.NewRouter().StrictSlash(true)
}
