package account_sgx

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

// test方便测试 (后期通过配置文件/填写参数形式)
var (
	// sgx服务地址
	URL    = "http://127.0.0.1:8080"
	Ping   = "/ping"
	Create = "/create"
	Sign   = "/sign"
	Verify = "/verify"

	PingMethod   = "GET"
	CreateMethod = "GET"
	SignMethod   = "POST"
	VerifyMethod = "POST"
)

// 响应结果
type Response struct {
	Code int    `json:"code"` // 错误码
	Msg  string `json:"msg"`  // 信息提示
	Data []byte `json:"data"` // 返回数据
}

// 通过接口提供相关服务
type ApiClient interface {
	Ping(method string, args map[string]interface{}) (*Response, error)
	Create(method string, args map[string]interface{}) (*Response, error)
	Sign(method string, args map[string]interface{}) (*Response, error)
	Verify(method string, args map[string]interface{}) (*Response, error)
}

// 封装请求api服务的结构体
type ApiClientXuperchain struct {
	URL string // 服务地址
}

func NewApiClientXuperchain(url string) ApiClient {
	return &ApiClientXuperchain{
		URL: url,
	}
}

func (c *ApiClientXuperchain) Ping(method string, args map[string]interface{}) (*Response, error) {
	return request(c.URL+Ping, method, args)
}

func (c *ApiClientXuperchain) Create(method string, args map[string]interface{}) (*Response, error) {
	return request(c.URL+Create, method, args)
}

func (c *ApiClientXuperchain) Sign(method string, args map[string]interface{}) (*Response, error) {
	return request(c.URL+Sign, method, args)
}

func (c *ApiClientXuperchain) Verify(method string, args map[string]interface{}) (*Response, error) {
	return request(c.URL+Verify, method, args)
}

// 请求服务封装
func request(url, method string, args map[string]interface{}) (*Response, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源
	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod(method)

	req.SetRequestURI(url)

	// 请求体
	if args != nil {
		requestBody, err := json.Marshal(&args)
		if err != nil {
			return nil, err
		}
		req.SetBody(requestBody)
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp) // 用完需要释放资源

	if err := fasthttp.Do(req, resp); err != nil {
		return nil, err
	}

	// 处理返回结果
	res := &Response{}
	err := json.Unmarshal(resp.Body(), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
