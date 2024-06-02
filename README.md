[![Go Reference](https://pkg.go.dev/badge/github.com/jackc/envconf.svg)](https://pkg.go.dev/github.com/jackc/envconf)
[![Build Status](https://github.com/jackc/envconf/actions/workflows/ci.yml/badge.svg)](https://github.com/jackc/envconf/actions/workflows/ci.yml)

# Envconf

Envconf is a simple, zero dependency library for managing configuration from the environment. It adds default values,
variable descriptions, and the ability to inject a environment fetching function.

Example usage:

```go
config := envconf.New()

// Register configuration items.
config.Register(envconf.Item{
  Name:        "FOO",
  Default:     "default-foo",
  Description: "the foo",
})

config.Register(envconf.Item{
  Name:        "BAR",
  Default:     "default-bar",
  Description: "the bar",
})

foo := config.Value("FOO") // returns ENV["FOO"] or "default-foo"
bar = config.Value("BAR") // returns ENV["BAR"] or "default-bar"
```
