package notifications

import "github.com/gen2brain/beeep"

// ShowNotification displays desktop notifications
func ShowNotification(header string, body string) {
	err := beeep.Notify(header, body, "assets/information.png")
	if err != nil {
		panic(err)
	}
}
