package kapis3_1_1

import (
	"encoding/json"
	"log"
	"testing"
)

func TestAuth(t *testing.T) {
	ksInfo := NewKubeSphereInfo("admin", "Marsone-2021", "http://192.168.1.177:32517")
	info, err := ksInfo.GetTokenInfo()
	if err != nil {
		t.Fatal("err:", err.Error())
	}
	marshal, _ := json.Marshal(info)
	log.Println("resultStr: ", string(marshal))
}
