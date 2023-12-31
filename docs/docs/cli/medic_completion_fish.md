## medic completion fish

Generate the autocompletion script for fish

### Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	medic completion fish | source

To load completions for every new session, execute once:

	medic completion fish > ~/.config/fish/completions/medic.fish

You will need to start a new shell for this setup to take effect.


```
medic completion fish [flags]
```

### Options

```
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
      --config string      config file (default is $PWD/config.yaml)
      --grpc-host string   Server host (default "staging.stacklok.dev")
      --grpc-insecure      Allow establishing insecure connections
      --grpc-port int      Server port (default 443)
```

### SEE ALSO

* [medic completion](medic_completion.md)	 - Generate the autocompletion script for the specified shell

