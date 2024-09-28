package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
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
	Raw          bool
}

func main() {
	host := flag.String("h", "8.8.8.8", "DNS server hostname/ip to use")
	port := flag.String("p", "53", "Port to connect on")
	tcpMode := flag.Bool("tcp", false, "use TCP")
	tlsMode := flag.Bool("tls", false, "use TLS (DoT)")
	questionType := flag.String("t", "A", "question dns.Type, ex: A, AAAA, NS, etc.")
	fqdn := flag.String("fqdn", "google.com", "fqdn to lookup")
	noColor := flag.Bool("nc", false, "no color")
	raw := flag.Bool("raw", false, "show raw response")
	flag.Parse()

	if *noColor {
		color.NoColor = true // disables colorized output
	}

	proto := UDP
	if *tcpMode {
		proto = TCP
	} else if *tlsMode {
		proto = TLS
		if *port == "53" {
			*port = "853"
		}
	}
	qc := &QueryConfig{Host: *host, Port: *port, Mode: proto, FQDN: *fqdn, QuestionType: *questionType, Raw: *raw}
	os.Exit(Run(qc))
}

func Run(qc *QueryConfig) int {
	fmt.Printf("Host: %v, Port: %v, Proto: %v, FQDN: %v, Question Type: %v\n",
		color.GreenString(qc.Host),
		color.GreenString(qc.Port),
		color.GreenString(ProtoString[qc.Mode]),
		color.CyanString(qc.FQDN),
		color.YellowString(qc.QuestionType))
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
		fmt.Printf("Id: %v, Opcode: %v, AA: %v, TC: %v, RD: %v, RA: %v, Z: %v, RCODE: %v\n",
			r.MsgHdr.Id,
			r.MsgHdr.Opcode,
			r.MsgHdr.Authoritative,
			r.MsgHdr.Truncated,
			r.MsgHdr.RecursionDesired,
			r.MsgHdr.RecursionAvailable,
			r.MsgHdr.Zero,
			color.RedString(dns.RcodeToString[r.MsgHdr.Rcode]))
		// fmt.Println("Rcode:", color.RedString(dns.RcodeToString[r.Rcode]))
		return 1
	}
	if r.Truncated && len(r.Answer) == 0 && qc.Mode == UDP {
		fmt.Println("WARNING: truncated response, will retry with tcp instead")
		return doLookup(qc, true)
	} else {
		timeElapsed := time.Since(timeNow)
		fmt.Printf("Id: %v, Opcode: %v, AA: %v, TC: %v, RD: %v, RA: %v, Z: %v, RCODE: %v\n",
			r.MsgHdr.Id,
			r.MsgHdr.Opcode,
			r.MsgHdr.Authoritative,
			r.MsgHdr.Truncated,
			r.MsgHdr.RecursionDesired,
			r.MsgHdr.RecursionAvailable,
			r.MsgHdr.Zero,
			color.GreenString(dns.RcodeToString[r.MsgHdr.Rcode]))
		fmt.Printf("QUERY: %v; ANSWER: %v; AUTHORITY: %v; ADDITIONAL: %v\n\n", len(r.Question), len(r.Answer), len(r.Ns), len(r.Extra))

		if len(r.Answer) == 0 && len(r.Ns) > 0 {
			for _, n := range r.Ns {
				fmt.Printf("%+v\n", n)
			}
		} else if len(r.Answer) > 0 {
			for _, a := range r.Answer {
				fmt.Printf("%+v\n", a)
			}
		}
		fmt.Printf("\nBYTES RECEIVED: %v, IN: %v\n", color.GreenString(strconv.Itoa(r.Len())), color.CyanString(timeElapsed.String()))

		if qc.Raw {
			packedMsg, err := r.Pack()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("\n%v", hex.Dump(packedMsg))
		}
		return 0
	}
}
