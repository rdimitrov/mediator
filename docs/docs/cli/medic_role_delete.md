## medic role delete

delete a role within a mediator controlplane

### Synopsis

The medic role delete subcommand lets you delete roles within a
mediator control plane.

```
medic role delete [flags]
```

### Options

```
  -f, --force           Force deletion of role, even if it's protected or has associated users (WARNING: removing a protected role may cause loosing mediator access)
  -h, --help            help for delete
  -r, --role-id int32   id of role to delete
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

