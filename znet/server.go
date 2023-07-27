package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

// Server :IServer接口实现,定义一个Server的服务器模块
type Server struct {
	//	服务器的名称
	Name string
	//服务器绑定的ip版本
	IPVersion string
	//服务器监听的端IP
	IP string
	//服务器监听的端口
	Port string
}

func (s *Server) Start() {
	fmt.Println("[Start] Server Listenner at IP : %s,Port %d, is starting\n", s.IP, s.Port)
	// 用go做一个异步
	go func() {
		//	1. 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addt error :", err)
		}
		//2. 监听服务器的地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}
		fmt.Println("start Zinx server succ", s.Name, "succ, Listenning..")

		// 3.阻塞的等待的客户端链接,处理客户端链接业务(读写)
		for {
			//	如果有客户端连接过来,阻塞会返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			// 以及与客户端建立连接,做一下业务,做一个最大512字节的回显业务
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err", err)
						continue
					}

					// 回显功能
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err", err)
						continue
					}
				}
			}()
		}
	}()

}

func (s *Server) Stop() {
	// TODO将一些服务器的资源,状态或一些已经开辟的连接信息,进行停止或者回收

}

func (s *Server) Serve() {
	//启动server的服务功能
	s.Start()

	// TODO 做一些启动服务器之后的额外业务

	//阻塞状态:如果不阻塞,上面Start结束后,整个server就结束了,服务器就没了
	select {}
}

/*
NewServer 初始化Server模块的方法
*/
func NewServer(name string) ziface.IServer {
	s := &Server{Name: name, IPVersion: "tcp4", IP: "0.0.0.0", Port: "8999"}
	return s
}
