package auth

import (
	"encoding/json"
	"kubesphere_sdk/ks"
	"log"
	"testing"
)

func TestAuth(t *testing.T) {
	ksInfo := ks.NewKubeSphereInfo("admin", "Marsone-2021", "http://192.168.1.177:32517")
	info, err := GetTokenInfo(ksInfo)
	if err != nil {
		t.Fatal("err:", err.Error())
	}
	marshal, _ := json.Marshal(info)
	log.Println("result: ", string(marshal))
}
