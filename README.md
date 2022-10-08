# Test server

nslookup google.com localhost -port=53

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
