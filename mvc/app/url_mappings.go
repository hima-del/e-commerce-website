package app

import "../controllers"

func mapUrls() {
	r.HandleFunc("/signup", controllers.Signup).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	//r.HandleFunc("/products/", controllers.GetProducts).Methods("GET")
	r.HandleFunc("/products/", controllers.CreateProduct).Methods("POST")
	//r.HandleFunc("/products/{id:[0-9]+}", controllers.UpdateProduct).Methods("PUT")
	//r.HandleFunc("/products/{id:[0-9]+}", controllers.GetSingleProduct).Methods("GET")
	//r.HandleFunc("/products/{id:[0-9]+}", controllers.DeleteProduct).Methods("DELETE")
	//r.HandleFunc("/logout", controllers.Logout).Methods("POST")
}
