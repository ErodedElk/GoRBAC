package service

import(
	"server/protocal"
)

func SignUp(v interface{}) interface{} {
	return SignUpService(*v.(*protocol.ReqSignUp))
}

func Login(v interface{}) interface{} {
	return LoginService(*v.(*protocol.ReqLogin))
}

func GetUserData(v interface{}) interface{} {
	return GetUserDataService(*v.(*protocol.ReqUserData))
}

func Admin(v interface{}) interface{} {
	return AdminService(*v.(*protocol.ReqAdmin))
}

func SignUpService(req protocol.ReqSignUp) (resp protocol.RespSignUp) {
	if req.UserName == "" || req.Password == "" {
		resp.Ret = 1
		return
	}
	// if err := mysql.CreateAccount(req.UserName, req.Password); err != nil {
	// 	resp.Ret = 2
	// 	log.Errorf("tcp.signUp: mysql.CreateAccount failed. usernam:%s, err:%q", req.UserName, err)
	// 	return
	// }
	// if err := mysql.CreateProfile(req.UserName, req.NickName); err != nil {
	// 	resp.Ret = 2
	// 	log.Errorf("tcp.signUp: mysql.CreateProfile failed. usernam:%s, err:%q", req.UserName, err)
	// 	return
	// }
	resp.Ret = 0
	return
}

func LoginService(req protocol.ReqLogin) (resp protocol.RespLogin) {
	// ok, err := mysql.LoginAuth(req.UserName, req.Password)
	// if err != nil {
	// 	resp.Ret = 2
	// 	log.Errorf("tcp.login: mysql.LoginAuth failed. usernam:%s, err:%q", req.UserName, err)
	// 	return
	// }
	// //账号或密码不正确.
	// if !ok {
	// 	resp.Ret = 1
	// 	return
	// }
	// token := utils.GetToken(req.UserName)
	// err = redis.SetToken(req.UserName, token, int64(config.TokenMaxExTime))
	// if err != nil {
	// 	resp.Ret = 2
	// 	log.Errorf("tcp.login: redis.SetToken failed. usernam:%s, token:%s, err:%q", req.UserName, token, err)
	// 	return
	// }
	// resp.Ret = 0
	// resp.Token = token
	// log.Infof("tcp.login: login done. username:%s", req.UserName)
	return
}


func GetUserDataService(req protocol.ReqUserData) (resp protocol.RespUserData) {
	return
}

func AdminService(req protocol.ReqAdmin) (resp protocol.RespAdmin) {
	return
}