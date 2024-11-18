# Workflow

- [x] Starts a DNS TCP server on port 53

- [x] Forwards DNS Requests to CloudFlare DNS-over-TlS (DoT)

- [x] Reads responses and sends them back to the client

- [x] Uses [RFC 1035](https://www.rfc-editor.org/rfc/rfc1035#section-4.1.1) error codes for error handling

# TODO

- [x] Allow multiple incoming requests at the same time

- [ ] Also handle UDP requests, while still querying tcp on the other side

## Research

- [ ] **Security Concerns**: Assess potential issues with deploying the proxy in infrastructure:
  - [ ] Determine the risks of exposing the service to the outside world (e.g., DDOS or DNS Spoofing due to unencrypted client-to-proxy communication).

- [ ] **Integration in Distributed Architecture**:
  - [ ] Explore deploying the service as a controller for multi-cluster DNS (service discovery), similar to `ExternalDNS` or `CoreDNS`.
  - [ ] Configure application pods to include the proxy in their start-up settings, overriding existing DNS lookup servers.

- [ ] **Potential Improvements**:
  - [ ] Implement caching for DNS Queries/Responses.
  - [ ] Add access control mechanisms using policy enforcement tools like Open Policy Agent (OPA).

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
