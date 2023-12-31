## medic policy_type get

Get details for a policy type within a mediator control plane

### Synopsis

The medic policy_type get subcommand lets you retrieve details for a policy type within a
mediator control plane.

```
medic policy_type get [flags]
```

### Options

```
  -d, --default_schema    Only get the default schema in a readable format
  -h, --help              help for get
  -p, --provider string   Provider for the policy type
  -s, --schema            Only get the json schema in a readable format
  -t, --type string       Type of the policy
```

### Options inherited from parent commands

```
      --config string      config file (default is $PWD/config.yaml)
      --grpc-host string   Server host (default "localhost")
      --grpc-port int      Server port (default 8090)
```

### SEE ALSO

* [medic policy_type](medic_policy_type.md)	 - Manage policy types within a mediator control plane

