package handler

//import (
//	"errors"
//	"html/template"
//	"mimi/basketball/session"
//	"net/http"
//)
//
//func Login(w http.ResponseWriter, r *http.Request) {
//	//fmt.Println("method:", r.Method) //获取请求的方法
//	//fmt.Println("path", r.URL.Path)
//	if r.Method == "GET" {
//		if authentication, _, _ := CheckLogin(w, r); authentication {
//			http.Redirect(w, r, "/", http.StatusFound)
//			return
//		}
//		t, _ := template.ParseFiles("html/template/login.html")
//		t.Execute(w, nil)
//	} else {
//		values := make(map[string]interface{})
//		if err := LoginAction(w, r); err == nil {
//			http.Redirect(w, r, "/", http.StatusFound)
//			return
//		} else {
//			values["error"] = err.Error()
//		}
//		t, _ := template.ParseFiles("html/template/login.html")
//		t.Execute(w, values)
//	}
//}
//
//func Logout(w http.ResponseWriter, r *http.Request) {
//	values := make(map[string]interface{})
//	if err := LogoutAction(w, r); err != nil {
//		values["error"] = err.Error()
//	}
//	t, _ := template.ParseFiles("html/template/login.html")
//	t.Execute(w, values)
//}
//
////检验登录信息，若当前账号已登录，不做处理，若另有账号已登录，则删除旧账号会话
//func LoginAction(w http.ResponseWriter, r *http.Request) error {
//	loginName := r.FormValue("loginName")
//	authentication, session, err := CheckLogin(w, r)
//	if err != nil {
//		return err
//	} else {
//		oldName, err := session.Get("loginName")
//		if err != nil {
//			return err
//		}
//		if oldName == loginName && oldName != "" {
//			if authentication {
//				return nil
//			} else {
//				if err := session.Set("authentication", true); err != nil {
//					return err
//				} else {
//					return nil
//				}
//			}
//		}
//	}
//	password := r.FormValue("password")
//	if err := checkLoginNamePassword(loginName, password); err != nil {
//		return err
//	} else {
//		newSn := NewSession(w)
//		if err := newSn.Set("loginName", loginName); err != nil {
//			return err
//		} else if err := newSn.Set("authentication", true); err != nil {
//			return err
//		} else if err := session.Del(); err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//func checkLoginNamePassword(name, password string) error {
//	if name == "" {
//		return errors.New("登录名不能为空")
//	}
//	if password == "" {
//		return errors.New("密码不能为空")
//	}
//	if ("mimi" != name && "bluemimi" != name) || "123123" != password {
//		return errors.New("账号密码不匹配")
//	}
//	return nil
//}
//
////获取会话（自动初始化），检查是否已经登录
//func CheckLogin(w http.ResponseWriter, r *http.Request) (bool, *session.Session, error) {
//	if session, err := GetOrInitSession(w, r); err != nil {
//		return false, session, err
//	} else {
//		if authentication, err := session.GetBool("authentication"); err != nil {
//			return false, session, err
//		} else {
//			return authentication, session, err
//		}
//	}
//}
//
////通过cookie获取会话ID
//func GetCurSessionId(r *http.Request) (string, error) {
//	if cookie, err := r.Cookie("sessionId"); err != nil && err != http.ErrNoCookie {
//		return "", err
//	} else {
//		if cookie == nil {
//			return "", nil
//		}
//		return cookie.Value, nil
//	}
//}
//
////获取会话，不存在则创建一个
//func GetOrInitSession(w http.ResponseWriter, r *http.Request) (*session.Session, error) {
//	if sessionId, err := GetCurSessionId(r); err != nil {
//		return nil, err
//		//return NewSession(w), nil
//	} else if sessionId != "" {
//		return session.Get(sessionId), err
//	} else {
//		return NewSession(w), err
//	}
//}
//
////新建一个会话并保存到cookie
//func NewSession(w http.ResponseWriter) *session.Session {
//	sessionId := session.NewSessionId()
//	//expiration := time.Now()
//	//expiration = expiration.AddDate(1, 0, 0)
//	cookie := http.Cookie{Name: "sessionId", Value: sessionId, Path: "/"}
//	http.SetCookie(w, &cookie)
//	return session.Get(sessionId)
//}
//
////检查登录状态并标记为未登录
//func LogoutAction(w http.ResponseWriter, r *http.Request) error {
//	if authentication, session, err := CheckLogin(w, r); err != nil {
//		return err
//	} else if authentication {
//		if err := session.Set("authentication", false); err != nil {
//			return err
//		}
//	}
//	return nil
//}
