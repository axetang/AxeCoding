/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.ListenConfig
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type ListenConfig struct {
     // If Control is not nil, it is called after creating the network
     // connection but before binding it to the operating system.
     //
     // Network and address parameters passed to Control method are not
     // necessarily the ones passed to Listen. For example, passing "tcp" to
     // Listen will cause the Control function to be called with "tcp4" or "tcp6".
     Control func(network, address string, c syscall.RawConn) error
 }
 ListenConfig contains options for listening to an address.
 func (lc *ListenConfig) Listen(ctx context.Context, network, address string) (Listener, error)
 Listen announces on the local network address.
 See func Listen for a description of the network and address parameters.
 func (lc *ListenConfig) ListenPacket(ctx context.Context, network, address string) (PacketConn, error)
 ListenPacket announces on the local network address.
 See func ListenPacket for a description of the network and address parameters.
 ------------------------------------------------------------------------------------
 **Description:
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
