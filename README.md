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
Parameters can also be set as flags or as environment variables, with flags
attempting to take precedence.

| Flag | Env Var | Description |
| --- | --- | --- |
| `-cert` | `CERT` | The path to the client certificate; **required** |
| `-key` | `KEY` | The path to the client certificate key; **required** |
| `-url` | `URL` | The `hostname:port` to which the proxy fowards requests; **required** |
| `-port` | `PORT` | The port on which the proxy listens on; default: `8080` |
| `-skip-ssl` | `SKIP_SSL` | If set to `true`, the proxy sill skip server certificate verification; default: `false` |
| `-silent` | `SILENT` | If set to `true`, the proxy will not log proxied events; default: `false` |

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

## Development

Use the following:
```bash
make          # to format and validate changes
make build    # to build the binary
```

See the [Makefile](Makefile) for additional options.

