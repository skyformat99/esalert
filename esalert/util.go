package esalert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/smtp"
)

// ToBuffer 转换一个对象为byte[]
func ToBuffer(i interface{}) *bytes.Buffer {
	json, err := json.Marshal(i)
	if err != nil {
		return nil
	}
	return bytes.NewBuffer(json)
}

// QueryToJSON 转化query为json byte[]
func QueryToJSON(query interface{}) ([]byte, error) {
	query = cleanupMapValue(query)
	return json.Marshal(query)
}

// Mail 邮件实体
type Mail struct {
	Host, Port, Username, Password, From, ReplyTo string
}

// Send 发送邮件
func (mail Mail) Send(to []string, subject string, msg []byte) error {
	if mail.From == "" {
		mail.From = mail.Username
	}
	// 如果msg中夹带空格，msg与header中间需要有一个分行smtp服务器才能识别内容
	msg = append([]byte("\r\n"), msg...)
	server := net.JoinHostPort(mail.Host, mail.Port)
	auth := smtp.PlainAuth("", mail.Username, mail.Password, mail.Host)
	from := []byte("From:" + mail.From)
	replyTo := []byte("Reply-To:" + mail.ReplyTo)
	sub := []byte("Subject:" + subject)
	return smtp.SendMail(server, auth, mail.From, to, bytes.Join([][]byte{from, replyTo, sub, msg}, []byte("\r\n")))
}

func cleanupInterfaceArray(in []interface{}) []interface{} {
	res := make([]interface{}, len(in))
	for i, v := range in {
		res[i] = cleanupMapValue(v)
	}
	return res
}

func cleanupInterfaceMap(in map[interface{}]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	for k, v := range in {
		res[fmt.Sprintf("%v", k)] = cleanupMapValue(v)
	}
	return res
}

func cleanupMapValue(v interface{}) interface{} {
	switch v := v.(type) {
	case []interface{}:
		return cleanupInterfaceArray(v)
	case map[interface{}]interface{}:
		return cleanupInterfaceMap(v)
	case string:
		return v
	default:
		return fmt.Sprintf("%v", v)
	}
}
