package services

import (
	"gawkbox-assignment/twitch"
	"gawkbox-assignment/store"

	"reflect"
	"gawkbox-assignment/models"
	"fmt"
)

func GetUserInfo(username string) models.User {
	userInfo,_ := twitch.GetUserInfo(username)
	if !reflect.DeepEqual(models.User{}, userInfo) {
		store.Put(username,userInfo)
	}
	return userInfo
}


func GetChannelViews(username string ) models.Channel {
	var channel models.Channel = models.Channel{}

	user := store.Get(username);
	var userId string
	if !reflect.DeepEqual(models.User{}, user) {
		fmt.Println("Cache hit")
		userId = user.Id
	} else {
		fmt.Println("Cache miss")
		user,_:= twitch.GetUserInfo(username)
		if !reflect.DeepEqual(models.User{}, user) {
			store.Put(username,user)
		} else {
			return channel
		}
		userId = user.Id
	}

	channel,_ = twitch.GetUserChannelInfo(userId)
	return  channel
}

