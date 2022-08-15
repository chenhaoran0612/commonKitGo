package e

//User模块类型
const (
	UserTypeClient = 10 // 用户属于客户端用户
	UserTypeAdmin  = 20 // 用户属于管理端用户
)

// 客户端平台类型
const (
	//Mobile ...
	Mobile = "mobile"
	//Web ...
	Web = "web"
)

// 多语言Enum
const (
	LOCALE_EN = "en"
	LOCALE_CN = "cn"
)

// Context 相关定义
const (
	CtxUserId   = "user_id"
	CtxPlatform = "platform"
	CtxLocale   = "locale"
)

//HTTPContextKey ...
type HTTPContextKey int64

//HTTPContext缓存key值
const (
	HTTPContextUserID HTTPContextKey = iota
	HTTPContextUser
	HTTPContextLocale
	HTTPContextPlatform
)

func (k HTTPContextKey) String() string {
	switch k {
	case HTTPContextUserID:
		return "userId"
	case HTTPContextUser:
		return "user"
	case HTTPContextLocale:
		return "locale"
	case HTTPContextPlatform:
		return "platform"
	default:
		return ""
	}
}
