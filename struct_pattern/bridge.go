package struct_pattern

type AbstractMessage interface {
	SendMessage() error
}

type CommonMessage interface {
	Send() error
}

type MessageSMS struct {

}

func (m *MessageSMS) Send() error {
	return nil
}

type MessageEmail struct {

}

func (m *MessageEmail) Send() error {
	return nil
}

type NormalMessage struct {
	CommonMessage
}

func NewNormalMessage(m CommonMessage) AbstractMessage {
	return &NormalMessage{
		m,
	}
}

func (m *NormalMessage) SendMessage() error {
	m.CommonMessage.Send()
	return nil
}

type UrgentMessage struct {
	CommonMessage
}

func NewUrgentMessage(m CommonMessage) AbstractMessage {
	return &UrgentMessage{
		m,
	}
}

func (m *UrgentMessage) SendMessage() error {
	m.CommonMessage.Send()
	return nil
}
