package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Name string `json:"name"`
	DisplayName string `json:"display_name"`
	Bio string `json:"bio"`
	AccCreationDate int64 `json:"created_at"`
	Id string `json:"_id"`
}

type Channel struct {
	Views int `json:"views"`
	Followers int `json:"followers"`
	Game string `json:"game"`
	Status bool `json:"streamingStatus"`
	Language string `json:"language"`
	Id string `json:"_id"`

}

func (channel *Channel) UnmarshalJSON(b []byte) error {

	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	for key, value := range objMap {

		if key == "_id" {
			if value != nil {
				var val string
				err = json.Unmarshal(*value, &val)
				if err != nil {
					return err
				}
				channel.Id = val
			}

		}
		if key == "game" {
			if value != nil {
				var val string
				err = json.Unmarshal(*value, &val)
				if err != nil {
					return err
				}
				channel.Game = val
			}

		}
		if key == "language" {
			if value != nil {
				var val string
				err = json.Unmarshal(*value, &val)
				if err != nil {
					return err
				}
				channel.Language = val
			}

		}
		if key == "views" {
			if value != nil {
				var val int
				err = json.Unmarshal(*value, &val)
				if err != nil {
					return err
				}
				channel.Views = val
			}

		}
		if key == "followers" {
			if value != nil {
				var val int
				err = json.Unmarshal(*value, &val)
				if err != nil {
					return err
				}
				channel.Followers = val
			}

		}
		if key == "status" {
			if value != nil {
				channel.Status = true
			} else {
				channel.Status = false
			}

		}

	}
	return nil
}

func (user *User) UnmarshalJSON(b []byte) error {

	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	var rawUserMsg []*json.RawMessage
	err = json.Unmarshal(*objMap["users"], &rawUserMsg)
	if err != nil {
		return err
	}


	var m map[string]string
	for _, rawMessage := range rawUserMsg {
		err = json.Unmarshal(*rawMessage, &m)
		if err != nil {
			return err
		}
		if m["_id"] != "" {
			user.Id = m["_id"]
		}
		if m["created_at"] != "" {
			user.AccCreationDate,_ = timeFromStr(m["created_at"])
		}
		if m["bio"] != "" {
			user.Bio = m["bio"]
		}
		if m["name"] != "" {
			user.Name = m["name"]
		}
		if m["display_name"] != "" {
			user.DisplayName = m["display_name"]
		}
	}
	return nil
}

func timeFromStr(timeStr string) (int64,error) {
	t, err := time.Parse(time.RFC3339Nano, timeStr)

	if err != nil {
		fmt.Println("Error in formatting time",err)
		return -1,err
	}
	return t.UnixNano()/1000000,err
}