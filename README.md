# twfconf

A thin wrapper around [spf13/cobra](https://github.com/spf13/cobra) to enable simplistic, overridable configuration for Go applications.

Inspired by "Twelve Factor" configuration, this package makes it easy to configure program configuration input via environmental variables that may be overridden by CLI parameters.

Each parameter is defined with an environment variable, CLI parameter, and default value.  The defaults may be any type which can be coerced into a string.  Once a definition has be made, input values will all be read in, overrides resolved, and final values returned in a map of key/value pairs.  Every parsed value is coerced into a string.

*NOTE:* Currently, sub-commands are not supported and are better handled by Cobra anyway.

## Example

~~~~
package main

import (
  "github.com/mikattack/twfconf"
)

config := twfconf.NewArgConf{
  Usage: "example [args]",
  Description: "API for processing example data",
}

args.NewArg(
  "log",                  // CLI parameter
  "LOG",                  // Environmental variable
  "/var/log/example.log", // Default value
  "Path to logging file." // Input parameter description
)
args.NewArg("port", "PORT", 1234, "Port to listen to HTTP requests on.")
args.NewArg("debug", "DEBUG", false, "Whether to additionally log to STDOUT")

config.GetArgValues()
~~~~

This will yield a `map[string]string`, given the following input:

~~~~
Environmental Variables:
  LOG="/var/tmp/test.log"
  PORT=4321

CLI Parameters:
  --log="/var/log/example/api.log"

Resulting Values:
  ["log"]   = "/var/log/example/api.log"
  ["port"]  = "4321"
  ["debug"] = "false"
~~~~