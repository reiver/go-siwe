package siwe

func (receiver *Message) AppendResource(value string) {
	receiver.Resources = append(receiver.Resources, value)
}

func (receiver *Message) UnsetResources(value string) {
	receiver.Resources = nil
}
