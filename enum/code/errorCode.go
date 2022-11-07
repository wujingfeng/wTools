package code

// Success success
const (
	Success = 200
)

// redirect
const (
	Moved = 301
	Found = 302
)

// client error
const (
	// BadRequest 请求失败
	BadRequest = 400
	// Unauthorized 授权失败
	Unauthorized = 401
	// Forbidden 服务器拒绝
	Forbidden = 403
	// NotFound 资源地址错误
	NotFound = 404
	// MethodNotAllowed 客户端请求中的方法被禁止
	MethodNotAllowed = 405
)

// server error
const (
	// ServerError 服务器内部错误
	ServerError = 500
	// BadGateway 无效响应
	BadGateway = 502
	// GatewayTimeout 超时
	GatewayTimeout = 504
)
