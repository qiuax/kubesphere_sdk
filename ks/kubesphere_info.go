package ks

// KsInfo kubesphere 的信息
type KsInfo struct {
	Username string // 用户名 初始化注入
	Password string // 密码 初始化注入
	URL      string // 域名/ip  初始化注入
}

type OauthTokenResp struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"` // Bearer
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"` // 7200 秒
}

func NewKubeSphereInfo(username, password, url string) *KsInfo {
	return &KsInfo{
		Username: username,
		Password: password,
		URL:      url,
	}
}
