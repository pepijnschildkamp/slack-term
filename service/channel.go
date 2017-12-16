package service

import (
	"fmt"
	"html"

	"github.com/erroneousboat/slack-term/components"
)

const (
	PresenceAway   = "away"
	PresenceActive = "active"
)

type Channel struct {
	ID           string
	Name         string
	Topic        string
	Type         string
	UserID       string
	Presence     string
	Notification bool
}

// ToString will set the label of the channel, how it will be
// displayed on screen. Based on the type, different icons are
// shown, as well as an optional notification icon.
func (c Channel) ToString(stlPrefix string, stlIcon string, stlName string) string {
	var prefix string
	if c.Notification {
		prefix = components.IconNotification
	} else {
		prefix = " "
	}

	var icon string
	switch c.Type {
	case ChannelTypeChannel:
		icon = components.IconChannel
	case ChannelTypeGroup:
		icon = components.IconGroup
	case ChannelTypeIM:
		switch c.Presence {
		case PresenceActive:
			icon = components.IconOnline
		case PresenceAway:
			icon = components.IconOffline
		default:
			icon = components.IconIM
		}
	}

	label := fmt.Sprintf(
		"[%s](%s) [%s](%s) [%s](%s)",
		prefix, stlPrefix,
		icon, stlIcon,
		c.Name, stlName,
	)

	return label
}

// GetChannelName will return a formatted representation of the
// name of the channel
func (c Channel) GetChannelName() string {
	var channelName string
	if c.Topic != "" {
		channelName = fmt.Sprintf("%s - %s",
			html.UnescapeString(c.Name),
			html.UnescapeString(c.Topic),
		)
	} else {
		channelName = c.Name
	}
	return channelName
}
