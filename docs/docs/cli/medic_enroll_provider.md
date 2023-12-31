## medic enroll provider

Enroll a provider within the mediator control plane

### Synopsis

The medic enroll provider command allows a user to enroll a provider
such as GitHub into the mediator control plane. Once enrolled, users can perform
actions such as adding repositories.

```
medic enroll provider [flags]
```

### Options

```
  -g, --group-id int32    ID of the group for enrolling the provider
  -h, --help              help for provider
  -n, --provider string   Name for the provider to enroll
  -t, --token string      Personal Access Token (PAT) to use for enrollment
```

### Options inherited from parent commands

```
      --config string      config file (default is $PWD/config.yaml)
      --grpc-host string   Server host (default "localhost")
      --grpc-port int      Server port (default 8090)
```

### SEE ALSO

* [medic enroll](medic_enroll.md)	 - Manage enrollments within a mediator control plane

