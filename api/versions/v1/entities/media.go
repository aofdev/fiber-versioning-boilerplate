package entities

import (
	"time"
)

type metaPost struct {
	CommentCount int    `json:"comment_count" bson:"comment_count" example:"100"` // Number of Comments of this media.
	LikeCount    int    `json:"like_count" bson:"like_count" example:"2000"`      // Number of Likes of this media.
	ShareCount   int    `json:"share_count" bson:"share_count" example:"100"`     // Number of Shares of this media.
	Type         string `json:"type" bson:"type" example:"post"`                  // Indicates whether this media represents a type post.
}

type metaComment struct {
	ReplyCommentCount int    `json:"reply_comment_count" bson:"reply_comment_count" example:"1000"`           // Number of Reply Comments of this media.
	LikeCount         int    `json:"like_count" bson:"like_count" example:"1000"`                             // Number of Likes of this media.
	ParentPostID      string `json:"parent_post_id" bson:"parent_post_id" example:"61605ba183de6c07967b102d"` // The parent post id identifier string.
	Type              string `json:"type" bson:"type" example:"comment"`                                      // Indicates whether this media represents a type comment.
}

type metaReplyComment struct {
	LikeCount       int    `json:"like_count" bson:"like_count" example:"1000"`                                   // Number of Likes of this media.
	ParentCommentID string `json:"parent_comment_id" bson:"parent_comment_id" example:"61605ba183de6c07967b102d"` // The parent comment id identifier string.
	ParentPostID    string `json:"parent_post_id" bson:"parent_post_id" example:"61605ba183de6c07967b102d"`       // The parent post id identifier string.
	Type            string `json:"type" bson:"type" example:"reply-comment"`                                      // Indicates whether this media represents a type reply-comment.
}

// templateMedia defines the fields which should be the template struct of Media.
type templateMedia struct {
	ID          string    `json:"_id,omitempty" bson:"_id" example:"61605ba183de6c07967b102d"`     // The identifier string of the account in the database.
	Author      string    `json:"author" bson:"author" example:"aof"`                              // The user identifier string.
	Permalink   string    `json:"permalink" bson:"permalink" example:"https://example.com/222"`    // The URL in the format Media by the user.
	PostTime    time.Time `json:"post_time" bson:"post_time" example:"2021-02-14T09:20:32.000Z"`   // Creation time of the Media.
	Text        string    `json:"text" bson:"text" example:"hello world"`                          // The content of the Media.
	DataVersion string    `json:"data_version" bson:"data_version" example:"v1.0"`                 // The data version identifier string of the post in the database.
	CreatedAt   time.Time `json:"created_at" bson:"created_at" example:"2021-01-27T05:44:15.337Z"` // Last this document was created.
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at" example:"2021-01-28T05:44:15.337Z"` // Last this document was updated.
}

// MediaType defines the field which should be checked type Media.
type MediaType struct {
	Meta struct {
		Type string `json:"type" bson:"type"`
	} `json:"meta" bson:"meta"`
}

// Post struct contains the media post field which should be returned in a response.
type Post struct {
	templateMedia `bson:",inline"`
	Meta          metaPost `json:"meta" bson:"meta"`
}

// Comment struct contains the media comment field which should be returned in a response.
type Comment struct {
	templateMedia `bson:",inline"`
	Meta          metaComment `json:"meta" bson:"meta"`
}

// ReplyComment struct contains the media reply-comment field which should be returned in a response.
type ReplyComment struct {
	templateMedia `bson:",inline"`
	Meta          metaReplyComment `json:"meta" bson:"meta"`
}

// RequestFindByListID defines the field which should be the request payload.
type RequestFindByListID struct {
	IDS []string `json:"ids" validate:"required,min=1,max=100" minLength:"1" maxLength:"100" example:"6929046813118876930,6929047870637441025"` // Request body `IDS` are array
}
