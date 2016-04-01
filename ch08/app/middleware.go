package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
	"strings"
	"time"
)

type Middleware func(next HandlerFunc) HandlerFunc

func logHandler(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		// next(c)를 실행하기 전에 tiemstamp 가록
		t := time.Now()

		// 다음 핸들러 수행
		next(c)

		// 웹 요청 정보와 전체 소요시간을 로그로 남김
		log.Printf("[%s] %q %v\n",
			c.Request.Method,
			c.Request.URL.String(),
			time.Now().Sub(t))
	}
}

func recoverHandler(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(c.ResponseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next(c)
	}
}

func parseFormHandler(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		c.Request.ParseForm()
		for k, v := range c.Request.PostForm {
			if len(v) > 0 {
				c.Params[k] = v[0]
			}
		}
		next(c)
	}
}

func parseJsonBodyHandler(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		var m map[string]interface{}
		if json.NewDecoder(c.Request.Body).Decode(&m); len(m) > 0 {
			for k, v := range m {
				c.Params[k] = v
			}
		}
		next(c)
	}
}

func staticHandler(next HandlerFunc) HandlerFunc {
	var (
		dir       = http.Dir(".")
		indexFile = "index.html"
	)
	return func(c *Context) {
		// HTTP Method가 GET이나 HEAD가 아닌 경우 바로 다음 핸들러 수행
		if c.Request.Method != "GET" && c.Request.Method != "HEAD" {
			next(c)
			return
		}

		file := c.Request.URL.Path
		// URL 경로에 해당하는 파일 open 시도
		f, err := dir.Open(file)
		if err != nil {
			// URL 경로에 해당하는 파일 open이 실패하면 바로 다음 한들러 수행
			next(c)
			return
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			// 파일의 상태가 정상인 아니면 바로 다음 한들러 수행
			next(c)
			return
		}

		// URL 경로가 디렉토리인 경우 indexFile을 사용
		if fi.IsDir() {
			// 디렉토리 경로를 URL로 사용하는 경우, 경로의 마지막에 "/"를 붙혀야 함
			if !strings.HasSuffix(c.Request.URL.Path, "/") {
				http.Redirect(c.ResponseWriter, c.Request, c.Request.URL.Path+"/", http.StatusFound)
				return
			}

			// 디렉토리를 가르키는 URL 경로에 indexFile명을 붙혀 전체 파일 경로 생성
			file = path.Join(file, indexFile)

			// indexFile open 시도
			f, err = dir.Open(file)
			if err != nil {
				next(c)
				return
			}
			defer f.Close()

			fi, err = f.Stat()
			if err != nil || fi.IsDir() {
				// indexFile 상태가 정상이 아니면 바로 다음 핸들러 수행
				next(c)
				return
			}
		}

		// file의 내용 전달(next 핸들러로 제어권을 넘기지 않고 요청 처리를 종료함)
		http.ServeContent(c.ResponseWriter, c.Request, file, fi.ModTime(), f)
	}
}
