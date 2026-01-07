package authhttp

import (
	"log"
	"net/http"

	"rc-forum-backend/internal/auth"
	"rc-forum-backend/internal/json"
	"rc-forum-backend/internal/users"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type handler struct {
	service Service
	userService users.Service
	tokenMaker *auth.JWTMaker
}

func NewHandler(service Service, userService users.Service, tokenMaker *auth.JWTMaker) *handler {
	return &handler{
		service:    service,
		userService: userService,
		tokenMaker:  tokenMaker,
	}
}

func (h *handler) HandleRegister (w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload
	if err := json.Read(r, &payload); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// check if user exists
	_, err := h.userService.GetUserByEmail(r.Context(), payload.Email)
	if err == nil {
		http.Error(w, "user already exists", http.StatusBadRequest)
		return
	}
	// hash password
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to hash password", http.StatusInternalServerError)
		return 
	}
	// convert payload to user
	tempUser := users.CreateUserParams{
		Name: payload.Name,
		Email: payload.Email,
		Password: hashedPassword,
		IsAdmin: payload.IsAdmin,
	}
	// create user
	userID, err := h.userService.CreateUser(r.Context(), tempUser)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return 
	}
	json.Write(w, http.StatusOK, userID)
}

func (h *handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var payload LoginUserPayload
	if err := json.Read(r, &payload); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// check if email exists
	user, err := h.userService.GetUserByEmail(r.Context(), payload.Email)
	if err != nil {
		http.Error(w, "Not registered email", http.StatusBadRequest)
		return
	}
	// check password with hashed password
	if !auth.ComparePassword(user.Password, []byte(payload.Password)) {
		http.Error(w, "Invalid email or password", http.StatusBadRequest)
		return 
	}
	// create JWT token
	accessToken, accessClaims, err := h.tokenMaker.CreateToken(int32(user.ID), user.Name, user.Email, user.IsAdmin, 15*time.Minute)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to create token", http.StatusInternalServerError)
		return
	}
	refreshToken, refreshClaims, err := h.tokenMaker.CreateToken(int32(user.ID), user.Name, user.Email, user.IsAdmin, 24*time.Hour)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to create token", http.StatusInternalServerError)
		return
	}
	// create session
	session, err := h.service.CreateSession(r.Context(), CreateSessionParams{
		UserID:      int32(user.ID),
		RefreshToken: refreshToken,
		IsRevoked:    false,
		ExpiresAt:    pgtype.Timestamp{Time: refreshClaims.ExpiresAt.Time, Valid: true},
		CreatedAt:    pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to create session", http.StatusInternalServerError)
		return
	}
	// return response
	response := SessionResponse{
		SessionID:            session.ID,
		AccessToken:          accessToken,
		RefreshToken:         refreshToken,
		AccessTokenExpiresAt: accessClaims.ExpiresAt.Time,
		RefreshTokenExpiresAt: refreshClaims.ExpiresAt.Time,
		User:                 toUser(user),
	}
	json.Write(w, http.StatusOK, response)
}

func (h *handler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	// get session ID from token
	claims := r.Context().Value(auth.AuthKey{}).(*auth.UserClaims)
	if claims == nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	err := h.service.RevokeSession(r.Context(), claims.RegisteredClaims.ID)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to revoke session", http.StatusInternalServerError)
		return
	}
	json.Write(w, http.StatusOK, "logged out successfully")
}

func (h *handler) RenewAccessToken(w http.ResponseWriter, r *http.Request) {
	var payload RenewAccessTokenPayload
	if err := json.Read(r, &payload); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	refreshClaims, err := h.tokenMaker.VerifyToken(payload.RefreshToken)
	if err != nil {
		log.Println(err)
		http.Error(w, "invalid refresh token", http.StatusUnauthorized)
		return
	}

	session, err := h.service.GetSession(r.Context(), refreshClaims.RegisteredClaims.ID)
	if err != nil {
		log.Println(err)
		http.Error(w, "error getting session", http.StatusNotFound)
		return
	}
	if session.IsRevoked {
		http.Error(w, "session has been revoked", http.StatusUnauthorized)
		return
	}
	if session.Email != refreshClaims.Email {
		http.Error(w, "token email does not match session email", http.StatusUnauthorized)
		return
	}

	accessToken, accessClaims, err := h.tokenMaker.CreateToken(session.UserID, refreshClaims.Name, refreshClaims.Email, refreshClaims.IsAdmin, 15*time.Minute)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to create access token", http.StatusInternalServerError)
		return
	}
	
	response := RenewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessClaims.ExpiresAt.Time,
	}

	json.Write(w, http.StatusOK, response)
}
