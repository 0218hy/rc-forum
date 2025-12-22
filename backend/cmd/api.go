package main

import (
	"log"
	"net/http"
	repo "rc-forum-backend/db/sqlc"
	"rc-forum-backend/internal/products"
	"rc-forum-backend/internal/users"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
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

	usersService := users.NewService(repo.New(app.db))
	usersHandler := users.NewHandler(usersService)
	r.Get("/users/{id}", usersHandler.GetUserProfile)

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
	db    *pgxpool.Pool
}

type config struct {
	addr string // server address
	db   dbConfig
}

type dbConfig struct {
	dsn string // data source name
}
