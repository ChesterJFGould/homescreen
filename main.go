package main

import (
	"net/http"
	"os"
	"os/exec"

	"github.com/zserge/webview"
)

var (
	port = "3141"
	webV webview.WebView
	dir  string
)

func main() {
	dir = os.Args[1]
	if len(os.Args) > 2 {
		port = os.Args[2]
	}

	webV = webview.New(webview.Settings{
		Title:     "homeScreen",
		URL:       "http://127.0.0.1:" + port + "/",
		Resizable: true,
	})

	go startServer()

	webV.Dispatch(func() {
		webV.SetFullscreen(true)
	})

	webV.Run()
}

func handle(e error) {
	if e != nil {
		panic(e)
	}
}

func startServer() {
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/reboot", reboot)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/quit", quit)
	http.HandleFunc("/images/", images)
	http.HandleFunc("/", index)

	http.ListenAndServe("127.0.0.1:"+port, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/chester/Projects/GUI/homeScreen/index.html")
}

func images(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/images/", http.FileServer(http.Dir(dir+"/images"))).ServeHTTP(w, r)
}

func shutdown(w http.ResponseWriter, r *http.Request) {
	println("shutdown")
	cmd := exec.Command(dir + "/shutdown.sh")
	err := cmd.Run()
	handle(err)
}

func reboot(w http.ResponseWriter, r *http.Request) {
	println("reboot")
	cmd := exec.Command(dir + "/reboot.sh")
	err := cmd.Run()
	handle(err)
}

func logout(w http.ResponseWriter, r *http.Request) {
	println("logout")
	cmd := exec.Command(dir + "/logout.sh")
	err := cmd.Run()
	handle(err)
}

func quit(w http.ResponseWriter, r *http.Request) {
	println("quit")
	webV.Terminate()
}
