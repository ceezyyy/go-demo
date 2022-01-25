package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type TabInfo struct {
	TabID    int    `json:"tab_id"`
	Name     string `json:"name"`     // 礼物墙名称
	TabType  int    `json:"tab_type"` // 礼物墙类型 同mysql配置
	TabInfo  string `json:"tab_info"` // 礼物墙额外信息，同mysql配置
	Priority int    `json:"Priority"` // 礼物墙tab优先级，优先级越高，在礼物墙上展示越靠前。
}

func main() {

	var model TabInfo

	data := "{\"tab_id\":1,\"name\":\"\xe6\x8e\xa8\xe8\x8d\x90\",\"tab_type\":0,\"tab_info\":\"\",\"Priority\":999999}"

	reader := strings.NewReader(data)
	decoder := json.NewDecoder(reader)

	decoder.Decode(&model)
	fmt.Println(model)

}
