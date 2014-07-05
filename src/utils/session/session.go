package session

import (
	"net/http"
	"sync"
	"time"
	"utils/tools"
)

type Session struct {
	kv             map[string]interface{}
	lastAccessTime time.Time
}

func (this *Session) IsTimeout(expires time.Duration) bool {
	return this.lastAccessTime.Add(expires).Before(time.Now())
}

func (this *Session) UpdateAccessTime() {
	this.lastAccessTime = time.Now()
}

func (this *Session) Get(k string) interface{} {
	if _, ok := this.kv[k]; ok {
		return this.kv[k]
	}
	return nil
}
func (this *Session) Set(k string, v interface{}) {
	this.kv[k] = v
}

func (this *Session) Delete(k string) {
	delete(this.kv, k)
}

type Sessions struct {
	cookieName string
	expires    time.Duration
	lock       sync.Mutex
	sessions   map[string]*Session
}

//
func NewSessions(cookieName string, expires time.Duration) *Sessions {
	return &Sessions{
		cookieName: cookieName,
		expires:    expires,
		sessions:   make(map[string]*Session),
	}
}

// 获取请求中sessionid对应的session，如果不存在或超时，则新建一个并写到响应流
func (this *Sessions) Prepare(w http.ResponseWriter, r *http.Request) *Session {
	cookie, err := r.Cookie(this.cookieName)
	if err != nil {
		sessionId := tools.GetUUID()
		cookie = &http.Cookie{
			Name:  this.cookieName,
			Value: sessionId,
			Path:  "/",
			//HttpOnly: true,
		}
		http.SetCookie(w, cookie)
		r.AddCookie(cookie)

		this.set(sessionId)
		return this.get(sessionId)
	}
	session := this.get(cookie.Value)
	if session == nil {
		sessionId := tools.GetUUID()
		cookie = &http.Cookie{
			Name:  this.cookieName,
			Value: sessionId,
			Path:  "/",
			//HttpOnly: true,
		}
		http.SetCookie(w, cookie)
		r.AddCookie(cookie)
		this.set(sessionId)
		return this.get(sessionId)
	}
	return session
}

// 定期清除超时的session
func (this *Sessions) GC() {
	this.lock.Lock()
	defer this.lock.Unlock()
	for k, v := range this.sessions {
		if v.lastAccessTime.Add(this.expires).Before(time.Now()) {
			delete(this.sessions, k)
		}
	}
	time.AfterFunc(this.expires, func() { this.GC() })
}

func (this *Sessions) get(sessionId string) *Session {
	this.lock.Lock()
	defer this.lock.Unlock()
	if _, ok := this.sessions[sessionId]; ok {
		if !this.sessions[sessionId].IsTimeout(this.expires) {
			this.sessions[sessionId].UpdateAccessTime()
			return this.sessions[sessionId]
		}
	}
	return nil
}

func (this *Sessions) set(sessionId string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.sessions[sessionId] = &Session{
		kv:             make(map[string]interface{}),
		lastAccessTime: time.Now(),
	}
}
