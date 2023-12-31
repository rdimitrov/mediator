## medic org create

Create an organization within a mediator control plane

### Synopsis

The medic org create subcommand lets you create new organizations
within a mediator control plane.

```
medic org create [flags]
```

### Options

```
  -c, --company string           Company name of the organization
  -d, --create-default-records   Create default records for the organization
  -h, --help                     help for create
  -n, --name string              Name of the organization
```

### Options inherited from parent commands

```
      --config string      config file (default is $PWD/config.yaml)
      --grpc-host string   Server host (default "staging.stacklok.dev")
      --grpc-insecure      Allow establishing insecure connections
      --grpc-port int      Server port (default 443)
```

### SEE ALSO

* [medic org](medic_org.md)	 - Manage organizations within a mediator control plane

