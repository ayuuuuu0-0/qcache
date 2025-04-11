package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"qcache/client"
	"time"
)

const defaultListenAddr = ":5001";

type Config struct {
	ListenAddr string
}

type Message struct {
	data []byte
	peer *Peer
}

type Server struct {
	Config
	peers       map[*Peer] bool
	ln          net.Listener
	addPeerCh   chan *Peer
	quitCh      chan struct{}
	msgCh       chan Message
	kv          *KV
}

func NewServer(cfg Config) *Server{
	if len(cfg.ListenAddr) == 0{
		cfg.ListenAddr = defaultListenAddr
	}
	return &Server{
		Config:          cfg,
		peers:           make(map[*Peer]bool),
		addPeerCh:       make(chan *Peer ), 
		quitCh:          make(chan struct{}),
		msgCh:           make(chan []byte),
		kv:              NewKV(),
	}
}

func (s *Server) Start() error{
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	s.ln = ln

	go s.loop()

	slog.Info("server running", "listenAddr", s.ListenAddr)

	return s.acceptLoop()
	
}


func (s* Server) handleMessage(Msg Message) error{
	cmd, err := parseCommand(string(Msg.data))
	if err != nil {
		return err
	}
	switch v := cmd.(type){
	case SetCommand:
		return s.kv.Set(v.key, v.val)
	case GetCommand:
		val, ok := s.kv.Get(v.key)
		if !ok {
			return fmt.Errorf("key not found")
		}
		msg.peer.conn.
		}
	return nil
}


func(s *Server) loop(){
	for{
		select{
		case rawMsg := <- s.msgCh:
			if err := s.handleRawMessage(rawMsg); err != nil {
				slog.Error("raw message error", "err", err)
			}
		case <- s.quitCh:
			return 
		case peer := <- s.addPeerCh:
			s.peers[peer] = true
		}
	}
}

func (s *Server) acceptLoop() error {
    for {
        conn, err := s.ln.Accept() 
        if err != nil {
            slog.Error("accept error", "err", err)
            continue
        }
        go s.handleConn(conn)
    }
}


func (s *Server) handleConn (conn net.Conn){
    peer := NewPeer(conn, s.msgCh)
	s.addPeerCh <- peer

	slog.Info("new peer created", "remoteAddr", conn.RemoteAddr())
	if err :=  peer.readLoop(); err != nil {
		slog.Error("peer read error", "err", err, "remoteAddr", conn.RemoteAddr())
	}
}

func main() {
	server := NewServer(Config{})
	go func(){
	log.Fatal(server.Start())
}()
time.Sleep(time.Second)
for i := 0; i<10; i++{
c := client.New("localhost:5001");
if err := c.Set(context.TODO(), 
            fmt.Sprint("foo_%id", i), 
               fmt.Sprintf("bar_%id", i)); err != nil{
				log.Fatal(err)
			   }
 val, err := c.Get(context.TODO(), 
            fmt.Sprint("foo_%id", i), 
               );
 if err != nil{
	log.Fatal(err);
};
fmt.Println(val)
}
fmt.Println(server.kv.data)
time.Sleep(time.Second * 2)
}

