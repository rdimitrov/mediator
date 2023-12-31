## medic user create

Create a user within a mediator control plane

### Synopsis

The medic user create subcommand lets you create new users for a role
within a mediator control plane.

```
medic user create [flags]
```

### Options

```
  -e, --email string            E-mail for the user
  -f, --firstname string        User's first name
  -s, --force                   Skip confirmation
  -g, --groups string           Comma separated list of groups
  -h, --help                    help for create
  -i, --is-protected            Is the user protected
  -l, --lastname string         User's last name
  -c, --needs-password-change   Does the user need to change their password (default true)
  -o, --org-id int32            Organization ID for the user
  -p, --password string         Password for the user
  -r, --roles string            Comma separated list of roles
  -u, --username string         Username
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

