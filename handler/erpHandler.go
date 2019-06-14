package handler

import (
	"fmt"
	"net/http"
)

func helloworld(w http.ResponseWriter, r *http.Request)  {
	mr,err := r.MultipartReader()
	if err != nil{
		fmt.Println("r.MultipartReader() err,",err)
		return
	}
	form ,_ := mr.ReadForm(128)

	getFormData(form)

	if out,err := execcmd("echo hello world s1");err != nil{
		fmt.Fprintf(w,fmt.Sprint(err))
	}else{
		fmt.Fprintf(w,out)
	}
}

func run(w http.ResponseWriter, r *http.Request)  {
	var (
		message string
		commond string
		err error
		imageName string
	)

	catalog := r.FormValue("catalog")
	name := r.FormValue("name")
	ctype := r.FormValue("type")
	version := r.FormValue("version")
	port := r.FormValue("port")

	imageName = fmt.Sprintf("docker.scn.weilian.cn/%s/%s-%s:%s",catalog,name,ctype,version)

	if message,err = stopContainer(catalog,name,ctype);err != nil{
		goto ERROR
	}else if message != "" {
		fmt.Fprintf(w,"Last container has been deleted. \n")
	}

	commond = fmt.Sprintf("docker run -d -e TZ=\"Asia/Shanghai\" -v /etc/localtime:/etc/localtime:ro -v /opt/docker/data/rocketmq/namesrv/store:/usr/local/rocketmq/store -v /opt/docker/data/rocketmq/namesrv/logs:/usr/local/rocketmq/logs --net=host --privileged=true --name rocketmq-namesrv %s",imageName)
	if message,err = execcmd(commond);err != nil {
		goto ERROR
	}else {
		fmt.Fprintf(w,"dns run containid:%s \n",message)
	}

	if err = openPort(9876,"tcp");err != nil{
		goto ERROR
	}
	fmt.Fprintf(w,"rocketmq-namesrv started")
	return
ERROR:
	fmt.Fprintf(w,fmt.Sprint(err))
}

func stopContainer(catalog string,name string,ctype string)(string,error)  {
	var(
		result string
		err error
	)
	projectName := fmt.Sprintf("%s-%s-%s",catalog,name,ctype)
	if result,err = execcmd(fmt.Sprintf("docker ps -a|grep %s|awk '{ print $(NF-0) }'",projectName));err != nil {
		return "",err
	}
	if result == ""{
		return "nil",nil
	}
	if _,err = execcmd(fmt.Sprintf("docker stop %s && docker rm %s",result,result));err != nil{
		return "",err
	}
	return "success",nil
}