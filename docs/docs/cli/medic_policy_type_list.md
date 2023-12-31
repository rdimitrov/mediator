## medic policy_type list

List policy types for a provider within a mediator control plane

### Synopsis

The medic policy_type list subcommand lets you list policy types within a
mediator control plane for an specific provider.

```
medic policy_type list [flags]
```

### Options

```
  -h, --help              help for list
  -o, --output string     Output format (json or yaml)
  -p, --provider string   Provider to list policies for
```

### Options inherited from parent commands

```
      --config string      config file (default is $PWD/config.yaml)
      --grpc-host string   Server host (default "localhost")
      --grpc-port int      Server port (default 8090)
```

### SEE ALSO

* [medic policy_type](medic_policy_type.md)	 - Manage policy types within a mediator control plane

