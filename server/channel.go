package server

import (
	"fmt"
	"time"
)

type Channel struct {
	Name      string
	Password  string
	CreatedAt time.Time
}

// Channel list with channel id.
type ChannelList map[string]*Channel

type ChannelManager struct {
	List ChannelList
}

type ChannelListActivity interface {
	AddChannel(channel *Channel) (string, error)
	RemoveChannel(channelName string) error
	CreateChannelList()
	ShowChannels()
}

// Can be added more logic in future.
func (c *ChannelManager) CreateChannelList() {
	c.List = make(ChannelList)
}

// Adding channel in channel list.
func (c *ChannelManager) AddChannel(channel *Channel) (string, error) {
	if _, exists := c.List[channel.Name]; exists {
		return channel.Name, fmt.Errorf("this %v already in the list", channel.Name)
	}
	c.List[channel.Name] = channel
	return channel.Name, nil

}

// Remove channel from channel list.
func (c *ChannelManager) RemoveChannel(channelName string) error {
	if _, exists := c.List[channelName]; exists {
		return fmt.Errorf("channel with name %v not found", channelName)
	}
	delete(c.List, channelName)
	return nil
}

// Iterating in list of channels.
func (c *ChannelManager) ShowChannels() {
	fmt.Print("Channels list of this server: \n")
	for channelName, channelInfo := range c.List {
		fmt.Printf("%v, info \n %v", channelName, channelInfo)
	}
}
