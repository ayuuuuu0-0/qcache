package main

import (
	"bufio"
	"net"
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
	reader := bufio.NewReader(p.conn)
	//buf := make([]byte, 1024)
	for{
		line, err := reader.ReadBytes('\n')
		if err != nil {
			return err
		}
		// msgBuf := make([]byte, n)
		// copy(msgBuf, buf[:n])
		// p.msgCh <- msgBuf

		p.msgCh <- line
	}
}