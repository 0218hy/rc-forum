package main

import (
	"log"
	"net/http"
	repo "rc-forum-backend/db/sqlc"
	"rc-forum-backend/internal/auth/authhttp"
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

	// create token maker
	tokenMaker := auth.NewJWTMaker(secretKey)

	// health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good!\n"))
  	})

	// For users
	usersService := users.NewService(repo.New(app.db))
	usersHandler := users.NewHandler(usersService)
	// For admin users
	r.Group(func(r chi.Router) {
		r.Use(auth.GetAdminMiddlewareFunc(tokenMaker))
		r.Get("/users/{id}", usersHandler.GetUserByID)
		r.Get("/users/{email}", usersHandler.GetUserByEmail)
		r.Delete("/users/{id}", usersHandler.DeleteUserByID)
	})
	// For auth users
	r.Group(func(r chi.Router) {
		r.Use(auth.GetAuthMiddlewareFunc(tokenMaker))
		r.Get("/users/me", usersHandler.GetMyProfile)
	})

	// For auth
	authService := authhttp.NewService(repo.New(app.db))
	authHandler := authhttp.NewHandler(authService, usersService, tokenMaker)
	// For public
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", authHandler.HandleRegister)
		r.Post("/login", authHandler.HandleLogin)
		r.Post("/renew", authHandler.RenewAccessToken)
	})
	//For auth users
	r.Group(func(r chi.Router) {
		r.Use(auth.GetAuthMiddlewareFunc(tokenMaker))
		r.Post("/logout", authHandler.HandleLogout)
	})

	// For posts
	postsService := posts.NewService(repo.New(app.db), app.db)
	postsHandler := posts.NewHandler(postsService)
	r.Get("/posts", postsHandler.GetAllPosts)
	r.Get("/posts/{id}", postsHandler.GetPostByID)
	// Auth user 
	r.Group(func(r chi.Router) {
		r.Use(auth.GetAuthMiddlewareFunc(tokenMaker))
		r.Patch("/posts/{id}", postsHandler.UpdatePostCore)
		r.Post("/posts", postsHandler.CreatePost)
	})

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
