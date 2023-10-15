package main

import "fmt"

func main() {

	var bldr = NewNotificationBuilder()
	bldr.SetTitle("Title")
	bldr.SetSubTitle("Subtitle")
	bldr.SetMessage("Message")
	bldr.SetImage("Image")
	bldr.SetIcon("Icon")
	bldr.SetPriority(9)

	notif, err := bldr.Build()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("No Error", notif)
	}
}
