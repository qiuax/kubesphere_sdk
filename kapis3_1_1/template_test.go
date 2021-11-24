package kapis3_1_1

import (
	"encoding/json"
	"log"
	"testing"
)

func TestKSInfo_GetActiveAppTemplates(t *testing.T) {
	ksInfo := NewKubeSphereInfo("admin", "2021", "http://127.0.0.1:32517")
	info, err1 := ksInfo.GetTokenInfo2()
	if err1 != nil {
		t.Fatal("err:", err1.Error())
	}
	ksInfo.Token = info.AccessToken
	req := &GetAppTemplatesReq{
		Page:     "",
		PageSize: "",
	}
	result, err := ksInfo.GetActiveAppTemplates(req)
	if err != nil {
		t.Fatal("=========>", err.Error())
	}
	marshal, _ := json.Marshal(result)
	log.Println("------>", string(marshal))

}
