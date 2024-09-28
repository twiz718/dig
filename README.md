# A lightweight dig clone

## Example usage

### Building

```
go build
```

### Running

```
dig -fqdn slackb.com -t ANY -h 192.168.254.10
Host: 192.168.254.10, Port: 53, Proto: udp, FQDN: slackb.com, Question Type: ANY
WARNING: truncated response, will retry with tcp instead
ANSWERS: 28, BYTES RECEIVED: 1873, IN: 86.9883ms

slackb.com.     60      IN      A       3.91.140.69
slackb.com.     60      IN      A       44.216.98.239
slackb.com.     60      IN      A       34.202.253.140
slackb.com.     60      IN      A       3.90.158.208
slackb.com.     60      IN      A       52.200.46.145
slackb.com.     60      IN      A       34.195.221.192
slackb.com.     60      IN      A       54.236.104.103
slackb.com.     60      IN      A       52.3.167.79
slackb.com.     60      IN      A       44.205.171.153
slackb.com.     60      IN      A       18.213.32.120
slackb.com.     60      IN      RRSIG   A 13 2 60 20240928193622 20240928173522 15779 slackb.com. bEaLkV0yxacJVJXO2NU/77oma5R5a01s+bqJokC7EFLI+a60yvmo5X+zPArh6m4HckzFTmcLZvLGcFfr4f5P/A==
slackb.com.     21600   IN      NS      ns-352.awsdns-44.com.
slackb.com.     21600   IN      NS      ns-1654.awsdns-14.co.uk.
slackb.com.     21600   IN      NS      ns-1347.awsdns-40.org.
slackb.com.     21600   IN      NS      ns-542.awsdns-03.net.
slackb.com.     21600   IN      RRSIG   NS 13 2 172800 20240930193522 20240928173522 15779 slackb.com. iPmWB6R3nBAq3zl+S4qj3sGaXbrWL5Rws22WlJ2B9aN0iG7K1wUXeZV0Apo7BywHBdaLEjg66h3C/8a2evJTeg==
slackb.com.     900     IN      SOA     ns-1654.awsdns-14.co.uk. awsdns-hostmaster.amazon.com. 1 7200 900 1209600 86400
slackb.com.     900     IN      RRSIG   SOA 13 2 900 20240928195022 20240928173522 15779 slackb.com. erSXySXoE5yUg/OeNh2s0CAFPWDt6HJR2v9ySUSTIWz6ndUtKv/iut43ott+XHY6W+JFbqw35skPgucyilhSMw==
slackb.com.     300     IN      TXT     "v=spf1 -all"
slackb.com.     300     IN      RRSIG   TXT 13 2 300 20240928194022 20240928173522 15779 slackb.com. DM+lIlDYMQQsjtKwfiLp/J7oDO++75aj53bP4jrSKTWWBR/HVWQN62V26BCiKZ/xniwsouI7Cu5ZKefgAkcCbA==
slackb.com.     3600    IN      DNSKEY  257 3 13 bSr2fT2hBdGDUltm3fPrF4HB6x+2k2Ol51pImwom4jGa1J66Bo6g4x/2a2iAdjzejf8d4cCeXkuQvHkeDpLV+w==
slackb.com.     3600    IN      DNSKEY  256 3 13 g4jECOQIRAufN1jKoEHGoR9ZSld0jNy/AqbiMJPh1YlNz4k7eAPjS/YuyhYOw7784LfzcGGXDIS+04Q6J4D0mQ==
slackb.com.     3600    IN      DNSKEY  256 3 13 cV5+Jxz0KlpwnBS009OppqW/wEAGECGElUSAsny6Dk5NNSK6M1gGOqXql6Yke2JqMlP30xycgTIVcgvJOQtMeg==
slackb.com.     3600    IN      RRSIG   DNSKEY 13 2 3600 20240929020000 20240928150000 55613 slackb.com. Nl0p33o7u3yQrLxekC23YpnpBTOrbrcYAiaM2mGahxcvOFaLCPVr1vS6fV5riW4wWxToyfTHtYqgTkljGahhlQ==
slackb.com.     300     IN      CAA     0 issue "digicert.com; account=455b15ed272bed097725c6ea50e89921fc57379c5338a29305c05771df65fff1"
slackb.com.     300     IN      CAA     0 iodef "mailto:hostmaster@slack-corp.com"
slackb.com.     300     IN      CAA     0 issue "letsencrypt.org; accounturi=https://acme-v02.api.letsencrypt.org/acme/acct/1532134906"
slackb.com.     300     IN      RRSIG   CAA 13 2 300 20240928194022 20240928173522 15779 slackb.com. cHAhe0WTjjtC1wt5LsTqAFWJlWWbaTZIWWS0XeLVSXZKYyKIWMdzuybSmy3pRdCtKN14TQKH7AAoh4s70AFs3Q==
```

### Help

```
Usage of dig:
  -fqdn string
        fqdn to lookup (default "google.com")
  -h string
        DNS server hostname/ip to use (default "8.8.8.8")
  -nc
        no color
  -p string
        Port to connect on (default "53")
  -t string
        question dns.Type, ex: A, AAAA, NS, etc. (default "A")
  -tcp
        use TCP
  -tls
        use TLS (DoT)
```