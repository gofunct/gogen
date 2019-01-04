package template

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Init9d044759d52dcc3985aeb8f73ab21158401be92a = "package {{.Go.Package }}\n{{if .Methods}}\nimport (\n\t\"context\"\n\t\"testing\"\n{{range .Go.TestImports}}\n\t\"{{.}}\"\n{{- end}}\n\n\t{{.PbGo.PackageName}} \"{{ .PbGo.PackagePath }}\"\n)\n{{$go := .Go -}}\n{{$pbGo := .PbGo -}}\n{{- range .Methods}}\nfunc Test_{{$go.ServerName}}_{{.Method}}(t *testing.T) {\n\tsvr := New{{$go.ServerName}}()\n\n\tctx := context.Background()\n\treq := &{{.RequestGo $pbGo.PackageName}}{}\n\n\tresp, err := svr.{{.Method}}(ctx, req)\n\n\tif err != nil {\n\t\tt.Errorf(\"returned an error %v\", err)\n\t}\n\n\tif resp == nil {\n\t\tt.Error(\"response should not nil\")\n\t}\n}\n{{end -}}\n{{end -}}\n"
var _Inita3426f4af9db68f6bf4705d905eb9527e5207334 = "language: go\n\ngo: '1.11'\n\nenv:\n  global:\n  - DEP_RELEASE_TAG=v0.5.0\n  - FILE_TO_DEPLOY=\"dist/*\"\n\n  # GITHUB_TOKEN\n  - secure: \"MhmvXAAzOA5HY7koCfagX1wJi2mBVQsVF5cCMaNx73l+7uDgNzGYfTn4OGKmckduiGB/mp5bTJ1DeMbPq+TDX1n/RE6kndu/Q/1vw4pbxm9BsmO9b3DizIFoWlnG+EABdAZa9igbCAfv+Jj57a0WjKGaiLazylj1mb7AYj6Vao+1zvm2ufoZvpKJcnKPqcWTsx/enJD3wx0LbqTpN5a/EdynJF9kj9Z97cGk9lS/hQHqmYVUYLYG5ZIvPjkuc6ho6pYaerupZ8aQgwraupRrNAzh70C3QgxnrCK+6RRmBMchhBsHOZq1MGhbN48ttlSMKow2NyVp8mK8+wLUnQgxEvYjVNJBXf5iKMmCTBiTO8IqgAKkkMgLaB3H0UpkeOoUQNTACPxR42+FJcwObmxYRSekTGFPwAAwnZV/1BuPrpxpT7JHa9ELlShz2OVEDz9aK/WC28/oEmtYKN8s9koKr1sx4OT5c0F/XG+er2idgCWwvfK5A0Om7Fudur+bbp1a38QWb00cAu8dPTIONe01vGXQ04d+NyohS2bcvK3iehVpa+WZ4CHkjRRuv6vQGvFMNCtwwQjXopBM99+yAykLm7yqOewbzbxFI7nCHNBc1zHvI13j7yniEoI/vdWk43e2H3Az0OOtdVASNmmp5Avwo/UWzjVACvlyNK1CST4pqYQ=\"\n  # REVIEWDOG_GITHUB_API_TOKEN\n  - secure: \"HIpuAXhIivyVkMKnWucjuFWJcDnGsvBPm4lQmpCnDOWFFWgblhBzojqN9q0DK91Sc7MEeZPDD3yhZAUOYK2mcRthLZYhbblCjZsE742i8dVB9Y8+PiMb/CHRdERCQUNvQKo+fiXJ4QWE42zx9ehTkKRQHGZkHx8cQVgtSnTyMD2lxxBJWRHUQ8OS+v0yKZCmERisClccbcm37vBQQe3/n6RhuhzxIBlA5G5MEt2ig3noocMcRjApl4Qz3eV/IqVrNs8iQeSm87N3eVxuqMS07SxpOBDhyq6tlU0Ab3VD6peY8aiQxqKLCNU5w0yL5ap9jLiHAV4TDYblS7wUAJLabp+Qdj4/5C2di+jfyn1ZITcKJu8H4kAr8hZqQXpAIQ9K6e/SUztyTfVlsPl9BBO56mx4FB2ZN2voAiJSE4ZUzXyp+zIPk2eiWfclPKiPyvPgFDF0RPV0n/EQXybGoJaLgEnZ4Tx7n2WTWCnZROZkw8EuldIY60D0qJiYYTDfhk2W3XBZUJ4isqrTYKdAP/SGcBLPRmWA2/Aaq7XaP6oHa9+jPIkmhyIALtarWESRbwzWtstXXBjPaUSStZx/J/lvJW2gpmbt8e74GKEEOv9FiX2NOglwN5vPwl7ZErPMdlEeMjOx+HOIts4BPfYwtGFD7Ws0WI4oiSf4PXmuvvzjf4s=\"\n\ncache:\n  directories:\n  - $GOPATH/pkg/dep\n  - $GOPATH/pkg/mod\n  - $HOME/.cache/go-build\n\njobs:\n  include:\n  - name: lint\n    install: make setup\n    script: make lint\n    if: type = 'pull_request'\n\n  - &test\n    install: make setup\n    script: make test\n    if: type != 'pull_request'\n\n  - &test-e2e\n    name: \"E2E test (go 1.11.2)\"\n    language: bash\n    sudo: required\n    services:\n    - docker\n    env:\n    - GO_VERSION=1.11.2\n    script: make test-e2e\n    if: type != 'pull_request'\n\n  - <<: *test\n    name: \"coverage (go 1.11)\"\n    script: make cover\n    after_success: bash <(curl -s https://codecov.io/bash)\n\n  - <<: *test\n    go: master\n\n  - <<: *test\n    go: '1.10'\n\n  - <<: *test-e2e\n    name: \"E2E test (go 1.10.5)\"\n    env:\n    - GO_VERSION=1.10.5\n\n  - stage: deploy\n    install: make setup\n    script: make packages -j4\n    deploy:\n    - provider: releases\n      skip_cleanup: true\n      api_key: $GITHUB_TOKEN\n      file_glob: true\n      file: $FILE_TO_DEPLOY\n      on:\n        tags: true\n    if: type != 'pull_request'\n"
var _Inite2ca5073ca8c5591f536e9cdd49ac54c34fbf83e = "gen: # generate a server.key and server.pem file\n\topenssl genrsa -out server.key 2048\n\topenssl req -new -x509 -key server.key -out server.pem -days 3650\n\nhelp: ## help\n\t@awk 'BEGIN {FS = \":.*?## \"} /^[a-zA-Z_-]+:.*?## / {printf \"\\033[36m%-30s\\033[0m %s\\n\", $$1, $$2}' $(MAKEFILE_LIST) | sort"
var _Init8b9f72785fe81f92c86db268d2b1dee952dcc4eb = "// Copyright © 2019\n//\n// Permission is hereby granted, free of charge, to any person obtaining a copy\n// of this software and associated documentation files (the \"Software\"), to deal\n// in the Software without restriction, including without limitation the rights\n// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell\n// copies of the Software, and to permit persons to whom the Software is\n// furnished to do so, subject to the following conditions:\n//\n// The above copyright notice and this permission notice shall be included in\n// all copies or substantial portions of the Software.\n//\n// THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\n// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,\n// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE\n// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER\n// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,\n// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN\n// THE SOFTWARE.\n\npackage cmd\n\nimport (\n\t\"fmt\"\n\n\t\"github.com/spf13/cobra\"\n\t\"{{ .importPath }}/server\"\n)\n\nfunc init() {\n\trootCmd.AddCommand(runCmd)\n}\n\n// runCmd represents the run command\nvar runCmd = &cobra.Command{\n\tUse:   \"run\",\n\tShort: \"start a grpc server\",\n\tRunE: func(cmd *cobra.Command, args []string) error {\n        return server.Run()\n\t},\n}\n"
var _Inita2e3eb52bd2fece8f799a66e5e157007908dfdf2 = "The MIT License (MIT)\n\nCopyright © 2019 {{.developer}} <{{.email}}>\n\nPermission is hereby granted, free of charge, to any person obtaining a copy\nof this software and associated documentation files (the \"Software\"), to deal\nin the Software without restriction, including without limitation the rights\nto use, copy, modify, merge, publish, distribute, sublicense, and/or sell\ncopies of the Software, and to permit persons to whom the Software is\nfurnished to do so, subject to the following conditions:\n\nThe above copyright notice and this permission notice shall be included in\nall copies or substantial portions of the Software.\n\nTHE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\nIMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,\nFITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE\nAUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER\nLIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,\nOUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN\nTHE SOFTWARE.\n"
var _Inite0883af55dcd49f75702d4fdac85d33a19a73cbd = "// Code generated by github.com/gofunct/common/proto/service. DO NOT EDIT.\n\npackage {{.Go.Package }}\n\nimport (\n\t\"context\"\n\n\t\"github.com/grpc-ecosystem/grpc-gateway/runtime\"\n\t\"google.golang.org/grpc\"\n\n\t{{.PbGo.PackageName}} \"{{ .PbGo.PackagePath }}\"\n)\n\n// RegisterWithServer implements runtime.Server.RegisterWithServer.\nfunc (s *{{.Go.StructName}}) RegisterWithServer(grpcSvr *grpc.Server) {\n\t{{.PbGo.PackageName}}.Register{{.Go.ServerName}}(grpcSvr, s)\n}\n\n// RegisterWithHandler implements runtime.Server.RegisterWithHandler.\nfunc (s *{{.Go.StructName}}) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {\n\treturn {{.PbGo.PackageName}}.Register{{.ServiceName}}ServiceHandler(ctx, mux, conn)\n}\n"
var _Initfb0d1ebd50ff117e5c293e61a9603d88d31bbc4f = "package {{.Go.Package }}\n\nimport (\n\t\"context\"\n{{range .Go.Imports}}\n\t\"{{.}}\"\n{{- end}}\n\n\t{{.PbGo.PackageName}} \"{{ .PbGo.PackagePath }}\"\n)\n\n// {{.Go.ServerName}} is a composite interface of {{.PbGo.PackageName }}.{{.Go.ServerName}} and grapiserver.Server.\ntype {{.Go.ServerName}} interface {\n\t{{.PbGo.PackageName }}.{{.Go.ServerName}}\n\tgrapiserver.Server\n}\n\n// New{{.Go.ServerName}} creates a new {{.Go.ServerName}} instance.\nfunc New{{.Go.ServerName}}() {{.Go.ServerName}} {\n\treturn &{{.Go.StructName}}{}\n}\n\ntype {{.Go.StructName}} struct {\n}\n{{$go := .Go -}}\n{{$pbGo := .PbGo -}}\n{{- range .Methods}}\nfunc (s *{{$go.StructName}}) {{.Method}}(ctx context.Context, req *{{.RequestGo $pbGo.PackageName}}) (*{{.ResponseGo $pbGo.PackageName}}, error) {\n\t// TODO: Not yet implemented.\n\treturn nil, status.Error(codes.Unimplemented, \"TODO: You should implement it!\")\n}\n{{end -}}\n"
var _Init560bfd3a40b63143fcc3be11ec4a6978feb0fd14 = "// Copyright © 2019\n//\n// Permission is hereby granted, free of charge, to any person obtaining a copy\n// of this software and associated documentation files (the \"Software\"), to deal\n// in the Software without restriction, including without limitation the rights\n// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell\n// copies of the Software, and to permit persons to whom the Software is\n// furnished to do so, subject to the following conditions:\n//\n// The above copyright notice and this permission notice shall be included in\n// all copies or substantial portions of the Software.\n//\n// THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\n// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,\n// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE\n// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER\n// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,\n// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN\n// THE SOFTWARE.\n\npackage cmd\n\nimport (\n\t\"fmt\"\n\t\"os\"\n\t\"github.com/spf13/cobra\"\n\t\"github.com/spf13/viper\"\n\t\"github.com/gofunct/gogen/gocloud\"\n\t\"github.com/gofunct/common/logging\"\n)\n\nvar (\n    cfgFile string\n    config = viper.New()\n    )\n\n// rootCmd represents the base command when called without any subcommands\nvar rootCmd = &cobra.Command{\n\tUse:   \"temp\",\n\tShort: \"A brief description of your application\",\n\tLong: `A longer description that spans multiple lines and likely contains\nexamples and usage of using your application. For example:\n\nCobra is a CLI library for Go that empowers applications.\nThis application is a tool to generate the needed files\nto quickly create a Cobra application.`,\n\t// Uncomment the following line if your bare application\n\t// has an action associated with it:\n\t//\tRun: func(cmd *cobra.Command, args []string) { },\n}\n\n// Execute adds all child commands to the root command and sets flags appropriately.\n// This is called by main.main(). It only needs to happen once to the rootCmd.\nfunc Execute() {\n\tif err := rootCmd.Execute(); err != nil {\n\t\tfmt.Println(err)\n\t\tos.Exit(1)\n\t}\n}\n\nfunc init() {\n\tcobra.OnInitialize(initConfig)\n\tlogging.AddFlags(rootCmd)\n\trootCmd.PersistentFlags().StringVar(&cfgFile, \"config\", \"\", \"config file (default is $HOME/.temp.yaml)\")\n\trootCmd.Flags().BoolP(\"toggle\", \"t\", false, \"Help message for toggle\")\n\trootCmd.AddCommand(gocloud.NewGocloudCommand)\n}\n\n// initConfig reads in config file and ENV variables if set.\nfunc initConfig() {\n\tconfig.AutomaticEnv()\n\tif cfgFile != \"\" {\n\t\t// Use config file from the flag.\n\t\tconfig.SetConfigFile(cfgFile)\n\t} else {\n\t\tconfig.AddConfigPath(\".\")\n        config.SetConfigName(\"config\")\n\n\t}\n\n\t// If a config file is found, read it in.\n\tif err := config.ReadInConfig(); err == nil {\n\t\tfmt.Println(\"Using config file:\", viper.ConfigFileUsed())\n\t}\n}\n"
var _Init2f8b03d8b9e9a1bf72da6ea9424da9a231fef8c2 = ".travis.yml\n.reviewdog.yml\n_tests\nbin\nvendor\nREADME.md\nLICENSE\n"
var _Initf08871a4ac25f3ef36434d141f852b90f2f2bc45 = "coverage:\n  status:\n    project:\n      default:\n        threshold: 1%\n    patch: off\n"
var _Init8d21956ba8abe388f964e47be0f7e5d170a2fce5 = ""
var _Init4a622176ca163dbfa742f4811d5365ee115f3311 = "// +build tools\n\npackage main\n\n// tool dependencies\nimport (\n\t_ \"github.com/golang/mock/mockgen\"\n\t_ \"github.com/golang/protobuf/protoc-gen-go\"\n\t_ \"github.com/google/wire/cmd/wire\"\n\t_ \"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway\"\n\t_ \"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger\"\n\t_ \"github.com/haya14busa/reviewdog/cmd/reviewdog\"\n\t_ \"github.com/izumin5210/gex/cmd/gex\"\n\t_ \"github.com/jessevdk/go-assets-builder\"\n\t_ \"github.com/kisielk/errcheck\"\n\t_ \"github.com/mitchellh/gox\"\n\t_ \"github.com/srvc/wraperr/cmd/wraperr\"\n\t_ \"golang.org/x/lint/golint\"\n\t_ \"honnef.co/go/tools/cmd/megacheck\"\n\t_ \"mvdan.cc/unparam\"\n)\n"
var _Initf8f5781f016ec71243395dbbb7bb65f373df2373 = "runner:\n  golint:\n    cmd: golint $(go list ./... | grep -v /vendor/)\n    format: golint\n  govet:\n    cmd: go vet $(go list ./... | grep -v /vendor/)\n    format: govet\n  errcheck:\n    cmd: errcheck -asserts -ignoretests -blank ./...\n    errorformat:\n      - \"%f:%l:%c:%m\"\n  wraperr:\n    cmd: wraperr ./...\n    errorformat:\n      - \"%f:%l:%c:%m\"\n  megacheck:\n    cmd: megacheck ./...\n    errorformat:\n      - \"%f:%l:%c:%m\"\n  unparam:\n    cmd: unparam ./...\n    errorformat:\n      - \"%f:%l:%c: %m\"\n"
var _Init77f37dc48057e4a71e6c5bdb397ec647111d06f1 = "syntax = \"proto3\";\npackage echopb;\n\nimport \"google/api/annotations.proto\";\n\nmessage EchoMessage {\n string value = 1;\n}\n\nservice EchoService {\n  rpc Echo(EchoMessage) returns (EchoMessage) {\n    option (google.api.http) = {\n      post: \"/v1/echo\"\n      body: \"*\"\n    };\n  }\n}"
var _Init109bef6bea90b5fdea7ed03f3c2b9a8f3a039ed4 = ""
var _Initcdd0be5fbec344ebcfaa6210bc26793a425b02f1 = ""
var _Init840caece5084c7325cd0babad92224d9ac8a13ce = "# {{.packageName}}\n\n* Developer: {{.developer}}\n* Email: {{.email}}\n* Language: Golang\n* Download: `go get github.com/replace-me`\n* Summary: {.summary}}\n## Table of Contents\n\n- [{{.packageName}}](#{{.packageName}})\n  * [Table of Contents](#table-of-contents)."
var _Init033f9d45db331f8a4c2cfc3abbf3009df9a1538a = "// Copyright © 2019 Coleman Word <coleman.word@gofunct.com>\n//\n// Permission is hereby granted, free of charge, to any person obtaining a copy\n// of this software and associated documentation files (the \"Software\"), to deal\n// in the Software without restriction, including without limitation the rights\n// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell\n// copies of the Software, and to permit persons to whom the Software is\n// furnished to do so, subject to the following conditions:\n//\n// The above copyright notice and this permission notice shall be included in\n// all copies or substantial portions of the Software.\n//\n// THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\n// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,\n// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE\n// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER\n// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,\n// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN\n// THE SOFTWARE.\n\npackage init\n\nimport \"{{ .importPath }}/cmd\"\n\nfunc main() {\n\tcmd.Execute()\n}\n"
var _Initbe517013ae530f11bebf119ec23f1f0d1a7b873c = "*.key\n*.pem"
var _Init709127a14334557e0f1f54a6a3ce570fb512c73c = ""
var _Init23b808cac963edf44a497827f2a6eff5ddac970f = ""
var _Init38e76c5db8962fa825cf2bd8b23a2dc985c4513e = "*.so\n/vendor\n/bin\n/tmp\n/.idea\n"
var _Init08a891419f06f3f4aa9b20a663cd7ef7e6523e88 = "package server\n\nimport (\n\t\"os\"\n)\n\nfunc Run() error {\n\treturn Run()\n}\n"
var _Init4b73f9b96c39121a8cbcb9892e9de616bc945621 = "package = \"{{.packageName}}\"\n\n[gogen]\nserver_dir = \"./app/server\"\n\n[protoc]\nprotos_dir = \"./api/protos\"\nout_dir = \"./api\"\nimport_dirs = [\n  \"./api/protos\",\n  \"./vendor/github.com/grpc-ecosystem/grpc-gateway\",\n  \"./vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis\",\n]\n\n  [[protoc.plugins]]\n  name = \"go\"\n  args = { plugins = \"grpc\", paths = \"source_relative\" }\n\n  [[protoc.plugins]]\n  name = \"grpc-gateway\"\n  args = { logtostderr = true, paths = \"source_relative\" }\n\n  [[protoc.plugins]]\n  name = \"swagger\"\n  args = { logtostderr = true }\n"

// Init returns go-assets FileSystem
var Init = assets.NewFileSystem(map[string][]string{"/api/protos": []string{"service.proto.tmpl", ".keep.tmpl"}, "/": []string{"Gopkg.toml.tmpl", "LICENSE", ".dockerignore.tmpl", ".gitignore.tmpl", ".codecov.yml.tmpl", ".travis.yml.tmpl", "main.go.tmpl", ".reviewdog.yml.tmpl", "README.md.tmpl", "tools.go.tmpl", "config.toml.tmpl"}, "/cmd": []string{"root.go.tmpl", ".keep.tmpl", "run.go.tmpl"}, "/server": []string{"register.go.tmpl", "server_test.go.tmpl", "server.go.tmpl", ".keep.tmpl", "run.go.tmpl"}, "/certs": []string{"Makefile.tmpl", ".gitignore.tmpl"}, "/api": []string{".keep.tmpl"}}, map[string]*assets.File{
	"/cmd": &assets.File{
		Path:     "/cmd",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1546571922, 1546571922227386910),
		Data:     nil,
	}, "/cmd/.keep.tmpl": &assets.File{
		Path:     "/cmd/.keep.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546143200, 1546143200000000000),
		Data:     []byte(_Init109bef6bea90b5fdea7ed03f3c2b9a8f3a039ed4),
	}, "/server/.keep.tmpl": &assets.File{
		Path:     "/server/.keep.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546143200, 1546143200000000000),
		Data:     []byte(_Initcdd0be5fbec344ebcfaa6210bc26793a425b02f1),
	}, "/README.md.tmpl": &assets.File{
		Path:     "/README.md.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1546562416, 1546562416171112173),
		Data:     []byte(_Init840caece5084c7325cd0babad92224d9ac8a13ce),
	}, "/main.go.tmpl": &assets.File{
		Path:     "/main.go.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1546569663, 1546569663276738627),
		Data:     []byte(_Init033f9d45db331f8a4c2cfc3abbf3009df9a1538a),
	}, "/certs/.gitignore.tmpl": &assets.File{
		Path:     "/certs/.gitignore.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1546571751, 1546571751050275497),
		Data:     []byte(_Initbe517013ae530f11bebf119ec23f1f0d1a7b873c),
	}, "/server": &assets.File{
		Path:     "/server",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1546570199, 1546570199971144882),
		Data:     nil,
	}, "/api/.keep.tmpl": &assets.File{
		Path:     "/api/.keep.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546143200, 1546143200000000000),
		Data:     []byte(_Init709127a14334557e0f1f54a6a3ce570fb512c73c),
	}, "/Gopkg.toml.tmpl": &assets.File{
		Path:     "/Gopkg.toml.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546143200, 1546143200000000000),
		Data:     []byte(_Init23b808cac963edf44a497827f2a6eff5ddac970f),
	}, "/.gitignore.tmpl": &assets.File{
		Path:     "/.gitignore.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546562094, 1546562094856946096),
		Data:     []byte(_Init38e76c5db8962fa825cf2bd8b23a2dc985c4513e),
	}, "/server/run.go.tmpl": &assets.File{
		Path:     "/server/run.go.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546570199, 1546570199970488825),
		Data:     []byte(_Init08a891419f06f3f4aa9b20a663cd7ef7e6523e88),
	}, "/certs": &assets.File{
		Path:     "/certs",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1546571780, 1546571780516538287),
		Data:     nil,
	}, "/config.toml.tmpl": &assets.File{
		Path:     "/config.toml.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546562155, 1546562155431827752),
		Data:     []byte(_Init4b73f9b96c39121a8cbcb9892e9de616bc945621),
	}, "/server/server_test.go.tmpl": &assets.File{
		Path:     "/server/server_test.go.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546548309, 1546548309000000000),
		Data:     []byte(_Init9d044759d52dcc3985aeb8f73ab21158401be92a),
	}, "/.travis.yml.tmpl": &assets.File{
		Path:     "/.travis.yml.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1545633496, 1545633496000000000),
		Data:     []byte(_Inita3426f4af9db68f6bf4705d905eb9527e5207334),
	}, "/certs/Makefile.tmpl": &assets.File{
		Path:     "/certs/Makefile.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1546571780, 1546571780513798465),
		Data:     []byte(_Inite2ca5073ca8c5591f536e9cdd49ac54c34fbf83e),
	}, "/cmd/run.go.tmpl": &assets.File{
		Path:     "/cmd/run.go.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1546570199, 1546570199963425290),
		Data:     []byte(_Init8b9f72785fe81f92c86db268d2b1dee952dcc4eb),
	}, "/LICENSE": &assets.File{
		Path:     "/LICENSE",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1546562554, 1546562554424344554),
		Data:     []byte(_Inita2e3eb52bd2fece8f799a66e5e157007908dfdf2),
	}, "/server/register.go.tmpl": &assets.File{
		Path:     "/server/register.go.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546548309, 1546548309000000000),
		Data:     []byte(_Inite0883af55dcd49f75702d4fdac85d33a19a73cbd),
	}, "/server/server.go.tmpl": &assets.File{
		Path:     "/server/server.go.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546548309, 1546548309000000000),
		Data:     []byte(_Initfb0d1ebd50ff117e5c293e61a9603d88d31bbc4f),
	}, "/tools.go.tmpl": &assets.File{
		Path:     "/tools.go.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1546571978, 1546571978590544741),
		Data:     []byte(_Init4a622176ca163dbfa742f4811d5365ee115f3311),
	}, "/cmd/root.go.tmpl": &assets.File{
		Path:     "/cmd/root.go.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1546571922, 1546571922221362753),
		Data:     []byte(_Init560bfd3a40b63143fcc3be11ec4a6978feb0fd14),
	}, "/.dockerignore.tmpl": &assets.File{
		Path:     "/.dockerignore.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1545633496, 1545633496000000000),
		Data:     []byte(_Init2f8b03d8b9e9a1bf72da6ea9424da9a231fef8c2),
	}, "/.codecov.yml.tmpl": &assets.File{
		Path:     "/.codecov.yml.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546562554, 1546562554432824735),
		Data:     []byte(_Initf08871a4ac25f3ef36434d141f852b90f2f2bc45),
	}, "/api": &assets.File{
		Path:     "/api",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1546570168, 1546570168530812862),
		Data:     nil,
	}, "/api/protos/.keep.tmpl": &assets.File{
		Path:     "/api/protos/.keep.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546143200, 1546143200000000000),
		Data:     []byte(_Init8d21956ba8abe388f964e47be0f7e5d170a2fce5),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1546571978, 1546571978596436036),
		Data:     nil,
	}, "/.reviewdog.yml.tmpl": &assets.File{
		Path:     "/.reviewdog.yml.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1545633496, 1545633496000000000),
		Data:     []byte(_Initf8f5781f016ec71243395dbbb7bb65f373df2373),
	}, "/api/protos": &assets.File{
		Path:     "/api/protos",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1546571618, 1546571618288328213),
		Data:     nil,
	}, "/api/protos/service.proto.tmpl": &assets.File{
		Path:     "/api/protos/service.proto.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546571618, 1546571618287649317),
		Data:     []byte(_Init77f37dc48057e4a71e6c5bdb397ec647111d06f1),
	}}, "")
