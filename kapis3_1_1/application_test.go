package kapis3_1_1

import (
	"encoding/base64"
	"encoding/json"
	"testing"
)

func TestCreateApplication(t *testing.T) {
	ksInfo := NewKubeSphereInfo("admin", "2021", "http://192.168.111.177:32517")
	info, err := ksInfo.GetTokenInfo()
	if err != nil {
		t.Fatal("err:", err.Error())
	}
	ksInfo.Token = info.AccessToken
	decodeString, _ := base64.StdEncoding.DecodeString("IyBEZWZhdWx0IHZhbHVlcyBmb3IgZWRnZS4KIyBUaGlzIGlzIGEgWUFNTC1mb3JtYXR0ZWQgZmlsZS4KIyBEZWNsYXJlIHZhcmlhYmxlcyB0byBiZSBwYXNzZWQgaW50byB5b3VyIHRlbXBsYXRlcy4KCnJlcGxpY2FDb3VudDogMQoKaW1hZ2U6CiAgd2ViOgogICAgcmVwb3NpdG9yeTogZG9ja2VyLmlvL21hcnNvbmUwMDEvbWFyc29uZQogICAgdGFnOiAiMS4wLjAiCiAgICBob3N0UG9ydDogODAwMAogIG15c3FsOgogICAgcmVwb3NpdG9yeTogcmVnaXN0cnkuY24tc2hlbnpoZW4uYWxpeXVuY3MuY29tL21pbnRlZGdlL21pbnRlZGdlLW1xbDUuNwogICAgdGFnOiAiMS4wLjAiIAogICAgZGF0YURpcjogIi9kYXRhL215c3FsIgogICAgaG9zdFBvcnQ6IDMzMDYKCm5vZGVOYW1lOiAiZWRnZW5vZGUtcmR0MiIK")
	request := &CreateApplicationReq{
		Workspaces: "mint",
		Namespaces: "web",
		AppID:      "app-8l1mqy74mo3qkw-store",
		Name:       "edge-n9ueyu",
		VersionID:  "appv-mn6jlkx730pk7p",
		Conf:       string(decodeString),
	}
	if err := ksInfo.CreateApplication(request); err != nil {
		t.Error("==============>", err.Error())
	}
}

func TestDeleteApplication(t *testing.T) {
	ksInfo := NewKubeSphereInfo("admin", "2021", "http://192.168.1111.177:32517")
	info, err1 := ksInfo.GetTokenInfo()
	if err1 != nil {
		t.Fatal("err:", err1.Error())
	}
	ksInfo.Token = info.AccessToken
	req := &DeleteApplicationReq{
		Workspaces:   "testapi",
		Namespaces:   "getapi",
		Cluster:      "default",
		AppClusterID: "rls-wown7y2y0qll49",
	}
	err := ksInfo.DeleteApplication(req)
	if err != nil {
		t.Fatal("----------->", err.Error())
	}
}

func TestGetRunningApplication(t *testing.T) {
	ksInfo := NewKubeSphereInfo("admin", "12021", "http://192.168.1111.177:32517")
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
func TestGetApplicationFile(t *testing.T) {
	ksInfo := NewKubeSphereInfo("admin", "2021", "http://192.168.1111.177:32517")
	info, err1 := ksInfo.GetTokenInfo()
	if err1 != nil {
		t.Fatal("err:", err1.Error())
	}
	ksInfo.Token = info.AccessToken

	ksInfo.GetApplicationFile("app-8l1mqy74mo3qkw-store", "appv-mn6jlkx730pk7p")

}

func TestGetRunningApplicationByName(t *testing.T) {
	ksInfo := NewKubeSphereInfo("admin", "Marsone-2021", "http://192.168.1.177:32517")
	info, err1 := ksInfo.GetTokenInfo()
	if err1 != nil {
		t.Fatal("err:", err1.Error())
	}
	ksInfo.Token = info.AccessToken
	req := &GetRunningApplicationReq{
		//Workspaces:      "mint",
		//Namespaces:      "edge",
		//ApplicationName: "eba",
		Workspaces:      "testapi",
		Namespaces:      "getapi",
		ApplicationName: "nginx",
	}
	result, _ := ksInfo.GetRunningApplicationByName(req)
	marshal, _ := json.Marshal(result)
	t.Log("=============>", string(marshal))
}
