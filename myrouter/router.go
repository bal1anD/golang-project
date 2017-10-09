package myrouter

import (
	"net/http"
	"fmt"
	"strings"
	"gawkbox-assignment/controllers"
)


func init() {
	fmt.Println("Initializing Router API...")
}

func Route(w http.ResponseWriter, r *http.Request) {


	p := strings.Split(r.URL.Path, "/")

	if(p[1]=="users" && p[3]=="channel") {
		controllers.GetChannelInfo(w,p[2])
	}else if (p[1]=="users" && p[3]=="info") {
		controllers.GetUserInfo(w,p[2])
	}
}




