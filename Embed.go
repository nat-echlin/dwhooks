package main

import "time"

type Embed struct {
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Url         string  `json:"url,omitempty"`
	Timestamp   string  `json:"timestamp,omitempty"`
	Colour      int     `json:"color,omitempty"`
	Fields      []Field `json:"fields,omitempty"`
}

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

func NewEmbed() Embed {
	return Embed{"", "", "", "", 0, []Field{}}
}

func (emb *Embed) SetTitle(title string) {
	emb.Title = title
}

func (emb *Embed) SetDescription(desc string) {
	emb.Description = desc
}

func (emb *Embed) SetUrl(url string) {
	emb.Url = url
}

// ensure that the timestamp ts is in seconds (eg NOT milliseconds)
func (emb *Embed) SetTimestamp(ts int64) {
	// convert timestamp to time.time, then to correctly formatted string
	t := time.Unix(ts, 0)
	isoTime := t.UTC().Format(time.RFC3339)

	emb.Timestamp = isoTime
}

func (emb *Embed) SetColour(colour int) {
	emb.Colour = colour
}

func (emb *Embed) AddField(name string, value string, inline bool) {
	emb.Fields = append(emb.Fields, Field{name, value, inline})
}
