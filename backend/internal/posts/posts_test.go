package posts_test
/*
import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	repo "rc-forum-backend/db/sqlc"
	"rc-forum-backend/internal/posts"
	"testing"
)

type mockPostService struct {}

func (m *mockPostService) GetAllPosts(ctx context.Context) ([]repo.ListAllPostsRow, error) {
	return []repo.ListAllPostsRow{
		{ID: 42, Title: "Hello", Body: "World", AuthorID: 1},
	}, nil
}

func (m *mockPostService) GetPostByID(ctx context.Context, postID int32) (repo.FindPostByIDRow, error) {
	return repo.FindPostByIDRow{ID: postID, Title: "Hello", Body: "World", AuthorID: 1}, nil
}

func (m *mockPostService) DeletePostByID(ctx context.Context, postID int32) error { return nil }

func (m *mockPostService) CreatePost(ctx context.Context, req posts.CreatePostRequest) (int32, error) {
	// Just return fixed post ID for testing
	return 42, nil
}

func (m *mockPostService) UpdatePostCore(ctx context.Context, postID int32, req posts.UpdatePostCoreRequest) error {
	return nil
}

/* func TestPostsHandlers(t *testing.T) {
	service := &mockPostService{}
	handler := posts.NewHandler(service)

	t.Run("GetAllPosts explicit output", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/posts", nil)
		rec := httptest.NewRecorder()

		handler.GetAllPosts(rec, req)

		var got []posts.ListAllPostsRow
		if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
			t.Fatalf("failed to decode: %v", err)
		}

		want := []posts.ListAllPostsRow{
			{ID: 42, Title: "Hello", Body: "World", AuthorID: 1},
		}

		if len(got) != len(want) || got[0] != want[0] {
			t.Fatalf("expected %+v, got %+v", want, got)
		}
	})

	t.Run("CreatePost with subtypes", func(t *testing.T) {
		subtypes := []struct {
			name     string
			jsonBody string
		}{
			{
				name: "announcement",
				jsonBody: `{
					"author_id": 1,
					"title": "Announcement",
					"body": "Important",
					"type": "announcement",
					"announcement": {"expires_at": "2026-01-12T00:00:00Z"}
				}`,
			},
			{
				name: "report",
				jsonBody: `{
					"author_id": 1,
					"title": "Report",
					"body": "Report body",
					"type": "report",
					"report": {"status":"open","urgency":"high"}
				}`,
			},
			{
				name: "marketplace",
				jsonBody: `{
					"author_id":1,
					"title":"Marketplace",
					"body":"Selling item",
					"type":"marketplace",
					"marketplace":{"listing":"item","price":"10.0","quantity":5,"listing_status":"active"}
				}`,
			},
			{
				name: "openjio",
				jsonBody: `{
					"author_id":1,
					"title":"Openjio",
					"body":"Event",
					"type":"openjio",
					"openjio":{"activity_category":"sports","location":"park","event_date":"2026-01-15","start_time":"10:00:00","end_time":"12:00:00"}
				}`,
			},
		}

		for _, tt := range subtypes {
			t.Run(tt.name, func(t *testing.T) {
				req := httptest.NewRequest(http.MethodPost, "/posts", strings.NewReader(tt.jsonBody))
				rec := httptest.NewRecorder()

				handler.CreatePost(rec, req)

				var got int32
				if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
					t.Fatalf("failed to decode: %v", err)
				}

				want := int32(42)
				if got != want {
					t.Fatalf("expected post ID %d, got %d", want, got)
				}
			})
		}
	})
}
*/


