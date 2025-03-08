package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/leaanthony/clir"
	"github.com/miekg/dns"
	"github.com/shynome/doh-client"
)

type Protocol int

const (
	UDP Protocol = iota
	TCP
	TLS
	DOHPOST // DoH POST
	DOHGET  // DoH GET
)

var ProtoString = map[Protocol]string{
	UDP:     "udp",
	TCP:     "tcp",
	TLS:     "tcp-tls",
	DOHPOST: "doh-post",
	DOHGET:  "doh", // doh = DoH GET
}

type QueryConfig struct {
	Host               string
	Port               string
	Mode               Protocol
	QuestionType       string
	FQDN               string
	Raw                bool
	PrintRequestBase64 bool
	BinRequestToFile   string
	BinResponseToFile  string
	DnsSec             bool
	JsonResponseToFile string
}

func main() {
	cli := clir.NewCli("dig", "A lightweight dig replacement", "v0.0.2")
	cli.LongDescription("ex: dig @8.8.4.4 google.com -t MX +dnssec")
	host := "8.8.8.8"
	port := "53"
	questionType := "A"
	var fqdn string

	tcpMode := false
	tlsMode := false
	dohGetMode := false
	dohPostMode := false
	noColor := false
	raw := false
	printRequestBase64 := false
	dnssec := false

	binRequestToFile := ""
	binResponseToFile := ""
	jsonResponseToFile := ""

	cli.StringFlag("host", "DNS server hostname/ip to use", &host)
	cli.StringFlag("port", "port to connect on", &port)
	cli.BoolFlag("tcp", "use TCP", &tcpMode)
	cli.BoolFlag("tls", "use TLS (DoT)", &tlsMode)
	cli.BoolFlag("doh", "use DoH (GET) json format", &dohGetMode)
	cli.BoolFlag("doh-post", "use DoH via HTTP POST wire format", &dohPostMode)
	cli.StringFlag("t", "question type, ex: A, NS, MX, etc.", &questionType)
	cli.BoolFlag("nc", "disable ansi colors", &noColor)
	cli.BoolFlag("raw", "show raw response", &raw)
	cli.BoolFlag("print-request-base64", "print request base64", &printRequestBase64)
	cli.StringFlag("bin-request-to-file", "print request binary to file", &binRequestToFile)
	cli.StringFlag("bin-response-to-file", "print response binary to file", &binResponseToFile)
	cli.StringFlag("json-response-to-file", "print response json to file", &jsonResponseToFile)

	cli.Action(func() error {
		for _, arg := range cli.OtherArgs() {
			if arg[0] == '@' && host == "8.8.8.8" {
				host = arg[1:]
				continue
			}
			if strings.EqualFold(arg, "+dnssec") {
				dnssec = true
				continue
			}
			if fqdn == "" {
				fqdn = arg
			}
		}
		if noColor {
			color.NoColor = true // disables colorized output
		}

		proto := UDP
		if tcpMode {
			proto = TCP
		} else if tlsMode {
			proto = TLS
			if port == "53" {
				port = "853"
			}
		} else if dohGetMode {
			proto = DOHGET
			port = "443"
		} else if dohPostMode {
			proto = DOHPOST
			port = "443"
		}

		if fqdn == "" {
			return errors.New("No fqdn provided")
		}
		qc := &QueryConfig{
			Host:               host,
			Port:               port,
			Mode:               proto,
			FQDN:               fqdn,
			QuestionType:       strings.ToUpper(questionType),
			Raw:                raw,
			PrintRequestBase64: printRequestBase64,
			BinRequestToFile:   binRequestToFile,
			BinResponseToFile:  binResponseToFile,
			DnsSec:             dnssec,
			JsonResponseToFile: jsonResponseToFile,
		}
		Run(qc)
		return nil
	})

	if err := cli.Run(); err != nil {
		fmt.Println(err)
		cli.PrintHelp()
		os.Exit(1)
	}
	os.Exit(0)
}

func Run(qc *QueryConfig) int {
	return doLookup(qc, false)
}

func doLookup(qc *QueryConfig, trunc bool) int {
	timeNow := time.Now()
	questionStringToType := make(map[string]uint16, len(dns.TypeToString))
	for t, q := range dns.TypeToString {
		questionStringToType[q] = t
	}

	if trunc && qc.Mode == UDP {
		qc.Mode = TCP
	}
	m := new(dns.Msg)
	m.Compress = true
	fqdn := qc.FQDN
	if string(qc.FQDN[len(qc.FQDN)-1]) != "." {
		fqdn = fqdn + "."
	}
	m.SetQuestion(fqdn, questionStringToType[qc.QuestionType])
	m.RecursionDesired = true
	if qc.DnsSec {
		o := &dns.OPT{
			Hdr: dns.RR_Header{
				Name:   fqdn,
				Rrtype: dns.TypeOPT,
			},
		}
		o.SetDo()
		o.SetUDPSize(dns.DefaultMsgSize)
		m.Extra = append(m.Extra, o)
	}

	// Do we only want to show the base64 of the request? (useful for manual DoH wireformat requests)
	if qc.PrintRequestBase64 {
		packedMsg, err := m.Pack()
		if err != nil {
			fmt.Println(err)
			return 1
		}
		fmt.Printf("Request Base64 encoded: %v\n", color.GreenString(base64.StdEncoding.EncodeToString(packedMsg)))
		return 0
	}

	// Do we want to save the request in binary format to a file?
	if qc.BinRequestToFile != "" {
		packedMsg, err := m.Pack()
		if err != nil {
			fmt.Println(err)
			return 1
		}
		err = os.WriteFile(qc.BinRequestToFile, packedMsg, 0644)
		if err != nil {
			fmt.Println(err)
			return 1
		}
		fmt.Printf("Wrote request binary to file: %v\n", color.GreenString(qc.BinRequestToFile))
	}

	fmt.Printf("Host: %v, Port: %v, Proto: %v, DNSSEC: %v, FQDN: %v, Question Type: %v\n",
		color.GreenString(qc.Host),
		color.GreenString(qc.Port),
		color.GreenString(ProtoString[qc.Mode]),
		color.GreenString(strconv.FormatBool(qc.DnsSec)),
		color.CyanString(fqdn),
		color.YellowString(qc.QuestionType))

	var r *dns.Msg
	var err error
	// for non DOH queries
	if qc.Mode != DOHGET && qc.Mode != DOHPOST {
		c := new(dns.Client)
		c.Net = ProtoString[qc.Mode]
		r, _, err = c.Exchange(m, net.JoinHostPort(qc.Host, qc.Port))
		if err != nil {
			fmt.Println(err)
			return 1
		}
	} else if qc.Mode == DOHGET {
		co := &dns.Conn{Conn: doh.NewConn(nil, nil, qc.Host)}
		if err = co.WriteMsg(m); err != nil {
			fmt.Println(err)
			return 1
		}
		r, err = co.ReadMsg()
		if err != nil {
			fmt.Println(err)
			return 1
		}
	} else if qc.Mode == DOHPOST {
		dohConn := doh.NewConn(nil, nil, qc.Host)
		dohConn.HttpGet = false
		co := &dns.Conn{Conn: dohConn}
		if err = co.WriteMsg(m); err != nil {
			fmt.Println(err)
			return 1
		}
		r, err = co.ReadMsg()
		if err != nil {
			fmt.Println(err)
			return 1
		}
	} else {
		fmt.Println("Unknown dns request mode, exiting.")
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
				// TODO: Verify RRSIGs?
				// if a.Header().Rrtype == dns.TypeRRSIG {
				// 	s := a.(*dns.RRSIG)
				// 	fmt.Printf("algorithm: %+v\n", s.Algorithm)
				// 	err := s.Verify(&dns.DNSKEY{}, r.Answer)
				// 	fmt.Println(err)
				// }
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

		// Do we want to save the response in binary format to a file?
		if qc.BinResponseToFile != "" {
			packedMsg, err := r.Pack()
			if err != nil {
				fmt.Println(err)
				return 1
			}
			err = os.WriteFile(qc.BinResponseToFile, packedMsg, 0644)
			if err != nil {
				fmt.Println(err)
				return 1
			}
			fmt.Printf("Wrote response binary to file: %v\n", color.GreenString(qc.BinResponseToFile))
		}
		// Do we want to save the response in json format to a file?
		if qc.JsonResponseToFile != "" {
			jStr, err := json.MarshalIndent(r, "", "    ")
			if err != nil {
				fmt.Println(err)
				return 1
			}
			err = os.WriteFile(qc.JsonResponseToFile, jStr, 0644)
			if err != nil {
				fmt.Println(err)
				return 1
			}
			fmt.Printf("Wrote response json to file: %v\n", color.GreenString(qc.JsonResponseToFile))
		}
		return 0
	}
}
