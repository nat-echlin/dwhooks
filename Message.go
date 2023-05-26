package dwh

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Content   string  `json:"content,omitempty"`
	Embeds    []Embed `json:"embeds,omitempty"`
	Username  string  `json:"username,omitempty"`
	AvatarURL string  `json:"avatar_url,omitempty"`
}

func NewMessage(content string) Message {
	return Message{content, []Embed{}, "", ""}
}

func (msg *Message) SetUsername(username string) {
	msg.Username = username
}

func (msg *Message) SetEmbed(embed Embed) {
	msg.Embeds = []Embed{embed}
}

func (msg *Message) SetAvatarURL(avatarURL string) {
	msg.AvatarURL = avatarURL
}

func (msg *Message) SetEmbeds(embeds []Embed) {
	msg.Embeds = embeds
}

func (msg *Message) AddEmbed(embed Embed) {
	msg.Embeds = append(msg.Embeds, embed)
}

func (msg Message) toJSON() ([]byte, error) {
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return []byte(""), fmt.Errorf("failed to marshall json, %v", err)
	}
	return jsonData, nil
}
