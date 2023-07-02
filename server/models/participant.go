// models/participant.go

package models

import "github.com/gorilla/websocket"

type Participant struct {
	Conn *websocket.Conn `json:"-"`
}

func NewParticipant(conn *websocket.Conn) *Participant {
	return &Participant{
		Conn: conn,
	}
}

func (p *Participant) GetConn() *websocket.Conn {
	return p.Conn
}
