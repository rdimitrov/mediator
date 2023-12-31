## medic user update

Update a user within a mediator control plane

### Synopsis

The medic user update subcommand allows to modify the details of an user.

```
medic user update [flags]
```

### Options

```
  -e, --email string                   Email for your profile
  -f, --firstname string               First name for your profile
  -h, --help                           help for update
  -l, --lastname string                Last name for your profile
  -p, --password string                Password
  -c, --password_confirmation string   Password confirmation
```

### Options inherited from parent commands

```
      --config string      config file (default is $PWD/config.yaml)
      --grpc-host string   Server host (default "staging.stacklok.dev")
      --grpc-insecure      Allow establishing insecure connections
      --grpc-port int      Server port (default 443)
```

### SEE ALSO

* [medic user](medic_user.md)	 - Manage users within a mediator control plane

