package main
import(
	"server/server"
	"server/config"
	"fmt"
)

func main() {
	server := server.Server()
	fmt.Println("%v",server)
	err := config.Parse_config("./example.yaml")
	if err!=nil{
		
	}
	fmt.Println(config.Data)
}