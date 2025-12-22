package main

import (
	"log"
	"net/http"
	"rc-forum-backend/internal/products"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

// mount
func (app *application) mount() http.Handler {

	r := chi.NewRouter()

	// A good base middleware stack (from chi documentation)
	r.Use(middleware.RequestID) 
	r.Use(middleware.RealIP) 
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second)) 

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good\n"))
  	})

	productService := products.NewService(repo.New(app.db))
	productHandler := products.NewHandler(productService)
	r.Get("/products", productHandler.ListProducts)

	orderService := orders.NewService(repo.New(app.db), app.db)
	ordersHandler := orders.NewHandler(orderService)
	r.Post("/orders", ordersHandler.PlaceOrder)

	return r
}

//run
func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:    app.config.addr,
		Handler: h,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second, 	
		IdleTimeout: time.Minute, 
	}

	log.Printf("Starting server on %s", app.config.addr)

	return srv.ListenAndServe()
}

type application struct {
	config config
	// logger
	db    *pgx.Conn
}

type config struct {
	addr string // server address
	db   dbConfig
}

type dbConfig struct {
	dsn string // data source name
}
