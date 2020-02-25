/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.Listener
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type Listener interface {
     // Accept waits for and returns the next connection to the listener.
     Accept() (Conn, error)
     // Close closes the listener.
     // Any blocked Accept operations will be unblocked and return errors.
     Close() error
     // Addr returns the listener's network address.
     Addr() Addr
 }
 A Listener is a generic network listener for stream-oriented protocols.
 Multiple goroutines may invoke methods on a Listener simultaneously.
 func FileListener(f *os.File) (ln Listener, err error)
 FileListener returns a copy of the network listener corresponding to the open file f. It is the caller's responsibility to close ln when finished. Closing ln does not affect f, and closing f does not affect ln.
 func Listen(network, address string) (Listener, error)
 Listen announces on the local network address.
 The network must be "tcp", "tcp4", "tcp6", "unix" or "unixpacket".
 For TCP networks, if the host in the address parameter is empty or a literal unspecified IP address, Listen listens on all available unicast and anycast IP addresses of the local system. To only use IPv4, use network "tcp4". The address can use a host name, but this is not recommended, because it will create a listener for at most one of the host's IP addresses. If the port in the address parameter is empty or "0", as in "127.0.0.1:" or "[::1]:0", a port number is automatically chosen. The Addr method of Listener can be used to discover the chosen port.
 See func Dial for a description of the network and address parameters.
 ------------------------------------------------------------------------------------
 **Description:
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
