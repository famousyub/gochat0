package ws


import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "log"
    "net/http"
)

// ClientManager is a websocket manager
type ClientManager struct {
    Clients    map[string]*Client
    Broadcast  chan []byte
    Register   chan *Client
    Unregister chan *Client
}

// Client is a websocket client
type Client struct {
    ID     string
    Socket *websocket.Conn
    Send   chan []byte
}

// Message is return msg
type Message struct {
    Sender    string `json:"sender,omitempty"`
    Recipient string `json:"recipient,omitempty"`
    Content   string `json:"content,omitempty"`
}

// Manager define a ws server manager
var Manager = ClientManager{
    Broadcast:  make(chan []byte),
    Register:   make(chan *Client),
    Unregister: make(chan *Client),
    Clients:    make(map[string]*Client),
}

//Start is before the project runs, the program starts start > go Manager.Start ()
func (manager *ClientManager) Start() {
    for {
        log.Println ("< --- pipeline communication -- >")
        select {
        case conn := <-Manager.Register:
            log.Printf ("new user joined in% v"), conn.ID )
            Manager.Clients[conn.ID] = conn
            jsonMessage, _ := json.Marshal(&Message{Content: "Successful connection to socket service"})
            conn.Send <- jsonMessage
        case conn := <-Manager.Unregister:
            log.Printf ("user left% v"), conn.ID )
            if _, ok := Manager.Clients[conn.ID]; ok {
                jsonMessage, _ := json.Marshal(&Message{Content: "A socket has disconnected"})
                conn.Send <- jsonMessage
                close(conn.Send)
                delete(Manager.Clients, conn.ID)
            }
        case message := <-Manager.Broadcast:
            MessageStruct :=Message{}
            json.Unmarshal(message, &MessageStruct)
            for id, conn := range Manager.Clients {
                if id!=creatId(MessageStruct.Recipient,MessageStruct.Sender){
                    continue
                }
                select {
                case conn.Send <- message:
                default:
                    close(conn.Send)
                    delete(Manager.Clients, conn.ID)
                }
            }
        }
    }
}
func creatId(uid,touid string) string {
    return uid+"_"+touid
}
func (c *Client) Read() {
    defer func() {
        Manager.Unregister <- c
        c.Socket.Close()
    }()

    for {
        c.Socket.PongHandler()
        _, message, err := c.Socket.ReadMessage()
        if err != nil {
            Manager.Unregister <- c
            c.Socket.Close()
            break
        }
        log.Printf ("message read to client: s", string (message))
        Manager.Broadcast <- message
    }
}

func (c *Client) Write() {
    defer func() {
        c.Socket.Close()
    }()

    for {
        select {
        case message, ok := <-c.Send:
            if !ok {
                c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
                return
            }
            log.Printf ("message sent to client: s", string (message))

            c.Socket.WriteMessage(websocket.TextMessage, message)
        }
    }
}

//Testhandler socket connection middleware function: upgrade protocol, user authentication, user-defined information, etc
func WsHandler(c *gin.Context) {
    uid := c.Query("uid")
    touid := c.Query("to_uid")
    conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        http.NotFound(c.Writer, c.Request)
        return
    }
    //User information authentication can be added
    client := &Client{
        ID:    creatId(uid,touid),
        Socket: conn,
        Send:   make(chan []byte),
    }
    Manager.Register <- client
    go client.Read()
    go client.Write()
}
