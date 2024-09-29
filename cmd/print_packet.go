package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"regexp"
	"strings"

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

		// try unpacking as is first, hopefully file provided to us is in binary format
		err = msg.Unpack(bytes)
		if err != nil {
			// fmt.Println(err)
			// try to see if we can convert this file from a bad dig text w/ hex output to binary
			// for example, file can be this and we can try to convert it to binary and then Unpack it
			//
			// ;; Got bad packet: extra input data
			// 58 bytes
			// b6 0d 81 80 00 01 00 01 00 00 00 00 06 73 6c 61          .............sla
			// 63 6b 62 03 63 6f 6d 00 00 41 00 01 c0 0c 00 41          ckb.com..A.....A
			// 00 01 00 00 00 1e 00 12 00 00 00 00 01 00 03 02          ................
			// 68 32 00 04 00 04 2d fd 83 e2                            h2....-...

			convertedBytes, err := convertTextWithHexBytesToBinary(bytes)
			if err != nil {
				fmt.Println(err)
				// finally we're out of possible ways to use this data, exit the program
				os.Exit(1)
			}
			err = msg.Unpack(convertedBytes)
			if err != nil {
				// if we still cannot Unpack these bytes to a valid msg, exit the program
				fmt.Println(err)
				os.Exit(1)
			}
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

func convertTextWithHexBytesToBinary(bytes []byte) ([]byte, error) {
	var sb strings.Builder
	s := string(bytes)
	for _, line := range strings.Split(s, "\n") {
		r := regexp.MustCompile(`(?i)([0-9a-f]{2}) ([0-9a-f]{2})|([0-9a-f]){2}$`)
		matches := r.FindAllStringSubmatch(line, -1)
		if len(matches) > 0 {
			for _, mb := range matches {
				if mb[1] != "" && mb[2] != "" {
					sb.WriteString(mb[1] + mb[2])
				} else {
					sb.WriteString(mb[0])
				}
			}
		}
	}
	// fmt.Println(sb.String())
	convertedBytes, err := hex.DecodeString(sb.String())
	if err != nil {
		return []byte{}, err
	}
	return convertedBytes, nil
}
