/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: http
 **Element: http.convar
 **Type: const,var
 ------------------------------------------------------------------------------------
 **Definition:
 const (
     MethodGet     = "GET"
     MethodHead    = "HEAD"
     MethodPost    = "POST"
     MethodPut     = "PUT"
     MethodPatch   = "PATCH" // RFC 5789
     MethodDelete  = "DELETE"
     MethodConnect = "CONNECT"
     MethodOptions = "OPTIONS"
     MethodTrace   = "TRACE"
 )
 Common HTTP methods.
 Unless otherwise noted, these are defined in RFC 7231 section 4.3.
 const (
     StatusContinue           = 100 // RFC 7231, 6.2.1
     StatusSwitchingProtocols = 101 // RFC 7231, 6.2.2
     StatusProcessing         = 102 // RFC 2518, 10.1
     StatusOK                   = 200 // RFC 7231, 6.3.1
     StatusCreated              = 201 // RFC 7231, 6.3.2
     StatusAccepted             = 202 // RFC 7231, 6.3.3
     StatusNonAuthoritativeInfo = 203 // RFC 7231, 6.3.4
     StatusNoContent            = 204 // RFC 7231, 6.3.5
     StatusResetContent         = 205 // RFC 7231, 6.3.6
     StatusPartialContent       = 206 // RFC 7233, 4.1
     StatusMultiStatus          = 207 // RFC 4918, 11.1
     StatusAlreadyReported      = 208 // RFC 5842, 7.1
     StatusIMUsed               = 226 // RFC 3229, 10.4.1
     StatusMultipleChoices  = 300 // RFC 7231, 6.4.1
     StatusMovedPermanently = 301 // RFC 7231, 6.4.2
     StatusFound            = 302 // RFC 7231, 6.4.3
     StatusSeeOther         = 303 // RFC 7231, 6.4.4
     StatusNotModified      = 304 // RFC 7232, 4.1
     StatusUseProxy         = 305 // RFC 7231, 6.4.5
     StatusTemporaryRedirect = 307 // RFC 7231, 6.4.7
     StatusPermanentRedirect = 308 // RFC 7538, 3
     StatusBadRequest                   = 400 // RFC 7231, 6.5.1
     StatusUnauthorized                 = 401 // RFC 7235, 3.1
     StatusPaymentRequired              = 402 // RFC 7231, 6.5.2
     StatusForbidden                    = 403 // RFC 7231, 6.5.3
     StatusNotFound                     = 404 // RFC 7231, 6.5.4
     StatusMethodNotAllowed             = 405 // RFC 7231, 6.5.5
     StatusNotAcceptable                = 406 // RFC 7231, 6.5.6
     StatusProxyAuthRequired            = 407 // RFC 7235, 3.2
     StatusRequestTimeout               = 408 // RFC 7231, 6.5.7
     StatusConflict                     = 409 // RFC 7231, 6.5.8
     StatusGone                         = 410 // RFC 7231, 6.5.9
     StatusLengthRequired               = 411 // RFC 7231, 6.5.10
     StatusPreconditionFailed           = 412 // RFC 7232, 4.2
     StatusRequestEntityTooLarge        = 413 // RFC 7231, 6.5.11
     StatusRequestURITooLong            = 414 // RFC 7231, 6.5.12
     StatusUnsupportedMediaType         = 415 // RFC 7231, 6.5.13
     StatusRequestedRangeNotSatisfiable = 416 // RFC 7233, 4.4
     StatusExpectationFailed            = 417 // RFC 7231, 6.5.14
     StatusTeapot                       = 418 // RFC 7168, 2.3.3
     StatusMisdirectedRequest           = 421 // RFC 7540, 9.1.2
     StatusUnprocessableEntity          = 422 // RFC 4918, 11.2
     StatusLocked                       = 423 // RFC 4918, 11.3
     StatusFailedDependency             = 424 // RFC 4918, 11.4
     StatusTooEarly                     = 425 // RFC 8470, 5.2.
     StatusUpgradeRequired              = 426 // RFC 7231, 6.5.15
     StatusPreconditionRequired         = 428 // RFC 6585, 3
     StatusTooManyRequests              = 429 // RFC 6585, 4
     StatusRequestHeaderFieldsTooLarge  = 431 // RFC 6585, 5
     StatusUnavailableForLegalReasons   = 451 // RFC 7725, 3
     StatusInternalServerError           = 500 // RFC 7231, 6.6.1
     StatusNotImplemented                = 501 // RFC 7231, 6.6.2
     StatusBadGateway                    = 502 // RFC 7231, 6.6.3
     StatusServiceUnavailable            = 503 // RFC 7231, 6.6.4
     StatusGatewayTimeout                = 504 // RFC 7231, 6.6.5
     StatusHTTPVersionNotSupported       = 505 // RFC 7231, 6.6.6
     StatusVariantAlsoNegotiates         = 506 // RFC 2295, 8.1
     StatusInsufficientStorage           = 507 // RFC 4918, 11.5
     StatusLoopDetected                  = 508 // RFC 5842, 7.2
     StatusNotExtended                   = 510 // RFC 2774, 7
     StatusNetworkAuthenticationRequired = 511 // RFC 6585, 6
 )
 HTTP status codes as registered with IANA. See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
 const DefaultMaxHeaderBytes = 1 << 20 // 1 MB
 DefaultMaxHeaderBytes is the maximum permitted size of the headers in an HTTP request. This can be overridden by setting Server.MaxHeaderBytes.
 const DefaultMaxIdleConnsPerHost = 2
 DefaultMaxIdleConnsPerHost is the default value of Transport's MaxIdleConnsPerHost.
 const TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"
 TimeFormat is the time format to use when generating times in HTTP headers. It is like time.RFC1123 but hard-codes GMT as the time zone. The time being formatted must be in UTC for Format to generate the correct format.
 For parsing this time format, see ParseTime.
 const TrailerPrefix = "Trailer:"
 TrailerPrefix is a magic prefix for ResponseWriter.Header map keys that, if present, signals that the map entry is actually for the response trailers, and not the response headers. The prefix is stripped after the ServeHTTP call finishes and the values are sent in the trailers.
 This mechanism is intended only for trailers that are not known prior to the headers being written. If the set of trailers is fixed or known before the header is written, the normal Go trailers mechanism is preferred:
 https://golang.org/pkg/net/http/#ResponseWriter
 https://golang.org/pkg/net/http/#example_ResponseWriter_trailers
 Variables
 var (
     // ErrNotSupported is returned by the Push method of Pusher
     // implementations to indicate that HTTP/2 Push support is not
     // available.
     ErrNotSupported = &ProtocolError{"feature not supported"}
     // Deprecated: ErrUnexpectedTrailer is no longer returned by
     // anything in the net/http package. Callers should not
     // compare errors against this variable.
     ErrUnexpectedTrailer = &ProtocolError{"trailer header without chunked transfer encoding"}
     // ErrMissingBoundary is returned by Request.MultipartReader when the
     // request's Content-Type does not include a "boundary" parameter.
     ErrMissingBoundary = &ProtocolError{"no multipart boundary param in Content-Type"}
     // ErrNotMultipart is returned by Request.MultipartReader when the
     // request's Content-Type is not multipart/form-data.
     ErrNotMultipart = &ProtocolError{"request Content-Type isn't multipart/form-data"}
     // Deprecated: ErrHeaderTooLong is no longer returned by
     // anything in the net/http package. Callers should not
     // compare errors against this variable.
     ErrHeaderTooLong = &ProtocolError{"header too long"}
     // Deprecated: ErrShortBody is no longer returned by
     // anything in the net/http package. Callers should not
     // compare errors against this variable.
     ErrShortBody = &ProtocolError{"entity body too short"}
     // Deprecated: ErrMissingContentLength is no longer returned by
     // anything in the net/http package. Callers should not
     // compare errors against this variable.
     ErrMissingContentLength = &ProtocolError{"missing ContentLength in HEAD response"}
 )
 var (
     // ErrBodyNotAllowed is returned by ResponseWriter.Write calls
     // when the HTTP method or response code does not permit a
     // body.
     ErrBodyNotAllowed = errors.New("http: request method or response status code does not allow body")
     // ErrHijacked is returned by ResponseWriter.Write calls when
     // the underlying connection has been hijacked using the
     // Hijacker interface. A zero-byte write on a hijacked
     // connection will return ErrHijacked without any other side
     // effects.
     ErrHijacked = errors.New("http: connection has been hijacked")
     // ErrContentLength is returned by ResponseWriter.Write calls
     // when a Handler set a Content-Length response header with a
     // declared size and then attempted to write more bytes than
     // declared.
     ErrContentLength = errors.New("http: wrote more than the declared Content-Length")
     // Deprecated: ErrWriteAfterFlush is no longer returned by
     // anything in the net/http package. Callers should not
     // compare errors against this variable.
     ErrWriteAfterFlush = errors.New("unused")
 )
 Errors used by the HTTP server.
 var (
     // ServerContextKey is a context key. It can be used in HTTP
     // handlers with context.WithValue to access the server that
     // started the handler. The associated value will be of
     // type *Server.
     ServerContextKey = &contextKey{"http-server"}
     // LocalAddrContextKey is a context key. It can be used in
     // HTTP handlers with context.WithValue to access the local
     // address the connection arrived on.
     // The associated value will be of type net.Addr.
     LocalAddrContextKey = &contextKey{"local-addr"}
 )
 var DefaultClient = &Client{}
 DefaultClient is the default Client and is used by Get, Head, and Post.
 var DefaultServeMux = &defaultServeMux
 DefaultServeMux is the default ServeMux used by Serve.
 var ErrAbortHandler = errors.New("net/http: abort Handler")
 ErrAbortHandler is a sentinel panic value to abort a handler. While any panic from ServeHTTP aborts the response to the client, panicking with ErrAbortHandler also suppresses logging of a stack trace to the server's error log.
 var ErrBodyReadAfterClose = errors.New("http: invalid Read on closed Body")
 ErrBodyReadAfterClose is returned when reading a Request or Response Body after the body has been closed. This typically happens when the body is read after an HTTP Handler calls WriteHeader or Write on its ResponseWriter.
 var ErrHandlerTimeout = errors.New("http: Handler timeout")
 ErrHandlerTimeout is returned on ResponseWriter Write calls in handlers which have timed out.
 var ErrLineTooLong = internal.ErrLineTooLong
 ErrLineTooLong is returned when reading request or response bodies with malformed chunked encoding.
 var ErrMissingFile = errors.New("http: no such file")
 ErrMissingFile is returned by FormFile when the provided file field name is either not present in the request or not a file field.
 var ErrNoCookie = errors.New("http: named cookie not present")
 ErrNoCookie is returned by Request's Cookie method when a cookie is not found.
 var ErrNoLocation = errors.New("http: no Location header in response")
 ErrNoLocation is returned by Response's Location method when no Location header is present.
 var ErrServerClosed = errors.New("http: Server closed")
 ErrServerClosed is returned by the Server's Serve, ServeTLS, ListenAndServe, and ListenAndServeTLS methods after a call to Shutdown or Close.
 var ErrSkipAltProtocol = errors.New("net/http: skip alternate protocol")
 ErrSkipAltProtocol is a sentinel error value defined by Transport.RegisterProtocol.
 var ErrUseLastResponse = errors.New("net/http: use last response")
 ErrUseLastResponse can be returned by Client.CheckRedirect hooks to control how redirects are processed. If returned, the next request is not sent and the most recent response is returned with its body unclosed.
 var NoBody = noBody{}
 NoBody is an io.ReadCloser with no bytes. Read always returns EOF and Close always returns nil. It can be used in an outgoing client request to explicitly signal that a request has zero bytes. An alternative, however, is to simply set Request.Body to nil.
 ------------------------------------------------------------------------------------
 **Description:
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
