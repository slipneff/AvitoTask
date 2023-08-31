package server

import (
	_ "avito/docs"
	"avito/internal/services"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func NewHTTPServer() {
	r := mux.NewRouter()
	r.HandleFunc("/user", services.CreateUser).Methods("POST")
	r.HandleFunc("/user/segment", services.GetUserWithSegments).Methods("POST")
	r.HandleFunc("/user/addSegment", services.AddSegmentsToUser).Methods("POST")
	r.HandleFunc("/user/all", services.GetUsers).Methods("GET")
	r.HandleFunc("/segment", services.CreateSegment).Methods("POST")
	r.HandleFunc("/segment", services.DeleteSegment).Methods("DELETE")
	r.HandleFunc("/history", services.GenerateReportCSV).Methods("POST")
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	srv := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	srv.ListenAndServe()
}
