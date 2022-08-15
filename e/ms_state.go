package e

const Ok = 0
const Error = 500

//Api返回代码
const (
	ApiOk           = 200 // 正常
	ApiErrData      = 300 // 处理业务失败 例如注册手机号 返回手机号已被注册
	ApiAuthRefresh  = 401 // token已失效请重新登录
	ApiAuthRefuse   = 402 // token被其他用户挤掉了
	ApiParameter    = 403 // 入参问题
	ApiPrivilege    = 404 // 权限问题
	ApiDoubleReqErr = 499 // 连续请求
	ApiErrServer    = 500 // 服务接口处理异常报错
)
