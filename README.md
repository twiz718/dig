# A lightweight dig clone

## Example usage

### Building

```
go build
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

### Help

`dig.exe --help`:

```
dig v0.0.1 - A lightweight dig replacement

ex: dig @8.8.4.4 google.com -t MX

Flags:

  -help
        Get help on the 'dig' command.
  -host string
        DNS server hostname/ip to use (default "8.8.8.8")
  -nc
        disable ansi colors
  -port string
        port to connect on (default "53")
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