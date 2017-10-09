package store

import "gawkbox-assignment/models"

var USERNAME_ID_CACHE map[string]models.User

func init() {
	USERNAME_ID_CACHE = make(map[string]models.User)
}



func Put(key string, userRecord models.User) {
	USERNAME_ID_CACHE[key] = userRecord
}

func Get(key string) models.User {
	return USERNAME_ID_CACHE[key]
}