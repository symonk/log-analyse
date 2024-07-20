<img src="https://github.com/symonk/log-analyse/blob/main/.github/images/logo.png" border="1" width="275" height="275"/>

[![GoDoc](https://pkg.go.dev/badge/github.com/symonk/log-analyse)](https://pkg.go.dev/github.com/symonk/log-analyse)
[![Build Status](https://github.com/symonk/log-analyse/actions/workflows/go_test.yml/badge.svg)](https://github.com/symonk/log-analyse/actions/workflows/go_test.yml)
[![codecov](https://codecov.io/gh/symonk/log-analyse/branch/main/graph/badge.svg)](https://codecov.io/gh/symonk/log-analyse)
[![Go Report Card](https://goreportcard.com/badge/github.com/symonk/log-analyse)](https://goreportcard.com/report/github.com/symonk/log-analyse)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://github.com/symonk/log-analyse/blob/master/LICENSE)


> [!CAUTION]
> log-analyse is currently in alpha and not fit for production level use.


# Log Analyse

`log-analyse` allows scanning hundreds of log files for pre-determined pattern matches.
The aim of `log-analyse` is to allow teams to store an array of patterns that may be
of interest in an assortment of log files and be notified when various thresholds around
those patterns are met.


> [!IMPORTANT]
> log-analyse will only ever need read permissions on the files it is monitoring

-----

## Planned Features

`log-analyse` hops to support the following:

 * tail (daemon) mode that is rotation aware etc
 * notification integrations for alerting
 * highly performant scanning of log files


## Quick start

`log-analyse` by default will look for a configuration file in `~/.loganalyse/loganalyse.yaml`, however you can provide
an explicit absolute path to a yaml file via the `-c` file.

An example of the current configuration (changing rapidly):

```yaml
---
files:
  # A glob based folder lookup
  - loc: "~/folder1"
    threshold:
      glob: "*.log"
      times: 5
      period: 30s
      patterns:
        - ".*FATAL.*"
        - ".*payment failed.*"
      notify: "email"
  # An explicit log file
  - loc: "~/folder2/explicit.log"
    threshold:
      times: 1
      period: 1m
    patterns:
      - ".*disk space low.*"
    notify: "slack"

## TBD
integrations:
    slack:
        webhook: ""
    email:
        to:
            - "alerts@bar.com"
            - "admin@bar.com"
```

-----

## Running Log-analyse

Running log analyse on your system is as easy as:

```bash
# ensure to use the minimum permissions necessary for the below:
go install github.com/symonk/log-analyse
mkdir ~/.loganalyse/loganalyse.yaml
# populate loganalyse.yaml with your configuration
log-analyse
```