package main

import (
	"fmt"
	"os"

	"github.com/miekg/dns"
)

func main() {
	if len(os.Args) > 1 {
		bytes, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var msg dns.Msg
		err = msg.Unpack(bytes)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Rcode: ", dns.RcodeToString[msg.MsgHdr.Rcode])
		fmt.Println("HEADER:")
		fmt.Printf("%+v\n", msg.MsgHdr)
		fmt.Println("\nQUESTION:", len(msg.Question))
		for _, q := range msg.Question {
			fmt.Printf("Name [%v] Class [%v] Type [%v]\n", q.Name, q.Qclass, dns.TypeToString[q.Qtype])
		}
		fmt.Println("\nANSWER:", len(msg.Answer))
		for _, a := range msg.Answer {
			fmt.Println(a)
		}
		fmt.Println("\nAUTHORITATIVE:", len(msg.Ns))
		for _, n := range msg.Ns {
			fmt.Println(n)
		}
		fmt.Println("\nEXTRA:", len(msg.Extra))
		for _, e := range msg.Extra {
			fmt.Println(e)
		}
	} else {
		fmt.Println("missing packet filename as cmd line argument")
	}
}
