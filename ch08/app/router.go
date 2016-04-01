package main

import (
	"net/http"
	"strings"
)

type router struct {
	// key: HTTP Method
	// value: URL Pattern별로 실행 할 HandlerFunc
	handlers map[string]map[string]HandlerFunc
}

func (r *router) HandleFunc(method, pattern string, h HandlerFunc) {
	// method로 등록된 맵이 있는지 확인
	m, ok := r.handlers[method]
	if !ok {
		// 등록된 맵이 없는 경우 새로운 맵 생성
		m = make(map[string]HandlerFunc)
		r.handlers[method] = m
	}
	// method르 등록된 맵에 URL pattern과 핸들러 함수 등록
	m[pattern] = h
}

func (r *router) handler() HandlerFunc {
	return func(c *Context) {
		// HTTP Method에 맞는 모든 handers를 반복하면서 요청 url에 해당하는 handler를 찾음
		for pattern, handler := range r.handlers[c.Request.Method] {
			if ok, params := match(pattern, c.Request.URL.Path); ok {
				for k, v := range params {
					c.Params[k] = v
				}
				// 요청 url에 해당하는 handler 수행
				handler(c)
				return
			}
		}

		// 요청 url에 해당하는 handler를 찾지 못한 경우 NotFound 에러 처리
		http.NotFound(c.ResponseWriter, c.Request)
		return
	}
}

func match(pattern, path string) (bool, map[string]string) {
	// 패턴과 패쓰가 정확히 일치하는 경우 바로 true를 리턴
	if pattern == path {
		return true, nil
	}

	// 패턴과 패쓰를 '/' 단위로 구분
	patterns := strings.Split(pattern, "/")
	paths := strings.Split(path, "/")

	// 패턴과 패쓰를 '/'로 구분한 후 부분 문자열 집합의 갯수가 다르면 false를 리턴
	if len(patterns) != len(paths) {
		return false, nil
	}

	// 패턴에 일치하는 url 파라미터를 담기 위한 params 맵 생성
	params := make(map[string]string)

	// '/'로 구분된 패턴/패쓰의 각 문자열을 하나씩 비교
	for i := 0; i < len(patterns); i++ {
		switch {
		case patterns[i] == paths[i]:
			// 패턴과 패쓰의 부분 문자열이 일치하는 경우
			// => 바로 다음 루프 수행
		case len(patterns[i]) > 0 && patterns[i][0] == ':':
			// 패턴이 ':' 문자로 시작하는 경우
			// => params에 url params를 담은 후 다음 루프 수행
			params[patterns[i][1:]] = paths[i]
		default:
			// 일치하는 경우가 없으면 false를 리턴
			return false, nil
		}
	}

	// true와 params를 리턴
	return true, params
}
