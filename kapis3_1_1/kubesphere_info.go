package kapis3_1_1

// ksInfo kubesphere 的信息
type KSInfo struct {
	Username string // 用户名 初始化注入
	Password string // 密码 初始化注入
	URL      string // 域名/ip  初始化注入
	Token    string
}

type OauthTokenResp struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"` // Bearer
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"` // 7200 秒
}

func NewKubeSphereInfo(username, password, url string) *KSInfo {
	return &KSInfo{
		Username: username,
		Password: password,
		URL:      url,
	}
}
