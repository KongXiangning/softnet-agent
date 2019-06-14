package handler

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

func HandlerInit() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/restart/",restart)
	mux.HandleFunc("/sleep/",addProject)
	mux.HandleFunc("/hello/",helloworld)
	mux.HandleFunc("/init/postgres",postgres)
	mux.HandleFunc("/init/postgres/erpone",postgres_erpone)
	mux.HandleFunc("/init/postgres/erptwo",postgres_erptwo)
	mux.HandleFunc("/init/dns",dns)
	mux.HandleFunc("/init/zookeeper",zookeeper)
	mux.HandleFunc("/init/redis",redis)
	mux.HandleFunc("/init/mqnamesrv",mqNamesrv)
	mux.HandleFunc("/init/mqbroker",mqBroker)
	return mux
}

func restart(w http.ResponseWriter, r *http.Request)  {
	var err error
	tag := r.FormValue("tag")
	path := fmt.Sprintf("https://files.scn.weilian.cn/o/softnet/%s/softnet-agent",tag)

	if _,err = execcmd("rm -f /opt/agent");err != nil {
		goto ERROR
	}

	if _,err = execcmd(fmt.Sprintf("wget -O /opt/agent %s",path));err != nil{
		goto ERROR
	}

	if _,err = execcmd("chmod +x /opt/agent");err != nil{
		goto ERROR
	}

	if _,err = execcmd(fmt.Sprintf("kill -USR2 %d",os.Getpid()));err != nil{
		goto ERROR
	}
	fmt.Fprintf(w,"update success")
	return

	ERROR:
		fmt.Fprintf(w,fmt.Sprint(err))
}

func addProject(w http.ResponseWriter, r *http.Request)  {
	now := time.Now()
	duration, err := time.ParseDuration(r.FormValue("duration"))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	time.Sleep(duration)
	fmt.Fprintf(
		w,
		"started111 at %s slept for %d nanoseconds from pid %d.\n",
		now,
		duration.Nanoseconds(),
		os.Getpid(),
	)
}

func execcmd(commond string) (string,error)  {
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)
	cmd := exec.Command("sh", "-c", commond)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "",fmt.Errorf(fmt.Sprintf("error:%s:%s,error commond:%s",fmt.Sprint(err),stderr.String(),commond))
	}
	return stdout.String(),nil
}

func openPort(port int,t string) error {
	if _,err := execcmd(fmt.Sprintf("iptables -C INPUT -p %s --dport %d -j ACCEPT",t,port));err != nil {
		if strings.Contains(err.Error(),"Bad rule") {
			if _,err = execcmd(fmt.Sprintf("iptables -I INPUT -p %s --dport %d -j ACCEPT",t,port));err != nil {
				return err
			}
		}else{
			return err
		}
	}
	return nil
}


func getFormData(form *multipart.Form) {

	for k, v := range form.Value {

		fmt.Println("value,k,v = ", k, ",", v)

	}
}