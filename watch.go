// Copyright 2024 Matthew P. Dargan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Watch runs a command each time any file in the current directory is written.
//
// Usage:
//
//	watch [-r] cmd [args...]
//
// The -r flag causes watch to monitor the current directory and all
// subdirectories for modifications.
//
// Examples:
//
// Run tests on file changes in the current directory:
//
//	$ watch go test ./...
//
// Run tests on file changes recursively from the current directory:
//
//	$ watch -r go test ./...
package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var recursive = flag.Bool("r", false, "watch all subdirectories recursively")

func usage() {
	fmt.Fprintf(os.Stderr, "usage: watch [-r] cmd [args...]\n")
	os.Exit(2)
}

func main() {
	log.SetPrefix("watch: ")
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() == 0 {
		usage()
	}
	args := flag.Args()
	var prevMod, currMod time.Time
	for {
		err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !*recursive && d.IsDir() && path != "." {
				return filepath.SkipDir
			}
			info, err := d.Info()
			if err != nil {
				return err
			}
			modTime := info.ModTime()
			if modTime.After(currMod) {
				currMod = modTime
			}
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
		if currMod.After(prevMod) {
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			_ = cmd.Run()
			prevMod = currMod
		}
		time.Sleep(1 * time.Second)
	}
}
