package main

import (
	"fmt"
	"log"
	"time"

	"github.com/valyala/fasthttp"
)

func main() {
	port := 3000
	t := time.Now().Unix() % 10000
	pw := fmt.Sprintf("%04s", fmt.Sprint(t))
	fmt.Printf("Password: %s \n", pw)

	IPs := GetIP()
	for _, ip := range IPs {
		fmt.Printf("Link: http://%s:%s\n", ip, fmt.Sprint(port))
	}

	err := fasthttp.ListenAndServe(":3000", func(ctx *fasthttp.RequestCtx) {
		if IsLogin(ctx, pw) {
			RequestHandler(ctx)
			return
		}
		ctx.Response.Header.SetContentType("text/html")
		ctx.WriteString(LoginPage())
		Login(ctx, pw)
	})
	if err != nil {
		log.Fatal(err)
	}
}
