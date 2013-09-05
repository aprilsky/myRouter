/**
暂时不支持 filter 仅仅支持mux
 */
package main

import (
	."controller"
	"mux"
	"net/http"
	"log"
)

func initRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/topics{view:(|/popular|/no_reply|/last)}", TopicsHandler)
	router.HandleFunc("/topics/{tid:[0-9]+}", TopicDetailHandler)
	router.HandleFunc("/topics/new{json:(|.json)}", NewTopicHandler)
	// 登录
	router.HandleFunc("/account/login", LoginHandler)
	router.HandleFunc("/account/logout", LogoutHandler)
	// 用户相关
	router.HandleFunc("/user/{username:\\w+}", UserHomeHandler)
	// 404页面
	//router.HandleFunc("/{*}", NotFoundHandler)
	return router
}
func main() {
	router := initRouter()
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
