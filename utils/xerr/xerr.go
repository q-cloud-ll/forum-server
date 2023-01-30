package xerr

const (
	OK                            = "SUCCESS"
	SERVER_COMMON_ERROR           = "服务器开小差啦,稍后再来试一试"
	REUQEST_PARAM_ERROR           = "请求参数错误"
	TOKEN_EXPIRE_ERROR            = "token失效，请重新登陆"
	TOKEN_GENERATE_ERROR          = "生成token失败"
	DB_ERROR                      = "数据库繁忙,请稍后再试"
	DB_UPDATE_AFFECTED_ZERO_ERROR = "更新数据影响行数为0"
)
