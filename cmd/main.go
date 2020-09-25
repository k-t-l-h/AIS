package main

import "AIS/internal/doi"

func main() {

	doi.GetHrefs()
	/*
	r := mux.NewRouter()

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

	})

	srv := &http.Server{
		Handler:      r,
		Addr:         ":5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Print("Server running at ", srv.Addr)
	log.Fatal(srv.ListenAndServe())

	 */
}
