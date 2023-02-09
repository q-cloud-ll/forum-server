package forum

import (
	frmReq "forum-server/model/forum/request"

	uuid "github.com/satori/go.uuid"
)

type CommentService struct{}

func (cs *CommentService) FrmCreateComment(c *frmReq.Comment, uid uuid.UUID) (err error) {

	return
}
