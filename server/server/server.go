package server

import(
	"net"
	"fmt"
	"reflect"
	"encoding/json"
	"strconv"
	"io"
	"errors"
)
const PackMaxSize int = 4
type serverFunc func(interface{}) interface{}

type rpcHandler struct {
	handler    serverFunc
	argsType   reflect.Type //handler函数的参数类型.
	replysType reflect.Type //handler函数的返回值类型.
}

type RPCServer struct {
	router map[string]rpcHandler
}
type request struct {
	Name string `json:"name"`
	Data []byte `json:"data"`
}

func Server() RPCServer {
	return RPCServer{make(map[string]rpcHandler)}
}


func (r *RPCServer) Register(name string, handler serverFunc, service interface{}) error {
	return r.register(name, handler, service)
}
func (r *RPCServer) register(name string, handler serverFunc, service interface{}) error {
	serviceType := reflect.TypeOf(service)
	err := r.checkHandlerType(serviceType)
	if err != nil {
		return err
	}
	argsType := serviceType.In(0)
	replysType := serviceType.Out(0)
	r.router[name] = rpcHandler{handler: handler, argsType: argsType, replysType: replysType}
	return nil
}
func (r *RPCServer) checkHandlerType(handlerType reflect.Type) error {
	// 判断是否是函数类型.
	if handlerType.Kind() != reflect.Func {
		return errors.New("rpc.Register: handler is not func")
	}
	// 判断参数数量.
	if handlerType.NumIn() != 1 {
		return errors.New("rpc.Register: handler input parameters number is wrong, need one")
	}
	// 判断返回值数量.
	if handlerType.NumOut() != 1 {
		return errors.New("rpc.Register: handler output parameters number is wrong, need one")
	}
	// 判断参数和返回值类型.
	if handlerType.In(0).Kind() != reflect.Struct || handlerType.Out(0).Kind() != reflect.Struct {
		return errors.New("rpc.Register: parameters must be Struct")
	}
	return nil
}

func (r *RPCServer) ListenAndServe(address string) error {
	listener, err := r.listen(address)
	if err != nil {
		return err
	}
	err = r.accept(listener)
	if err != nil {
		return err
	}
	return nil
}
func (r *RPCServer) listen(address string) (*net.TCPListener, error) {
	laddr, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		return nil, err
	}
	listener, err := net.ListenTCP("tcp4", laddr)
	if err != nil {
		return nil, err
	}
	return listener, nil
}
func (r *RPCServer) accept(listener *net.TCPListener) error {
	defer listener.Close()
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			return err
		}
		defer conn.Close()
		go r.handle(conn)
	}
}

func (r *RPCServer) handle(conn *net.TCPConn) {
	if conn == nil {
	}
	dataLen := make([]byte, PackMaxSize)
	for {
		n, err := conn.Read(dataLen)
		if err != nil && err != io.EOF {
		}
		if n <= 0 {
		}
		len, err := strconv.ParseInt(string(dataLen[:PackMaxSize]), 10, 64)
		if err != nil {
		}
		buff := make([]byte, len)
		n, err = conn.Read(buff)
		if err != nil {
		}
		if n <= 0 {
		}

		rsp, err := r.dispatcher(buff)
		if err != nil {
		}
		rspBytes, err := r.packResponse(rsp)
		if err != nil {
		}
		conn.Write(rspBytes)
	}
}
func (r *RPCServer) dispatcher(req []byte) (interface{}, error) {
	var cReq request
	if err := json.Unmarshal(req, &cReq); err != nil {
		return nil, err
	}
	rh, ok := r.router[cReq.Name]
	if !ok {
		return nil, fmt.Errorf("rpc.ListenAndServe: can't find handler named %s", cReq.Name)
	}
	args := reflect.New(rh.argsType).Interface()
	if err := json.Unmarshal(cReq.Data, args); err != nil {
		return nil, err
	}
	return rh.handler(args), nil
}

func (r *RPCServer) packResponse(v interface{}) ([]byte, error) {
	return pack(v)
}

func main(){
	fmt.Println("aaa")
}
func pack(v interface{}) ([]byte, error) {
	rspBytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	rspLen := len(rspBytes)
	rspLenStr := strconv.Itoa(rspLen)
	intLen := len(rspLenStr)

	if intLen > PackMaxSize {
		return nil, errors.New("rpc: package is out of size")
	}

	tb := make([]byte, PackMaxSize+rspLen)
	zerob := []byte("0")
	intLen--
	for i := PackMaxSize - 1; i >= 0; i-- {
		if intLen >= 0 {
			tb[i] = []byte(rspLenStr)[intLen]
			intLen--
		} else {
			tb[i] = zerob[0]
		}
	}
	for i := 0; i < rspLen; i++ {
		tb[PackMaxSize+i] = rspBytes[i]
	}

	return tb, nil
}
