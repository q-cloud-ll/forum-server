package response

import (
	"forum/model/forum"
)

type LoginResponse struct {
	User      forum.FrmUser `json:"user"`
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expiresAt"`
}
