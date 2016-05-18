package deeplink

import (
	"fmt"
	"github.com/michaelfemi81/go-inliner"

	"os"
	"strings"
	"net/http"
	"html/template"

	//"reflect"
	"errors"
)
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
type DeepLink struct {
	Android_package_name string
	Title                string
	Ios_store_link       string
	Fallback             string
	Url string
}

func (d *DeepLink) Init(f string, t... string) (error) {
	d.Fallback = f
	if (strings.Trim(f, " ") == "") {
		return errors.New("Invalid Fallback");
	}
	if (len(t) < 1) {
		return errors.New("Set At Least Android Package");
	} else {
		if (len(t) == 1) {
			d.Android_package_name = t[0]
			d.Ios_store_link = ""
			d.Title = ""
		} else if (len(t) == 2) {
			d.Android_package_name = t[0]
			d.Ios_store_link = t[1]
			d.Title = ""
		} else {
			d.Android_package_name = t[0]
			d.Ios_store_link = t[1]
			d.Title = t[2]
		}
	}




return nil
}
func(d DeepLink)DoDeepLink(w http.ResponseWriter, r *http.Request)(err error){
	url:=r.URL.Query()["url"];
	//fmt.Println(d.Title)
	if(url==nil){
		err=errors.New("no query uri");
		fmt.Fprintln(w,"Invalid Query String");
		return
	}else{

		d.Url=url[0]

		t := template.New("index.html")
		t,err:=t.ParseFiles("./public/index.html");
		fil, err := os.OpenFile("./public/aa.html", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			//  os.Exit(1)
		}
		err=t.Execute(fil,d);
		if err != nil {
			fmt.Println(err)
			//  os.Exit(1)
		}else{
			inliner.RenderToHttp("./public/aa.html",w,r);

		}
		defer remove("./public/aa.html")
		defer fil.Close()



}


return
}
func remove(aa string){
	err3 := os.Remove(aa)
	if err3 != nil {
		fmt.Println(err3)
	}
}
