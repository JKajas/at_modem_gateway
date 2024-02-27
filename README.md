# at_modem_gateway

## About

Package was created for fast deployement simple sms gateway.
To achive that purpose, package uses docker(with compose plugin), gRPC protocol and protobuffer serialization.

## Requirements

- Modem supports AT commands

## Quick start

Application is easy to connect with your projects.\
All you have to do is run container and implement expected request with protobuffer. By default application uses USB serial port.

### Steps

1. Change needed AREA_CODE in env variables in docker-compose.yaml
2. Run container with exec file which runs gRPC server and listens on port 8000. Command: `docker compose up`
3. Serialize request and response for your project with protobuffer file located in /pb directory or connect to container directly with gRPC client.
