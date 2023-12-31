## medic role get

Get details for an role within a mediator control plane

### Synopsis

The medic role get subcommand lets you retrieve details for an role within a
mediator control plane.

```
medic role get [flags]
```

### Options

```
  -h, --help           help for get
  -i, --id int32       ID for the role to query
  -n, --name string    Name for the role to query
  -o, --org-id int32   Organization ID
```

### Options inherited from parent commands

```
      --config string      config file (default is $PWD/config.yaml)
      --grpc-host string   Server host (default "staging.stacklok.dev")
      --grpc-insecure      Allow establishing insecure connections
      --grpc-port int      Server port (default 443)
```

### SEE ALSO

* [medic role](medic_role.md)	 - Manage roles within a mediator control plane

