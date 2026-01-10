package users_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	repo "rc-forum-backend/db/sqlc"
	"rc-forum-backend/internal/auth"
	"rc-forum-backend/internal/users"

	"github.com/go-chi/chi/v5"
)

// Mock Service

type mockUserService struct{}

func (m *mockUserService) GetUserByID(ctx context.Context, userID int32) (repo.User, error) {
	if userID == 1 {
		return repo.User{ID: 1, Name: "Hayoung", Email: "hayoung@test.com", IsAdmin: false}, nil
	}
	// simulate user not found
	return repo.User{}, sql.ErrNoRows
}

func (m *mockUserService) GetUserByEmail(ctx context.Context, email string) (repo.User, error) {
	if email == "hayoung@test.com" {
		return repo.User{ID: 1, Name: "Hayoung", Email: email, IsAdmin: false}, nil
	}
	// simulate user not found
	return repo.User{}, sql.ErrNoRows
}

func (m *mockUserService) CreateUser(ctx context.Context, tempUser users.CreateUserParams) (int32, error) {
	return 2, nil
}

func (m *mockUserService) DeleteUserByID(ctx context.Context, userID int32) error {
	if userID == 1 {
		return nil
	}
	return sql.ErrNoRows
}

// Handlers Tests

func TestGetMyProfile(t *testing.T) {
	service := &mockUserService{}
	handler := users.NewHandler(service)

	t.Run("GetMyProfile_JWT", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/me", nil)
		rec := httptest.NewRecorder()

		// Simulate JWT auth 
		claims := &auth.UserClaims{
			ID:    1,
			IsAdmin: false,
		}
		ctx := context.WithValue(req.Context(), auth.AuthKey{}, claims)
		req = req.WithContext(ctx)

		handler.GetMyProfile(rec, req)

		var received repo.User
		if err := json.NewDecoder(rec.Body).Decode(&received); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		want := repo.User{ID: 1, Name: "Hayoung", Email: "hayoung@test.com", IsAdmin: false}
		if received != want {
			t.Errorf("GetMyProfile() = %v, want %v", received, want)
		}
	})

	t.Run("GetUserByID_ValidID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
		rec := httptest.NewRecorder()

		// manually set chi URL param
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		handler.GetUserByID(rec, req)

		var received repo.User
		if err := json.NewDecoder(rec.Body).Decode(&received); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		want := repo.User{ID: 1, Name: "Hayoung", Email: "hayoung@test.com", IsAdmin: false}
		if received != want {
			log.Println("received:", received, "\n", "want:", want)
			t.Errorf("GetUserByID() = %v, want %v", received, want)
		}
	})

	t.Run("GetUserByID_InvalidID_notinteger", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/abc", nil)
		rec := httptest.NewRecorder()

		// manually set chi URL param
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("id", "abc")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		handler.GetUserByID(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("GetUserByID() status = %v, want %v", rec.Code, http.StatusBadRequest)
		}
	})

	t.Run("GetUserByID_InvalidID_notfound", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/999", nil)
		rec := httptest.NewRecorder()

		// manually set chi URL param
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("id", "999")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		handler.GetUserByID(rec, req)

		if rec.Code != http.StatusNotFound {
			t.Errorf("GetUserByID() status = %v, want %v", rec.Code, http.StatusNotFound)
		}
	})

	t.Run("GetUserByEmail_ValidEmail", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/email/hayoung@test.com", nil)
		rec := httptest.NewRecorder()

		// manually set chi URL param
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("email", "hayoung@test.com")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		handler.GetUserByEmail(rec, req)

		var received repo.User
		if err := json.NewDecoder(rec.Body).Decode(&received); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		want := repo.User{ID: 1, Name: "Hayoung", Email: "hayoung@test.com", IsAdmin: false}
		if received != want {
			t.Errorf("GetUserByEmail() = %v, want %v", received, want)
		}
	})

	t.Run("GetUserByEmail_InvalidEmail", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/email/invalid-email", nil)
		rec := httptest.NewRecorder()

		handler.GetUserByEmail(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("GetUserByEmail() status = %v, want %v", rec.Code, http.StatusBadRequest)
		}
	})

	t.Run("DeleteUserByID_ValidID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
		rec := httptest.NewRecorder()

		// manually set chi URL param
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		handler.DeleteUserByID(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("DeleteUserByID() status = %v, want %v", rec.Code, http.StatusOK)
		}
	})

	t.Run("DeleteUserByID_InvalidID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/users/999", nil)
		rec := httptest.NewRecorder()

		// manually set chi URL param
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("id", "999")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		handler.DeleteUserByID(rec, req)

		if rec.Code != http.StatusNotFound {
			t.Errorf("DeleteUserByID() status = %v, want %v", rec.Code, http.StatusNotFound)
		}
	})
}