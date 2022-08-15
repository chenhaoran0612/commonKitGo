package base

import (
	_ "github.com/apache/dubbo-go-hessian2"
)

type Response struct {
	Code int
}

func (r Response) JavaClassName() string {
	return "com.ikurento.user.Response"
}
