package xerr

/*
 --------------通用业务(1xxx)----------------
*/

const (
	OK                            = "SUCCESS"
	SERVER_COMMON_ERROR           = "服务器开小差啦,稍后再来试一试"
	REUQEST_PARAM_ERROR           = "请求参数错误"
	TOKEN_EXPIRE_ERROR            = "token失效，请重新登陆"
	TOKEN_GENERATE_ERROR          = "生成token失败"
	DB_ERROR                      = "数据库繁忙,请稍后再试"
	DB_UPDATE_AFFECTED_ZERO_ERROR = "更新数据影响行数为0"
)

/*
	第三方相关(3xxx)
*/
const (
	/*
	   3001-3020 微信公众号
	*/
	// 微信公众号JSSDK获取access_token失败
	CodeWxGzhAccessTokenFail = "微信公众号JSSDK获取access_token失败"
	// 微信公众号JSSDK获取jsapi_ticket失败
	CodeWxGzhJsApiTicketFail = "微信公众号JSSDK获取jsapi_ticket失败"
	// 微信公众号JSSDK获取SIGN失败
	CodeWxGzhSignFail = "微信公众号JSSDK获取SIGN失败"
	// 微信wxCode为空
	CodeWxEmpty = "微信wxCode为空"
	// 微信wxCode失效或不正确请重新获取
	CodeWxOutTime = "微信wxCode失效或不正确请重新获取"
	// 微信生成二维码失败
	CodeWxTicketFail = "微信生成二维码失败"
)
