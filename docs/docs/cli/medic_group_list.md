## medic group list

Get list of groups within a mediator control plane

### Synopsis

The medic group list subcommand lets you list groups within
a mediator control plane.

```
medic group list [flags]
```

### Options

```
  -h, --help            help for list
  -l, --limit int32     Limit the number of results returned (default -1)
  -f, --offset int32    Offset the results returned
  -i, --org-id int32    org id to list groups for
  -o, --output string   Output format
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

