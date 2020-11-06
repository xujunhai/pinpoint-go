package httpclient

import (
	context2 "context"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"pinpoint-go/common/trace"
	"pinpoint-go/profiler"
	"pinpoint-go/profiler/context"
	"time"
)

const (
	PpCtx = "PinpointContext"
)

type ppHttpConfig struct {
	TimeOut               time.Duration
	MaxIdleConnPerHost    int
	IdleConnTimeout       time.Duration
	TLSHandshakeTimeout   time.Duration
	ExpectContinueTimeout time.Duration
	DialTimeout           time.Duration
	DialKeepAlive         time.Duration
	TlsConfig             *tls.Config
}

//设置默认配置
func (p *ppHttpConfig) setDefault() {
	p.TimeOut = 0
	p.MaxIdleConnPerHost = 500
	p.IdleConnTimeout = 90 * time.Second
	p.TLSHandshakeTimeout = 10 * time.Second
	p.ExpectContinueTimeout = 1 * time.Second
	p.DialTimeout = 30 * time.Second
	p.DialKeepAlive = 30 * time.Second
}

type httpConfigure func(config *ppHttpConfig)

//请求超时设置
func WithHttpTimeOut(timeOut time.Duration) httpConfigure {
	return func(config *ppHttpConfig) {
		config.TimeOut = timeOut
	}
}

//每个host最大空闲链接
func WithHttpMaxIdleConnsPerHost(count int) httpConfigure {
	return func(config *ppHttpConfig) {
		config.MaxIdleConnPerHost = count
	}
}

//空闲链接多少释放
func WithHttpIdleConnTimeout(timeout time.Duration) httpConfigure {
	return func(config *ppHttpConfig) {
		config.IdleConnTimeout = timeout
	}
}

//TLS握手超时
func WithHttpTLSHandshakeTimeout(timeout time.Duration) httpConfigure {
	return func(config *ppHttpConfig) {
		config.TLSHandshakeTimeout = timeout
	}
}

func WithHttpExpectContinueTimeout(timeout time.Duration) httpConfigure {
	return func(config *ppHttpConfig) {
		config.ExpectContinueTimeout = timeout
	}
}

func WithHttpDialTimeout(timeout time.Duration) httpConfigure {
	return func(config *ppHttpConfig) {
		config.DialTimeout = timeout
	}
}

func WithHttpDialKeepAlive(timeout time.Duration) httpConfigure {
	return func(config *ppHttpConfig) {
		config.DialKeepAlive = timeout
	}
}

func WithHttpTlsClientConfig(tlsConfig *tls.Config) httpConfigure {
	return func(config *ppHttpConfig) {
		config.TlsConfig = tlsConfig
	}
}

var DefaultTransport http.RoundTripper = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).DialContext,
	MaxIdleConnsPerHost:   200,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

type Transport struct {
	http.RoundTripper
}

//http调用通过ppCtx传输 trace.Context
func (p *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	rt := p.RoundTripper
	if rt == nil {
		rt = DefaultTransport
	}

	//获取trace上下文
	traceContext, _ := req.Context().Value(PpCtx).(*context.Context)

	//说明是个根请求,尝试创建一个trace
	if traceContext == nil {
		traceContext = profiler.StartTrace(nil)
		defer profiler.FinishTrace(traceContext)

		//说明不采集
		if traceContext == nil {
			return rt.RoundTrip(req)
		}
		traceContext.Span.SetAPIID(profiler.GAgent.GetApiID("http." + req.Method))
		traceContext.Span.SetRpc(req.URL.Path)
	}

	spanEvent := traceContext.StartTraceSpanEvent()
	spanEvent.SetApiID(profiler.GAgent.GetApiID("http." + req.Method))
	spanEvent.SetServiceType(trace.GoMethod)
	spanEvent.AddAnnotation(context.NewStringAnnotation(trace.AkArgs0, req.URL.Path))

	traceHead := traceContext.GetNextSpanInfo()
	for k, v := range traceHead {
		req.Header.Set(k, v)
	}
	resp, err := rt.RoundTrip(req)

	if resp != nil {
		spanEvent.AddAnnotation(context.NewInt32Annotation(trace.AkHttpStatusCode, int32(resp.StatusCode)))
	}

	if err != nil {
		spanEvent.SetExceptionInfo(profiler.GAgent.GetStrID(profiler.StrError), err.Error())
	}
	spanEvent.Finish()
	return resp, err
}

func NewRequest(ctx interface{}, method, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	cx, _ := ctx.(*context.Context)

	//echoCtx, _ := ctx.(echo.Context)
	//if cx == nil && echoCtx != nil {
	//	cx, _ = echoCtx.Get(agent.PP_CTX).(*agent.TraceContext)
	//}

	if cx == nil {
		return request, nil
	}

	parentCtx := request.Context()
	request = request.WithContext(context2.WithValue(parentCtx, PpCtx, cx))
	return request, err
}

func NewPPHttpClient(configures ...httpConfigure) *http.Client {
	httpConfig := new(ppHttpConfig)
	httpConfig.setDefault()

	//更新配置
	for _, configure := range configures {
		configure(httpConfig)
	}

	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   httpConfig.DialTimeout,
			KeepAlive: httpConfig.DialKeepAlive,
		}).DialContext,
		MaxIdleConnsPerHost:   httpConfig.MaxIdleConnPerHost,
		IdleConnTimeout:       httpConfig.IdleConnTimeout,
		TLSHandshakeTimeout:   httpConfig.TLSHandshakeTimeout,
		ExpectContinueTimeout: httpConfig.ExpectContinueTimeout,
		TLSClientConfig:       httpConfig.TlsConfig,
	}

	return &http.Client{Transport: &Transport{RoundTripper: transport}, Timeout: httpConfig.TimeOut}
}
