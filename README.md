# bindman-dns-file-agent

# How it works

The bindmand-dns-file-agent is an agent responsible for persisting in bindman-dns-manager by a file.

### Dependencies

bindman-dns-file-agent requires [Golang](https://golang.org/dl/) v1.13, a [BIND9 DNS Server](https://www.isc.org/bind/) or [Dockerize BIND9 DNS Derver](https://github.com/labbsr0x/docker-dns-bind9)  and a [bindman-dns-bind9](https://github.com/labbsr0x/bindman-dns-bind9) service to run.


### Configurations

Some parameters are required to running the api, these parameters can be passed via the command line or environment variables as described below


| ENV                              | Command | Required | Default  | Description                                        |
|----------------------------------|---------|----------|----------|----------------------------------------------------|
| BINDMAN_DNS_MANAGER_ADDR         | -d      | true     | null     | DNS Manager Address                                |
| BINDMAN_DNS_REVERSE_PROXY_ADDR   | -r      | true     | null     | DNS Reverse Proxy Address                          |
| BINDMAN_DNS_CONFIG_FILE          | -f      | true     | null     | DNS Config File                                    |


## How to run

With all dependencies installed, start the server.

Run in localhost

```sh
$ cd bindman-dns-file-agent
$ go run main.go agent \
-d=bindman-dns-bind9:7070 \
-r=0.0.0.0 \
-f=agent.conf
```

For production environments

```sh
$ go run main.go agent
$ ENV BINDMAN_DNS_MANAGER_ADDR="bindman-dns-bind9:7070"
$ ENV BINDMAN_DNS_REVERSE_PROXY_ADDR="0.0.0.0"
$ ENV BINDMAN_DNS_CONFIG_FILE="agent.conf"
```