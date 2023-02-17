package request

type Register struct {
	Mobile     string `json:"mobile" binding:"required"`
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
