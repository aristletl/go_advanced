package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"log"
)

//1.总结几种 socket 粘包的解包方式: fix length/delimiter based/length field based frame decoder。尝试举例其应用
//    (1)fix length：无论收到多少报文，都按照设定好的固定长度LENGTH进行解码，累计读取一个LENGTH的报文之后就认为读取到了一个完整的消息。然后将计数器复位，重新开始读下一个数据报文。
//    (2)delimiter based：基于分隔符对字节分割，比如回车或者其他自定义的分隔符。 如读取文本协议时吗，一般用 \r\n 做分隔符，以实现按行读取
//    (3)length field based frame decoder：通过指定长度来标识整包消息，这样就可以自动的处理黏包和半包消息，只要传入正确的参数，就可以轻松解决“读半包”的问题。
//    类似 goim 、NSQ、ECTD等协议都是通过这种方式。
//2.实现一个从 socket connection 中解码出 goim 协议的解码器

type Protocol struct {
	Version    uint16
	Operation  uint32
	SequenceId uint32
	Body       []byte
}

const (
	LeastLen int = 14
)

func Encoder(w io.Writer, pro Protocol) (int, error) {
	total := 0
	packetLen := len(pro.Body)
	buf := make([]byte, LeastLen) //包长度+版本+操作+序列号
	binary.BigEndian.PutUint32(buf[:4], uint32(packetLen))
	binary.BigEndian.PutUint16(buf[4:6], pro.Version)
	binary.BigEndian.PutUint32(buf[6:10], pro.Operation)
	binary.BigEndian.PutUint32(buf[10:14], pro.SequenceId)
	n, err := w.Write(buf)
	if err != nil {
		return total, err
	}
	total += n
	if packetLen < 1 {
		return total, err
	}
	n, err = w.Write(pro.Body)
	if err != nil {
		return total, err
	}
	total += n
	return total, nil
}
func Decoder(data []byte) (Protocol, error) {
	var pro Protocol
	if len(data) < LeastLen {
		return pro, errors.New("data error")
	}
	packetLen := binary.BigEndian.Uint32(data[:4])
	pro.Version = binary.BigEndian.Uint16(data[4:6])
	pro.Operation = binary.BigEndian.Uint32(data[6:10])
	pro.SequenceId = binary.BigEndian.Uint32(data[10:14])
	body := make([]byte, packetLen)
	copy(body[:], data[14:])
	pro.Body = body
	return pro, nil
}
func main() {
	var buf bytes.Buffer
	pro := Protocol{
		Version:    1,
		Operation:  2,
		SequenceId: 1234,
		Body:       []byte("Hello World"),
	}
	total, err := Encoder(&buf, pro)
	if err != nil {
		log.Println("encoder error:", err)
		return
	}
	data := buf.Bytes()
	log.Println("write len:", total, "data:", data)
	pro2, err := Decoder(data)
	if err != nil {
		log.Println("decoder error:", err)
		return
	}
	log.Println(pro2.Version, pro2.Operation, pro2.SequenceId, string(pro2.Body))
}
