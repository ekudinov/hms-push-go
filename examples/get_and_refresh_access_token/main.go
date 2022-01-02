package main

import (
	"fmt"
	"time"

	"github.com/ekudinov/hms-push-go/examples/common"
)

func main() {

	client := common.GetPushClient()

	fmt.Printf("%+v", client.GetToken())

	timeout := 10

	fmt.Println("\nWait", timeout, "seconds before refresh.")

	time.Sleep(time.Duration(timeout) * time.Second)

	client.RefreshToken()

	fmt.Println("\nRefresh token ok")

	fmt.Printf("%+v", client.GetToken())

}
