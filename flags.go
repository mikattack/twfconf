/* 
 * Library to allow applications to define their input configuration through
 * environmental variables, define defaults for them, and override those
 * settings via command-line parameters.
 * 
 * Configuration variables are referenced by their CLI name (by convention
 * lowercase alphanumeric with dashes or underscores).  Their value is
 * determined according to precedence, from highest to lowest:
 * 
 *    1. CLI parameter
 *    2. Environmental variable
 *    3. Default value
 * 
 * All values are coerced into strings to simplify option definition,
 * parsing, and output.
 */
package main

import (
  "os"
  "github.com/spf13/cobra"
)


/*
 * Defines an argument configuration layout and usage information for the
 * program.  Eventually contains the program's parsed input values.
 */
type ArgConf struct {
  envs      map[string]string
  config    map[string]string
  desc      map[string]string
  command   cobra.Command
}


/* 
 * Initialize a new input configuration.
 */
func NewArgConf(usage string, description string) ArgConf {
  cmd := cobra.Command{
    Use:    "misato-provision-api",
    Short:  "Cloud resource provisioning API",
    Run:    func (cmd *cobra.Command, args []string) {
      // No operation, but necessary for Cobra to recognize a command
    },
  }
  return ArgConf{
    config:   map[string]string{},
    envs:     map[string]string{},
    desc:     map[string]string{},
    command:  cmd,
  }
}


/* 
 * Add argument/environmental variable to the input configuration.
 */
func (ac *ArgConf) NewArg(cli string, env string, initial interface{}, description string) {
  value := initial.(string)
  ac.config[cli] = value
  ac.envs[env] = cli
  ac.desc[cli] = description
}


/* 
 * Read values from the environment and from a given argument collection
 * (typically "os.Args[1:]") according to a configuration.  Arguments
 * from the argument collection override those from the environment.
 */
func (ac *ArgConf) GetArgValues(args []string) map[string]string {

  // Override default values with environmental variables, if they exist
  for opt, mapping := range ac.envs {
    value := os.Getenv(opt)
    if len(value) > 0 {
      ac.config[mapping] = value
    }
  }

  // Override environmental variables with commandline flags, if they exist
  flags := map[string]*string{}
  ac.command.SetArgs(args)  // Passed in to enable unit testing
  for opt, value := range ac.config {
    flags[opt] = ac.command.Flags().String(opt, value, ac.desc[opt])
  }
  ac.command.Execute()
  for opt, ptr := range flags {
    if len(*ptr) > 0 {
      ac.config[opt] = *ptr
    }
  }

  return ac.config
}
