package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// https://github.com/anthdm/gstream

const (
	Scheme string = "http://"
)

const (
	NumberOfClients    = 8
	NumberOfConnection = 8
)

var QM *QueueManager

var (
	TickLength uint64 = 4
	Rate       uint64 = 1000
)

var (
	indexConsumerRR   uint64 = 0
	indexConnectionRR uint64 = 0
	indexClientRR     uint64 = 0
)

type Message struct {
	Body []byte
	Path string
	Port string
	// Next *Message
	// Prev *Message
}

type Connection struct {
	Ip         string
	ClientList []*http.Client
}

func (c *Connection) getClientRR() *http.Client {
	indexClientRR++
	return c.ClientList[indexClientRR%uint64(len(c.ClientList))]
}

type ConsumerInfo struct {
	Ip             string
	ConnectionList []*Connection
}

func (c *ConsumerInfo) getConnectionRR() *Connection {
	indexConnectionRR++
	return c.ConnectionList[indexConnectionRR%uint64(len(c.ConnectionList))]
}

type Queue interface {
	InitConsumerInfo(string)
	Producer(Message) error
	Consumer() error
	Len() uint64
}

type QueueManager struct {
	QueueMap map[string]Queue
	MuLock   *sync.Mutex
}

func NewQueueManager() *QueueManager {
	return &QueueManager{
		QueueMap: make(map[string]Queue),
	}
}

type GPRCQueue struct {
	Name         string
	MessageQueue []Message
}

func NewGPRCQueue(name string) *GPRCQueue {
	return &GPRCQueue{
		Name: name,
	}
}

func (gq *GPRCQueue) Producer(mess Message) error {

	return nil
}

func (gq *GPRCQueue) Consumer() error {

	return nil
}

func (gq *GPRCQueue) Len() uint64 {
	return uint64(len(gq.MessageQueue))
}

func (gq *GPRCQueue) InitConsumerInfo(ip string) {

}

type HTTP2Queue struct {
	Name         string
	MessageQueue []*Message
	ConsumerList []*ConsumerInfo
}

func NewHTTP2Queue(name string) *HTTP2Queue {
	return &HTTP2Queue{
		Name: name,
	}
}

func (hq *HTTP2Queue) Producer(mess Message) error {
	hq.MessageQueue = append(hq.MessageQueue, &mess)
	return nil
}

func (hq *HTTP2Queue) Consumer() error {
	for {
		if hq.Len() == 0 {
			continue
		}
		rate := min(Rate, hq.Len())
		for i := uint64(0); i < TickLength; i++ {
			doneFlag := false
			for j := uint64(0); j < rate/TickLength+1; j++ {
				if hq.Len() == 0 {
					doneFlag = true
					break
				}
				go hq.DeQueue()
			}
			if doneFlag {
				break
			}
			time.Sleep(time.Duration(rate/TickLength) * time.Millisecond)
		}
	}
}

func (hq *HTTP2Queue) getConsumerRR() *ConsumerInfo {
	indexConsumerRR++
	return hq.ConsumerList[indexConsumerRR%uint64(len(hq.ConsumerList))]
}

func (hq *HTTP2Queue) DeQueue() {
	mess := hq.MessageQueue[0]
	if !strings.HasPrefix(mess.Path, "/") {
		mess.Path = "/" + mess.Path
	}
	hq.MessageQueue = hq.MessageQueue[1:]
	connection := hq.getConsumerRR().getConnectionRR()
	url := Scheme + connection.Ip + ":" + mess.Port + mess.Path
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(mess.Body))
	if err != nil {
		fmt.Println("Cannot create request: ", err)
		return
	}
	_, err = connection.getClientRR().Do(req)
	if err != nil {
		fmt.Println("Cannot send request, mess go back to queue, err: ", err)
		hq.Producer(*mess)
		return
	}
}

func (hq *HTTP2Queue) Len() uint64 {
	return uint64(len(hq.MessageQueue))
}

func (hq *HTTP2Queue) InitConsumerInfo(ip string) {
	var connectionList []*Connection
	for i := 0; i < NumberOfConnection; i++ {
		var clientList []*http.Client
		for j := 0; j < NumberOfClients; j++ {
			clientList = append(clientList, &http.Client{
				Transport: &http2.Transport{
					AllowHTTP:          true,
					DisableCompression: true,
				},
				Timeout: 10 * time.Second,
			})
		}
		connectionList = append(connectionList, &Connection{
			Ip:         ip,
			ClientList: clientList,
		})
	}

	hq.ConsumerList = append(hq.ConsumerList, &ConsumerInfo{
		Ip:             ip,
		ConnectionList: connectionList,
	})
}

func (hq *HTTP2Queue) RemoveConsumerInfo(ip string) {

}

func getMessHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	qName := r.Header.Get("queue_name")
	if qName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing queue name in header"))
		return
	}
	port := r.Header.Get("port")
	if port == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing port in header"))
		return
	}
	path := r.Header.Get("path")
	if path == "" {
		path = "/"
	}

	mess := Message{
		Body: body,
		Path: path,
		Port: port,
	}

	_, ok := QM.QueueMap[qName]
	if !ok {
		QM.QueueMap[qName].(*HTTP2Queue).Name = qName
	} else {
		err = QM.QueueMap[qName].Producer(mess)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("received message"))
}

func registerConsumer(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	qName := r.Header.Get("queue_name")
	if qName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing queue name in header"))
		return
	}
	_, ok := QM.QueueMap[qName]
	if !ok {
		QM.QueueMap[qName].(*HTTP2Queue).Name = qName
	} else {
		QM.QueueMap[qName].InitConsumerInfo(string(body))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("register consumer success"))
}

func newMuxRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/send-req", getMessHandler)
	mux.HandleFunc("/register/consumer", registerConsumer)

	return mux
}

func newServerHTTP() *http.Server {
	h2s := http2.Server{}
	mux := newMuxRouter()
	return &http.Server{
		Addr:    ":8457",
		Handler: h2c.NewHandler(mux, &h2s),
	}
}

func StartHttpServer() {
	server := newServerHTTP()
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func main() {
	QM = NewQueueManager()
	go StartHttpServer()
}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}
