package libneosay

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/matrix-org/gomatrix"
)

type Config struct {
	HomeserverURL string            `json:"homeserverURL"`
	UserID        string            `json:"userID"`
	AccessToken   string            `json:"accessToken"`
	Rooms         map[string]string `json:"rooms"`
}

type Neosay struct {
	client *gomatrix.Client
	config *Config
}

const maxMessageSize = 4000
const messageDelay = 2 * time.Second

func NewNeosay(configFile string) (*Neosay, error) {
	// load the config into json
	config := &Config{}
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	// make a new client for matrix
	client, err := gomatrix.NewClient(config.HomeserverURL, config.UserID, config.AccessToken)
	if err != nil {
		return nil, err
	}

	return &Neosay{
		client: client,
		config: config,
	}, nil
}

func (n *Neosay) SendMessage(roomName, message string) error {
	roomID, ok := n.config.Rooms[roomName]
	if !ok {
		return fmt.Errorf("room %s not found in config", roomName)
	}

	// send the message to the room in the config
	_, err := n.client.SendText(roomID, message)
	if err != nil {
		return err
	}

	time.Sleep(messageDelay)
	return nil
}
