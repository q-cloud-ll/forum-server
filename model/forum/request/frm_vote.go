package request

type FrmVoteData struct {
	PostId    int64 `json:"post_id"`
	Direction int8  `json:"direction,string"`
}
