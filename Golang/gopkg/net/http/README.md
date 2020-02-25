# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package http
## import "net/http"

Package http provides HTTP client and server implementations.
Get, Head, Post, and PostForm make HTTP (or HTTPS) requests:
```
resp, err := http.Get("http://example.com/")
...
resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
...
resp, err := http.PostForm("http://example.com/form",
	url.Values{"key": {"Value"}, "id": {"123"}})
The client must close the response body when finished with it:

resp, err := http.Get("http://example.com/")
if err != nil {
	// handle error
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
// ...
For control over HTTP client headers, redirect policy, and other settings, create a Client:

client := &http.Client{
	CheckRedirect: redirectPolicyFunc,
}

resp, err := client.Get("http://example.com")
// ...

req, err := http.NewRequest("GET", "http://example.com", nil)
// ...
req.Header.Add("If-None-Match", `W/"wyzzy"`)
resp, err := client.Do(req)
// ...
For control over proxies, TLS configuration, keep-alives, compression, and other settings, create a Transport:

tr := &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}
client := &http.Client{Transport: tr}
resp, err := client.Get("https://example.com")
```
Clients and Transports are safe for concurrent use by multiple goroutines and for efficiency should only be created once and re-used.

ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil, which means to use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux:
```
http.Handle("/foo", fooHandler)

http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})

log.Fatal(http.ListenAndServe(":8080", nil))
More control over the server's behavior is available by creating a custom Server:

s := &http.Server{
	Addr:           ":8080",
	Handler:        myHandler,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}
log.Fatal(s.ListenAndServe())
```
Starting with Go 1.6, the http package has transparent support for the HTTP/2 protocol when using HTTPS. Programs that must disable HTTP/2 can do so by setting Transport.TLSNextProto (for clients) or Server.TLSNextProto (for servers) to a non-nil, empty map. Alternatively, the following GODEBUG environment variables are currently supported:
```
GODEBUG=http2client=0  # disable HTTP/2 client support
GODEBUG=http2server=0  # disable HTTP/2 server support
GODEBUG=http2debug=1   # enable verbose HTTP/2 debug logs
GODEBUG=http2debug=2   # ... even more verbose, with frame dumps
```
The GODEBUG variables are not covered by Go's API compatibility promise. Please report any issues before disabling HTTP/2 support: https://golang.org/s/http2bug

The http package's Transport and Server both automatically enable HTTP/2 support for simple configurations. To enable HTTP/2 for more complex configurations, to use lower-level HTTP/2 features, or to use a newer version of Go's http2 package, import "golang.org/x/net/http2" directly and use its ConfigureTransport and/or ConfigureServer functions. Manually configuring HTTP/2 via the golang.org/x/net/http2 package takes precedence over the net/http package's built-in HTTP/2 support.

## Index
```
Constants
Variables
func CanonicalHeaderKey(s string) string
func DetectContentType(data []byte) string
func Error(w ResponseWriter, error string, code int)
func Handle(pattern string, handler Handler)
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
func ListenAndServe(addr string, handler Handler) error
func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
func MaxBytesReader(w ResponseWriter, r io.ReadCloser, n int64) io.ReadCloser
func NotFound(w ResponseWriter, r *Request)
func ParseHTTPVersion(vers string) (major, minor int, ok bool)
func ParseTime(text string) (t time.Time, err error)
func ProxyFromEnvironment(req *Request) (*url.URL, error)
func ProxyURL(fixedURL *url.URL) func(*Request) (*url.URL, error)
func Redirect(w ResponseWriter, r *Request, url string, code int)
func Serve(l net.Listener, handler Handler) error
func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker)
func ServeFile(w ResponseWriter, r *Request, name string)
func ServeTLS(l net.Listener, handler Handler, certFile, keyFile string) error
func SetCookie(w ResponseWriter, cookie *Cookie)
func StatusText(code int) string
type Client
func (c *Client) CloseIdleConnections()
func (c *Client) Do(req *Request) (*Response, error)
func (c *Client) Get(url string) (resp *Response, err error)
func (c *Client) Head(url string) (resp *Response, err error)
func (c *Client) Post(url, contentType string, body io.Reader) (resp *Response, err error)
func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error)
type CloseNotifier
type ConnState
func (c ConnState) String() string
type Cookie
func (c *Cookie) String() string
type CookieJar
type Dir
func (d Dir) Open(name string) (File, error)
type File
type FileSystem
type Flusher
type Handler
func FileServer(root FileSystem) Handler
func NotFoundHandler() Handler
func RedirectHandler(url string, code int) Handler
func StripPrefix(prefix string, h Handler) Handler
func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler
type HandlerFunc
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
type Header
func (h Header) Add(key, value string)
func (h Header) Del(key string)
func (h Header) Get(key string) string
func (h Header) Set(key, value string)
func (h Header) Write(w io.Writer) error
func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error
type Hijacker
type ProtocolError
func (pe *ProtocolError) Error() string
type PushOptions
type Pusher
type Request
func NewRequest(method, url string, body io.Reader) (*Request, error)
func ReadRequest(b *bufio.Reader) (*Request, error)
func (r *Request) AddCookie(c *Cookie)
func (r *Request) BasicAuth() (username, password string, ok bool)
func (r *Request) Context() context.Context
func (r *Request) Cookie(name string) (*Cookie, error)
func (r *Request) Cookies() []*Cookie
func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
func (r *Request) FormValue(key string) string
func (r *Request) MultipartReader() (*multipart.Reader, error)
func (r *Request) ParseForm() error
func (r *Request) ParseMultipartForm(maxMemory int64) error
func (r *Request) PostFormValue(key string) string
func (r *Request) ProtoAtLeast(major, minor int) bool
func (r *Request) Referer() string
func (r *Request) SetBasicAuth(username, password string)
func (r *Request) UserAgent() string
func (r *Request) WithContext(ctx context.Context) *Request
func (r *Request) Write(w io.Writer) error
func (r *Request) WriteProxy(w io.Writer) error
type Response
func Get(url string) (resp *Response, err error)
func Head(url string) (resp *Response, err error)
func Post(url, contentType string, body io.Reader) (resp *Response, err error)
func PostForm(url string, data url.Values) (resp *Response, err error)
func ReadResponse(r *bufio.Reader, req *Request) (*Response, error)
func (r *Response) Cookies() []*Cookie
func (r *Response) Location() (*url.URL, error)
func (r *Response) ProtoAtLeast(major, minor int) bool
func (r *Response) Write(w io.Writer) error
type ResponseWriter
type RoundTripper
func NewFileTransport(fs FileSystem) RoundTripper
type SameSite
type ServeMux
func NewServeMux() *ServeMux
func (mux *ServeMux) Handle(pattern string, handler Handler)
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
type Server
func (srv *Server) Close() error
func (srv *Server) ListenAndServe() error
func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error
func (srv *Server) RegisterOnShutdown(f func())
func (srv *Server) Serve(l net.Listener) error
func (srv *Server) ServeTLS(l net.Listener, certFile, keyFile string) error
func (srv *Server) SetKeepAlivesEnabled(v bool)
func (srv *Server) Shutdown(ctx context.Context) error
type Transport
func (t *Transport) CancelRequest(req *Request)
func (t *Transport) CloseIdleConnections()
func (t *Transport) RegisterProtocol(scheme string, rt RoundTripper)
func (t *Transport) RoundTrip(req *Request) (*Response, error)
```

## Package Files

client.go cookie.go doc.go filetransport.go fs.go h2_bundle.go header.go http.go jar.go method.go request.go response.go roundtrip.go server.go sniff.go socks_bundle.go status.go transfer.go transport.go

  