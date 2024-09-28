package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/miekg/dns"
)

type Protocol int

const (
	UDP Protocol = iota
	TCP
	TLS
)

var ProtoString = map[Protocol]string{
	UDP: "udp",
	TCP: "tcp",
	TLS: "tcp-tls",
}

type QueryConfig struct {
	Host         string
	Port         string
	Mode         Protocol
	QuestionType string
	FQDN         string
}

func main() {
	host := flag.String("h", "8.8.8.8", "DNS server hostname/ip to use")
	port := flag.String("p", "53", "Port to connect on")
	tcpMode := flag.Bool("tcp", false, "use TCP")
	tlsMode := flag.Bool("tls", false, "use TLS (DoT)")
	questionType := flag.String("t", "A", "question dns.Type, ex: A, AAAA, NS, etc.")
	fqdn := flag.String("fqdn", "google.com", "fqdn to lookup")

	flag.Parse()

	proto := UDP
	if *tcpMode {
		proto = TCP
	} else if *tlsMode {
		proto = TLS
		if *port == "53" {
			*port = "853"
		}
	}
	qc := &QueryConfig{Host: *host, Port: *port, Mode: proto, FQDN: *fqdn, QuestionType: *questionType}
	os.Exit(Run(qc))
}

func Run(qc *QueryConfig) int {
	fmt.Printf("Host: %v, Port: %v, Proto: %v, FQDN: %v, Question Type: %v\n", qc.Host, qc.Port, ProtoString[qc.Mode], qc.FQDN, qc.QuestionType)
	return doLookup(qc, false)
}

func doLookup(qc *QueryConfig, trunc bool) int {
	timeNow := time.Now()
	questionStringToType := make(map[string]uint16, len(dns.TypeToString))
	for t, q := range dns.TypeToString {
		questionStringToType[q] = t
	}

	c := new(dns.Client)
	if trunc && qc.Mode == UDP {
		qc.Mode = TCP
	}
	c.Net = ProtoString[qc.Mode]
	m := new(dns.Msg)
	m.Compress = true
	m.SetQuestion(qc.FQDN+".", questionStringToType[qc.QuestionType])
	m.RecursionDesired = true
	r, _, err := c.Exchange(m, net.JoinHostPort(qc.Host, qc.Port))
	if err != nil {
		fmt.Println(err)
		return 1
	}
	if r.Rcode != dns.RcodeSuccess {
		fmt.Println("Rcode:", dns.RcodeToString[r.Rcode])
		return 1
	}
	if r.Truncated && len(r.Answer) == 0 && qc.Mode == UDP {
		fmt.Println("WARNING: truncated response, will retry with tcp instead")
		return doLookup(qc, true)
	} else {
		timeElapsed := time.Since(timeNow)
		fmt.Printf("ANSWERS: %v, BYTES RECEIVED: %v, IN: %v\n\n", len(r.Answer), r.Len(), timeElapsed)
		for _, a := range r.Answer {
			fmt.Printf("%+v\n", a)
		}
		return 0
	}
}
