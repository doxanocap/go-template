package ws

import (
	"app/internal/manager/interfaces"
	"encoding/json"
	"go.uber.org/zap"
	"time"
)

type PoolService struct {
	broadcast   chan []byte
	register    chan *Client
	unregister  chan *Client
	clients     map[*Client]bool
	chatHistory []BroadcastMsg
	log         *zap.Logger
}

func initPoolService(manager interfaces.IManager, log *zap.Logger) *PoolService {
	return &PoolService{
		broadcast:   make(chan []byte),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		clients:     map[*Client]bool{},
		chatHistory: []BroadcastMsg{},
		log:         log,
	}
}

type BroadcastMsg struct {
	SenderId   int64     `json:"sender_id"`
	ReceiverId int64     `json:"receiver_id"`
	Message    string    `json:"message"`
	Timestamp  time.Time `json:"timestamp"`
}

func (p *PoolService) Start() {
	for {
		select {
		case client := <-p.register:
			p.clients[client] = true

		case client := <-p.unregister:
			if _, ok := p.clients[client]; ok {
				delete(p.clients, client)
				close(client.send)
			}

		case message := <-p.broadcast:
			var msg BroadcastMsg
			if err := json.Unmarshal(message, &msg); err != nil {
				p.log.Error("bad request", zap.Error(err))
				continue
			}

			p.recordMessage(&msg)

			for client := range p.clients {
				if client.id == msg.SenderId || client.id == msg.ReceiverId {
					select {
					case client.send <- message:
					default:
						close(client.send)
						delete(p.clients, client)
					}
				}
			}
		}
	}
}

func (p *PoolService) recordMessage(broadcast *BroadcastMsg) {
	p.chatHistory = append(p.chatHistory, *broadcast)
}
