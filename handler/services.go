package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

func postgres_erpone(w http.ResponseWriter,r *http.Request)  {
	r.Form.Add("name","postgre_5432")
	r.Form.Add("port","5432")
	postgres(w,r)
}

func postgres_erptwo(w http.ResponseWriter,r *http.Request)  {
	r.Form.Add("name","postgres_5433")
	r.Form.Add("port","5433")
	postgres(w,r)
}

func postgres(w http.ResponseWriter, r *http.Request)  {
	var (
		message string
		commond string
		err error
		imageName string
		name string
		port int
	)

	name = r.FormValue("name")
	if name == "" {
		name = "postgres"
	}
	if r.FormValue("port") == "" {
		port = 5432
	}else {
		if port,err = strconv.Atoi(r.FormValue("port"));err != nil{
			goto ERROR
		}
	}
	imageName = r.FormValue("image")
	if imageName == "" {
		imageName = "docker.scn.weilian.cn/library/postgres:9.6.3"
	}

	commond = fmt.Sprintf("docker run -d -e TZ=\"Asia/Shanghai\" -v /etc/localtime:/etc/localtime:ro -v /opt/docker/data/postgres/%s:/var/lib/postgresql/data  --net=host --restart always --name %s --privileged=true %s",name,name,imageName)
	if message,err = execcmd(commond);err != nil {
		goto ERROR
	}else {
		fmt.Fprintf(w,"dns run containid:%s \n",message)
	}

	if err = openPort(port,"tcp");err != nil{
		goto ERROR
	}
	fmt.Fprintf(w,"postgres started")
	return
ERROR:
	fmt.Fprintf(w,fmt.Sprint(err))
}

func dns(w http.ResponseWriter, r *http.Request)  {
	var (
		dnsmap []string
		message string
		commond string
		err error
		imageName string
	)

	if mr,err := r.MultipartReader();err != nil{
		goto  ERROR
	}else {
		form ,_ := mr.ReadForm(128)
		for k, v := range form.Value {
			if k == "dns" {
				dnsmap = v
			}else if k == "image" {
				imageName = v[0]
			}
		}
	}

	for i, v := range dnsmap{
		if _,err = execcmd(fmt.Sprintf("echo %s >> /opt/docker/data/dnsmasq/dnsmasq.hosts",v));err != nil {
			if i != 0{
				execcmd("echo '' > /opt/docker/data/dnsmasq/dnsmasq.hosts")
			}
			goto ERROR
		}
	}

	if imageName == "" {
		imageName = "docker.scn.weilian.cn/library/dnsmasq:v1"
	}
	commond = fmt.Sprintf("docker run -d -e TZ=\"Asia/Shanghai\" -v /etc/localtime:/etc/localtime:ro -v /opt/docker/data/dnsmasq/dnsmasq.conf:/etc/dnsmasq.conf -v /opt/docker/data/dnsmasq/dnsmasq.hosts:/etc/dnsmasq.hosts -v /opt/docker/data/dnsmasq/resolv.dnsmasq:/etc/resolv.dnsmasq --cap-add=NET_ADMIN --net=host --name dnsmasq %s",imageName)
	if message,err = execcmd(commond);err != nil {
		goto ERROR
	}else {
		fmt.Fprintf(w,"dns run containid:%s \n",message)
	}

	if err = openPort(53,"udp");err != nil{
		goto ERROR
	}
	fmt.Fprintf(w,"dns started")
	return
	ERROR:
		fmt.Fprintf(w,fmt.Sprint(err))
}

func zookeeper(w http.ResponseWriter, r *http.Request)  {
	var (
		message string
		commond string
		err error
	)

	imageName := r.FormValue("image")
	if imageName == "" {
		imageName = "docker.scn.weilian.cn/library/zookeeper:v1"
	}
	commond = fmt.Sprintf("docker run -d -e TZ=\"Asia/Shanghai\" -v /etc/localtime:/etc/localtime:ro -v /opt/docker/data/zookeeper/data:/data -v /opt/docker/data/zookeeper/datalog:/datalog -v /opt/docker/data/zookeeper/conf/zoo.cfg:/conf/zoo.cfg --privileged=true --net=host --restart always --name zookeeper %s",imageName)
	if message,err = execcmd(commond);err != nil {
		goto ERROR
	}else {
		fmt.Fprintf(w,"dns run containid:%s \n",message)
	}

	if err = openPort(12233,"tcp");err != nil{
		goto ERROR
	}
	fmt.Fprintf(w,"zookeeper started")
	return
	ERROR:
		fmt.Fprintf(w,fmt.Sprint(err))
}

func redis(w http.ResponseWriter, r *http.Request)  {
	var (
		message string
		commond string
		err error
		imageName string
	)

	imageName = r.FormValue("image")
	if imageName == "" {
		imageName = "docker.scn.weilian.cn/library/redis:v1"
	}
	commond = fmt.Sprintf("docker run -d -e TZ=\"Asia/Shanghai\" -v /etc/localtime:/etc/localtime:ro -v /opt/docker/data/redis/data:/data -v /opt/docker/data/redis/conf/:/usr/local/etc/redis/--net=host --name redis --restart always --privileged=true %s",imageName)
	if message,err = execcmd(commond);err != nil {
		goto ERROR
	}else {
		fmt.Fprintf(w,"dns run containid:%s \n",message)
	}

	if err = openPort(6379,"tcp");err != nil{
		goto ERROR
	}
	fmt.Fprintf(w,"redis started")
	return
ERROR:
	fmt.Fprintf(w,fmt.Sprint(err))
}

func mqNamesrv(w http.ResponseWriter, r *http.Request)  {
	var (
		message string
		commond string
		err error
		imageName string
	)

	imageName = r.FormValue("image")
	if imageName == "" {
		imageName = "docker.scn.weilian.cn/library/rocketmq-namesrv:v1"
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

func mqBroker(w http.ResponseWriter, r *http.Request)  {
	var (
		message string
		commond string
		err error
		imageName string
	)

	ip := r.FormValue("ip")
	imageName = r.FormValue("image")
	if imageName == "" {
		imageName = "docker.scn.weilian.cn/library/rocketmq-broker:v1"
	}
	commond = fmt.Sprintf("sed -i 's/{#ip}/%s/g' /opt/docker/data/rocketmq/broker/config/broker.properties",ip)
	if message,err = execcmd(commond);err != nil{
		goto ERROR
	}
	commond = fmt.Sprintf("docker run -d -e TZ=\"Asia/Shanghai\" -v /etc/localtime:/etc/localtime:ro -v /opt/docker/data/rocketmq/namesrv/store:/usr/local/rocketmq/store -v /opt/docker/data/rocketmq/namesrv/logs:/usr/local/rocketmq/logs --net=host --privileged=true --name rocketmq-namesrv %s",imageName)
	if message,err = execcmd(commond);err != nil {
		goto ERROR
	}else {
		fmt.Fprintf(w,"dns run containid:%s \n",message)
	}

	if err = openPort(10909,"tcp");err != nil{
		goto ERROR
	}
	if err = openPort(10911,"tcp");err != nil{
		goto ERROR
	}
	fmt.Fprintf(w,"rocketmq-brokerr started")
	return
ERROR:
	fmt.Fprintf(w,fmt.Sprint(err))
}
