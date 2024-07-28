<img src="https://github.com/symonk/log-analyse/blob/main/.github/images/logo.png" border="1" width="275" height="275"/>

[![GoDoc](https://pkg.go.dev/badge/github.com/symonk/log-analyse)](https://pkg.go.dev/github.com/symonk/log-analyse)
[![Build Status](https://github.com/symonk/log-analyse/actions/workflows/go_test.yml/badge.svg)](https://github.com/symonk/log-analyse/actions/workflows/go_test.yml)
[![codecov](https://codecov.io/gh/symonk/log-analyse/branch/main/graph/badge.svg)](https://codecov.io/gh/symonk/log-analyse)
[![Go Report Card](https://goreportcard.com/badge/github.com/symonk/log-analyse)](https://goreportcard.com/report/github.com/symonk/log-analyse)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://github.com/symonk/log-analyse/blob/master/LICENSE)


> [!CAUTION]
> log-analyse is currently in alpha and not fit for production level use.


# Log Analyse

`log-analyse` is a tool for asynchronously monitoring log files for pre defined pattern
matches and causing a trigger when matches are found based on arbitrary options. It can
easily monitoring thousands of individual files for `Write` events.

`log-analyse` can be leveraged as a tool for basic visibility and alerting, aswell as a
security utility.


> [!IMPORTANT]
> log-analyse will only ever need read permissions on the files it is monitoring

-----

## Planned Features

`log-analyse` aims to support the following:

 * tail mode - live monitoring of log files with rotation support etc.
 * trigger system for dispatching actions
 * highly performant (and configurable) scanning of log files.

-----

## Triggers

for now `log-analyse` allows the following (basic) triggers:

  * `trigger:slack`: Dispatch a notification to slack.
  * `trigger:teams`: Dispatch a notification to teams. 
  * `trigger:cloud_watch`: Publish a metric to cloudwatch.
  * `trigger:shell (experimental)`: Invoke a shell script with context args.
  * `trigger:print`: Print violations to stdout.

-----


## Quick start

`log-analyse` by default will look for a configuration file in `~/.loganalyse/loganalyse.yaml`, however you can provide
an explicit absolute path to a yaml file via the `--config` file.

An example of the current configuration (changing rapidly):

```yaml
---
files:
  - glob: ~/logs/*.log
    options:
      active: false
      hits: 5
      period: 30s
      notify: email
      patterns:
        - .*FATAL.*
        - .*payment failed.*

  - glob: ~/logs/foo.log
    options:
      active: true
      hits: 1
      period: 1h10s
      notify: slack
      patterns:
        - .*critical error.*
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

----

## Configuring log-analyse

Log analyse can be configured on a per `glob` basis.  It is possible with overlapping globs
that the same file on disk may be traversed, this behaviour is controlled by the `strict`
flag at the top level and duplicate files can cause an exit during the collection phase.

The config is composed of an array of objects, each of which currently supports the following:

* `glob`: A glob pattern for file collection.
* `options`: An object of object for all files matching the glob.
  * `active`: If the glob is enabled and should be monitored.
  * `hits`: How many matches before alerting.
  * `period`: Over what period should hits be considered before alerting.
  * `patterns`: Per line regex patterns for lines of interest.
  * `trigger`: Which notification mechanism to fire for detections.


-----

## Benchmarks

`log-analyse` aims in the first pass to perform faster than `grep` with all the extra functionality baked in.
more information on benchmarks will follow later.