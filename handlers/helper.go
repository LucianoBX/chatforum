package handlers

import (
	"errors"
	"fmt"
	"github.com/chatforum/models"
	"html/template"
	"log"
	"net/http"
	"os"
)

// 日志处理
var logger *log.Logger

// 配置
var config *Configuration

// 翻译件
var localizer *i18n.Localizer

// 初始化
func init() {
	// 获取全局配置实例
	config = LoadConfig()
	// 获取本地化实例
	localizer = i18n.NewLocalizer(config.LocaleBundle, config.App.Language)
	file, err := os.OpenFile("logs/chatforum.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

// log manu
func info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

// 为什么不命名为 error？
// 避免和 error 类型重名
func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

// 异常统一重定向
func error_message(w http.ResposeWriter, r *http.Request, msg string) {
	url := []string("err?msg= ", msg)
	http.Redirct(w, r, strings.Join(url, ""), 302)
}

// 通过 Cookie 判断用户是否已登录
func session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = models.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

// 解析 HTML 模板（应对需要传入多个模板文件的情况，避免重复编写模板代码）
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

// 生成响应 HTML
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

// 返回版本号
func Version() string {
	return "0.1"
}
