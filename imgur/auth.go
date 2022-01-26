package imgur

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/kkdai/LineBotTemplate/httphelper"
)

const (
	clinetID     = "a3faf220f21f34b"
	clinetSecret = "0b94c94f23984e6d6fb592623a865023a8afe1a"
	tokenType    = "bearer"
	refreshToken = "68b2fbcc0241da8b7d2bd8e1268994f1673ad41c"
	accessToken  = "0e409e446ed66f4a965af3bc32a5da78cb88d064"
)

type ResponseWarrper struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Status  int         `json:"status"`
}

var client *ImgurClient

type ImgurClient struct {
	AccessToken string
	RefrshToken string
}

func init() {
	client = new(ImgurClient)
	client.refreshToken()
}

func (i ImgurClient) getHeaders() map[string]string {
	var theResult = make(map[string]string)
	theResult["Content-Type"] = "application/x-www-form-urlencoded"
	theResult["Authorization"] = fmt.Sprintf("Bearer %v", accessToken)
	return theResult
}

func (i *ImgurClient) refreshToken() error {

	param := url.Values{}
	param.Set("refresh_token", i.RefrshToken)
	param.Set("client_id", clinetID)
	param.Set("client_secret", clinetSecret)
	param.Set("grant_type", "refresh_token")

	b, err := httphelper.Do("POST", "https://api.imgur.com/oauth2/token", param, i.getHeaders())
	if err != nil {
		return err
	}

	var theData ResponseWarrper
	err = json.Unmarshal(b, &theData)

	return err
}

type ImageInfo struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Link string `json:"link"`
}

func (i ImageInfo) GetLinkInfo() (string, string) {
	if strings.Contains(i.Type, "video") {
		var baseLink = fmt.Sprintf("https://i.imgur.com/%v.%v", i.ID, "jpg")
		return baseLink, i.Link
	} else {
		return i.Link, i.Link
	}
}

func (i ImgurClient) GetAlbumImagesInfo(inAlbumName string) (interface{}, error) {

	b, err := httphelper.Do("GET", fmt.Sprintf("https://api.imgur.com/3/album/%v/images", inAlbumName), nil, i.getHeaders())
	if err != nil {
		return nil, err
	}

	var response ResponseWarrper
	response.Data = []ImageInfo{}

	if err = json.Unmarshal(b, &response); err != nil {
		return nil, err
	}

	if response.Status != http.StatusOK || !response.Success {
		return nil, fmt.Errorf("error happenps")
	}

	return response.Data, err
}
