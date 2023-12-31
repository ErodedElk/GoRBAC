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

func SignUpService(req protocol.ReqSignUp) (resp protocol.RespSignUp) {
	if req.UserName == "" || req.Password == "" {
		resp.Ret = 1
		return
	}

	if err := mysql.CreateAccount(req.UserName, req.Password); err != nil {
		resp.Ret = 2
		log.Errorf("tcp.signUp: mysql.CreateAccount failed. usernam:%s, err:%q", req.UserName, err)
		return
	}
	if err := mysql.CreateProfile(req.UserName, req.NickName); err != nil {
		resp.Ret = 2
		log.Errorf("tcp.signUp: mysql.CreateProfile failed. usernam:%s, err:%q", req.UserName, err)
		return
	}

	resp.Ret = 0
	return
}