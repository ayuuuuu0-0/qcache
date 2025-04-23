package main

import (
	// "bytes"
	// "fmt"
	// "io"
	"net"

	//"github.com/tidwall/resp"
)

type Peer struct {
	conn net.Conn
	msgCh chan Message
}

func (p* Peer) Send(msg []byte) (int, error) {
	return p.conn.Write(msg)
}

func NewPeer(conn net.Conn, msgCh chan Message) *Peer{
	return &Peer{
		conn: conn,
		msgCh: msgCh,
	}
}

func (p *Peer) readLoop() error{
	//rd := resp.NewReader(p.conn)
	buf := make([]byte, 1024)
	for{
		n, err := p.conn.Read(buf)
		if err != nil{
			return err
		}
		msgBuf := make([]byte, n)
		copy(msgBuf, buf[:n])
		p.msgCh <- Message{
			data: msgBuf,
			peer: p,
		}
	}
}
	// for{
	// 	value,_, err := rd.ReadValue()
	// 	if err != nil {
	// 		if err == io.EOF{
	// 			return fmt.Errorf("Connection Closed")
	// 		}
	// 		return err
	// 	}
	// 	var buf bytes.Buffer
	// 	if err:= resp.NewWriter(&buf).WriteValue(value); err!=nil{
	// 		return err
	// 	}
		