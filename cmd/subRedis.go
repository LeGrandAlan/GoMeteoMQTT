package main

import (
	"./pubsub/subscribers"
	"./pubsub/utils"
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	uri := utils.GetURIFromConf()
	subscriberClient := utils.Connect(uri, 234)
	subscriberClient.Subscribe("/+/+", 2, subscribers.OnReceive)

	fmt.Print("\nEnter :q to quit\n")

	for {
		time.Sleep(100 * time.Millisecond)
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		if text == ":q\n" {
			os.Exit(0)
		}
	}

}
