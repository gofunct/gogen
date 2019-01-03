# GoGen

**A utility for generating boilerplate code for cli tools that integrate with cloud providers**

**by: Coleman Word**

## Gogen Commands

### `make help`

```commandline
make [command]

all                            generate all bins
clean                          clean bin
cover                          test coverage
fmt                            go install all programs
gen                            go generate
help                           help
install                        go install all programs
lint                           lint
packages                       packages
setup                          setup
test                           test all


```

### `gogen`

```commandline
Usage:
  gogen [command]

Available Commands:
  build       Build commands
  destroy     Destroy an existing new code
  generate    Generate a new code
  gocloud     cloud opts
  gogen       
  help        Help about any command
  init        Initialize a gogen application
  protoc      Run protoc
  version     Print the version information

Flags:
      --debug     Debug level output
  -h, --help      help for gogen
  -v, --verbose   Verbose level output

Use "gogen [command] --help" for more information about a command.

```

### `gogen version`

```commandline
gogen v0.1.1 (go1.11.4 darwin/amd64)
```

## Generated Commands

### `{appName} gocloud`

```commandline
Usage:
  gogen gocloud [command]

Available Commands:
  init        initialize

Flags:
      --bucket string               what bucket name do you want to setup?
      --cloud_sql_region string     region of the Cloud SQL instance (GCP only)
      --db_host string              what is the database host or Cloud SQL instance name?
      --db_name string              what is the database name?
      --db_password string          what is the database user password?
      --db_user string              what is the database username? (default "guestbook")
      --env string                  what is environment do you want to run under?(gcp or aws) (default "local")
  -h, --help                        help for gocloud
      --listen string               what port do you want to listen on? (default ":8080")
      --runtime_config string       runtime Configurator config resource (GCP only)
      --runtime_var string          what is the runtime variable location?
      --runtime_var_wait duration   polling frequency of message of the day (default 5s)

Global Flags:
      --debug     Debug level output
  -v, --verbose   Verbose level output

Use "gogen gocloud [command] --help" for more information about a command.

```


### `{appName} version`

```commandline
{appName} v0.1.1 (go1.11.4 darwin/amd64)
```
