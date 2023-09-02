package mail

import (
	"bytes"
	"fmt"
	"strings"
)

type Mail struct {
	To      []string
	Subject string
	Body    string
}

func (m *Mail) Build(from string) []byte {
	var buff = bytes.Buffer{}
	buff.WriteString(fmt.Sprintf("Content-Type: %s\r\n", "text/html"))
	buff.WriteString(fmt.Sprintf("From: %s\r\n", from))
	if len(m.To) > 0 {
		buff.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(m.To, ";")))
	}

	buff.WriteString(fmt.Sprintf("Subject: %s\r\n", m.Subject))
	buff.WriteString("\r\n\n" + m.Body)

	return buff.Bytes()
}
