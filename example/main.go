package main

import (
	"fmt"
	"time"

	"github.com/dtm-labs/dtmdriver"
	pd "github.com/dtm-labs/dtmdriver-polaris"
)

func main() {
	err := dtmdriver.Use(pd.Name)
	fmt.Println("got err", err)
	err = dtmdriver.GetDriver().RegisterService("polaris://0.0.0.0:8080/your.service?namespace=Test",
		"YOUR_TOKEN")
	fmt.Println("got err", err)
	time.Sleep(30 * time.Second)
}
