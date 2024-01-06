package protocol

// SignUp
type ReqSignUp struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
type RespSignUp struct {
	Ret int `json:"ret"` // 结果码 0:成功 1:用户名或密码为空 2:用户名重复或创建失败
}

// Login
type ReqLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
type RespLogin struct {
	Ret   int    `json:"ret"`   // 结果码 0:成功 1:用户名或密码错误 2:登录失败
	Token string `json:"token"` // token
}

// Request Data
type ReqUserData struct {
	UserName string `json:"user_name"`
	Token    string `json:"token"`
	Count uint64 `json:"count"` //数量
}
type RespUserData struct {
	Ret      int    `json:"ret"`       // 结果码 0:成功 1:token校验失败 2:数据为空 3:获取失败
	UserName string `json:"user_name"` // 用户名，不为空
	Data string `json:"data"`	//用户名及交易记录 —— base64+json
}

// Admin
type ReqAdmin struct {
	UserName string `json:"user_name"`
	Token    string `json:"token"`
	Action string `json:"action"`
	Args string `json:"args"` // 
}
type RespAdmin struct {
	Ret      int    `json:"ret"`       // 结果码 0:成功 1:token校验失败 2:失败
}