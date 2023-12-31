## medic auth revoke_provider

Revoke access tokens for provider

### Synopsis

It can revoke access tokens for specific provider.

```
medic auth revoke_provider [flags]
```

### Options

```
  -a, --all               Revoke all tokens
  -h, --help              help for revoke_provider
  -n, --provider string   Name for the provider to revoke tokens for
```

### Options inherited from parent commands

```
      --config string      config file (default is $PWD/config.yaml)
      --grpc-host string   Server host (default "staging.stacklok.dev")
      --grpc-insecure      Allow establishing insecure connections
      --grpc-port int      Server port (default 443)
```

### SEE ALSO

* [medic auth](medic_auth.md)	 - Authorize and manage accounts within a mediator control plane

