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
        * [Display Records for a Different Month](#display-records-for-a-different-month)
        * [Display Records for a Different Year](#display-records-for-a-different-year)
        * [Combining Optional Flags](#combining-optional-flags)
* [How to Set the Storage Option](#how-to-set-the-storage-option)

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

This tool is capable of storing shift data into a SQLite instance or into CSV spreadsheets. **The default is storing into CSV spreadsheets**, but you can configure which storage option you would like to use in the `.shiftconfig.yml` dotfile. The [How to Set the Storage Option](#how-to-set-the-storage-option) provides information for how to do so.

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
    └── TABLE `2021`
        └── TABLE `july`
```

# How To Use `shift`

## Clocking In

**usage: `shift in`**

Use this command to clock-in. The clock-in time is then written to the `CURRENT_MONTH.csv` file.

You can add a message corresponding to your clock in by including the `-m` flag:

**usage: `shift in -m "YOUR MESSAGE HERE"`**

The message will also be written to the timesheet.

## Clocking Out

**usage: `shift out`**

Use this command to clock-out. The clock-in time is then written to the `CURRENT_MONTH.csv` file.

You can add a message corresponding to your clock in by including the `-m` flag:

**usage: `shift out -m "YOUR MESSAGE HERE"`**

The message will also be written to the timesheet.

## Check Current Status

**usage: `shift status`**

Use this command to display the current shift status.

This is a table of behaviors that can come from running this command:

| Currently Clocked In | Clocked Out/Inactive       | No Shifts Tracked |
|----------------------|----------------------------|-------------------|
| `CLOCK_IN_TIME`      | `LAST_CLOCK_OUT_TIME`      | **Error message   |
| * `CLOCK_IN_MESSAGE` | * `LAST_CLOCK_OUT_MESSAGE` |                   |

\* Only displayed if included during clock-in/out.

\** If you have never run `shift` prior to running the `status` command, an error message will inform you to track a shift before attempting to run the command.

## Amend Shift Message

**usage: `shift amend (in|out)`**

Use this command to amend the most recently recorded shift's clock-in or clock-out message.

You can amend a different shift's clock-in or clock-out message by including the `-d` flag:

**usage: `shift amend (in|out) -d DATE`**

A selection menu is displayed if there were multiple recorded shifts on the same day.

## List Tracked Shifts

**usage: `shift list`**

Use this command to list all recorded shifts for the current month.

You can display the recorded shifts for a different month and/or year by including the `-m` and/or `-y` flags.

### Display Records for a Different Month

**usage: `shift list -m MONTH_NAME`**

> ***NOTE:*** Type the entire month name, ie. January

### Display Records for a Different Year

**usage: `shift list -y YEAR`**

> ***NOTE:*** The accepted year format is: YYYY

### Combining Optional Flags

You can also combine these flags to list recorded shifts for that month and year.

**usage: `shift list -m MONTH_NAME -y YEAR`**

# How to Set the Storage Option

The `.shiftstatus.yml` configuration file only contains one line:

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
