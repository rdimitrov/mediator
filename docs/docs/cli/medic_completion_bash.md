## medic completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(medic completion bash)

To load completions for every new session, execute once:

#### Linux:

	medic completion bash > /etc/bash_completion.d/medic

#### macOS:

	medic completion bash > $(brew --prefix)/etc/bash_completion.d/medic

You will need to start a new shell for this setup to take effect.


```
medic completion bash
```

### Options

```
  -h, --help              help for bash
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

