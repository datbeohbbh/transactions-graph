package amqpClient

func (c *AmqpClient) GetChannel() error {
	ch, err := c.client.Channel()
	if err != nil {
		return err
	}
	c.channel = ch
	return nil
}
