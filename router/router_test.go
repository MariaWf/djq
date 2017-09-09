package router

import (
	"fmt"
	"io/ioutil"
	"mimi/djq/util"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestAdminList(t *testing.T) {
	resp, err := http.Get("http://djq.tunnel.qydev.com/mi/admin/?name=mimi&targetPage=2&pageSize=2")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func TestAdminAdd(t *testing.T) {
	password, _ := util.EncryptPassword("123123")
	roleIds := "d3490d47eaee4e7c85077baa9542908b"
	fmt.Println(password)
	fmt.Println(util.DecryptPassword(password))
	resp, err := http.PostForm("http://localhost:8080/mi/admin/?name=mimi",
		url.Values{"mobile": {"12222222222"}, "name": {"mimiLogin1"}, "password": {password}, "locked": {"false"}, "roleIds": {roleIds}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
func TestAdminLogin(t *testing.T) {
	password, _ := util.EncryptPassword("123123")
	fmt.Println(util.BuildPassword4DB("123123"))
	resp, err := http.PostForm("http://localhost:8080/mi/login",
		url.Values{"name": {"mimiLogin1"}, "password": {password}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
func httpGet() {
	resp, err := http.Get("http://djq.tunnel.qydev.com/mi/admin/?name=mimi&targetPage=2&pageSize=2")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
func httpPost() {
	resp, err := http.Post("http://www.01happy.com/demo/accept.php",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
func httpPostForm() {
	password, _ := util.EncryptPassword("s,3214.")
	roleIds := "d3490d47eaee4e7c85077baa9542908b,,"
	fmt.Println(password)
	fmt.Println(util.DecryptPassword(password))
	resp, err := http.PostForm("http://djq.tunnel.qydev.com/mi/admin/?name=mimi",
		url.Values{"mobile": {"12222222222"}, "name": {"mimiWithRoles4"}, "password": {password}, "locked": {"true"}, "roleIds": {roleIds}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}
func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://www.01happy.com/demo/accept.php", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
