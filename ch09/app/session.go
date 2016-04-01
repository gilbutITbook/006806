package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/goincremental/negroni-sessions"
)

const (
	currentUserKey  = "oauth2_current_user" // 세션에 저장되는 CurrentUser의 키
	sessionDuration = time.Hour             // 로그인 세션 유지 시간
)

type User struct {
	Uid       string    `bson:"uid" json:"uid"`
	Name      string    `bson:"name" json:"name"`
	Email     string    `bson:"email" json:"user"`
	AvatarUrl string    `bson:"avatar_url" json:"avatar_url"`
	Expired   time.Time `bson:"-" json:"expired"`
}

func (u *User) Valid() bool {
	// 현재 시간 기준으로, 만료시간 확인
	return u.Expired.Sub(time.Now()) > 0
}

func (u *User) Refresh() {
	// 만료시간 시간 연장
	u.Expired = time.Now().Add(sessionDuration)
}

func GetCurrentUser(r *http.Request) *User {
	// 세션으로부터 CurrentUser 정보를 가져옴
	s := sessions.GetSession(r)

	if s.Get(currentUserKey) == nil {
		return nil
	}

	data := s.Get(currentUserKey).([]byte)
	var u User
	json.Unmarshal(data, &u)
	return &u
}

func SetCurrentUser(r *http.Request, u *User) {
	if u != nil {
		// CurrentUser 만료 시간 갱신
		u.Refresh()
	}

	// 세션에 CurrentUser 정보를 json으로 저장
	s := sessions.GetSession(r)
	val, _ := json.Marshal(u)
	s.Set(currentUserKey, val)
}
