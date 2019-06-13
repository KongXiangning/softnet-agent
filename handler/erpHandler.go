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
