# graphikctl

command line interface for graphikDB

`git clone git@github.com:graphikDB/graphik.git`

[Documentation](./docs/graphikctl.md)

```text
A command line utility for graphikDB

---
env-prefix: GRAPHIKCTL

Usage:
  graphikctl [command]

Available Commands:
  auth        authentication/authorization subcommands (login)
  config      configuration subcommands (get)
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.graphikctl.yaml)
  -h, --help            help for graphikctl
  -v, --version         version for graphikctl

Use "graphikctl [command] --help" for more information about a command.

```

## Example Config

```yaml
auth:
  open_id: https://accounts.google.com/.well-known/openid-configuration
  client_id: ${uuid}.apps.googleusercontent.com
  client_secret: ${client_secret}
```