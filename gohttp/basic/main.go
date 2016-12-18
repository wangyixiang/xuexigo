package main

import (
	"net/http"
	"bytes"
	"log"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"strconv"
	"github.com/mholt/binding"
	"io/ioutil"
	"encoding/json"
	"os"
	"io"
	"time"
	"context"
	"sync"
	"github.com/rs/xhandler"
	oldcontext "golang.org/x/net/context"
	"net"
)

func main() {
	var err error
	var way = "httprouterway"
	switch way {
	case "nethttpway":
		err = nethttpway()
	case "httprouterway":
		err = httprouterway()
	}
	if err != nil {
		log.Fatal("Starting failed on error: ", err)
	}
}

var myhandle = &myHandle{}

type myHandle struct {

}

func (mh *myHandle) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte(req.URL.Path + "\n"))
	resp.Write([]byte(req.Host + "\n"))
	resp.Write([]byte(req.RemoteAddr + "\n"))
	resp.Write([]byte(req.UserAgent()))
}

func (mh *myHandle) httprouterServeHTTP(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	resp.Write([]byte("From httprouter\n"))
	mh.ServeHTTP(resp, req)
}

// 这里传过来的是application/x-www-form-urlencoded, 常规的表单数据, 用parseForm后, 用Form和PostForm来进一步处理.
func (mh *myHandle) hrFormPostForm(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	req.ParseForm()
	resp.Write([]byte("req.Form: " + strconv.Itoa(len(req.Form)) + "\n"))
	for k := range req.Form {
		resp.Write([]byte("----wangyixiang----\n"))
		s := fmt.Sprintf("%v = %v\n", k, req.FormValue(k))
		resp.Write([]byte(s))
	}
	resp.Write([]byte("req.PostForm: " + strconv.Itoa(len(req.PostForm)) + "\n"))
	for k := range req.PostForm {
		resp.Write([]byte("----wangyixiang----\n"))
		s := fmt.Sprintf("%v = %v\n", k, req.PostForm.Get(k))
		resp.Write([]byte(s))
	}

}

type user struct {
	id   int
	name string
}

func (u *user) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&u.id: "uid",
		&u.name: binding.Field{
			Form: "uname",
			Required: true,
		},
	}
}
func (mh *myHandle) hrBindUser(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	u := new(user)
	err := binding.Bind(req, u)
	if err.Handle(resp) {
		return
	}
	resp.Write([]byte("userid: " + strconv.Itoa(u.id) + "\n"))
	resp.Write([]byte("username: " + u.name + "\n"))
}

type user2 struct {
	id   int `json:"uid"`
	Name string `json:"uname"`
}

// 而如果是别的数据, 比如是json数据的话, 则要读取body,来进一步处理, 在使用json.Unmarshal时, 由于user2会传给json package,
// 其里面定义的field要给json访问得到的权限, 所以就必须要把给json访问的字段export, 在golang里,就是要大写. 比如我这个, id就永远是0,
// Name是可以正常拿到json里的值.
func (mh *myHandle) hrJsonUser(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	jsonData, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		resp.Write([]byte(fmt.Sprint(err) + "\n"))
		resp.WriteHeader(500)
		return
	}
	u := new(user2)
	fmt.Println(string(jsonData))
	if err = json.Unmarshal(jsonData, &u); err != nil {
		resp.Write([]byte(fmt.Sprint(err) + "\n"))
		resp.WriteHeader(500)
		return
	}
	resp.Write([]byte("userid: " + strconv.Itoa(u.id) + "\n"))
	resp.Write([]byte("username: " + u.Name + "\n"))
}

// 上传文件, golang来读代码就可以看得清清楚楚. no black box, that's for unix-like programmer
func (mh *myHandle) hrUploadFile(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	err := req.ParseMultipartForm(32 << 20)
	fmt.Println(req.Header.Get("Content-Type"))
	if err != nil {
		// Write会设置status code为200, 所以在其后调用WriteHeader是会有一个warning, 并且无效,
		// warning是multiple response.WriteHeader calls
		resp.Write([]byte(fmt.Sprint(err) + "\n"))
		resp.WriteHeader(500)
		return
	}

	resp.Write([]byte("req.MultipartForm.File: " + strconv.Itoa(len(req.MultipartForm.File)) + "\n"))
	for k, v := range req.MultipartForm.File {
		resp.Write([]byte("----wangyixiang----\n"))
		s := fmt.Sprintf("filename = %v filecount = %v\n", k, len(v))
		resp.Write([]byte(s))
	}

	resp.Write([]byte("req.MultipartForm.Value: " + strconv.Itoa(len(req.MultipartForm.Value)) + "\n"))
	for k := range req.MultipartForm.Value {
		resp.Write([]byte("----wangyixiang----\n"))
		s := fmt.Sprintf("filename = %v\n", k)
		resp.Write([]byte(s))
	}

	file, filehead, err := req.FormFile("wyxfile")
	if err != nil {
		resp.Write([]byte(fmt.Sprint(err) + "\n"))
		resp.WriteHeader(500)
		return
	}

	filetype := filehead.Header.Get("Content-Type")
	resp.Write([]byte("uploaded file type is " + filetype + "\n"))
	f, err := os.Create("upload.file")
	defer f.Close()
	if err != nil {
		resp.Write([]byte(fmt.Sprint(err) + "\n"))
		resp.WriteHeader(500)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		resp.Write([]byte("uploaded failed!"))
		resp.Write([]byte(fmt.Sprint(err) + "\n"))
		resp.WriteHeader(500)
		return
	}
	resp.Write([]byte("uploaded successfully."))
}

func (mh *myHandle) hrCookie(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	cookie := &http.Cookie{
		Name: "author",
		Value: "wangyixiang",
		Expires: time.Now().Add(1000 * time.Second),
		MaxAge: 1000,
	}
	http.SetCookie(resp, cookie)
	resp.Write([]byte("cookie setup successfully."))

}

// a captcha

// https://blog.golang.org/context
// At Google, we developed a context package that makes it easy to pass request-scoped values,
// cancelation signals, and deadlines across API boundaries to all the goroutines involved in
// handling a request.
// 上面这段话摘自blog.golang.org上的介绍context的blog, 很明确地指出了context的应用场景, 就是一个request的
// 声明周期.
// context 传递数据
var wg = sync.WaitGroup{}

func (mh *myHandle) hrCHandler1(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	err := req.ParseForm()
	if err != nil {
		internalError(resp, err, "")
		return
	}
	if len(req.Form) > 0 {
		for k, v := range req.Form {
			ctx := context.WithValue(req.Context(), k, v[0])
			wg.Add(1)
			go mh.hrCHandler2(resp, req.WithContext(ctx), ps, k)
		}
		wg.Wait()
	}
}

func (mh *myHandle) hrCHandler2(resp http.ResponseWriter, req *http.Request, ps httprouter.Params, key interface{}) {
	defer wg.Done()
	if v, ok := req.Context().Value(key).(string); ok {
		resp.Write([]byte(fmt.Sprintf("context.%v=%v\n", key, v)))
		return
	}
	resp.Write([]byte(fmt.Sprintf("context.%v is not a string\n", key)))
}

// context 超时机制

func (mh *myHandle) hrCTimeout(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	ctx, cancelFn := context.WithTimeout(req.Context(), time.Second)
	defer cancelFn()

	// strCh := make(chan string)
	// defer close(strCh)
	// 上面这段code有一个不是那么明显的bug, 就是由于这个method必然是以timeout结束的,在method结束后, strCh会被close掉,
	// 而在下面的goroutine中, 在5秒后会向strCh这个channel里send 数据, 而在golang中向close的channel里send数据,是会
	// panic: send on closed channel. 所以就不能用defer close.
	// 但是只是去掉defer close就可以了吗? 答案是否定的, 因为如果不defer close, 那么在下面的goroutine中, 向strCh发送
	// 这个动作将会永远被block, 因为没有别的地方会去接收她, 而这样就会导致goroutine leak.
	// 所以最好的办法就是给这个channel一个长度为1的buff, 确保goroutine不会block, 然后strCh没有再被引用, golang GC将会在
	// 合适的时候来回收她.
	strCh := make(chan string, 1)
	go func() {
		time.Sleep(5 * time.Second)
		strCh <- "golang"
	}()
	select {
	case <-ctx.Done():
		resp.Write([]byte("ctx.Done\n"))
	case str := <-strCh:
		resp.Write([]byte(str + "\n"))
	}

}

// context 被中间件使用
type myMiddleWare struct {
	next xhandler.HandlerC
}
// xhandler还在使用老的在golang.org/x/net/context下的context, 所以在不改写xhandler的情况下,我就用oldcontext来引用她.
func (mm *myMiddleWare) ServeHTTPC(ctx oldcontext.Context, resp http.ResponseWriter, req *http.Request) {
	ctx = oldcontext.WithValue(ctx, "2ndPart", "Golang")
	mm.next.ServeHTTPC(ctx, resp, req)
}

func setupMiddleWare(hr *httprouter.Router) {
	c := xhandler.Chain{}
	c.UseC(xhandler.CloseHandler)
	c.UseC(func(next xhandler.HandlerC) xhandler.HandlerC {
		return &myMiddleWare{
			next: next,
		}
	})
	xh := xhandler.HandlerFuncC(
		func(ctx oldcontext.Context, resp http.ResponseWriter, req *http.Request) {
			v, ok := ctx.Value("2ndPart").(string)
			if !ok {
				resp.Write([]byte("no 2ndPart here.\n"))
				return
			}
			resp.Write([]byte("Hello " + v + "\n"))
		})
	h := c.Handler(xh)
	hr.GET("/middle/", func(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		h.ServeHTTP(resp, req)
	})

}


// Hijack其实就是把http请求中的tcp connection拿出来耍
// 而这个handle的结果要在有websocket支持的client才能看到.
func (mh *myHandle) hrHijack(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	hj, ok := resp.(http.Hijacker)
	if !ok {
		resp.Write([]byte("Hijack Failed.\n"))
		return
	}
	conn, buf, err := hj.Hijack()
	if err != nil {
		internalError(resp, err, "")
		return
	}
	defer conn.Close()

	resp.Write([]byte("it won't be writen to client, it's hijacked."))
	buf.WriteString("it's hijacked, pay or die.\n")
	buf.Flush()
}

func internalError(resp http.ResponseWriter, err error, desc string) {
	resp.WriteHeader(500)
	resp.Write([]byte(desc))
	resp.Write([]byte(fmt.Sprint(err) + "\n"))
}

func nethttpway() error {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write(bytes.NewBufferString("This is Root").Bytes())
	})
	http.Handle("/m/", &myHandle{})
	return http.ListenAndServe("0.0.0.0:80", nil)
}

func httprouterway() error {
	hrouter := httprouter.New()
	hrouter.GET("/h/*.php", myhandle.httprouterServeHTTP)
	hrouter.GET("/fpf/", myhandle.hrFormPostForm)
	hrouter.POST("/fpf/", myhandle.hrFormPostForm)
	hrouter.GET("/u/", myhandle.hrBindUser)
	hrouter.POST("/u/", myhandle.hrBindUser)
	hrouter.POST("/u/json/", myhandle.hrJsonUser)
	hrouter.POST("/file/", myhandle.hrUploadFile)
	hrouter.GET("/cookie/", myhandle.hrCookie)
	hrouter.GET("/context/p", myhandle.hrCHandler1)
	hrouter.GET("/context/t", myhandle.hrCTimeout)
	setupMiddleWare(hrouter)
	hrouter.GET("/hijack/", myhandle.hrHijack)

	// 对server设置超时
	// 其实http.ListenAndServe就是create一个简单的server实例, 然后在其上面调用ListenAndServe, 不如自己来, 还可以使用
	// server的高级属性
	server := &http.Server{
		Handler: hrouter,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
	}

	l, err := net.Listen("tcp", "0.0.0.0:80")
	if err != nil {
		return err
	}
	return server.Serve(l)
}