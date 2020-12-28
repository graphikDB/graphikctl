# graphikctl

command line interface for graphikDB

    git clone git@github.com:graphikDB/graphik.git

[Generated Documentation](./docs/graphikctl.md)

```text
A command line utility for graphikDB

---
env-prefix: GRAPHIKCTL
config-path: ~/.graphikctl.yaml

Usage:
  graphikctl [command]

Available Commands:
  auth        authentication/authorization subcommands (login)
  broadcast   graphikDB broadcast operations
  config      configuration subcommands (get, open)
  create      graphikDB create operations (doc, connection)
  edit        graphikDB edit operations (doc, connection)
  get         graphikDB get operations (doc, connection, schema)
  help        Help about any command
  search      graphikDB search operations  (docs, connections)
  stream      graphikDB stream operations
  traverse    graphikDB traversal operations

Flags:
      --config string   config file (default is $HOME/.graphikctl.yaml)
  -h, --help            help for graphikctl
  -v, --version         version for graphikctl

Use "graphikctl [command] --help" for more information about a command.

```

## Example Config (~/.graphikctl.yaml)

```yaml
auth:
  open_id: https://accounts.google.com/.well-known/openid-configuration
  client_id: ${uuid}.apps.googleusercontent.com
  client_secret: ${client_secret}
host: localhost:7820
server:
  port: :8080
```