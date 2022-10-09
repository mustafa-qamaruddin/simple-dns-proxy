# Test server

```
 dig +tcp @localhost -p 9953 google.com

 kdig -d @1.1.1.1 +tls-ca +tls-host=cloudflare-dns.com  example.com

```

# Workflow

[x] Starts a DNS TCP/UDP server on port 53

[x] Decodes DNS Requests and Forwards them to CloudFlare DNS-over-TlS (DoT)

[x] Reads responses and sends them back to the client

# Extra Features

[x] Allow multiple incoming requests at the same time

[ ] Also handle UDP requests, while still querying tcp on the other side

[ ] Any other improvements you can think of!

# QA

[ ] Imagine this proxy being deployed in an infrastructure. What would be the security concerns you would raise?

[ ] How would you integrate that solution in a distributed, microservices-oriented and containerized architecture?

[ ] What other improvements do you think would be interesting to add to the project?

# Docker

```
 docker build -t simple-dns-proxy simple-dns-proxy

 docker run -p 0.0.0.0:53:53 --name simple-dns-proxy simple-dns-proxy

 docker start -a simple-dns-proxy

 docker-compose up

```

```
 dig +tcp @0.0.0.0 -p 53 example.com
 kdig +tcp -d @0.0.0.0 -p 53 example.com
```
