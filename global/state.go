package global

const (
	STATE_ENABLE = 0  // 状态可见
	STATE_DELETE = 1  // 状态为已删除
)

const (
	EVENT_STATE_SENDED = iota
	EVENT_STATE_EXECUTED
	EVENT_STATE_SUCCEED
	EVENT_STATE_FAILED
)

const (
	OFF = 0
	ON = 1
)