package tools

import (
	resty "github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"time"
)

var DEBUG = false

func GetRestyClient() *resty.Client {
	client := resty.New()
	if DEBUG {
		return client.SetDebug(DEBUG).EnableTrace()
	}

	return client
}

func Get(url string, header map[string]string) (resp *resty.Response, err error, traceInfo resty.TraceInfo) {
	client := GetRestyClient()
	request := client.R()
	if header != nil {
		request = request.SetHeaders(header)
	}
	resp, err = request.Get(url)
	traceInfo = request.TraceInfo()
	return
}

func PostJSONTimeout(url string, body interface{}, header map[string]string, timeout time.Duration) (resp *resty.Response, err error, traceInfo resty.TraceInfo) {

	client := GetRestyClient()
	client.SetTimeout(timeout)
	request := client.R()
	if body != nil {
		request = request.SetBody(body)
	}

	if header != nil {
		request = request.SetHeaders(header)
	}

	request = request.SetHeader("Content-Type", "application/json")

	resp, err = request.Post(url)
	traceInfo = request.TraceInfo()

	return
}

func PostJSON(url string, body interface{}, header map[string]string) (resp *resty.Response, err error, traceInfo resty.TraceInfo) {
	zap.L().Info("cilent start ...")
	client := GetRestyClient()
	request := client.R()
	if body != nil {
		zap.L().Info("set body start ...", zap.Any("body", body))
		request = request.SetBody(body)
	}

	if header != nil {
		zap.L().Info("set headers start ...")
		request = request.SetHeaders(header)
	}

	request = request.SetHeader("Content-Type", "application/json")

	resp, err = request.Post(url)
	if err != nil {
		zap.L().Info(err.Error())
		return nil, err, traceInfo
	}
	traceInfo = request.TraceInfo()

	if resp != nil {
		zap.L().Info("cilent end ...", zap.Any("resp", resp))
	}

	return
}

func PostForm(url string, body map[string]string, header map[string]string) (resp *resty.Response, err error, traceInfo resty.TraceInfo) {

	client := GetRestyClient()
	request := client.R()
	if body != nil {
		request = request.SetFormData(body)
	}

	if header != nil {
		request = request.SetHeaders(header)
	}
	// application/x-www-form-urlencoded / multipart/form-data
	request = request.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	resp, err = request.Post(url)
	traceInfo = request.TraceInfo()

	return
}

/**
 * post发送二进制格式
 */
func PostBin(url string, body []byte, header map[string]string) (resp *resty.Response, err error, traceInfo resty.TraceInfo) {

	client := GetRestyClient()
	request := client.R()
	if body != nil {
		request = request.SetBody(body)
	}

	if header != nil {
		request = request.SetHeaders(header)
	}

	request = request.SetContentLength(true)

	request = request.SetHeader("Content-Type", "application/macbinary")
	//request.URL = url
	resp, err = request.Post(url)
	traceInfo = request.TraceInfo()

	return
}
