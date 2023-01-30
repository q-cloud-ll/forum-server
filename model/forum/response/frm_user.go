package response

import (
	"forum-server/model/forum"
)

type LoginResponse struct {
	User      forum.FrmUser `json:"user"`
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expiresAt"`
}
