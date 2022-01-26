package imgur

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

var videoURLs = make(map[string]string, 0) //video, pic
func init() {
	videoURLs["https://v16-webapp.tiktok.com/cc283c9a05c0e2d93d1876c2a82375c8/61f1739e/video/tos/alisg/tos-alisg-pve-0037/38ea2f650270497f93b2cd56a8153a74/?a=1988&br=2942&bt=1471&cd=0%7C0%7C1%7C0&ch=0&cr=0&cs=0&cv=1&dr=0&ds=3&er=&ft=pCjVgag3-In4cgpOc6&l=20220126101513010245242197200C200E&lr=tiktok&mime_type=video_mp4&net=0&pl=0&qs=0&rc=M3I6cDY6ZnI2NzMzODgzNEApZTw4OTdnODs5Nzc5ZWk2O2dkMmxpcjRvamFgLS1kLy1zcy5jLzMyNmE1Ml9eMmMuNjA6Yw%3D%3D&vl=&vr="] = "https://i.imgur.com/2STfnWI.jpg"
	//videoURLs["https://v16-webapp.tiktok.com/9fbd1337a9ec912d9b9c773551887e9/61f174f1/video/tos/alisg/tos-alisg-pve-0037/af2f751c716e4ee695a156a231d5c9f0/?a=1988&br=3754&bt=1877&cd=0%7C0%7C1%7C0&ch=0&cr=0&cs=0&cv=1&dr=0&ds=3&er=&ft=pCjVgag3-InnzHpOc6&l=20220126102050010245244178160D59A6&lr=tiktok&mime_type=video_mp4&net=0&pl=0&qs=0&rc=anRpNjpya2k6NDMzaTgzM0ApNDw7ODc7OjtoNzY6NzU1NGdgMGFvNDNebWRgLS00LzRzczBeXjFhLzViYl4yNC8xMDI6Yw%3D%3D&vl=&vr="] = "https://i.imgur.com/2STfnWI.jp"
	videoURLs["https://v16-webapp.tiktok.com/91e8f451cb8e63dd640501089156ad3a/61f1744c/video/tos/alisg/tos-alisg-pve-0037/d7bad96cdd9248ed9428d1f427e69a83/?a=1988&br=4278&bt=2139&cd=0%7C0%7C1%7C0&ch=0&cr=0&cs=0&cv=1&dr=0&ds=3&er=&ft=pCjVgag3-InmjgpOc6&l=20220126101803010245136209230B6F3E&lr=tiktok&mime_type=video_mp4&net=0&pl=0&qs=0&rc=am82aTo6ZnYzNzMzODgzNEApNmQ1PDk3M2U1NzM5PDw0O2cvZ2tpcjQwMG1gLS1kLy1zcy5gMS9gXy01Xi5fYjRfYmI6Yw%3D%3D&vl=&vr="] = "https://i.imgur.com/2STfnWI.jpg"
	videoURLs["https://v16-webapp.tiktok.com/ef3032445266048891e5761b1512267/61f1758d/video/tos/alisg/tos-alisg-pve-0037c001/4b6e3f7fcf924806b7cf66389960adbb/?a=1988&br=4022&bt=2011&cd=0%7C0%7C1%7C0&ch=0&cr=0&cs=0&cv=1&dr=0&ds=3&er=&ft=pCjVgag3-InK2HpOc6&l=20220126102331010244069074240D86A5&lr=tiktok&mime_type=video_mp4&net=0&pl=0&qs=0&rc=M3dmOmw5b252dzMzOzgzM0ApO2Y2Nzg6O2VnNzQ0ZjszaWdjLWVfaS9zZWhfLS1gLzRzc2JfLTIuX2FfLzJjLzE0YDU6Yw%3D%3D&vl=&vr="] = "https://i.imgur.com/2STfnWI.jpg"
	videoURLs["https://v16-webapp.tiktok.com/ccb85dc9a096b22c729821cd72fee141/61f17587/video/tos/alisg/tos-alisg-pve-0037c001/976ab42fb8d14550a1bb161c14ac3752/?a=1988&br=5996&bt=2998&cd=0%7C0%7C1%7C0&ch=0&cr=0&cs=0&cv=1&dr=0&ds=3&er=&ft=pCjVgag3-Inm2HpOc6&l=20220126102323010245248003270DF592&lr=tiktok&mime_type=video_mp4&net=0&pl=0&qs=0&rc=ajxmOGd1anhrdzMzMzgzM0ApZjc8NThkZTs7NzQ6NzVlZmdpaWpwY2k0Nm9fLS1hLzRzczRiMDZfNjVjYWM2NC9hMC46Yw%3D%3D&vl=&vr="] = "https://i.imgur.com/2STfnWI.jpg"
	videoURLs["https://v16-webapp.tiktok.com/2c292760240c5d6cedbc9cd773745545/61f1756c/video/tos/alisg/tos-alisg-pve-0037/01efa57d2ed14ae09d0edd83953735a9/?a=1988&br=2700&bt=1350&cd=0%7C0%7C1%7C0&ch=0&cr=0&cs=0&cv=1&dr=0&ds=3&er=&ft=pCjVgag3-In12HpOc6&l=202201261023000102450401060B0D2D63&lr=tiktok&mime_type=video_mp4&net=0&pl=0&qs=0&rc=M3NxPDQ6Zjt0OTMzODgzNEApNzVnNjU8N2VnNzg1NWhlaGdoYGJwcjRnMDRgLS1kLy1zczMyLmM1Yl4vLWM1MjUuXjY6Yw%3D%3D&vl=&vr="] = "https://i.imgur.com/2STfnWI.jpg"
}

const AlbumMeme = "Meme"
const AlbumGirl = "5Ae8BSu"

func GetRandAlbumLink(inAlbumName string) (string, string) {

	datas, err := client.GetAlbumImagesInfo(inAlbumName)
	if err != nil {
		return "", ""
	}

	b, err := json.Marshal(datas)
	if err != nil {
		fmt.Println(err.Error())
		return "", ""
	}

	var theImageInfoList []ImageInfo
	err = json.Unmarshal(b, &theImageInfoList)
	if err != nil {
		fmt.Println(err.Error())
		return "", ""
	}

	if len(theImageInfoList) == 0 {
		return "", ""
	}

	return theImageInfoList[rand.Intn(len(theImageInfoList))].GetLinkInfo()
}

func RandVideo() (string, string) {
	index := rand.Intn(len(videoURLs))

	i := 0
	for link, preview := range videoURLs {
		if i == index {
			return preview, link
		}
		i++
	}
	return "", ""
}
