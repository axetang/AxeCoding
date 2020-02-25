/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: smtp
 **Element: smtp.Client
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Client struct {
     // Text is the textproto.Conn used by the Client. It is exported to allow for
     // clients to add extensions.
     Text *textproto.Conn
     // contains filtered or unexported fields
 }
 func (c *Client) Auth(a Auth) error
 func (c *Client) Close() error
 func (c *Client) Data() (io.WriteCloser, error)
 func (c *Client) Extension(ext string) (bool, string)
 func (c *Client) Hello(localName string) error
 func (c *Client) Mail(from string) error
 func (c *Client) Noop() error
 func (c *Client) Quit() error
 func (c *Client) Rcpt(to string) error
 func (c *Client) Reset() error
 func (c *Client) StartTLS(config *tls.Config) error
 func (c *Client) TLSConnectionState() (state tls.ConnectionState, ok bool)
 func (c *Client) Verify(addr string) error
 ------------------------------------------------------------------------------------
 **Description:
 A Client represents a client connection to an SMTP server.
 Auth authenticates a client using the provided authentication mechanism. A failed
 authentication closes the connection. Only servers that advertise the AUTH extension
 support this function.
 Close closes the connection.
 Data issues a DATA command to the server and returns a writer that can be used to
 write the mail headers and body. The caller should close the writer before calling
 any more methods on c. A call to Data must be preceded by one or more calls to Rcpt.
 Extension reports whether an extension is support by the server. The extension name
 is case-insensitive. If the extension is supported, Extension also returns a string
 that contains any parameters the server specifies for the extension.
 Hello sends a HELO or EHLO to the server as the given host name. Calling this method
 is only necessary if the client needs control over the host name used. The client
 will introduce itself as "localhost" automatically otherwise. If Hello is called,
 it must be called before any of the other methods.
 Mail issues a MAIL command to the server using the provided email address. If the
 server supports the 8BITMIME extension, Mail adds the BODY=8BITMIME parameter. This
 initiates a mail transaction and is followed by one or more Rcpt calls.
 Noop sends the NOOP command to the server. It does nothing but check that the
 connection to the server is okay.
 Quit sends the QUIT command and closes the connection to the server.
 Rcpt issues a RCPT command to the server using the provided email address. A call
 to Rcpt must be preceded by a call to Mail and may be followed by a Data call or
 another Rcpt call.
 Reset sends the RSET command to the server, aborting the current mail transaction.
 StartTLS sends the STARTTLS command and encrypts all further communication. Only
 servers that advertise the STARTTLS extension support this function.
 TLSConnectionState returns the client's TLS connection state. The return values are
 their zero values if StartTLS did not succeed.
 Verify checks the validity of an email address on the server. If Verify returns nil,
 the address is valid. A non-nil return does not necessarily indicate an invalid
 address. Many servers will not verify addresses for security reasons.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Client代表一个连接到SMTP服务器的客户端；

*************************************************************************************/
package main

func main() {
}
