## medic rule_type test

test a rule type definition

### Synopsis

The 'rule_type test' subcommand allows you test a rule type definition

```
medic rule_type test [flags]
```

### Options

```
  -e, --entity string      YAML file containing the entity to test the rule against
  -h, --help               help for test
  -p, --policy string      YAML file containing a policy to test the rule against
  -r, --rule-type string   file to read rule type definition from
  -t, --token string       token to authenticate to the provider.Can also be set via the AUTH_TOKEN environment variable.
```

### Options inherited from parent commands

```
      --config string      config file (default is $PWD/config.yaml)
      --grpc-host string   Server host (default "staging.stacklok.dev")
      --grpc-insecure      Allow establishing insecure connections
      --grpc-port int      Server port (default 443)
```

### SEE ALSO

* [medic rule_type](medic_rule_type.md)	 - Manage rule types within a mediator control plane

