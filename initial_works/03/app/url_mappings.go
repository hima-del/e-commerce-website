package app

func mapUrls() {
	r.HandleFunc("/signup", signup).Methods("POST")
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/products/", getProducts).Methods("GET")
	r.HandleFunc("/products/", createProduct).Methods("POST")
	r.HandleFunc("/products/{id:[0-9]+}", updateProduct).Methods("PUT")
	r.HandleFunc("/products/{id:[0-9]+}", getSingleProduct).Methods("GET")
	r.HandleFunc("/products/{id:[0-9]+}", deleteProduct).Methods("DELETE")
	r.HandleFunc("/logout", logout).Methods("POST")
}
