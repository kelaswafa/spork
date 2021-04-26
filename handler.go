package main

import (
	"fmt"
	"os"

	"github.com/valyala/fasthttp"
)

func cookieKey() string {
	return "password"
}

func IsLogin(ctx *fasthttp.RequestCtx, pw string) bool {
	password := ctx.Request.Header.Cookie(cookieKey())
	if string(password) == pw {
		return true
	}
	return false
}

func Login(ctx *fasthttp.RequestCtx, pw string) {
	password := ctx.QueryArgs().Peek(cookieKey())
	if pw == string(password) {
		fmt.Println("-------")
		fmt.Println("Login Success !")
		cook := fasthttp.Cookie{}
		cook.SetKey(cookieKey())
		cook.SetValue(string(password))
		cook.HTTPOnly()
		cook.SetMaxAge(3600000)
		cook.SetPath(("/"))
		ctx.Response.Header.SetCookie(&cook)
		ctx.Redirect("/", fasthttp.StatusTemporaryRedirect)
	}
}

func RequestHandler(ctx *fasthttp.RequestCtx) {
	fmt.Printf("%s - %s \n", ctx.Method(), ctx.Path())
	file := string(ctx.Path())[1:]
	info, err := os.Stat(file)
	if err == nil || os.IsExist(err) {
		if info.IsDir() {
			files, _ := os.ReadDir(file)
			ResponseList(ctx, files, file)
		} else {
			fasthttp.ServeFileUncompressed(ctx, file)
		}
		return
	}
	files, _ := os.ReadDir(".")
	ResponseList(ctx, files, "/")
}
