package kapis3_1_1

import (
	"encoding/json"
	"testing"
)

func TestCreateApplication(t *testing.T) {
	ksInfo := NewKubeSphereInfo("admin", "2021", "http://127.0.0.1:32517")
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

func TestDeleteApplication(t *testing.T) {
	ksInfo := NewKubeSphereInfo("admin", "2021", "http://127.0.0.1:32517")
	info, err1 := ksInfo.GetTokenInfo()
	if err1 != nil {
		t.Fatal("err:", err1.Error())
	}
	ksInfo.Token = info.AccessToken
	req := &DeleteApplicationReq{
		Workspaces:   "testapi",
		Namespaces:   "getapi",
		Cluster:      "default",
		AppClusterID: "rls-wrn6o92jymll49",
	}
	err := ksInfo.DeleteApplication(req)
	if err != nil {
		t.Fatal("----------->", err.Error())
	}
}

func TestGetRunningApplication(t *testing.T) {
	ksInfo := NewKubeSphereInfo("admin", "2021", "http://127.0.0.1:32517")
	info, err1 := ksInfo.GetTokenInfo()
	if err1 != nil {
		t.Fatal("err:", err1.Error())
	}
	ksInfo.Token = info.AccessToken
	req := &GetRunningApplicationReq{
		Workspaces: "testapi",
		Namespaces: "getapi",
		Page:       "",
		PageSize:   "",
	}
	result, err1 := ksInfo.GetRunningApplication(req)
	if err1 != nil {
		t.Fatal("err: ", err1.Error())
	}
	marshal, _ := json.Marshal(result)
	t.Log("=============>", string(marshal))
}
