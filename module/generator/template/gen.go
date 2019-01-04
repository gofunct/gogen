//go:generate go-assets-builder -p template -s="/init" -o init.go -v Init init

package template

/*
Usage:
  go-assets-builder [OPTIONS] FILES...

Help Options:
  -h, --help=         Show this help message

Application Options:
  -p, --package=      The package name to generate the assets for (main)
  -v, --variable=     The name of the generated asset tree (Assets)
  -s, --strip-prefix= Strip the specified prefix from all paths
  -c, --compressed    Enable gzip compression of assets
  -o, --output=       File to write output to, or - to write to stdout (-)
 */
