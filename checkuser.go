package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

//文件路径
const (
	on = iota
	tw
	three
	four
	fiv
	path = "/Users/codoon/Desktop/query_result.csv"

)

func toString(ty int) string {
	switch ty {
	case on:
		return "0~20"
	case tw:
		return "20~40"
	case three:
		return "40~60"
	case four:
		return "60~80"
	case fiv:
		return ">=80"
	}
	return ""
}

type Rsp struct {
	Count     int         `json:"count"` //统计总数量
	ScoreType []ScoreType `json:"score_type"`
	Dangerous []Dangerous `json:"dangerous"`
}

type ScoreType struct {
	Type       string `json:"type"`
	Count      int    `json:"count"`
	Proportion string `json:"proportion"`
}

//高危具体信息
type Dangerous struct {
	UserId     string         `json:"user_id"`
	FinalScore int            `json:"final_score"`
	Status string `json:"status"`
	CreateTime string         `json:"create_time"`
	Score      []ScoreContent `json:"score_content"`
}

//得分具体细节
type ScoreContent struct {
	Score       int    `json:"score"`
	Description string `json:"description"`
}

type Content struct {
	FinalScore    int         `json:"final_score"`
	PolicySetName interface{} `json:"policy_set_name"`
	RiskType      interface{} `json:"risk_type"`
	FinalDecision string `json:"final_decision"`
	PolicySet     interface{} `json:"policy_set"`
	SpendTime     interface{} `json:"spend_time"`
	Success       interface{} `json:"success"`
	PolicyName    interface{} `json:"policy_name"`
	SeqId         interface{} `json:"seq_id"`
	HitRules      []HitRules  `json:"hit_rules"`
}

type HitRules struct {
	Score      int    `json:"score"`
	Decision   string `json:"decision"`
	Name       string `json:"name"`
	Id         string `json:"id"`
	Uuid       string `json:"uuid"`
	ParentUuid string `json:"parentUuuid"`
}

func main()  {
	ReadFile()
}

func ReadFile()  {
	file,err:=os.Open(path)
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer file.Close()

	csvReader:=csv.NewReader(file)
	csvReader.Read()//忽略第一行
	rsp:=GetMessage(csvReader)
	fmt.Println(rsp.Count)
	fmt.Println(rsp.ScoreType)
	for _,v:=range rsp.Dangerous{
		fmt.Println(v)
	}
}

//文件解析得出比例以及超过80分的id，时间
func GetMessage(csvReader *csv.Reader)Rsp {
	data:=Rsp{}
	dang :=make([]Dangerous,0)
	array:=make([]int,5)
	sum:=0
	for{
		record,err:=csvReader.Read()
		if err==io.EOF{
			break
		}else if err!=nil{
			fmt.Println("GetMessage error",err.Error())
			return data
		}
		message:=Convert(record[4])
		v:=message.FinalScore
			switch  {
			case v < 20:
				array[0]++
			case v >=20&&v<40:
				array[1]++
			case v >=40&&v<60:
				array[2]++
			case v >=60&&v<80:
				array[3]++
			case v >=80:
				array[4]++
			}
			//大于80分具体分数信息加载
			if v>=60&&v<80{
				mes:=Dangerous{
					Status: message.FinalDecision,
					FinalScore: message.FinalScore,
					Score: GetDangerousContent(message),
					UserId: record[1],
					CreateTime: record[5],
				}
				dang = append(dang,mes)
			}
			sum++
	}
	data.ScoreType = GetScoreType(array,sum)
	data.Dangerous = dang
	data.Count = sum
	return data
}

//Dangerous具体分数加载
func GetDangerousContent(message Content)[]ScoreContent  {
	var score []ScoreContent
	for _,v:=range message.HitRules{
		score = append(score,ScoreContent{
			Score: v.Score,
			Description: v.Name,
		})
	}
	return score
}

//计算比例
func GetScoreType(array []int,sum int)[]ScoreType {
	res:=make([]ScoreType,0)
	for i,v:=range array{
		score:=ScoreType{}
		score.Count = v
		score.Proportion = fmt.Sprintf("%v%%",100*float32(v)/float32(sum))
		score.Type =toString(i)
		res = append(res,score)
	}
	return res
}

//content 字段解析
func Convert(content string)(message Content)  {
	err:=json.Unmarshal([]byte(content),&message)
	if err!=nil{
		fmt.Println("Convert Unmarshal error",err.Error())
		return
	}
	return message
}
