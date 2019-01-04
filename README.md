![](https://github.com/gofunct/common/blob/master/logo/dark_logo_transparent_background.png?raw=true)
# Gogen
**A utility for generating boilerplate code for cli tools that integrate with cloud providers**

* Author: Coleman Word 
* Email: coleman.word@gofunct.com
* Language: Golang
* Download: `go get github.com/gofunct/gogen/...`

## Table of Contents

- [Gogen](#gogen)
  * [Table of Contents](#table-of-contents)
  * [Gogen Commands](#gogen-commands)
    + [`gogen`](#-gogen-)
  * [Project Generation](#project-generation)
    + [Features](#features)
    + [Files](#files)
  * [Generated Commands](#generated-commands)
    + [`{appName} gocloud`](#--appname--gocloud-)
    + [`{appName} version`](#--appname--version-)
  * [Generated Files](#generated-files)
    + [Makefile](#makefile)
  * [File Tree](#file-tree)


## Gogen Commands

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

## Project Roadmap

### Features
- [x] Generate fully functional grpc cloud service with one command(include grpc-json gateway)
- [x] OpenCensus Tracing
- [x] StackDriver Monitoring
- [x] Local Deployment
- [x] Gcloud Kubernetes Deployment
- [x] Gcloud AppEngine Deployment
- [x] AWS Deployment
- [x] Cobra cli base
- [x] Wlog user interface
- [x] Local Configuration
- [ ] Etcd Configuration
- [x] Runtime Variable Configuration
- [ ] Gcloud Kubernetes Deployment & manifest
- [x] Gcloud CloudSql
- [x] Gcloud Authorization
- [x] Gcloud Storage
- [x] Terraform
- [x] Version, Build Date, Revision, App Name
- [x] Wire Dependency Injection
- [x] Auto Generate 3rd party binaries to bin/
- [ ] Docker Build


### Files Generated

- [ ] .dockerignore
- [ ] .gitignore
- [ ] main.tf
- [ ] output.tf
- [ ] variable.tf
- [ ] {{.PATH}}_server.go
- [ ] {{.PATH}}_server_register_funcs.go
- [ ] {{.PATH}}_server_test.go
- [ ] main.go
- [ ] reviewdog.yaml
- [ ] {{.App}}-deployment.yaml
- [ ] schema.sql
- [ ] roles.sql
- [ ] wire.go
- [ ] .gcloudignore

## Issues
- [ ] change config file path to home not current directory
- [ ] could not find /Users/coleman/go/src/github.com/gofunct/tools.go




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
### `gogen gen`

```commandline

Usage:
  gogen generate [command]

Aliases:
  generate, g, gen

Available Commands:
  gen-command          
  gen-scaffold-service 
  gen-service          
  gen-type             

Flags:
  -h, --help   help for generate

Global Flags:
      --debug     Debug level output
  -v, --verbose   Verbose level output

Use "gogen generate [command] --help" for more information about a command.
```

### `{appName} version`

```commandline
{appName} v0.1.1 (go1.11.4 darwin/amd64)
```

## Generated Files

### Makefile

Input: `make help`

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

## File Tree

```commandline

├── LICENSE
├── Makefile
├── README.md
├── cmd
│   ├── gogen
│   │   ├── const.go
│   │   ├── gogen.yaml
│   │   └── main.go
│   └── root.go
├── cobrafs
│   ├── app.go
│   ├── command.go
│   ├── context.go
│   ├── executor.go
│   ├── generator.go
│   ├── injectors.go
│   ├── main.go
│   ├── options.go
│   ├── providers.go
│   └── wire_gen.go
├── context
│   ├── build.go
│   ├── config.go
│   ├── context.go
│   ├── dirs.go
│   ├── entry.go
│   ├── params.go
│   └── runfunc.go
├── exec
│   ├── cmd.go
│   ├── init.go
│   ├── init_test.go
│   └── version.go
├── go.mod
├── go.sum
├── gocloud
│   ├── app.go
│   ├── aws
│   │   ├── aws.go
│   │   ├── blob.go
│   │   ├── runtimevar.go
│   │   └── user.go
│   ├── bucket.go
│   ├── flags.go
│   ├── google
│   │   ├── app.go
│   │   ├── blob.go
│   │   ├── db.go
│   │   ├── gcloud.go
│   │   ├── kube.go
│   │   ├── run.go
│   │   ├── runtime_config.go
│   │   └── user.go
│   ├── healthcheck.go
│   ├── inject_aws.go
│   ├── inject_gcp.go
│   ├── inject_local.go
│   └── wire_gen.go
├── gogen
│   ├── cmd
│   │   ├── build.go
│   │   ├── cmd.go
│   │   ├── generators.go
│   │   ├── init.go
│   │   ├── protoc.go
│   │   └── user_defined.go
│   ├── context.go
│   └── inject
│       ├── injectors.go
│       ├── providers.go
│       └── wire_gen.go
├── gogen.yaml
├── main.go
├── module
│   ├── gen.go
│   ├── generator
│   │   ├── base.go
│   │   ├── generator.go
│   │   ├── project.go
│   │   ├── status.go
│   │   └── template
│   │       ├── gen.go
│   │       ├── init
│   │       │   ├── Gopkg.toml.tmpl
│   │       │   ├── api
│   │       │   │   └── protos
│   │       │   │       └── type
│   │       │   ├── app
│   │       │   │   ├── run.go.tmpl
│   │       │   │   └── server
│   │       │   ├── cmd
│   │       │   │   └── server
│   │       │   │       └── run.go.tmpl
│   │       │   └── grapi.toml.tmpl
│   │       └── init.go
│   ├── generator.go
│   ├── script
│   │   ├── loader.go
│   │   └── script.go
│   ├── script.go
│   └── testing
│       ├── generator_mock.go
│       └── script_mock.go
├── protoc
│   ├── config.go
│   ├── config_test.go
│   ├── plugin.go
│   ├── providers.go
│   └── wrapper.go
├── templates
│   ├── _data
│   │   ├── {{.ProtoDir}}
│   │   │   └── {{.Path}}.proto.tmpl
│   │   ├── {{.RootDir}}
│   │   │   ├── Makefile.tmpl
│   │   │   ├── main.go.tmpl
│   │   │   └── reviewdog.tmpl
│   │   ├── {{.ServerDir}}
│   │   │   ├── {{.Path}}_server.go.tmpl
│   │   │   ├── {{.Path}}_server_register_funcs.go.tmpl
│   │   │   └── {{.Path}}_server_test.go.tmpl
│   │   ├── {{.StaticDir}}
│   │   │   └── guestbook.html.tmpl
│   │   └── {{.TfDir}}
│   │       ├── main.tf.tmpl
│   │       ├── output.tf.tmpl
│   │       └── variable.tf.tmpl
│   ├── fs.go
│   ├── gen.go
│   ├── vfsgen.go
│   └── virtualfs_vfsdata.go
├── test
│   ├── docker_images.txt
│   ├── user.pb.go
│   └── user.proto
├── tools.go
├── usecase
    ├── init_config.go
    ├── init_config_test.go
    └── initialize_project_usecase.go
```
.