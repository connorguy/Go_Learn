package main

import (
	"log"

	"github.com/deckarep/gosx-notifier"
)

// Notifier is a example of using the gosx-notifier library for creating osx
// notifications.
func main() {
	//At a minimum specifiy a message to display to end-user.
	note := gosxnotifier.NewNotification("Looke Here")

	//Optionally, set a title
	note.Title = "Hello World"

	//Optionally, set a subtitle
	note.Subtitle = "Gopher"

	//Optionally, set a sound from a predefined set.
	//note.Sound = gosxnotifier.Default

	//Optionally, set a group which ensures only one notification is ever shown replacing previous notification of same group id.
	//note.Group = "com.unique.yourapp.identifier"

	//Optionally, set a sender (Notification will now use the Safari icon)
	//note.Sender = "com.apple.Safari"

	//Optionally, specifiy a url or bundleid to open should the notification be
	//clicked.
	//note.Link = "http://www.yahoo.com" //or BundleID like: com.apple.Terminal

	//Optionally, an app icon (10.9+ ONLY)
	//note.AppIcon = "gopher.png"

	//Optionally, a content image (10.9+ ONLY)
	//note.ContentImage = "gopher.png"

	//Then, push the notification
	err := note.Push()

	//If necessary, check error
	if err != nil {
		log.Println("Uh oh!")
	}
}
