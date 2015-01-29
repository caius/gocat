# gocat

`/bin/cat` implemented in go

## Behaviour

* Copies stdin to stdout given no arguments
* Given arguments, doesn't read from stdin
* Interprets arguments as filenames
* `-` is a special argument meaning stdin
* Copies contents of files given in order given

## License

MIT. See LICENSE for details.
