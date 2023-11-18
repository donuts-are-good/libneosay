![libneosay](https://github.com/donuts-are-good/libneosay/assets/96031819/5bce5427-6f92-4cc3-a32f-949c0510b43b)
![donuts-are-good's followers](https://img.shields.io/github/followers/donuts-are-good?&color=555&style=for-the-badge&label=followers) ![donuts-are-good's stars](https://img.shields.io/github/stars/donuts-are-good?affiliations=OWNER%2CCOLLABORATOR&color=555&style=for-the-badge) ![donuts-are-good's visitors](https://komarev.com/ghpvc/?username=donuts-are-good&color=555555&style=for-the-badge&label=visitors)

# libneosay

a library for neosay, a program that shouts into the matrix

## usage

you probably don't need this, but i need this, so that is why it got made.

generally you can use neosay by piping program output into it and configuring it to speak to a matrix channel, so that you can have notifications and things in your bash scripts that easily go to your devices of choice. 

i needed libneosay because i like using matrix channels to monitor web apps, and the one im working on has distinct subsections so libneosay has support for multiple chat channels in the config and you can message each channel easily like the example below.

```go
package main

import (
	"fmt"
	"github.com/donuts-are-good/libneosay"
)

func main() {
	neosay, err := libneosay.NewNeosay("path/to/config.json")
	if err != nil {
		fmt.Println("Error initializing Neosay:", err)
		return
	}

	rooms := []string{"rooma", "roomb", "roomc", "roomd"}

	for _, room := range rooms {
		message := fmt.Sprintf("Hello, %s!", room)
		err = neosay.SendMessage(room, message)
		if err != nil {
			fmt.Printf("Error sending message to %s: %v\n", room, err)
		} else {
			fmt.Printf("Message sent to %s\n", room)
		}
	}
}

```

this demo program will say hello in every room it has a config for. 

this is what a config looks like:

```json
{
  "HomeserverURL": "your_homeserver_url",
  "UserID": "your_user_id",
  "AccessToken": "your_access_token",
  "Rooms": {
      "rooma": "room_id_a",
      "roomb": "room_id_b",
      "roomc": "room_id_c",
      "roomd": "room_id_d"
  }
}
```

## license

MIT License 2023 donuts-are-good, for more info see license.md
