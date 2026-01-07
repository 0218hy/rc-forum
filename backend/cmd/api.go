package main

import (
	"log"
	"net/http"
	repo "rc-forum-backend/db/sqlc"
	"rc-forum-backend/internal/auth"
	"rc-forum-backend/internal/env"
	"rc-forum-backend/internal/posts"
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

	// secret key 
	var secretKey = env.GetString("secretKey", "01234567890123456789012345678901") // 32 chars
	if len(secretKey) < 32 {
		log.Fatal("secretKey must be at least 32 characters long")
	}

	// health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good\n"))
  	})

	// For users
	usersService := users.NewService(repo.New(app.db))
	usersHandler := users.NewHandler(usersService)
	r.Get("/users/{id}", usersHandler.GetUserByID)
	r.Get("/users/email/{email}", usersHandler.GetUserByEmail)

	// For auth
	authService := auth.NewService(repo.New(app.db))
	authHandler := auth.NewHandler(authService, usersService, secretKey)
	r.Post("/register", authHandler.HandleRegister)
	r.Post("/login", authHandler.HandleLogin)
	r.Post("/logout", authHandler.HandleLogout)
	r.Post("/renew_access_token", authHandler.RenewAccessToken)
	

	// For posts
	postsService := posts.NewService(repo.New(app.db), app.db)
	postsHandler := posts.NewHandler(postsService)
	r.Get("/posts", postsHandler.GetAllPosts)
	r.Get("/posts/{id}", postsHandler.GetPostByID)
	r.Delete("/posts/{id}", postsHandler.DeletePostByID)
	r.Post("/posts", postsHandler.CreatePost)

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
