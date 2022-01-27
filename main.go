// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/kkdai/LineBotTemplate/imgur"

	//"github.com/kkdai/LineBotTemplate/vendor/github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/kkdai/LineBotTemplate/wow"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var bot *linebot.Client

const (
	port               = "8000"
	channelSecret      = "ef013359722736ac09f50f90e882c27d"
	channelAccessToken = "nVTZGhQzCtStrobItATyfE3evValw6eFQBVmWOisn13jTFriwxiTmaDYObv5SwgGqJZOSkkr9dYoRFZPj+vpUMvzvIsm6VOlM65ccBQUe/Etzda2P52OwNWlmoIdIPTwWxAkfeDYIoi8Pj7m4KUQxAdB04t89/1O/w1cDnyilFU="
)
const (
	meUID    = "U4baaa53fe69d189e272d1aaa4deffe9e"
	feiyoUID = "U90ccc43a7575a173397b9d8d0b5f1e01"
	tataUID  = "Ua5d8aac70a1c84aa4d4c4850f1c0974c"
	tyUID    = "Ub4f2e9497822c8e93a8725be3100fcae"
	sisterID = "U331854135514f7054055c4d3a598cf77"
)

func main() {
	var err error
	//bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	bot, err = linebot.New(channelSecret, channelAccessToken)
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	log.Print("Server Start ...")
	if err = http.ListenAndServe(addr, nil); err == nil {
		log.Print("Server Start ... Error", err)
	}
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		fmt.Println("userId", event.Source.UserID)
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if replyMessages, err := eventHandler(event.Source.UserID, message.Text); err == nil {
					_, err := bot.ReplyMessage(event.ReplyToken, replyMessages...).Do()
					if err != nil {
						log.Print(err)
					}
				} else {
					// do nothing
				}
			}
		}
	}
}

func eventHandler(inUser, inMessage string) (results []linebot.SendingMessage, err error) {

	//err = fmt.Errorf("no event")

	events := strings.Split(inMessage, " ")
	if len(events) == 0 {
		return
	}

	switch events[0] {
	case "!help":
		helpInfo := `目前指令:
		1. 分身 {id} {伺服器} (查詢wow分身)
		2. 抽 (隨機美女圖)
		3. 吃 (隨機美食)
		4. 坦 (隨機梗圖)
		5. 扛 (隨機梗圖)
		`

		results = append(results, linebot.NewTextMessage(helpInfo))
	case "分身": //分身 名字 伺服器
		var characterInfos []string
		if characterInfos, err = wow.QueryCharacterOtherRole(events[1], events[2]); err != nil {
			return
		}

		mergeDatas := strings.Join(characterInfos, "\n\n")
		results = append(results, linebot.NewTextMessage(mergeDatas))

		// for _, data := range characterInfos {
		// 	results = append(results, linebot.NewTextMessage(data))
		// }

	case "抽":
		preview, link := imgur.GetRandAlbumLink(imgur.AlbumGirl)

		// if inUser == feiyoUID {
		// 	preview, link = imgur.GetRandAlbumLink(imgur.AlbumFeiyo)
		// }
		if inUser == tataUID {
			preview, link = imgur.GetRandAlbumLink(imgur.AlbumTATA)
		}
		//imageMsg := linebot.NewImageMessage("https://i.imgur.com/CMp5awi.png", "https://i.imgur.com/CMp5awi.png")

		imageMsg := linebot.NewImageMessage(preview, link)

		results = append(results, imageMsg)

	case "抽正妹":

		//case "抽正妹":
		// imageMsg := linebot.NewVideoMessage("https://v16-webapp.tiktok.com/b93c12d5b1df54a72489c4529f4a2e03/61efd2d7/video/tos/alisg/tos-alisg-pve-0037/38ea2f650270497f93b2cd56a8153a74/?a=1988&br=2942&bt=1471&cd=0%7C0%7C1%7C0&ch=0&cr=0&cs=0&cv=1&dr=0&ds=3&er=&ft=pCjVgag3-InYhx.Oc6&l=202201250436580102440552200605C176&lr=tiktok&mime_type=video_mp4&net=0&pl=0&qs=0&rc=M3I6cDY6ZnI2NzMzODgzNEApZTw4OTdnODs5Nzc5ZWk2O2dkMmxpcjRvamFgLS1kLy1zcy5jLzMyNmE1Ml9eMmMuNjA6Yw%3D%3D&vl=&vr=",
		// 	"https://p16-sign-sg.tiktokcdn.com/aweme/100x100/tiktok-obj/1630275762697218.jpeg?x-expires=1643173200&x-signature=kCcv2bKlmOVOzxWORyUAHAGvySU%3D")
		preview, link := imgur.RandVideo()
		videoMsg := linebot.NewVideoMessage(link, preview)

		results = append(results, videoMsg)
	case "坦", "扛":
		preview, link := imgur.GetRandAlbumLink(imgur.AlbumMeme)
		imageMsg := linebot.NewImageMessage(preview, link)

		results = append(results, imageMsg)
	case "吃":
		preview, link := imgur.GetRandAlbumLink(imgur.AlbumFood)
		imageMsg := linebot.NewImageMessage(preview, link)

		results = append(results, imageMsg)

	default:
		return
	}
	return
}
