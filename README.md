# labbsr0x/bindman-dns-file-agent

- [labbsr0x/bindman-dns-file-agent](#labbsr0xbindman-dns-file-agent)
- [How it works](#how-it-works)
  - [Contributing](#contributing)
  - [Issues](#issues)
- [Getting started](#getting-started)
  - [Dependencies](#dependencies)
  - [Configurations](#configurations)
  - [Installation](#installation)
  - [Quickstart](#quickstart)
  - [Persistence](#persistence)
- [Maintenance](#maintenance)
  - [Upgrading](#upgrading)
  - [Shell Access](#shell-access)
  - [How to run with GO](#how-to-run-with-go)

# How it works

The bindmand-dns-file-agent is an agent responsible for persisting in bindman-dns-manager by a file.

## Contributing

If you find this image useful here's how you can help:

- Send a pull request with your awesome features and bug fixes
- Help users resolve their [issues](../../issues?q=is%3Aopen+is%3Aissue).

## Issues

Before reporting your issue please try updating Docker to the latest version and check if it resolves the issue. Refer to the Docker [installation guide](https://docs.docker.com/installation) for instructions.

SELinux users should try disabling SELinux using the command `setenforce 0` to see if it resolves the issue.

If the above recommendations do not help then [report your issue](../../issues/new) along with the following information:

- Output of the `docker version` and `docker info` commands
- The `docker run` command or `docker-compose.yml` used to start the image. Mask out the sensitive bits.

# Getting started

## Dependencies

bindman-dns-file-agent requires [Golang](https://golang.org/dl/) v1.13, a [BIND9 DNS Server](https://www.isc.org/bind/) or [Dockerize BIND9 DNS Derver](https://github.com/labbsr0x/docker-dns-bind9)  and a [bindman-dns-bind9](https://github.com/labbsr0x/bindman-dns-bind9) service to run.

## Configurations

Some parameters are required to running the api, these parameters can be passed via the command line or environment variables as described below

| ENV                              | Command | Required | Default  | Description                                                                                       |
|----------------------------------|---------|----------|----------|---------------------------------------------------------------------------------------------------|
| BINDMAN_DNS_MANAGER_ADDR         | -d      | true     | null     | Bindman DNS Manager Address                                                                       |
| BINDMAN_DNS_REVERSE_PROXY_ADDR   | -r      | true     | null     | Bindman DNS Reverse Proxy Address                                                                 |
| BINDMAN_AGENT_CONFIG_PATH        | -c      | true     | null     | Bindman Agent Config Path                                                                         |
| BINDMAN_LOG_LEVEL                | -l      | false    | info     | Sets the Log Level to one of seven (trace, debug, info, warn, error, fatal, panic). Default: info |

## Installation

Automated builds of the image are available on [Dockerhub](https://hub.docker.com/r/labbsr0x/bindman-dns-file-agent) and is the recommended method of installation.

```bash
docker pull labbsr0x/bindman-dns-file-agent
```

Alternatively you can build the image yourself.

```bash
docker build -t labbsr0x/bindman-dns-file-agent github.com/labbsr0x/bindman-dns-file-agent
```

or 

```bash
make build
```

## Quickstart

With all dependencies running, start the agent.

Start BINDMAN-DNS-AGENT-FILE using:

```bash
docker run --rm --name bindman-dns-file-agent -d --volume ${PWD}/.agent:/data -e BINDMAN_DNS_MANAGER_ADDR=http://bindman-dns-bind9:7070 -e BINDMAN_DNS_REVERSE_PROXY_ADDR=0.0.0.0 -e BINDMAN_AGENT_CONFIG_PATH=/data/bindman_agent.json --network network-bind labbsr0x/bindman-dns-file-agent
```

or

```bash
make docker-run
```

*Alternatively, you can use the sample [docker-compose.yml](docker-compose.yml) file to start the container using [Docker Compose](https://docs.docker.com/compose/)*

## Persistence

For the bindman-dns-file-agent to preserve its state across container shutdown and startup you should mount a volume at `/data`.

> *The [Quickstart](#quickstart) command already mounts a volume for persistence.*

```bash
mkdir -p .agent
```

# Maintenance

## Upgrading

To upgrade to newer releases:

  1. Download the updated Docker image:

  ```bash
  docker pull labbsr0x/bindman-dns-file-agent
  ```

  2. Stop the currently running image:

  ```bash
  docker stop bindman-dns-file-agent
  ```

  or 

  ```bash
  make docker-stop
  ```

  3. Remove the stopped container

  ```bash
  docker rm labbsr0x/bindman-dns-file-agent
  ```

  and

  ```bash
  rm -rf .agent
  ```

  4. Start the updated image

  ```bash
  docker run -name bindman-dns-file-agent -d \
    [OPTIONS] \
    labbsr0x/bindman-dns-file-agent
  ```

## Shell Access

For debugging and maintenance purposes you may want access the containers shell. If you are using Docker version `1.3.0` or higher you can access a running containers shell by starting `bash` using `docker exec`:

```bash
docker exec -it bindman-dns-file-agent sh
```

or

```
make docker-exec
```

## How to run with GO

With all dependencies running, start the agent.

Run in localhost

```sh
$ go run main.go agent \
-d=bindman-dns-bind9:7070 \
-r=0.0.0.0 \
-c=bindman_agent.json
```

For production environments

```sh
$ go run main.go agent
$ ENV BINDMAN_DNS_MANAGER_ADDR="bindman-dns-bind9:7070"
$ ENV BINDMAN_DNS_REVERSE_PROXY_ADDR="0.0.0.0"
$ ENV BINDMAN_AGENT_CONFIG_PATH="bindman_agent.json"
```