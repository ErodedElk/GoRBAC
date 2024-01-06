package main

import(
	"fmt"
	"client/client"
	"client/protocal"
	"strings"
	"bufio"
	// "io"
	// "net/http"
	"os"
)


var rpcClient client.RPCClient
var cmd string
func main(){
	inputReader := bufio.NewReader(os.Stdin)
	rpcClient, err := client.Client(2000, "0.0.0.0:8099")
	fmt.Println(rpcClient)
	fmt.Println(err)

	for {
		fmt.Printf("$ ")
		tmp, err := inputReader.ReadString('\n')
		if err != nil {
			os.Exit(1)
		}
		cmd = strings.Replace(tmp,"\n","",-1)
		go handle_cmd(cmd)
		cmd=""
	}

}

func handle_cmd(input string){
	s := strings.Split(input," ")
	fmt.Printf("%v",s)
	switch s[0]{
	case "exit":
		os.Exit(0)
	case "signup":
		req := protocol.ReqSignUp{
			UserName: s[1],
			Password: s[2],
		}
		resp := protocol.RespSignUp{}
		rpcClient.Call("SignUp",req,&resp)
		fmt.Println(resp)
	}


}


// func SignUp(rw http.ResponseWriter, req *http.Request) {
// 	// 处理http post方法.
// 	if req.Method == "POST" {
// 		//获取请求各个字段值.
// 		userName := req.FormValue("username")
// 		password := req.FormValue("password")
// 		nickName := req.FormValue("nickname")

// 		if userName == "" || password == "" {
// 			rw.Write([]byte("Username and password couldn't be NULL!"))
// 			return
// 		}
// 		fmt.Printf("userName = %s, password = %s,nickName = %s\n", userName, password, nickName)
// 		req := protocol.ReqSignUp{
// 			UserName: userName,
// 			Password: password,
// 			NickName: nickName,
// 		}
// 		resp := protocol.RespSignUp{}
// 		//调用远程rpc服务, 将数据存入到数据库.
// 		if err := rpcClient.Call("SignUp", req, &resp); err != nil {
// 			log.Errorf("http.SignUp: Call SignUp failed. username:%s, err:%q", userName, err)
// 			rw.Write([]byte("创建账号失败！"))
// 			return
// 		}

// 		switch resp.Ret {
// 		case 0:
// 			rw.Write([]byte("创建账号成功！"))
// 		case 1:
// 			rw.Write([]byte("用户名或密码错误！"))
// 		default:
// 			rw.Write([]byte("创建账号失败！"))
// 		}
// 		log.Infof("http.SignUp: SignUp done. username:%s, ret:%d", userName, resp.Ret)
// 	}
// }