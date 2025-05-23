package commentquery

import (
	"github.com/gqlgen/_examples/mini-habr-with-subscriptions/internal/model"
)

type CommentQueryImp interface {
	GetCommentsBranch(postID int64, path string) ([]*model.Comment, error)
	GetCommentPath(parentID int64) (string, error)
}
