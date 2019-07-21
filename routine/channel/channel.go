package channel

type Message []uint8

type Channel struct {
	message Message
	channel chan Message
}

func (c *Channel) Push(mess Message) {

	c.channel <- mess
}

func (c *Channel) Pull() Message {
	mess := <-c.channel
	return mess
}

func Make() Channel {
	return Channel{
		channel: make(chan Message),
	}
}
