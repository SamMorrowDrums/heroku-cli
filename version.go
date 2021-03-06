package main

import (
	"fmt"
	"runtime"
)

var versionTopic = &Topic{
	Name:        "version",
	Description: "print the version",
}

var versionCmd = &Command{
	Topic:       "version",
	Description: "print the version",
	Help: `Shows the Heroku CLI version.

Example:

  $ heroku version
  heroku-toolbelt/1.2.3 (x86_64-darwin11.2.0) ruby/1.9.3`,
	Run: func(ctx *Context) {
		fmt.Printf("heroku-cli/%s (%s-%s) %s", Version, runtime.GOARCH, runtime.GOOS, runtime.Version())
		if Channel != "master" {
			fmt.Printf(" %s", Channel)
		}
		fmt.Println()
	},
}
