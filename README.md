# mtls-transmitter

mtls-transmitter is a simple reverse proxy that injects client certificate for mTLS communication. It creates a reverse proxy that injects client certificate to incomming http requests and routes them to the https endpoint.

## Building the binary

```bash
$ go build -o mtls-transmitter ./cmd/transmitter
```

## Running the binary

```bash
$ ./mtls-transmitter -cert=/path/to/cert.pem -key=/path/to/key.pem -url=desired.host.com [-port=8080] [-skip-ssl]
```

### Parameters

- cert - path to the client certifiacte, required
- key - path to the client certificate key, required
- url - hostname to which the proxy forwards the calls, required
- port - the port on which the proxy listens on, optional, default 8080
- skip-ssl - if defined the proxy will skip server certificate verification, optional, default false

## Running inside the docker container

### Building

```bash
$ docker build -t crabtree/mtls-transmitter .
```

### Running 

>**NOTE:** To run mtls-transmitter inside the docker container you need to provide your client certificate to your container.

```bash
$ docker run --rm -v /path/to/cert-dir:/cert -p 8080:8080 crabtree/mtls-transmitter -cert /cert/cert.pem -key /cert/key.pem -url desired.host.com
```