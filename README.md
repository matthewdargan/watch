# Watch

Watch runs a command each time any file in the current directory is written.

Usage:

    watch [-r] cmd [args...]

The `-r` flag causes watch to monitor the current directory and all
subdirectories for modifications.

## Examples

Run tests on file changes in the current directory:

```sh
$ watch go test ./...
```

Run tests on file changes recursively from the current directory:

```sh
$ watch -r go test ./...
```
