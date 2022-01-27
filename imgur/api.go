package imgur

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

var videoURLs = make(map[string]string, 0) //video, pic
func init() {
	videoURLs["https://v16-webapp.tiktok.com/317365a4ac21529e25ab5a8a4d822bac/61f1dda0/video/tos/alisg/tos-alisg-pve-0037/38ea2f650270497f93b2cd56a8153a74/?a=1988&br=2942&bt=1471&cd=0%7C0%7C1%7C0&ch=0&cr=0&cs=0&cv=1&dr=0&ds=3&er=&ft=pCjVgag3-InfQ6pOc6&l=202201261747310102450622022047CF9E&lr=tiktok&mime_type=video_mp4&net=0&pl=0&qs=0&rc=M3I6cDY6ZnI2NzMzODgzNEApZTw4OTdnODs5Nzc5ZWk2O2dkMmxpcjRvamFgLS1kLy1zcy5jLzMyNmE1Ml9eMmMuNjA6Yw%3D%3D&vl=&vr="] = "https://i.imgur.com/2STfnWI.jpg"
	videoURLs["https://v16-webapp.tiktok.com/2919d5b42ac3ea6abf0af5467214ecc2/61f1de86/video/tos/alisg/tos-alisg-pve-0037c001/976ab42fb8d14550a1bb161c14ac3752/?a=1988&br=5996&bt=2998&cd=0%7C0%7C1%7C0&ch=0&cr=0&cs=0&cv=1&dr=0&ds=3&er=&ft=pCjVgag3-InQi6pOc6&l=20220126175121010245242107014802BD&lr=tiktok&mime_type=video_mp4&net=0&pl=0&qs=0&rc=ajxmOGd1anhrdzMzMzgzM0ApZjc8NThkZTs7NzQ6NzVlZmdpaWpwY2k0Nm9fLS1hLzRzczRiMDZfNjVjYWM2NC9hMC46Yw%3D%3D&vl=&vr="] = "https://i.imgur.com/2STfnWI.jpg"
	videoURLs["https://v16-webapp.tiktok.com/21a1ab4d12dd37f26565130cc89ceb37/61f1dec2/video/n/v0102/1efb9252f7054ccda83443afe699b0c0/?a=1988&br=3438&bt=1719&cd=0%7C0%7C1%7C0&ch=0&cr=0&cs=0&cv=1&dr=0&ds=3&er=&ft=pCjVgag3-In.y6pOc6&l=202201261752220102440750480E4982FB&lr=tiktok&mime_type=video_mp4&net=0&pl=0&qs=0&rc=MzRlcnJmdng5bzMzaTgzM0ApNzhoOGY3ZTw4N2g4NTRoOGdoLXJvMm81aWFfLS1fLzRzc2ExXjEuLi0wNTNjMy1jMGE6Yw%3D%3D&vl=&vr="] = "https://i.imgur.com/2STfnWI.jpg"
	videoURLs["https://v16-webapp.tiktok.com/c0a97f3dc698237e2618df18d4ba18c5/61f1dcfd/video/tos/alisg/tos-alisg-pve-0037/01efa57d2ed14ae09d0edd83953735a9/?a=1988&br=2700&bt=1350&cd=0%7C0%7C1%7C0&ch=0&cr=0&cs=0&cv=1&dr=0&ds=3&er=&ft=pCjVgag3-InJq6pOc6&l=2022012617445301024504010110468302&lr=tiktok&mime_type=video_mp4&net=0&pl=0&qs=0&rc=M3NxPDQ6Zjt0OTMzODgzNEApNzVnNjU8N2VnNzg1NWhlaGdoYGJwcjRnMDRgLS1kLy1zczMyLmM1Yl4vLWM1MjUuXjY6Yw%3D%3D&vl=&vr="] = "https://i.imgur.com/2STfnWI.jpg"
	// videoURLs["https://v16-webapp.tiktok.com/ccb85dc9a096b22c729821cd72fee141/61f17587/video/tos/alisg/tos-alisg-pve-0037c001/976ab42fb8d14550a1bb161c14ac3752/?a=1988&br=5996&bt=2998&cd=0%7C0%7C1%7C0&ch=0&cr=0&cs=0&cv=1&dr=0&ds=3&er=&ft=pCjVgag3-Inm2HpOc6&l=20220126102323010245248003270DF592&lr=tiktok&mime_type=video_mp4&net=0&pl=0&qs=0&rc=ajxmOGd1anhrdzMzMzgzM0ApZjc8NThkZTs7NzQ6NzVlZmdpaWpwY2k0Nm9fLS1hLzRzczRiMDZfNjVjYWM2NC9hMC46Yw%3D%3D&vl=&vr="] = "https://i.imgur.com/2STfnWI.jpg"
	// videoURLs["https://v16-webapp.tiktok.com/2c292760240c5d6cedbc9cd773745545/61f1756c/video/tos/alisg/tos-alisg-pve-0037/01efa57d2ed14ae09d0edd83953735a9/?a=1988&br=2700&bt=1350&cd=0%7C0%7C1%7C0&ch=0&cr=0&cs=0&cv=1&dr=0&ds=3&er=&ft=pCjVgag3-In12HpOc6&l=202201261023000102450401060B0D2D63&lr=tiktok&mime_type=video_mp4&net=0&pl=0&qs=0&rc=M3NxPDQ6Zjt0OTMzODgzNEApNzVnNjU8N2VnNzg1NWhlaGdoYGJwcjRnMDRgLS1kLy1zczMyLmM1Yl4vLWM1MjUuXjY6Yw%3D%3D&vl=&vr="] = "https://i.imgur.com/2STfnWI.jpg"
}

const AlbumMeme string = "wsxIIJ2"
const AlbumGirl string = "5Ae8BSu"
const AlbumFeiyo string = "9unQFQS"
const AlbumTATA string = "fIQ5bLy"
const AlbumFood string = "6n5GHkV"

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
