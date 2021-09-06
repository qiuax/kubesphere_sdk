package kapis3_1_1

import (
	"testing"
)

func TestCreateApplication(t *testing.T) {
	ksInfo := NewKubeSphereInfo("admin", "admin", "http://127.0.0.1:32517")
	info, err := ksInfo.GetTokenInfo()
	if err != nil {
		t.Fatal("err:", err.Error())
	}
	ksInfo.Token = info.AccessToken
	request := &CreateApplicationReq{
		Workspaces: "testapi",
		Namespaces: "getapi",
		AppID:      "app-x3q3640n76rw47-store",
		Name:       "nginx-n9ueyu",
		VersionID:  "appv-x795nnrnlx8qmy",
		Conf:       "",
	}
	if err := ksInfo.CreateApplication(request); err != nil {
		t.Error("==============>", err.Error())
	}
}
