package ws

import (
	"app/internal/consts"
	"app/internal/manager/interfaces"
	"app/internal/manager/interfaces/processor/ws"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"sync"
	"time"
)

var ID = 0

type ClientService struct {
	manager interfaces.IManager
	pool    *PoolService
	log     *zap.Logger
}

func initClientService(manager interfaces.IManager, pool *PoolService, log *zap.Logger) *ClientService {
	return &ClientService{
		manager: manager,
		pool:    pool,
		log:     log,
	}
}

type Client struct {
	id int64
	mu sync.RWMutex

	conn *websocket.Conn
	pool *PoolService
	ctx  *gin.Context

	log  *zap.Logger
	send chan []byte
}

func (cs *ClientService) NewClient(ctx *gin.Context, conn *websocket.Conn) ws.IClient {
	return &Client{
		mu:   sync.RWMutex{},
		ctx:  ctx,
		pool: cs.pool,
		conn: conn,
		log:  cs.log,
		id:   int64(ID),
		send: make(chan []byte),
	}
}

func (c *Client) Reader() {
	defer func() {
		c.pool.unregister <- c
		if err := c.conn.Close(); err != nil {
			c.log.Error("failed to close connection: ", zap.Error(err))
			return
		}
	}()

	c.setReadParams()
	for {
		message, err := c.readMessage()
		if err != nil {
			return
		}

		c.pool.broadcast <- message
	}
}

func (c *Client) Writer() {
	ticker := time.NewTicker(consts.WebsocketPingPeriod)

	defer func() {
		ticker.Stop()
		if err := c.conn.Close(); err != nil {
			c.log.Error("failed to close connection: ", zap.Error(err))
			return
		}
	}()

	for {
		select {
		case message, ok := <-c.send:
			err := c.conn.SetWriteDeadline(time.Now().Add(consts.WebsocketWriteWait))
			if err != nil {
				c.log.Error("msg write deadline: ", zap.Error(err))
				return
			}

			if !ok {
				err := c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					c.log.Error("msg receive: ", zap.Error(err))
					return
				}
				return
			}

			err = c.writeMessage(message)
			if err != nil {
				c.log.Error("msg write: ", zap.Error(err))
				return
			}

		case <-ticker.C:
			err := c.conn.SetWriteDeadline(time.Now().Add(consts.WebsocketWriteWait))
			if err != nil {
				c.log.Error("ticker write deadline: ", zap.Error(err))
				return
			}

			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) readMessage() ([]byte, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, message, err := c.conn.ReadMessage()
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
			c.log.Error("unexpected close", zap.String("user-agent: %v", c.ctx.GetHeader("User-Agent")), zap.Error(err))
		}
		return nil, err
	}
	message = bytes.TrimSpace(bytes.Replace(message, consts.ByteNewLine, consts.ByteSpace, -1))
	return message, nil
}

func (c *Client) writeMessage(message []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	w, err := c.conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return err
	}

	_, err = w.Write(message)
	if err != nil {
		return err
	}

	n := len(c.send)
	for i := 0; i < n; i++ {
		msg := <-c.send
		msg = append(consts.ByteNewLine, msg...)
		if _, err := w.Write(message); err != nil {
			return err
		}
	}

	err = w.Close()
	return err
}

func (c *Client) setReadParams() {
	c.conn.SetReadLimit(consts.WebsocketMaxMessageSize)
	err := c.conn.SetReadDeadline(time.Now().Add(consts.WebsocketPongWait))
	if err != nil {
		c.log.Error("handshake read deadline: ", zap.Error(err))
		return
	}
	c.conn.SetPongHandler(func(string) error {
		err := c.conn.SetReadDeadline(time.Now().Add(consts.WebsocketPongWait))
		if err != nil {
			c.log.Error("pong read deadline: ", zap.Error(err))
			return err
		}
		return nil
	})
}
