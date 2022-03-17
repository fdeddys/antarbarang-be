package router

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"com.ddabadi.antarbarang/dto"
	handlers "com.ddabadi.antarbarang/handlers"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {

	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.Use(cekToken)
	r.Use(mux.CORSMethodMiddleware(r))
	pathPref := "/api"

	s := r.PathPrefix(pathPref + "/seller").Subrouter()
	s.HandleFunc("/{id:[0-9]+}", handlers.GetSellerByIDHandler).Methods(http.MethodGet)
	s.HandleFunc("/code/{code}", handlers.GetSellerByCodeHandler).Methods(http.MethodGet)
	s.HandleFunc("", handlers.SaveSellerHandler).Methods(http.MethodPost)

	s = r.PathPrefix(pathPref + "/customer").Subrouter()
	s.HandleFunc("/{id:[0-9]+}", handlers.GetCustomerByIDHandler).Methods(http.MethodGet)
	s.HandleFunc("/seller-id/{sellerId}", handlers.GetCustomerBySellerIdHandler).Methods(http.MethodGet)
	s.HandleFunc("/nama/{nama}", handlers.GetCustomerByNamaHandler).Methods(http.MethodGet)
	s.HandleFunc("/", handlers.CustomerCreateHandler).Methods(http.MethodPost)

	s = r.PathPrefix(pathPref + "/driver").Subrouter()
	s.HandleFunc("/{id:[0-9]+}", handlers.GetDriverByIdHandler).Methods(http.MethodGet)
	s.HandleFunc("/code/{code}", handlers.GetDriverByCodeHandler).Methods(http.MethodGet)
	s.HandleFunc("", handlers.DriverCreateHandler).Methods(http.MethodPost)

	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil && len(queriesTemplates) > 0 {
			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil && len(queriesRegexps) > 0 {
			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		fmt.Println()
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return r
}

func cekToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		dto.CurrUser = "system"
		if token != "" {
			// We found the token in our map
			log.Printf("Authenticated user ")
			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)
		} else {
			// Write an error and stop the handler chain
			log.Printf("Forbidden user ")
			next.ServeHTTP(w, r)
		}
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
