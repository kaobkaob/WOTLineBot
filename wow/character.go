package wow

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/kkdai/LineBotTemplate/httphelper"
)

var serverNameMap = make(map[string]string)

func init() {
	serverNameMap["世界之樹"] = "world-tree"
	serverNameMap["亞雷戈斯"] = "arygos"
	serverNameMap["冰霜之刺"] = "frostmane"
	serverNameMap["冰風崗哨"] = "chillwind-point"
	serverNameMap["地獄吼"] = "hellscream"
	serverNameMap["夜空之歌"] = "nightsong"
	serverNameMap["天空之牆"] = "skywall"
	serverNameMap["寒冰皇冠"] = "icecrown"
	serverNameMap["尖石"] = "spirestone"
	serverNameMap["屠魔山谷"] = "demon-fall-canyon"
	serverNameMap["巨龍之喉"] = "dragonmaw"
	serverNameMap["憤怒使者"] = "wrathbringer"
	serverNameMap["日落沼澤"] = "sundown-marsh"
	serverNameMap["暗影之月"] = "shadowmoon"
	serverNameMap["水晶之刺"] = "crystalpine-stinger"
	serverNameMap["狂熱之刃"] = "zealot-blade"
	serverNameMap["眾星之子"] = "queldorei"
	serverNameMap["米奈希爾"] = "menethil"
	serverNameMap["聖光之願"] = "lights-hope"
	serverNameMap["血之谷"] = "bleeding-hollow"
	serverNameMap["語風"] = "whisperwind"
	serverNameMap["銀翼要塞"] = "silverwing-hold"
	serverNameMap["阿薩斯"] = "arthas"
	serverNameMap["雲蛟衛"] = "order-of-the-cloud-serpent"
	serverNameMap["雷鱗"] = "stormscale"
}

func getServerName(inCode string) string {
	for name, code := range serverNameMap {
		if code == inCode {
			return name
		}
	}
	return inCode
}

type characterInfo struct {
	Camp            string `json:"camp"`   //"ALLIANCE"
	Clz             string `json:"clz"`    // "法師"
	Name            string `json:"name"`   // "哈酷納瑪塔塔"
	MythicPlusScore int    `json:"rating"` //: 2385
	//ratingD: [0, 0, 2385.2]
	Server string `json:"server"` //: "shadowmoon"
}

func (c characterInfo) String() string {
	var camp = "部落"
	if c.Camp == "ALLIANCE" {
		camp = "聯盟"
	}

	theResult := []string{
		"角色:" + c.Name,
		"職業:" + c.Clz,
		"M+Score:" + strconv.Itoa(int(c.MythicPlusScore)),
		"伺服器:" + getServerName(c.Server),
		"陣營:" + camp,
	}

	return strings.Join(theResult, "\n")
}
func QueryCharacterOtherRole(inName, inServer string) ([]string, error) {

	url := fmt.Sprintf("https://pricer0606.com/q/%v/%v", url.QueryEscape(inName), serverNameMap[inServer])
	body, err := httphelper.Do("GET", url, nil, nil)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	var theDatas []characterInfo
	err = json.Unmarshal(body, &theDatas)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	if len(theDatas) == 0 {
		return nil, fmt.Errorf("no Character")
	}

	// var theResult strings.Builder
	// for _, info := range theDatas {
	// 	theResult.WriteString(info.String() + "\n")
	// }

	// return theResult.String(), nil

	var theResult []string
	for _, info := range theDatas {
		theResult = append(theResult, info.String())
	}
	return theResult, nil
}
