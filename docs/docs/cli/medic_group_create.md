## medic group create

Create a group within a mediator control plane

### Synopsis

The medic group create subcommand lets you create new groups within
a mediator control plane.

```
medic group create [flags]
```

### Options

```
  -d, --description string   Description of the group
  -h, --help                 help for create
  -i, --is_protected         Is the group protected
  -n, --name string          Name of the group
      --org-id int32         Organization ID
```

### Options inherited from parent commands

```
      --config string      config file (default is $PWD/config.yaml)
      --grpc-host string   Server host (default "staging.stacklok.dev")
      --grpc-insecure      Allow establishing insecure connections
      --grpc-port int      Server port (default 443)
```

### SEE ALSO

* [medic group](medic_group.md)	 - Manage groups within a mediator control plane

