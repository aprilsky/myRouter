/**
 * @Author: caoxiao
 * @Date: 13-9-2 下午5:26
 */
package controller
import (
	"fmt"
	"net/http"
	"io"
)
// 首页
func IndexHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL)
	_,err:=io.WriteString(rw,req.URL.Path)
	if err!=nil{
		fmt.Println(err)
	}
}
// LoginHandler
func LoginHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL)
	_,err:=io.WriteString(rw,req.URL.Path)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("LoginHandler")
}
// LogoutHandler
func LogoutHandler(rw http.ResponseWriter, req *http.Request) {


}// NewTopicHandler
func NewTopicHandler(rw http.ResponseWriter, req *http.Request) {


}// TopicsHandler
func TopicsHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL)
	_,err:=io.WriteString(rw,req.URL.Path)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("TopicsHandler")
}// TopicDetailHandler
func TopicDetailHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL)
	_,err:=io.WriteString(rw,req.URL.Path)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("TopicDetailHandler")

}// UserHomeHandler
func UserHomeHandler(rw http.ResponseWriter, req *http.Request) {

}



