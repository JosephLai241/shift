         _   _ ___ _   
     ___| |_|_|  _| |_ 
    |_ -|   | |  _|  _|
    |___|_|_|_|_| |_|

[![Go Badge](https://img.shields.io/badge/-lang-blue?logo=go)][Go]

# Table of Contents

* [Introduction](#introduction)
* [Installation](#installation)
* [How `shift` Works](#how-shift-works)
    + [Storing Data Into Timesheets](#storing-data-into-timesheets)
    + [Storing Data Into a SQLite Instance](#storing-data-into-a-sqlite-instance)
* [How To Use `shift`](#how-to-use-shift)
    + [Clocking In](#clocking-in)
    + [Clocking Out](#clocking-out)
    + [Check Current Status](#check-current-status)
    + [Amend Shift Message](#amend-shift-message)
    + [List Tracked Shifts](#list-tracked-shifts)
    + [Delete a Shift](#delete-a-shift)
* [How to Set the Storage Option](#how-to-set-the-storage-option)

# Introduction

`shift` is a command-line tool for keeping track of shifts. Its primary audience is contractors (like myself) or remote workers who need to track their own hours, but it is still useful for logging time spent doing anything.

This project is based on Luke Schenk's Python CLI tool [`clck`][clck].

# Installation

You will need [Go][Go] installed on your computer to compile the source code.

```
git clone --depth=1 git@github.com:JosephLai241/shift.git
cd shift/
go build
```

> ***NOTE:*** This program initializes and reads from files in your current working directory. Run `shift` in a directory in which you would like all your records to be stored.

# How `shift` Works

This tool is capable of storing shift data into CSV spreadsheets or a local SQLite instance. **The default is storing into CSV spreadsheets**, but you can configure which storage option you would like to use in the `.shiftconfig.yml` dotfile. The [How to Set the Storage Option](#how-to-set-the-storage-option) provides information for how to do so.

## Storing Data Into Timesheets

The directory `shift_timesheets/CURRENT_YEAR` is created within the current working directory. User-entered information is then recorded into a CSV-format timesheet titled `CURRENT_MONTH.csv`, located in the `CURRENT_YEAR` directory.

This is an example of the `shift_timesheets` directory structure if you ran `shift` sometime during July 2021:

```
shift_timesheets/
└── 2021
    └── July.csv
```

The information that is always recorded into the timesheet includes:

* Clock-in time
* Clock-out time
* Shift duration

Optionally, an accompanying clock-in or clock-out message will also be written to the timesheet.

## Storing Data Into a SQLite Instance

A SQLite instance `shifts.db` is created within the current working directory. The data is then stored in a table labeled with the current year, which is then linked to a sub-table labeled with the current month.

This is an example of the database structure if you ran `shift` sometime during July 2021:

```
shifts.db
└── TABLE `year`
    └── TABLE `Y_2021`
        └── TABLE `M_July`
```

# How To Use `shift`

## Clocking In

```
shift in

    -m "OPTIONAL MESSAGE"
```

Use this command to clock-in. The clock-in time is then written to the timesheet or database.

You can record a message corresponding to your clock in by including the `-m` flag.

## Clocking Out

```
shift out

    -m "OPTIONAL MESSAGE"
```

Use this command to clock-out. The clock-in time is then written to the timesheet or database.

You can record a message corresponding to your clock in by including the `-m` flag.

## Check Current Status

```
shift status
```

Use this command to display the current shift status.

This is a table of behaviors that can come from running this command:

| Currently Clocked In | Clocked Out/Inactive       | No Shifts Tracked |
|----------------------|----------------------------|-------------------|
| `CLOCK_IN_TIME`      | `LAST_CLOCK_OUT_TIME`      | **Error message   |
| * `CLOCK_IN_MESSAGE` | * `LAST_CLOCK_OUT_MESSAGE` |                   |

\* Only displayed if included during clock-in/out.

\** If you have never run `shift` prior to running the `status` command, an error message will inform you to track a shift before attempting to run the command.

## Amend Shift Message

```
shift amend (in|out)

    -d DAY_OF_THE_WEEK
    -m MONTH
    -y YEAR
```

Use this command to amend the most recent shift's clock-in or clock-out message.

You can amend a different record's clock-in or clock-out message by including the `-d`, `-m`, and/or `-y` flags. Combine these flags to narrow your search.

> ***NOTE:*** Type the entire day of the week when using the `-d` flag, ie. Monday.

> ***NOTE:*** Type the entire month name when using the `-m` flag, ie. January.

> ***NOTE:*** Type the entire year in YYYY format when using the `-y` flag, ie. 2021.

## List Tracked Shifts

```
shift list

    -d DAY_OF_THE_WEEK
    -m MONTH
    -y YEAR
```

Use this command to list all recorded shifts for the current month.

You can display the recorded shifts for a different day, month, and/or year by including the `-d`, `-m`, and/or `-y` flags. Combine these flags to narrow your search.

> ***NOTE:*** Type the entire day of the week when using the `-d` flag, ie. Monday.

> ***NOTE:*** Type the entire month name when using the `-m` flag, ie. January.

> ***NOTE:*** Type the entire year in YYYY format when using the `-y` flag, ie. 2021.

## Delete a Shift

```
shift delete

    -d DAY_OF_THE_WEEK
    -m MONTH
    -y YEAR
```

Use this command to delete the most recent shift.

You can delete a different shift by including the `-d`, `-m`, and/or `-y` flags. Combine these flags to narrow your search.

> ***NOTE:*** Type the entire day of the week when using the `-d` flag, ie. Monday.

> ***NOTE:*** Type the entire month name when using the `-m` flag, ie. January.

> ***NOTE:*** Type the entire year in YYYY format when using the `-y` flag, ie. 2021.

# How to Set the Storage Option

The `.shiftconfig.yml` configuration file only contains one line:

```yaml
storage-type: timesheet
```

As mentioned before, storing shift data in CSV timesheets is the default storage option and is preset within the YAML file. There are two accepted values:

* `timesheet`
* `database`

> ***NOTE***: `shift` will process your change on your next clock in if the `storage-type` value is changed while you are clocked in. 
>
> For example, if `storage-type` is currently set to `timesheet` and you change the value to `database` while clocked in, your clock-out data will still be written to the current month's timesheet. Shift data will be written to the `shifts.db` SQLite instance on your next clock-in.

<!-- Links -->
[Go]: https://golang.org/
[clck]: https://github.com/LukeDSchenk/clck
