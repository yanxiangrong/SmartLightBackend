package network

import (
	"SmartLightBackend/pkg/logging"
	"net"
)

type HandlerFunc func([]byte)

var listener net.Listener
var conns []*net.Conn
var callBackHandler HandlerFunc

// SetCallBack 设置回调函数，在接收到数据后会调用改函数
func SetCallBack(callBack HandlerFunc) {
	callBackHandler = callBack
}

func Close() {
	logging.Info("Close listener")
	for i := 0; i < len(conns); i++ {
		logging.Debug("Close connect ", (*conns[i]).RemoteAddr())
		err := (*conns[i]).Close()
		if err != nil {
			logging.Error("Error close connect ", err.Error())
		}
	}
	err := listener.Close()
	if err != nil {
		logging.Error("Error close listener ", err.Error())
		return
	}
}

// ListenAndServe 监听
func ListenAndServe() {
	logging.Info("Listen TCP :9528")
	var err error
	listener, err = net.Listen("tcp", ":9528")
	if err != nil {
		logging.Fatal("Error listening ", err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			logging.Error("Error accepting ", err.Error())
			return
		}
		logging.Debug("Accept connect ", conn.RemoteAddr())
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	conns = append(conns, &conn)

	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)

		if err != nil {
			logging.Debug("Error reading ", err.Error())
			conn = nil
			cleanDeadConn()
			return
		}
		if callBackHandler != nil {
			logging.Debug("Receive Data ", conn.RemoteAddr().String()+" ", buf[:13])
			callBackHandler(buf)
		}
	}
}

// SendBroadcast 发送广播，给每一个客户端发送
func SendBroadcast(b []byte) {
	for i := 0; i < len(conns); i++ {
		_, err := (*conns[i]).Write(b)
		if err != nil {
			cleanDeadConn()
			logging.Error("Conn Write error ", err.Error())
		}
	}
}

// 清除已经失效的连接
func cleanDeadConn() {
	for i := 0; i < len(conns); i++ {
		if (*conns[i]) == nil {
			conns = append(conns[:i], conns[i+1:]...)
		}
	}
}
