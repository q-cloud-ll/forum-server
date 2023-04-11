package response

import "forum/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
