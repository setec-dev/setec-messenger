package main

import (
	"log"
	"os/exec"
	"os/user"
	"time"

	"github.com/webview/webview"
)

func main() {
	go run_node()
	time.Sleep(300)
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Setec-Secure-Messenger")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("http://localhost:5000")
	w.Run()
}

func run_node() {
	log.Println("[*] Running node:")
	user.Current()
	node := "/usr/bin/npm"
	arg1 := "run"
	arg2 := "dev"
	run_srv := exec.Command(node, arg1, arg2)
	stdout, err := run_srv.Output()

	if err != nil {
		log.Println("---- [*] ERROR:")
		log.Println(err.Error())
		return
	}

	log.Println(string(stdout))
}
