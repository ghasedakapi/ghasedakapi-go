package main

import(
"fmt"
"github.com/saeidmaroufi/ghasedakapi-go"
)

func main(){
   api := ghasedakapi.New("76657649624A704B50447571627839795779327566673D3D")
	sender := "30005088"                 
	receptor :="09373576025"
	message := "Hello Go!" 
	if res, err := api.SMS.Send( message,sender, receptor); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(r)
	}

}