package main

import (
	"bytes"
	"fmt"
	"io"
	"net"

	"github.com/tidwall/resp"
)

type Peer struct {
	conn net.Conn
	msgCh chan []byte
}

func NewPeer(conn net.Conn, msgCh chan []byte) *Peer{
	return &Peer{
		conn: conn,
		msgCh: msgCh,
	}
}

func (p *Peer) readLoop() error{
	rd := resp.NewReader(p.conn)
	//buf := make([]byte, 1024)
	for{
		value,_, err := rd.ReadValue()
		if err != nil {
			if err == io.EOF{
				return fmt.Errorf("Connection Closed")
			}
			return err
		}
		var buf bytes.Buffer
		if err:= resp.NewWriter(&buf).WriteValue(value); err!=nil{
			return err
		}
		p.msgCh <- buf.Bytes()
	}
}