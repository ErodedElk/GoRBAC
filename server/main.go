package main
import(
	"server/server"
	"server/config"
	"fmt"
	"server/service"
)

func main() {
	server := server.Server()
	err := config.Parse_config("./example.yaml")
	if err!=nil{
		
	}
	server.Register("SignUp", service.SignUp, service.SignUpService)
	server.Register("Login", service.Login, service.LoginService)
	server.Register("GetUserData", service.GetUserData, service.GetUserDataService)
	server.Register("Admin", service.Admin, service.AdminService)
	fmt.Println("%v",server)

	server.ListenAndServe(config.Origin_Conf.Serv.Addr)
}