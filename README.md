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
of interest in an assortment of log files and be notified when various options around
those patterns are met.

`log-analyse` can be leveraged as a tool for basic visibility and alerting, aswell as a
security utility.


> [!IMPORTANT]
> log-analyse will only ever need read permissions on the files it is monitoring

-----

## Planned Features

`log-analyse` aims to support the following:

 * tail mode - live monitoring of log files with rotation support etc.
 * analyse mode - retrospectively analyse log files.
 * notification integrations for alerting.
 * highly performant (and configurable) scanning of log files.
 * extensible plugin system to allow user defined behaviour on alerting.


## Quick start

`log-analyse` by default will look for a configuration file in `~/.loganalyse/loganalyse.yaml`, however you can provide
an explicit absolute path to a yaml file via the `-c` file.

An example of the current configuration (changing rapidly):

```yaml
---
files:
  # A glob based folder lookup
  - glob: "~/logs/*.txt"
    options:
      hits: 5
      period: 30s
      patterns:
        - ".*FATAL.*"
        - ".*payment failed.*"
      notify: email
      mode: sequential
  # An explicit log file
  - glob: "~/logs/foo.log"
    options:
      hits: 1
      period: 1m
      patterns:
        - ".*disk space low.*"
      notify: slack
      mode: reverse
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

-----

## Configuring log-analyse

Log analyse can be configured on a per `glob` basis.  It is possible with overlapping globs
that the same file on disk may be traversed, this behaviour is controlled by the `strict`
flag at the top level and duplicate files can cause an exit during the collection phase.

The config is composed of an array of objects, each of which currently supports the following:

* `glob`: A glob pattern for file collection.
* `options`: An object of object for all files matching the glob.
  * `hits`: How many matches before alerting.
  * `period`: Over what period should hits be considered before alerting.
  * `patterns`: Per line regex patterns for lines of interest.
  * `notify`: Which notification mechanism to fire for detections.
  * `mode`: Which strategy/mode to apply when scanning the files.


-----