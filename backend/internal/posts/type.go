package posts

import "github.com/jackc/pgx/v5/pgtype"

type CreatePostRequest struct {
	AuthorID  	 int32            `json:"author_id"`
	Title        string              `json:"title"`
	Body         string              `json:"body"`
	Type         PostType          `json:"type"`
	Announcement *CreateAnnouncement `json:"announcement,omitempty"`
	Report       *CreateReport       `json:"report,omitempty"`
	Marketplace  *CreateMarketplace  `json:"marketplace,omitempty"`
	Openjio      *CreateOpenjio       `json:"openjio,omitempty"`
}


type UpdatePostCoreRequest struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
}


type PostType string

const (
	PostTypeAnnouncement PostType = "announcement"
	PostTypeReport       PostType = "report"
	PostTypeMarketplace  PostType = "marketplace"
	PostTypeOpenjio      PostType = "openjio"
	PostTypeNormal       PostType = "normal"
)

type CreateAnnouncement struct {
	PostID    int32            `json:"post_id"`
	ExpiresAt pgtype.Timestamp `json:"expires_at"`
}

type CreateReport struct {
	PostID    int32            `json:"post_id"`
	Status  string `json:"status"`
	Urgency string `json:"urgency"`
}

type CreateMarketplace struct {
	PostID    int32            `json:"post_id"`
	Listing       string
	Price         pgtype.Numeric
	Quantity      int32
	ListingStatus string
}

type CreateOpenjio struct {
	PostID    int32            `json:"post_id"`
	ActivityCategory string
	Location         string
	EventDate        pgtype.Date
	StartTime        pgtype.Time
	EndTime          pgtype.Time
}


