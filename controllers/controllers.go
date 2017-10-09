package controllers

import (
	"net/http"
	"fmt"
	"gawkbox-assignment/services"
	"gawkbox-assignment/models"
	"encoding/json"
	"reflect"
)

func GetUserInfo(w http.ResponseWriter, username string) {
	fmt.Println("Getting userInfo for userName ",username)
	var userInfo = services.GetUserInfo(username)
	if reflect.DeepEqual(models.User{}, userInfo){
		w.WriteHeader(http.StatusNotFound)
		var e = fmt.Errorf("Displayname not found")
		prepareResp(nil,e,w)
	} else {
		jsonResp, err := json.Marshal(userInfo)
		prepareResp(jsonResp,err,w)
	}
}
func GetChannelInfo(w http.ResponseWriter, username string) {
	fmt.Println("Getting channel views ",username)
	channel := services.GetChannelViews(username)
	fmt.Println(channel)
	if reflect.DeepEqual(models.Channel{}, channel) {
		w.WriteHeader(http.StatusNotFound)
		var e = fmt.Errorf("User not found")
		prepareResp(nil,e,w)
	}else {
		jsonResp, err := json.Marshal(channel)
		prepareResp(jsonResp, err, w)
	}
}
func prepareResp( b []byte , err error, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write([] byte(b))
}

