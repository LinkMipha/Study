package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func main()  {
	str:="hello"
	encode:=hex.EncodeToString([]byte(str))
	fmt.Println(encode)

	sta,_:=hex.DecodeString(encode)
	fmt.Println(string(sta))
	//<-make(chan struct{})

	language := make(map[string]string)
	err := json.Unmarshal([]byte("{\"en\":\"新手欢迎礼包En\",\"zh\":\"新手欢迎礼包\",\"zh\":\"\",\"ja\":\"\",\"de\":\"\",\"es\":\"西语--新手礼包\",\"pt\":\"\"}	"), &language)
	fmt.Println(err,language)
	fmt.Printf("map : %v",language)

	testmap:=make(map[string]string)
	json.Unmarshal([]byte("{}"),&testmap)
	//json.Unmarshal([]byte("{\"desc\":\"备注\",\"identification\":\"身份证号\",\"other_user_info0\":\"性别\",\"other_user_info1\":\"紧急联系人\",\"other_user_info2\":\"紧急联系人电话\",\"other_user_info3\":\"所在城市\",\"other_user_info4\":\"近14天是否去过高风险区\",\"other_user_info5\":\"乘坐免费大巴还是自驾前往\",\"phone\":\"手机号\",\"real_name\":\"真实姓名\"}"),&testmap)
	fmt.Println(err)
	fmt.Printf("test:%v,%d",testmap,len(testmap))
}
