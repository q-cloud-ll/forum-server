package forum

import (
	"forum/global"
	"time"

	uuid "github.com/satori/go.uuid"
)

type FrmUser struct {
	global.GVA_MODEL
	UserId      uuid.UUID  `json:"user_id" gorm:"index;unique;not null;comment:用户user_id"`
	Birthday    *time.Time `json:"birthday"`
	Gender      int8       `json:"gender"   gorm:"size:1"`
	Type        int8       `json:"type"   gorm:"size:5;comment:是否wx登录"`
	Status      int        `json:"status"   gorm:"size:5"`
	Enable      int        `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
	CommentNum  uint32     `json:"comment_num"`
	ArticleNum  uint32     `json:"article_num"`
	LikeNum     int64      `json:"like_num"`
	Company     string     `json:"company"   gorm:"size:500"`
	WxOpenid    string     `json:"wx_openid"   gorm:"size:500"`
	RealName    string     `json:"real_name" gorm:"size:120"`
	Nickname    string     `json:"nickname" gorm:"size:120"`
	Username    string     `json:"username"  gorm:"size:120"`
	Password    string     `json:"-"  gorm:"size:120"`
	Mobile      string     `json:"mobile"  gorm:"size:120"`
	Email       string     `json:"email" gorm:"size:120"`
	Blog        string     `json:"facebook"   gorm:"size:3000"`
	Avatar      string     `json:"avatar" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"`
	Description string     `json:"description"  gorm:"default:Ta很懒，还没有添加简介"`
	Location    string     `json:"location"   gorm:"size:500"`
	School      string     `json:"school"   gorm:"size:500"`
}

type UserInfo struct {
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
}
