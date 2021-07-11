         _   _ ___ _   
     ___| |_|_|  _| |_ 
    |_ -|   | |  _|  _|
    |___|_|_|_|_| |_|

[![Go Badge](https://img.shields.io/badge/-lang-blue?logo=go)][Go]

# Table of Contents

* [Introduction](#introduction)
* [Installation](#installation)
* [How `shift` Works](#how-shift-works)
* [How To Use `shift`](#how-to-use-shift)
    + [Clocking In](#clocking-in)
    + [Clocking Out](#clocking-out)
    + [Check Current Status](#check-current-status)

# Introduction

`shift` is a command-line tool for keeping track of shifts. Its primary audience is contractors/remote workers who need to track their own hours, but it is still useful for logging time spent doing anything.

This project is based on Luke Schenk's Python CLI tool [`clck`][clck].

# Installation

You will need [Go][Go] installed on your computer to compile the source code.

```
git clone --depth=1 git@github.com:JosephLai241/shift.git
cd shift/
go build
```

# How `shift` Works

This tool will create a directory `shift_timesheets/CURRENT_YEAR` within the current working directory. User-entered information is then recorded into a CSV-format timesheet titled `CURRENT_MONTH.csv`, located in the `CURRENT_YEAR` directory.

This is an example of the `shift_timesheets` directory structure if you ran `shift` sometime during July 2021:

```
shift_timesheets/
└── 2021
    └── July.csv
```

The information that is always recorded into the timesheet includes:

* Clock-in time
* Clock-out time

Optionally, an accompanying clock-in or clock-out message will also be written to the timesheet.

# How To Use `shift`

## Clocking In

**usage: `shift in`**

Use this command to clock-in. The clock-in time to the `CURRENT_MONTH.csv` file.

The `message` will also be written if a value is included.

## Clocking Out

**usage: `shift out`**

Use this command to clock-out. The clock-in time to the `CURRENT_MONTH.csv` file.

The `message` will also be written if a value is included.

## Check Current Status

**usage: `shift status`**

Use this command to display the current shift status.

This is a table of behaviors that can come from running this command:

| Currently Clocked In      | Clocked Out/Inactive           | No Shifts Tracked |
|---------------------------|--------------------------------|-------------------|
| `CLOCK_IN_TIME`           | `LAST_CLOCK_IN`                | **Error message   |
| * `CLOCK_IN_MESSAGE`      | * `LAST_CLOCK_IN_MESSAGE`      |                   |

\* Only displayed if included during clock-in/out.

\** If you have never run `shift` prior to running the `status` command, an error message will inform you to track a shift before attempting to run the command.

<!-- Links -->
[Go]: https://golang.org/
[clck]: https://github.com/LukeDSchenk/clck
