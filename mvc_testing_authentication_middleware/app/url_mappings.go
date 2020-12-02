package app

import (
	"../controllers"
	"../middlewares"
)

func mapUrls() {
	r.HandleFunc("/signup", controllers.Signup).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/products/", middlewares.Middleware(controllers.GetProducts)).Methods("GET")
	r.HandleFunc("/products/", middlewares.Middleware(controllers.CreateProduct)).Methods("POST")
	r.HandleFunc("/products/{id:[0-9]+}", middlewares.Middleware(controllers.UpdateProduct)).Methods("PUT")
	r.HandleFunc("/products/{id:[0-9]+}", middlewares.Middleware(controllers.GetSingleProduct)).Methods("GET")
	r.HandleFunc("/products/{id:[0-9]+}", middlewares.Middleware(controllers.DeleteProduct)).Methods("DELETE")
	r.HandleFunc("/logout", middlewares.Middleware(controllers.Logout)).Methods("POST")
}
