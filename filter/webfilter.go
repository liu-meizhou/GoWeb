package filter

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	redisdb "redis"
)
//StaticSafe  12
type StaticSafe struct{
	root http.FileSystem
}
//SafeFileServer 安全
func SafeFileServer(root http.FileSystem) http.Handler {
	return &StaticSafe{root}
}

//StaticSafeHandler 安全验证处理
func (p *StaticSafe)ServeHTTP(w http.ResponseWriter, r *http.Request){
	//安全处理（过滤器）
	if !myfilter(r) {
		http.Redirect(w, r, "/Login.html", http.StatusTemporaryRedirect)
		return
	}
	http.FileServer(p.root).ServeHTTP(w,r)
}



//SafeHandler 安全验证处理
func SafeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		//编码处理

		//安全处理（过滤器）
		if !myfilter(r) {
			http.Redirect(w, r, "/Login.html", http.StatusTemporaryRedirect)
			return
		}

		//异常处理
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				log.Printf("WARN: panic in %v - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()

		fn(w,r)
	}
}

func myfilter(r *http.Request) bool {
	cookie, err := r.Cookie("MyCookie")

	if err != nil {
		fmt.Println("MyCookie不存在")
		return false
	}

	CookieValue := cookie.Value

	_,err = redisdb.GetDo(CookieValue)

	if err != nil{
		return false
	}
	return true
}
