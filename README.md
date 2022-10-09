# Workflow

[x] Starts a DNS TCP server on port 53

[x] Forwards DNS Requests to CloudFlare DNS-over-TlS (DoT)

[x] Reads responses and sends them back to the client

[x] Uses [RFC 1035](https://www.rfc-editor.org/rfc/rfc1035#section-4.1.1) error codes for error handling

# Bonus points

[x] Allow multiple incoming requests at the same time

[ ] Also handle UDP requests, while still querying tcp on the other side

[ ] Any other improvements you can think of!

# QA

**Imagine this proxy being deployed in an infrastructure. What would be the security concerns you would raise?**

* if the service is exposed to outside world, it could be subject to DDOS or DNS Spoofing, since the communication from the client to the proxy is unencrypted.

**How would you integrate that solution in a distributed, microservices-oriented and containerized architecture?**

* for multi-cluster dns (service discovery), we could deploy the service as a controller similar to `ExternalDNS` plugin or `CoreDNS`.
* for application pods, we could add the proxy to start-up configuration of apps to override the existing DNS lookup servers.

**What other improvements do you think would be interesting to add to the project?**

* Caching DNS Queries/Responses.
* Controlling access by applying policies, for example, using Open Policy Agent (OPA).

# Getting Started

## Go

```
 golangci-lint run
 go build .
 ./simple-dns-proxy 
```

## Docker

```
 docker build -t simple-dns-proxy simple-dns-proxy

 docker run -p 0.0.0.0:53:53 --name simple-dns-proxy simple-dns-proxy

 docker start -a simple-dns-proxy
```

## Docker Compose

```
 docker-compose up
```

## Dig

```
 dig +tcp @0.0.0.0 -p 53 example.com
 kdig +tcp -d @0.0.0.0 -p 53 example.com
```
