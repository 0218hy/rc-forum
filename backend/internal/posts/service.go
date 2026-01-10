package posts

import (
	"context"
	"errors"
	"fmt"
	repo "rc-forum-backend/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service interface {
	GetAllPosts(ctx context.Context) ([]repo.ListAllPostsRow, error)
	GetPostByID(ctx context.Context, postID int32) (repo.FindPostByIDRow, error)
	DeletePostByID(ctx context.Context, postID int32) error
	CreatePost(ctx context.Context, req CreatePostRequest) (int32, error)
	UpdatePostCore(ctx context.Context, postID int32, authorID int32, title string, body string) error
}

type svc struct {
	repo *repo.Queries
	db   *pgxpool.Pool
}

func NewService(repo *repo.Queries, db *pgxpool.Pool) Service {
	return &svc{
		repo: repo,
		db:   db,
	}
}

func (s *svc) GetAllPosts(ctx context.Context) ([]repo.ListAllPostsRow, error) {
	return s.repo.ListAllPosts(ctx)
}

func (s *svc) GetPostByID(ctx context.Context, postID int32) (repo.FindPostByIDRow, error) {
	return s.repo.FindPostByID(ctx, postID)
}

func (s *svc) DeletePostByID(ctx context.Context, postID int32) error {
	return s.repo.DeletePostByID(ctx, postID)
}

func (s *svc) CreatePost(ctx context.Context, req CreatePostRequest) (int32, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	qtx := s.repo.WithTx(tx)

	// create main post
	post, err := qtx.CreatePost(ctx, repo.CreatePostParams{
		AuthorID: req.AuthorID,
		Type:     repo.PostType(req.Type),
		Title:    req.Title,
		Body:     req.Body,
	})
	if err != nil {
		return 0, err
	}

	// create subtype post
	switch req.Type {
	case PostTypeAnnouncement:
		if req.Announcement == nil {
			return 0, fmt.Errorf("announcement data required")
		}
		err = qtx.CreateAnnouncement(ctx, repo.CreateAnnouncementParams{
			PostID:    post.ID,
			ExpiresAt: req.Announcement.ExpiresAt,
		})
	case PostTypeReport:
		if req.Report == nil {
			return 0, errors.New("report data required")
		}
		err = qtx.CreateReport(ctx, repo.CreateReportParams{
			PostID:  post.ID,
			Status:  repo.ReportStatus(req.Report.Status),
			Urgency: repo.UrgencyLevel(req.Report.Urgency),
		})
	case PostTypeMarketplace:
		if req.Marketplace == nil {
			return 0, errors.New("marketplace data required")
		}
		err = qtx.CreateMarketplace(ctx, repo.CreateMarketplaceParams{
			PostID:        post.ID,
			Listing:       repo.ListingType(req.Marketplace.Listing),
			Price:         req.Marketplace.Price,
			ListingStatus: repo.ListingStatusType(req.Marketplace.ListingStatus),
		})
	case PostTypeOpenjio:
		if req.Openjio == nil {
			return 0, errors.New("openjio data required")
		}
		err = qtx.CreateOpenjio(ctx, repo.CreateOpenjioParams{
			PostID:           post.ID,
			ActivityCategory: repo.ActivityCategoryType(req.Openjio.ActivityCategory),
			Location:         req.Openjio.Location,
			EventDate:        req.Openjio.EventDate,
			StartTime:        req.Openjio.StartTime,
			EndTime:          req.Openjio.EndTime,
		})
	}

	tx.Commit(ctx)

	return post.ID, nil
}

// MVP: only allow author to update title and body (future work: update subtypes)
func (s *svc) UpdatePostCore (ctx context.Context, postID int32, authorID int32, title string, body string) error {
	// 1. get existing post
	existingPost, err := s.repo.FindPostByID(ctx, postID)
	if err != nil {
		return err
	}

	// 2. check authorization
	if existingPost.AuthorID != authorID {
		return errors.New("unauthorized: only author can update the post")
	}

	return s.repo.UpdatePostCore(ctx, repo.UpdatePostCoreParams{
		ID:    postID,
		Title: title,
		Body:  body,
	})

}
