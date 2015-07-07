package twfconf

import (
  "os"
  "testing"
)



var envs []string = []string{ "LOG","PORT","DATA_DIR" }
var keys []string = []string{ "log","port","data_dir" }
var dflt []string = []string{ "/var/tmp/file.log","1234","/var/tmp/dir" }
var args ArgConf


func reset() {
  // Environment
  for _, env := range envs {
    os.Setenv(env, "")
  }

  // Argument configuration
  args = NewArgConf("testing", "Unit testing program")
  for index, opt := range keys {
    args.NewArg(opt, envs[index], dflt[index], "Inconsequential")
  }
}


// Ensure default options are what we expect
func TestConfigFlagDefaults(t *testing.T) {
  reset()

  config := args.GetArgValues([]string{})

  expected := map[string]string{
    "log":      "/var/tmp/file.log",
    "port":     "1234",
    "data_dir": "/var/tmp/dir",
  }

  for key, value := range expected {
    if (config[key] != value) {
      t.Errorf("[%s] expected to be [%s], found [%s]", key, value, config[key])
    }
  }
}


// Test that environmental variables successfully override defaults
func TestConfigEnv(t *testing.T) {
  reset()

  for _, env := range envs {
    os.Setenv(env, "8")
  }

  config := args.GetArgValues([]string{})

  for _, key := range keys {
    if (config[key] != "8") {
      t.Errorf("[%s] expected to be [8], found [%s]", key, config[key])
    }
  }
}


// Test that commandline flags successfully override defaults
func TestConfigCLI(t *testing.T) {
  reset()

  // Set commandline
  cli := []string{}
  for _, key := range keys {
    cli = append(cli, "--" + key + "=2")
  }

  config := args.GetArgValues(cli)

  for _, key := range keys {
    if (config[key] != "2") {
      t.Errorf("[%s] expected to be [2], found [%s]", key, config[key])
    }
  }
}


// Test that commandline flags successfully override environmental ones
func TestConfigOverride(t *testing.T) {
  reset()

  // Set environment
  for _, env := range envs {
    os.Setenv(env, "8")
  }

  // Set commandline
  cli := []string{}
  for _, key := range keys {
    cli = append(cli, "--" + key + "=2")
  }

  // Test
  config := args.GetArgValues(cli)
  for _, key := range keys {
    if (config[key] != "2") {
      t.Errorf("[%s] expected to be [2], found [%s]", key, config[key])
    }
  }
}
