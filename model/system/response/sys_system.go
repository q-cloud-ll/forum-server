package response

import "forum-server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
