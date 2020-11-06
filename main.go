package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"pinpoint-go/common/trace"
	"pinpoint-go/config"
	"pinpoint-go/plugins/httpclient"
	"pinpoint-go/profiler"
	"time"
)



func main() {
	address := "59.110.238.199"
	conf := &config.Config{
		AgentID:         "192.168.99.10",
		ApplicationName: "test-go",
		Pinpoint: struct {
			InfoAddr string
			StatAddr string
			SpanAddr string
		}{InfoAddr: address + ":9994", StatAddr: address + ":9995", SpanAddr: address + ":9996"},
		ServerType:trace.StandAlone,
	}

	config.InitConfig(conf)
	go profiler.GAgent.Start()


	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	ghttp := httpclient.NewPPHttpClient(httpclient.WithHttpTimeOut(10 * time.Second))
	url := "http://www.baidu.com"
	request, _ := httpclient.NewRequest(context.Background(), http.MethodGet, url, nil)
	fmt.Println(request)

	resp, err := ghttp.Do(request)
	fmt.Println(resp)
	fmt.Println(err)

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Println(err)
	} else {
		p := make([]byte, 1204)
		io.ReadFull(resp.Body, p)
		fmt.Println(string(p))
	}

	fmt.Fprintf(w, "hello world")
}