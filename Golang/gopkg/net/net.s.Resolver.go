/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.Resolver
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Resolver struct {
     // PreferGo controls whether Go's built-in DNS resolver is preferred
     // on platforms where it's available. It is equivalent to setting
     // GODEBUG=netdns=go, but scoped to just this resolver.
     PreferGo bool
     // StrictErrors controls the behavior of temporary errors
     // (including timeout, socket errors, and SERVFAIL) when using
     // Go's built-in resolver. For a query composed of multiple
     // sub-queries (such as an A+AAAA address lookup, or walking the
     // DNS search list), this option causes such errors to abort the
     // whole query instead of returning a partial result. This is
     // not enabled by default because it may affect compatibility
     // with resolvers that process AAAA queries incorrectly.
     StrictErrors bool
     // Dial optionally specifies an alternate dialer for use by
     // Go's built-in DNS resolver to make TCP and UDP connections
     // to DNS services. The host in the address parameter will
     // always be a literal IP address and not a host name, and the
     // port in the address parameter will be a literal port number
     // and not a service name.
     // If the Conn returned is also a PacketConn, sent and received DNS
     // messages must adhere to RFC 1035 section 4.2.1, "UDP usage".
     // Otherwise, DNS messages transmitted over Conn must adhere
     // to RFC 7766 section 5, "Transport Protocol Selection".
     // If nil, the default dialer is used.
     Dial func(ctx context.Context, network, address string) (Conn, error)
     // contains filtered or unexported fields
 }
 func (r *Resolver) LookupAddr(ctx context.Context, addr string) (names []string, err error)
 func (r *Resolver) LookupCNAME(ctx context.Context, host string) (cname string, err error)
 func (r *Resolver) LookupHost(ctx context.Context, host string) (addrs []string, err error)
 func (r *Resolver) LookupIPAddr(ctx context.Context, host string) ([]IPAddr, error)
 func (r *Resolver) LookupMX(ctx context.Context, name string) ([]*MX, error)
 func (r *Resolver) LookupNS(ctx context.Context, name string) ([]*NS, error)
 func (r *Resolver) LookupPort(ctx context.Context, network, service string) (port int, err error)
 func (r *Resolver) LookupSRV(ctx context.Context, service, proto, name string) (cname string, addrs []*SRV, err error)
 func (r *Resolver) LookupTXT(ctx context.Context, name string) ([]string, error)
 ------------------------------------------------------------------------------------
 **Description:
 A Resolver looks up names and numbers.
 A nil *Resolver is equivalent to a zero Resolver.
 LookupAddr performs a reverse lookup for the given address, returning a list of names mapping to that address.
 LookupCNAME returns the canonical name for the given host. Callers that do not care about the canonical name can call LookupHost or LookupIP directly; both take care of resolving the canonical name as part of the lookup.
 A canonical name is the final name after following zero or more CNAME records. LookupCNAME does not return an error if host does not contain DNS "CNAME" records, as long as host resolves to address records.
 LookupHost looks up the given host using the local resolver. It returns a slice of that host's addresses.
 LookupIPAddr looks up host using the local resolver. It returns a slice of that host's IPv4 and IPv6 addresses.
 LookupMX returns the DNS MX records for the given domain name sorted by preference.
 LookupNS returns the DNS NS records for the given domain name.
 LookupPort looks up the port for the given network and service.
 LookupSRV tries to resolve an SRV query of the given service, protocol, and domain name. The proto is "tcp" or "udp". The returned records are sorted by priority and randomized by weight within a priority.
 LookupSRV constructs the DNS name to look up following RFC 2782. That is, it looks up _service._proto.name. To accommodate services publishing SRV records under non-standard names, if both service and proto are empty strings, LookupSRV looks up name directly.
 LookupTXT returns the DNS TXT records for the given domain name.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
