package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

//。包httptest 可以让你能够模仿互联网资源的请求和响应
var mockResp = `<rss>
	<channel>
		<title>Go (Golang) Programming Blog - Ardan Labs on</title>
		<link>https://www.ardanlabs.com/blog/</link>
		<description>
			Recent content in Go (Golang) Programming Blog - Ardan Labs on
		</description>
		<generator>Hugo -- gohugo.io</generator>
		<language>en-us</language>
		<lastBuildDate>Thu, 10 Oct 2019 00:00:00 +0000</lastBuildDate>
		<atom:link href="https://www.ardanlabs.com/blog/index.xml" rel="self" type="application/rss+xml"/>
		<item>
			<title>Modules Part 01: Why And What</title>
			<link>
				https://www.ardanlabs.com/blog/2019/10/modules-01-why-and-what.html
			</link>
			<pubDate>Thu, 10 Oct 2019 00:00:00 +0000</pubDate>
			<guid>
				https://www.ardanlabs.com/blog/2019/10/modules-01-why-and-what.html
			</guid>
			<description>
				Introduction Modules provide an integrated solution for three key problems that have been a pain point for developers since Go’s initial release: Ability to work with Go code outside of the GOPATH workspace. Ability to version a dependency and identify the most compatible version to use. Ability to manage dependencies natively using the Go tooling. With the release of Go 1.13, these three problems are a thing of the past.
			</description>
		</item>
	</channel>
</rss>
`

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, mockResp)
	}
	//HandlerFunc 类型是一个适配器，允许常规函数作为HTTP 的处理函数使用。如果函数f 具有合适的签名，
	//HandlerFunc(f)就是一个处理HTTP 请求的Handler 对象，内部通过调用f 处理请求
	return httptest.NewServer(http.HandlerFunc(f))
}
