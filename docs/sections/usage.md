[Back to Table of Contents](../documentation.md)


## Yot

Yot (YAML Overlay Tool) is a YAML overlay tool.

### Synopsis

Yot (YAML Overlay Tool) is a YAML overlay tool which uses a YAML schema to 
	define overlay operations on a set of YAML documents. yot only produces valid YAML 
	documents on output, and can preserve and inject comments.

```
yot [flags]
```

### Examples

```
yot -i instructions.yaml -o /tmp/output
```

### Options

```
Available Commands:
  completion  Generate shell auto-completion scripts
  help        Help about any command

Flags:
  -h, --help                      Help for Yot.
  -I, --indent-level int          Number of spaces to be used for indenting YAML output (min: 2, max: 9) (default 2).
  -i, --instructions string       Path to instructions file (required).
  -o, --output-directory string   Path to directory for writing the YAML files which were operated on (default "./output").
  -s, --stdout                    Output YAML files which were operated on to stdout.
  -V, --verbose                   Verbose output.
  -v, --version                   Version for Yot.
```

[Back to Table of Contents](../documentation.md)  
[Next Up: Instructions File YAML Schema](instructionsFile.md)
