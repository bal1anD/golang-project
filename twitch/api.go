package twitch

import (
	"fmt"
	"net/http"
	"time"
	"encoding/json"
	"gawkbox-assignment/models"
	"io/ioutil"
)

func init() {
	fmt.Println("Initializing Twitch API...")
}


var TWITCH_CLIENT_ID string = "oaytvqxrucu959iqc000gcgykwetbl"
var TWITCH_BASE_URL = "https://api.twitch.tv/kraken"


func GetUserInfo(username string) (models.User, error) {

	var twitchUser models.User = models.User{}

	var url string = TWITCH_BASE_URL+"/users"+"?login="+username;

	client := &http.Client{Timeout: 10 * time.Second}


	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error in request ",err)
		return twitchUser,err;
	}


	req.Header.Add("Accept", "application/vnd.twitchtv.v5+json")
	req.Header.Add("Client-ID", TWITCH_CLIENT_ID)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error in request ",err)
		return twitchUser,err;
	}
	defer resp.Body.Close()
	if resp.StatusCode==200 {
		twitchUser = getJson(resp)
	}

	return twitchUser,err
}


func GetUserChannelInfo(userId string) (models.Channel, error) {
	var twitchChannel models.Channel = models.Channel{}

	var url string = TWITCH_BASE_URL+"/channels/"+userId;

	client := &http.Client{Timeout: 10 * time.Second}


	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error in request ",err)
		return twitchChannel,err;
	}


	req.Header.Add("Accept", "application/vnd.twitchtv.v5+json")
	req.Header.Add("Client-ID", TWITCH_CLIENT_ID)


	resp, err := client.Do(req)
	if err != nil {
		return twitchChannel,err;
	}
	defer resp.Body.Close()
	if resp.StatusCode==200 {
		twitchChannel = getChannel(resp)
	}

	return twitchChannel,err
}

func getChannel(r *http.Response)  models.Channel {
	var twitchUserModel *models.Channel = &models.Channel{}
	defer r.Body.Close()
	bodyInBytes, _ := ioutil.ReadAll(r.Body);
	err := json.Unmarshal(bodyInBytes, twitchUserModel)
	if err != nil {
		fmt.Println("Something went wrong")
		return *twitchUserModel
	}

	return *twitchUserModel

}

func getJson(r *http.Response)  models.User {
	var twitchUserModel *models.User = &models.User{}
	defer r.Body.Close()
	bodyInBytes, _ := ioutil.ReadAll(r.Body);
	err := json.Unmarshal(bodyInBytes, twitchUserModel)
	if err != nil {
		fmt.Println("Something went wrong")
		return *twitchUserModel
	}

	return *twitchUserModel

}
