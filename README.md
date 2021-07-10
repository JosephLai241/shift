         _   _ ___ _   
     ___| |_|_|  _| |_ 
    |_ -|   | |  _|  _|
    |___|_|_|_|_| |_|

[![Go Badge](https://img.shields.io/badge/language-Go-9cf?logo=go)][Go]

# Table of Contents

* [Introduction](#introduction)
* [How `shift` Works](#how-shift-works)
* [How To Use `shift`](#how-to-use-shift)
    + [Clocking In](#clocking-in)
    + [Clocking Out](#clocking-out)
    + [Check Current Status](#check-current-status)

# Introduction

`shift` is a command-line tool for keeping track of time spent during a shift. Designed for contractors/remote workers who need to track their own hours. 

This project is largely inspired by Luke Schenk's Python CLI tool [`clck`][clck].

# How `shift` Works

This tool will create a directory `shift_timesheets/CURRENT_YEAR` within the directory in which you ran `shift`. User-entered information is then recorded into a CSV file titled `CURRENT_MONTH.csv`, located in the `CURRENT_YEAR` directory.

The information that is always recorded into the CSV file includes:

* Clock-in time
* Clock-out time

Optional data may also be recorded for both clock-in/out records. This includes:

* An accompanying message
* Company name

# How To Use `shift`

## Clocking In

**usage: `shift in`**

Use this command to clock-in. `shift` will write the clock-in time to the 

## Clocking Out

**usage: `shift out`**

## Check Current Status

**usage: `shift status`**

<!-- Links -->
[Go]: https://golang.org/
[clck]: https://github.com/LukeDSchenk/clck
