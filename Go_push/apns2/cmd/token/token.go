package main

import (
	"fmt"
	"log"
	
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/token"
)

func main() {
// 	authKeyPath := flag.String("authKey", "", "Path to .p8 APNSAuthKey file (Required)")
// 	deviceToken := flag.String("token", "", "Push token (Required)")
// 	topic := flag.String("topic", "", "Topic (Required)")
// 	keyID := flag.String("keyID", "", "APNS KeyID (Required)")
// 	teamID := flag.String("teamID", "", "APNS TeamID (Required)")
// 	flag.Parse()

// 	if *authKeyPath == "" || *deviceToken == "" || *topic == "" || *keyID == "" || *teamID == "" {
// 		flag.PrintDefaults()
// 		os.Exit(1)
// }

	authKey, err := token.AuthKeyFromFile("D:/apns2/AuthKey_55Y9DTN8YF.p8")
	if err != nil {
		log.Fatal("token error:", err)
	}

	
	token := &token.Token{
		AuthKey: authKey,
		KeyID:   "55Y9DTN8YF",
		TeamID:  "E47U8TL8S2",
	}

	notification := &apns2.Notification{}
	notification.DeviceToken = "5fbce1ce8d959259d0cf999dc9085db1985ed73741cae657573010b1645a6c35f"
	notification.PushType=apns2.PushTypeVOIP
	notification.Topic = "com.uaround.uaroundSpecific.voip"
	notification.Payload = []byte(`{
			"aps" : {
				"alert" : "Hello Sergey!"
			}
		}
	`)

	client := apns2.NewTokenClient(token).Production()
	res, err := client.Push(notification)

	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}
