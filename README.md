# A lightweight dig clone

## Example usage

### Building

```
go build
go build -o pp cmd/print_packet.go
```

### Building for Windows
```
GOOS=windows GOARCH=386 go build -o dig.exe main.go
GOOS=windows GOARCH=386 go build -o pp.exe cmd/print_packet.go 
```

### Running

Example querying for `ANY` record type:
```
dig.exe slackb.com -t ANY
Host: 8.8.8.8, Port: 53, Proto: udp, FQDN: slackb.com, Question Type: ANY
WARNING: truncated response, will retry with tcp instead
Id: 15302, Opcode: 0, AA: false, TC: false, RD: true, RA: true, Z: false, RCODE: NOERROR
QUERY: 1; ANSWER: 28; AUTHORITY: 0; ADDITIONAL: 0

slackb.com.     60      IN      A       3.90.158.208
slackb.com.     60      IN      A       34.195.221.192
slackb.com.     60      IN      A       3.91.140.69
slackb.com.     60      IN      A       18.213.32.120
slackb.com.     60      IN      A       54.236.104.103
slackb.com.     60      IN      A       52.3.167.79
slackb.com.     60      IN      A       34.202.253.140
slackb.com.     60      IN      A       52.200.46.145
slackb.com.     60      IN      A       44.216.98.239
slackb.com.     60      IN      A       44.205.171.153
slackb.com.     60      IN      RRSIG   A 13 2 60 20240929004001 20240928223901 15779 slackb.com. zkF72VzQZz+OxKOACbG3uOax1wsMlxCSc1iaiR6Fls+M2+vHKyetVHkhapiuC2axz97uQQo206zimQ8/Z4Xs/w==
slackb.com.     21600   IN      NS      ns-352.awsdns-44.com.
slackb.com.     21600   IN      NS      ns-1347.awsdns-40.org.
slackb.com.     21600   IN      NS      ns-542.awsdns-03.net.
slackb.com.     21600   IN      NS      ns-1654.awsdns-14.co.uk.
slackb.com.     21600   IN      RRSIG   NS 13 2 172800 20241001003901 20240928223901 15779 slackb.com. KPnOoUVjTOFr7eEfbCA3eUpC/9CG1NTgZuWWnqaT1PmHlbQ+46mE8fM8x0U/lzWrn4y9e+Z39zkPbeZBnITKsA==
slackb.com.     900     IN      SOA     ns-1654.awsdns-14.co.uk. awsdns-hostmaster.amazon.com. 1 7200 900 1209600 86400
slackb.com.     900     IN      RRSIG   SOA 13 2 900 20240929005401 20240928223901 15779 slackb.com. BIenY3kEg3JCQvZ8z+qiPZbggY6pTcdi2fHhKYZebCsEedwBzmwb8Kq1Nc1p3zI0RecZZYlIQZRQEbx/a72smQ==
slackb.com.     300     IN      TXT     "v=spf1 -all"
slackb.com.     300     IN      RRSIG   TXT 13 2 300 20240929004401 20240928223901 15779 slackb.com. mGU0wrtCaFPvNyO7+IvJFr7uI44+m4ASb8TAb870yRZy+eOkcpsAdPBR91/0jI8U+cvTviKKOmb2pRHHgeQJ1Q==
slackb.com.     3600    IN      DNSKEY  256 3 13 cV5+Jxz0KlpwnBS009OppqW/wEAGECGElUSAsny6Dk5NNSK6M1gGOqXql6Yke2JqMlP30xycgTIVcgvJOQtMeg==
slackb.com.     3600    IN      DNSKEY  257 3 13 bSr2fT2hBdGDUltm3fPrF4HB6x+2k2Ol51pImwom4jGa1J66Bo6g4x/2a2iAdjzejf8d4cCeXkuQvHkeDpLV+w==
slackb.com.     3600    IN      DNSKEY  256 3 13 g4jECOQIRAufN1jKoEHGoR9ZSld0jNy/AqbiMJPh1YlNz4k7eAPjS/YuyhYOw7784LfzcGGXDIS+04Q6J4D0mQ==
slackb.com.     3600    IN      RRSIG   DNSKEY 13 2 3600 20240929020000 20240928150000 55613 slackb.com. nMTJ4QcSNgTKHZTZXbCJtLFqkpWuc+s3ugYL7VFt81QxL4XiLvX0jS6bvUJl5BSw9QO91tHbagKgLweGwj3VlQ==
slackb.com.     300     IN      CAA     0 iodef "mailto:hostmaster@slack-corp.com"
slackb.com.     300     IN      CAA     0 issue "letsencrypt.org; accounturi=https://acme-v02.api.letsencrypt.org/acme/acct/1532134906"
slackb.com.     300     IN      CAA     0 issue "digicert.com; account=455b15ed272bed097725c6ea50e89921fc57379c5338a29305c05771df65fff1"
slackb.com.     300     IN      RRSIG   CAA 13 2 300 20240929004401 20240928223901 15779 slackb.com. ET3FMsQZAReWWz0gHPPT9dMAQbPeDC1d3QNm2ZldqzCgaOJm9JxZJZQoFi4eSG4wLZrirUVH0GsHHmIiMi94OQ==

BYTES RECEIVED: 1873, IN: 27.5409ms
```

Example with `-raw` (hex dump) output included:
```
dig.exe slackb.com -t https -raw
Host: 8.8.8.8, Port: 53, Proto: udp, FQDN: slackb.com, Question Type: https
Id: 258, Opcode: 0, AA: false, TC: false, RD: true, RA: true, Z: false, RCODE: NOERROR
QUERY: 1; ANSWER: 0; AUTHORITY: 1; ADDITIONAL: 0

slackb.com.     340     IN      SOA     ns-1654.awsdns-14.co.uk. awsdns-hostmaster.amazon.com. 1 7200 900 1209600 86400

BYTES RECEIVED: 125, IN: 11.6849ms

00000000  01 02 81 80 00 01 00 00  00 01 00 00 06 73 6c 61  |.............sla|
00000010  63 6b 62 03 63 6f 6d 00  00 00 00 01 06 73 6c 61  |ckb.com......sla|
00000020  63 6b 62 03 63 6f 6d 00  00 06 00 01 00 00 01 54  |ckb.com........T|
00000030  00 4b 07 6e 73 2d 31 36  35 34 09 61 77 73 64 6e  |.K.ns-1654.awsdn|
00000040  73 2d 31 34 02 63 6f 02  75 6b 00 11 61 77 73 64  |s-14.co.uk..awsd|
00000050  6e 73 2d 68 6f 73 74 6d  61 73 74 65 72 06 61 6d  |ns-hostmaster.am|
00000060  61 7a 6f 6e 03 63 6f 6d  00 00 00 00 01 00 00 1c  |azon.com........|
00000070  20 00 00 03 84 00 12 75  00 00 01 51 80           | ......u...Q.|
```

Example of saving request & response in binary format to file(s) and using DoH POST wireformat for the DNS resolution:
```
./dig -doh-post @1.1.1.1 www.cnn.com -t A -bin-request-to-file cnn_req.bin -bin-response-to-file cnn_resp.bin
Wrote request binary to file: cnn_req.bin
Host: 1.1.1.1, Port: 443, Proto: doh-post, FQDN: www.cnn.com, Question Type: A
Id: 50278, Opcode: 0, AA: false, TC: false, RD: true, RA: true, Z: false, RCODE: NOERROR
QUERY: 1; ANSWER: 5; AUTHORITY: 0; ADDITIONAL: 0

www.cnn.com.	288	IN	CNAME	cnn-tls.map.fastly.net.
cnn-tls.map.fastly.net.	48	IN	A	151.101.3.5
cnn-tls.map.fastly.net.	48	IN	A	151.101.67.5
cnn-tls.map.fastly.net.	48	IN	A	151.101.195.5
cnn-tls.map.fastly.net.	48	IN	A	151.101.131.5

BYTES RECEIVED: 228, IN: 55.294083ms
Wrote response binary to file: cnn_resp.bin

$ ls -al *bin
-rw-r--r--   1 akhanin  staff       29 Mar  6 15:59 cnn_req.bin
-rw-r--r--   1 akhanin  staff      228 Mar  6 15:59 cnn_resp.bin

$ ./pp cnn_req.bin 
Rcode:  NOERROR
HEADER:
{Id:50278 Response:false Opcode:0 Authoritative:false Truncated:false RecursionDesired:true RecursionAvailable:false Zero:false AuthenticatedData:false CheckingDisabled:false Rcode:0}

QUESTION: 1
Name [www.cnn.com.] Class [1] Type [A]

ANSWER: 0

AUTHORITATIVE: 0

EXTRA: 0

$ ./pp cnn_resp.bin 
Rcode:  NOERROR
HEADER:
{Id:50278 Response:true Opcode:0 Authoritative:false Truncated:false RecursionDesired:true RecursionAvailable:true Zero:false AuthenticatedData:false CheckingDisabled:false Rcode:0}

QUESTION: 1
Name [www.cnn.com.] Class [1] Type [A]

ANSWER: 5
www.cnn.com.	288	IN	CNAME	cnn-tls.map.fastly.net.
cnn-tls.map.fastly.net.	48	IN	A	151.101.3.5
cnn-tls.map.fastly.net.	48	IN	A	151.101.67.5
cnn-tls.map.fastly.net.	48	IN	A	151.101.195.5
cnn-tls.map.fastly.net.	48	IN	A	151.101.131.5

AUTHORITATIVE: 0

EXTRA: 0

```

Using `dnssec` and saving the response to `json`:

```
$ ./dig www.verisign.com -t A +dnssec -json-response-to-file v.json
Host: 8.8.8.8, Port: 53, Proto: udp, DNSSEC: true, FQDN: www.verisign.com., Question Type: A
Id: 12816, Opcode: 0, AA: false, TC: false, RD: true, RA: true, Z: false, RCODE: NOERROR
QUERY: 1; ANSWER: 4; AUTHORITY: 0; ADDITIONAL: 1

www.verisign.com.       552     IN      CNAME   www.gslb.verisign.com.
www.verisign.com.       552     IN      RRSIG   CNAME 8 3 600 20250406231950 20250307231950 48940 verisign.com. Jsd5mhvQv1jPd0ppKGbB4bBBomqiuDU7TRteW4c+tnV02oI8bD4SIps+OYqsU0OavqDz0TzK3QLRlSpCp71ZBgigEpTBfdga/FUthOHrxWq9vTlwzoKHzrcpti7U4z82d2JJ6BJegK64WZDg68qVjz7hFk9ddtX9mHTkecwKqaU=
www.gslb.verisign.com.  12      IN      A       209.131.162.75
www.gslb.verisign.com.  12      IN      RRSIG   A 8 4 30 20250311180653 20250304170653 25746 gslb.verisign.com. hmxndf1ykuOhG6bg3hmaxLz6NxicC8QAyULe3td2BMaOBJO8scOrfE08SGCFSYNoWoCaANdVeGkjlrFTEx7NdNR5a5oJlmS6gr5m6IGf6HBLZL6hI7p+BL+jlYMPECcl4K8oOUlWEb8FoogWfN4oDfHT/RN2hhHerJn4qU5NSZiJY9VGc46S+YElx+91/25/NWBxXL/OcOQyYSgfF9BoGIGofa7FMyqSXgqDRGWjHNUf9PcG+JTJCgZ02aBh6eB2FTH8t65DoMpGu05ybUYoyMgmN9zFrXx6OWA0E85Kh7+Sb07IxoDyWATaOAva3b348zpC6jmPQV4HCPvCqHCLFg==

BYTES RECEIVED: 650, IN: 9.7809ms
Wrote response json to file: v.json


$ cat v.json
{
    "Id": 12816,
    "Response": true,
    "Opcode": 0,
    "Authoritative": false,
    "Truncated": false,
    "RecursionDesired": true,
    "RecursionAvailable": true,
    "Zero": false,
    "AuthenticatedData": true,
    "CheckingDisabled": false,
    "Rcode": 0,
    "Question": [
        {
            "Name": "www.verisign.com.",
            "Qtype": 1,
            "Qclass": 1
        }
    ],
    "Answer": [
        {
            "Hdr": {
                "Name": "www.verisign.com.",
                "Rrtype": 5,
                "Class": 1,
                "Ttl": 552,
                "Rdlength": 11
            },
            "Target": "www.gslb.verisign.com."
        },
        {
            "Hdr": {
                "Name": "www.verisign.com.",
                "Rrtype": 46,
                "Class": 1,
                "Ttl": 552,
                "Rdlength": 160
            },
            "TypeCovered": 5,
            "Algorithm": 8,
            "Labels": 3,
            "OrigTtl": 600,
            "Expiration": 1743981590,
            "Inception": 1741389590,
            "KeyTag": 48940,
            "SignerName": "verisign.com.",
            "Signature": "Jsd5mhvQv1jPd0ppKGbB4bBBomqiuDU7TRteW4c+tnV02oI8bD4SIps+OYqsU0OavqDz0TzK3QLRlSpCp71ZBgigEpTBfdga/FUthOHrxWq9vTlwzoKHzrcpti7U4z82d2JJ6BJegK64WZDg68qVjz7hFk9ddtX9mHTkecwKqaU="
        },
        {
            "Hdr": {
                "Name": "www.gslb.verisign.com.",
                "Rrtype": 1,
                "Class": 1,
                "Ttl": 12,
                "Rdlength": 4
            },
            "A": "209.131.162.75"
        },
        {
            "Hdr": {
                "Name": "www.gslb.verisign.com.",
                "Rrtype": 46,
                "Class": 1,
                "Ttl": 12,
                "Rdlength": 293
            },
            "TypeCovered": 1,
            "Algorithm": 8,
            "Labels": 4,
            "OrigTtl": 30,
            "Expiration": 1741716413,
            "Inception": 1741108013,
            "KeyTag": 25746,
            "SignerName": "gslb.verisign.com.",
            "Signature": "hmxndf1ykuOhG6bg3hmaxLz6NxicC8QAyULe3td2BMaOBJO8scOrfE08SGCFSYNoWoCaANdVeGkjlrFTEx7NdNR5a5oJlmS6gr5m6IGf6HBLZL6hI7p+BL+jlYMPECcl4K8oOUlWEb8FoogWfN4oDfHT/RN2hhHerJn4qU5NSZiJY9VGc46S+YElx+91/25/NWBxXL/OcOQyYSgfF9BoGIGofa7FMyqSXgqDRGWjHNUf9PcG+JTJCgZ02aBh6eB2FTH8t65DoMpGu05ybUYoyMgmN9zFrXx6OWA0E85Kh7+Sb07IxoDyWATaOAva3b348zpC6jmPQV4HCPvCqHCLFg=="
        }
    ],
    "Ns": null,
    "Extra": [
        {
            "Hdr": {
                "Name": ".",
                "Rrtype": 41,
                "Class": 512,
                "Ttl": 32768,
                "Rdlength": 0
            },
            "Option": null
        }
    ]
}
```

### Help

`dig.exe --help`:

```
$ ./dig -help
dig v0.0.2 - A lightweight dig replacement

ex: dig @8.8.4.4 google.com -t MX +dnssec

Flags:

  -bin-request-to-file string
        print request binary to file
  -bin-response-to-file string
        print response binary to file
  -doh
        use DoH (GET) json format
  -doh-post
        use DoH via HTTP POST wire format
  -help
        Get help on the 'dig' command.
  -host string
        DNS server hostname/ip to use (default "8.8.8.8")
  -json-response-to-file string
        print response json to file
  -nc
        disable ansi colors
  -port string
        port to connect on (default "53")
  -print-request-base64
        print request base64
  -raw
        show raw response
  -t string
        question type, ex: A, NS, MX, etc. (default "A")
  -tcp
        use TCP
  -tls
        use TLS (DoT)

```

`pp.exe` (Packet Print)

Example of a file called `bad3.txt` containing a mixed text & hex output from BIND's dig tool complaining it cannot parse a DNS response:

`type bad3.txt`:
```
;; Got bad packet: extra input data
58 bytes
0a 64 81 80 00 01 00 01 00 00 00 00 06 73 6c 61          .d...........sla
63 6b 62 03 63 6f 6d 00 00 41 00 01 c0 0c 00 41          ckb.com..A.....A
00 01 00 00 00 1e 00 12 00 00 00 00 01 00 03 02          ................
68 32 00 04 00 04 2d fd 83 e2                            h2....-...
```

We can pass it to `pp.exe` directly and it will attempt to extract the relevant parts to parse it as a DNS response:
```
pp bad3.txt
Rcode:  NOERROR
HEADER:
{Id:2660 Response:true Opcode:0 Authoritative:false Truncated:false RecursionDesired:true RecursionAvailable:true Zero:false AuthenticatedData:false CheckingDisabled:false Rcode:0}

QUESTION: 1
Name [slackb.com.] Class [1] Type [HTTPS]

ANSWER: 1
slackb.com.     30      IN      HTTPS   0 . alpn="h2" ipv4hint="45.253.131.226"

AUTHORITATIVE: 0

EXTRA: 0
```

### Testing with Wire format (for DoH)

Generating a base64 string to use for the Wire Format (POST via DoH). Ex: `google.com` for `A` record:
```
./dig -print-request-base64 google.com -t A
Request Base64 encoded: is8BAAABAAAAAAAABmdvb2dsZQNjb20AAAEAAQ==
```

Making the DoH Wire Format DNS query and capturing the response, replace with your own resolver hostname instead of `doh.resolver.to.use.com`:
```
echo -n 'is8BAAABAAAAAAAABmdvb2dsZQNjb20AAAEAAQ==' | base64 --decode | curl --header 'content-type: application/dns-message' --data-binary @- https://doh.resolver.to.use.com/dns-query --output - > google_a.bin
```

Inspecting the response:
```
./pp google_a.bin 
Rcode:  NOERROR
HEADER:
{Id:35535 Response:true Opcode:0 Authoritative:false Truncated:false RecursionDesired:true RecursionAvailable:true Zero:false AuthenticatedData:false CheckingDisabled:false Rcode:0}

QUESTION: 1
Name [google.com.] Class [1] Type [A]

ANSWER: 6
google.com.	267	IN	A	173.194.208.113
google.com.	267	IN	A	173.194.208.138
google.com.	267	IN	A	173.194.208.101
google.com.	267	IN	A	173.194.208.139
google.com.	267	IN	A	173.194.208.102
google.com.	267	IN	A	173.194.208.100

AUTHORITATIVE: 0

EXTRA: 0
```
