package session

import (
	"github.com/pkg/errors"
	"mimi/djq/db/redis"
	"mimi/djq/util"
	"net/http"
	"strconv"
	"time"
)

const (
	SessionNameMiAdminId = "miAdminId"
	SessionNameMiPermission = "miPermission"
	SessionNameMiAdminName = "miAdminName"

	SessionNameUiUserId = "uiUserId"
	SessionNameUiUserMobile = "uiUserMobile"
	SessionNameUiUserOpenId = "uiUserOpenId"
	SessionNameUiUserCaptcha = "uiUserCaptcha"
	SessionNameUiUserLoginCount = "uiUserLoginCount"

	SessionNameSiShopAccountId = "siShopAccountId"
	SessionNameSiShopAccountName = "siShopAccountName"
	SessionNameSiShopAccountOpenId = "siShopAccountOpenId"
)

func GetMi(w http.ResponseWriter, r *http.Request) (*Session, error) {
	return Get(w, r, "mi", 0, time.Hour * 1)
}

func GetUi(w http.ResponseWriter, r *http.Request) (*Session, error) {
	return Get(w, r, "ui", 60 * 60 * 24 * 365, time.Hour * 24 * 365)
}

func GetSi(w http.ResponseWriter, r *http.Request) (*Session, error) {
	return Get(w, r, "si", 0, time.Hour * 18)
}

func GetOpen(w http.ResponseWriter, r *http.Request) (*Session, error) {
	return Get(w, r, "ui", 0, time.Hour * 1)
}

func Get(w http.ResponseWriter, r *http.Request, sessionType string, cookieMaxAge int, redisExpires time.Duration) (*Session, error) {
	cookieName := sessionType + "SessionId"
	session := &Session{}
	if cookie, err := r.Cookie(cookieName); err != nil && err != http.ErrNoCookie {
		return nil, errors.Wrap(err, "获取cookie失败")
	} else if cookie == nil || cookie.Value == "" {
		session.Id = util.BuildUUID()
	} else {
		session.Id = cookie.Value
	}
	cookie := http.Cookie{Name: cookieName, Value: session.Id, Path: "/", MaxAge: cookieMaxAge}
	http.SetCookie(w, &cookie)

	session.CookieMaxAge = cookieMaxAge
	session.RedisExpires = redisExpires
	session.w = w
	session.r = r
	session.sessionType = sessionType
	session.cookieName = cookieName
	return session, nil
}

type Session struct {
	Id           string
	CookieMaxAge int
	RedisExpires time.Duration
	w            http.ResponseWriter
	r            *http.Request
	sessionType  string
	cookieName   string
}

func (sn *Session) Get(key string) (string, error) {
	conn := redis.Get()
	if exist, err := conn.HExists(sn.getKey(), key).Result(); err != nil {
		return "", err
	} else if !exist {
		return "", nil
	}
	if error := conn.Expire(sn.getKey(), sn.RedisExpires).Err(); error != nil {
		return "", error
	}
	return conn.HGet(sn.getKey(), key).Result()
}

func (sn *Session) GetBool(key string) (bool, error) {
	if result, err := sn.Get(key); err != nil || "" == result {
		return false, err
	} else {
		return strconv.ParseBool(result)
	}
}

func (sn *Session) Set(key string, value string) error {
	conn := redis.Get()
	var vMap map[string]interface{} = make(map[string]interface{})

	vMap[key] = value

	if error := conn.HMSet(sn.getKey(), vMap).Err(); error != nil {
		return error
	}
	if error := conn.Expire(sn.getKey(), sn.RedisExpires).Err(); error != nil {
		return error
	}
	return nil
}

func (sn *Session) SetTemp(key string, value string,expires time.Duration) error {
	conn := redis.Get()
	var vMap map[string]interface{} = make(map[string]interface{})

	vMap[key] = value

	if error := conn.HMSet(sn.getKey(), vMap).Err(); error != nil {
		return error
	}
	if error := conn.Expire(sn.getKey(), expires).Err(); error != nil {
		return error
	}
	return nil
}


func (sn *Session) Del() error {
	conn := redis.Get()
	if err := conn.Del(sn.getKey()).Err(); err != nil {
		return err
	}
	cookie := http.Cookie{Name: sn.cookieName, Value: sn.Id, Path: "/", MaxAge: 0}
	http.SetCookie(sn.w, &cookie)
	return nil
}

func (sn *Session) getKey() string {
	return sn.sessionType + ":" + sn.Id
}

//import (
//	"time"
//	"io"
//	"strconv"
//	"encoding/base64"
//	"net/http"
//	"net/url"
//	"sync"
//	"crypto/rand"
//)
//
///*Session会话管理*/
//type SessionMgr struct {
//	mCookieName  string       //客户端cookie名称
//	mLock        sync.RWMutex //互斥(保证线程安全)
//	mMaxLifeTime int64        //垃圾回收时间
//
//	mSessions map[string]*Session //保存session的指针[sessionID] = session
//}
//
////创建会话管理器(cookieName:在浏览器中cookie的名字;maxLifeTime:最长生命周期)
//func NewSessionMgr(cookieName string, maxLifeTime int64) *SessionMgr {
//	mgr := &SessionMgr{mCookieName: cookieName, mMaxLifeTime: maxLifeTime, mSessions: make(map[string]*Session)}
//
//	//启动定时回收
//	go mgr.GC()
//
//	return mgr
//}
//
////在开始页面登陆页面，开始Session
//func (mgr *SessionMgr) StartSession(w http.ResponseWriter, r *http.Request) string {
//	mgr.mLock.Lock()
//	defer mgr.mLock.Unlock()
//
//	//无论原来有没有，都重新创建一个新的session
//	newSessionID := url.QueryEscape(mgr.NewSessionID())
//
//	//存指针
//	var session *Session = &Session{mSessionID: newSessionID, mLastTimeAccessed: time.Now(), mValues: make(map[interface{}]interface{})}
//	mgr.mSessions[newSessionID] = session
//	//让浏览器cookie设置过期时间
//	cookie := http.Cookie{Name: mgr.mCookieName, Value: newSessionID, Path: "/", HttpOnly: true, MaxAge: int(mgr.mMaxLifeTime)}
//	http.SetCookie(w, &cookie)
//
//	return newSessionID
//}
//
////结束Session
//func (mgr *SessionMgr) EndSession(w http.ResponseWriter, r *http.Request) {
//	cookie, err := r.Cookie(mgr.mCookieName)
//	if err != nil || cookie.Value == "" {
//	return
//	} else {
//		mgr.mLock.Lock()
//		defer mgr.mLock.Unlock()
//
//		delete(mgr.mSessions, cookie.Value)
//
//		//让浏览器cookie立刻过期
//		expiration := time.Now()
//		cookie := http.Cookie{Name: mgr.mCookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
//		http.SetCookie(w, &cookie)
//	}
//}
//
////结束session
//func (mgr *SessionMgr) EndSessionBy(sessionID string) {
//	mgr.mLock.Lock()
//	defer mgr.mLock.Unlock()
//
//	delete(mgr.mSessions, sessionID)
//}
//
////设置session里面的值
//func (mgr *SessionMgr) SetSessionVal(sessionID string, key interface{}, value interface{}) {
//	mgr.mLock.Lock()
//	defer mgr.mLock.Unlock()
//
//	if session, ok := mgr.mSessions[sessionID]; ok {
//		session.mValues[key] = value
//	}
//}
//
////得到session里面的值
//func (mgr *SessionMgr) GetSessionVal(sessionID string, key interface{}) (interface{}, bool) {
//	mgr.mLock.RLock()
//	defer mgr.mLock.RUnlock()
//
//	if session, ok := mgr.mSessions[sessionID]; ok {
//		if val, ok := session.mValues[key]; ok {
//			return val, ok
//		}
//	}
//
//	return nil, false
//}
//
////得到sessionID列表
//func (mgr *SessionMgr) GetSessionIDList() []string {
//	mgr.mLock.RLock()
//	defer mgr.mLock.RUnlock()
//
//	sessionIDList := make([]string, 0)
//
//	for k, _ := range mgr.mSessions {
//		sessionIDList = append(sessionIDList, k)
//	}
//
//	return sessionIDList[0:len(sessionIDList)]
//}
//
////判断Cookie的合法性（每进入一个页面都需要判断合法性）
//func (mgr *SessionMgr) CheckCookieValid(w http.ResponseWriter, r *http.Request) string {
//	var cookie, err = r.Cookie(mgr.mCookieName)
//
//	if cookie == nil ||
//		err != nil {
//		return ""
//	}
//
//	mgr.mLock.Lock()
//	defer mgr.mLock.Unlock()
//
//	sessionID := cookie.Value
//
//	if session, ok := mgr.mSessions[sessionID]; ok {
//		session.mLastTimeAccessed = time.Now() //判断合法性的同时，更新最后的访问时间
//		return sessionID
//	}
//
//	return ""
//}
//
////更新最后访问时间
//func (mgr *SessionMgr) GetLastAccessTime(sessionID string) time.Time {
//	mgr.mLock.RLock()
//	defer mgr.mLock.RUnlock()
//
//	if session, ok := mgr.mSessions[sessionID]; ok {
//		return session.mLastTimeAccessed
//	}
//
//	return time.Now()
//}
//
////GC回收
//func (mgr *SessionMgr) GC() {
//	mgr.mLock.Lock()
//	defer mgr.mLock.Unlock()
//
//	for sessionID, session := range mgr.mSessions {
//		//删除超过时限的session
//		if session.mLastTimeAccessed.Unix()+mgr.mMaxLifeTime < time.Now().Unix() {
//			delete(mgr.mSessions, sessionID)
//		}
//	}
//
//	//定时回收
//	time.AfterFunc(time.Duration(mgr.mMaxLifeTime)*time.Second, func() { mgr.GC() })
//}
//
////创建唯一ID
//func (mgr *SessionMgr) NewSessionID() string {
//	b := make([]byte, 32)
//	if _, err := io.ReadFull(rand.Reader, b); err != nil {
//		nano := time.Now().UnixNano() //微秒
//		return strconv.FormatInt(nano, 10)
//	}
//	return base64.URLEncoding.EncodeToString(b)
//}
//
////——————————————————————————
///*会话*/
//type Session struct {
//	mSessionID        string                      //唯一id
//	mLastTimeAccessed time.Time                   //最后访问时间
//	mValues           map[interface{}]interface{} //其它对应值(保存用户所对应的一些值，比如用户权限之类)
//}
