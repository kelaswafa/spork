package main

import (
	"fmt"
	"io/fs"
	"path"
	"strings"

	"github.com/valyala/fasthttp"
)

func layout(body string) string {
	return fmt.Sprintf(`
	<html>
	<head>
		<title>Login</title>
	</head>
	<body>
	%s
	</body>
	</html>
	`, body)
}

func LoginPage() string {
	return layout(`
		<form action="/" method="GET">
			<input type="text" name="password" />
			<input type="submit" value="login" />
		</form>
	`)
}

func ResponseList(ctx *fasthttp.RequestCtx, s []fs.DirEntry, base string) {
	ctx.SetContentType("text/html")
	var links []string
	for _, file := range s {
		link := fmt.Sprintf("<a href='%s'>%s</a>", path.Join(base, file.Name()), file.Name())
		links = append(links, link)
	}

	ctx.WriteString(layout(strings.Join(links, "<br/>")))
}
