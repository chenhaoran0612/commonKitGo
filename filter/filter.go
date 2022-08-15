// @Title filter
// @Description 请求统一的过滤器，加入RequestId 信息方便追踪使用
// @Author chenhaoran
// @Datetime  2021/7/19 11:46
package filter

import (
	"context"
	"github.com/apache/dubbo-go/common/extension"
	"github.com/apache/dubbo-go/filter"
	"github.com/apache/dubbo-go/protocol"
)

const (
	Request   = "request"
	RequestID = "requestID"
)

func init() {
	extension.SetFilter(Request, getRequestFilter)
}

type RequestFilter struct {
}

// Invoke ...
func (sf *RequestFilter) Invoke(ctx context.Context, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
	xid := invocation.AttachmentsByKey(RequestID, "")
	if xid != "" {
		return invoker.Invoke(context.WithValue(ctx, RequestID, xid), invocation)
	}
	return invoker.Invoke(ctx, invocation)
}

// OnResponse ...
func (sf *RequestFilter) OnResponse(ctx context.Context, result protocol.Result, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
	return result
}

// getRequestFilter ...
func getRequestFilter() filter.Filter {
	return &RequestFilter{}
}
