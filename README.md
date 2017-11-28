# ghasedakapi-go
  Ghasedakapi Go library
  
 ## Installation
   go get github.com/ghasedakapi/ghasedakapi-go
 ## How to use?   
  On the import section of your Go project:
  
  import (
    ...
    "github.com/ghasedakapi/ghasedakapi-go"
    ...
  )
## Usage
 
 ### send SMS
  ```golang
package main
import(
"fmt"
"github.com/ghasedakapi/ghasedakapi-go"
)
const (
	ApiKey  = "your apikey"
)
func main(){
  c := ghasedakapi.New(ApiKey)
	lineNumber := "sender number"                 
	receptor := []string{"receptor1","receptor2"}
	message := "test message" 
	if res, err:= c.SMS.Send( message,lineNumber, receptor); err != nil {
		fmt.Println(err)
		} else {
			for _, r := range res.Items {
				fmt.Println(r)
			}
	}
}
```

 ### get account information
 ```golang
  package main
  import(
    "fmt"
    "github.com/ghasedakapi/ghasedakapi-go"
  )
  const (
	  ApiKey  = "your apikey"
  )
  func main(){
    c := ghasedakapi.New(ApiKey)
	  if result, err:= c.Account.GetInfo(); err != nil {
			fmt.Println(err)
		}
	  else {
		fmt.Println(result.Items.Balance)	
		fmt.Println(result.Items.Expire)	
	}
}



