package kapis3_1_1

import (
	"encoding/json"
	"log"
	"testing"
)

func TestAuth(t *testing.T) {
	ksInfo := NewKubeSphereInfo("admin", "2021", "http://192.168.1.17:32517")
	//ksInfo := NewKubeSphereInfo("admin", "Marsone-2021", "http://159.75.202.72")
	//info, err := ksInfo.GetTokenInfo()
	info, err := ksInfo.GetTokenInfo2()
	if err != nil {
		t.Fatal("err:", err.Error())
	}
	marshal, _ := json.Marshal(info)
	log.Println("resultStr: ", string(marshal))
}
