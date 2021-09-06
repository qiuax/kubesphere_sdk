package kapis3_1_1

import (
	"encoding/json"
	"log"
	"testing"
)

func TestAuth(t *testing.T) {
	ksInfo := NewKubeSphereInfo("admin", "admin", "http://127.0.0.1:32517")
	info, err := ksInfo.GetTokenInfo()
	if err != nil {
		t.Fatal("err:", err.Error())
	}
	marshal, _ := json.Marshal(info)
	log.Println("resultStr: ", string(marshal))
}
