         _   _ ___ _   
     ___| |_|_|  _| |_ 
    |_ -|   | |  _|  _|
    |___|_|_|_|_| |_|

[![Go Badge](https://img.shields.io/badge/-lang-blue?logo=go)][Go]
![Lines of code](https://img.shields.io/tokei/lines/github/JosephLai241/shift)
![License Badge](https://img.shields.io/github/license/JosephLai241/shift)

# Table of Contents

* [Introduction](#introduction)
    + [Compile `shift` From Source](#compile-shift-from-source)
        * [Verify/Run the Executable](#verifyrun-the-executable)
    + [Or Download a Binary](#or-download-a-binary)
* [Installation](#installation)
* [How `shift` Works](#how-shift-works)
    + [What Kind of Data Is Stored?](#what-kind-of-data-is-stored)
    + [Storing Data Into Timesheets](#storing-data-into-timesheets)
    + [Storing Data Into a SQLite Instance](#storing-data-into-a-sqlite-instance)
* [How To Use `shift`](#how-to-use-shift)
    + [Clocking In](#clocking-in)
    + [Clocking Out](#clocking-out)
    + [Check Current Status](#check-current-status)
        * [If `shift` Is Active](#if-shift-is-active)
        * [If `shift` Is Inactive](#if-shift-is-inactive)
    + [Amend Shift Message](#amend-shift-message)
    + [List Tracked Shifts](#list-tracked-shifts)
    + [Delete a Shift](#delete-a-shift)
* [Storage/How to Set the Storage Method](#storagehow-to-set-the-storage-method)
    + [Display Current Storage Method](#display-current-storage-method)
    + [Set New Storage Method](#set-new-storage-method)
* [Multiple "Instances" of `shift`](#multiple-instances-of-shift)

# Introduction

`shift` is a command-line tool for keeping track of shifts. Its primary audience is contractors (like myself) or remote workers who need to track their own hours, but it is still useful for logging time spent doing anything.

This project is inspired by Luke Schenk's Python CLI tool [`clck`][clck].

# Installation

## Compile `shift` From Source

You will need [Go][Go] installed on your computer to compile the source code.

```
git clone --depth=1 git@github.com:JosephLai241/shift.git
cd shift/
go build
```

The executable file `shift` (on Linux or Mac) or `shift.exe` (on Windows) is created in the `shift/` directory once compiling is done. 

### Verify/Run the Executable

Run the executable to verify `shift` compiled successfully.

On Linux or Mac:

```
./shift version
```

On Windows:

```
shift.exe version
```

## Or Download a Binary

If you do not want to compile `shift`, you can also download a binary attached to each release in the [Releases][Releases] section.

> ***NOTE:*** This program initializes and reads from files in your current working directory. Run `shift` in a directory in which you would like all your records and program-related files to be stored.

# How `shift` Works

This tool is capable of storing shift data into CSV spreadsheets or a local SQLite instance. **The default is CSV spreadsheets**. See the [Storage/How to Set the Storage Method](#storagehow-to-set-the-storage-method) section for information on how to configure this.

## What Kind of Data Is Stored?

The following data is recorded by `shift`:

* Date
* Day of the Week
* Clock-in Time
* Clock-in Message
* Clock-out Time
* Clock-out Message
* Shift Duration

## Storing Data Into Timesheets

The directory `shifts/CURRENT_YEAR` is created within the current working directory. Data is then recorded into a CSV-format timesheet titled `CURRENT_MONTH.csv`, located in the `CURRENT_YEAR` directory.

This is an example of the `shifts` directory structure if you ran `shift` sometime during July 2021:

```
shifts/
????????? 2021
    ????????? July.csv
```

## Storing Data Into a SQLite Instance

A SQLite instance `shifts.db` is created within the current working directory. A main `YEAR` table is created, containing the years that `shift` was run. The data is then stored in a table labeled with the current year, which is then linked to a sub-table labeled with the current month.

This is an example of the database's relationships if you ran `shift` sometime during July 2021:

```
shifts.db
????????? TABLE `YEAR`
    ????????? TABLE `Y_2021`
        ????????? TABLE `M_July`
```

# How To Use `shift`

## Clocking In

![Clock-In Demo][Clock-In Demo]

```
shift in

    [-m <"OPTIONAL MESSAGE">]
```

Use this command to clock-in. The clock-in time is then written to the timesheet or database.

You can record a message corresponding to your clock in by including the `-m` flag.

The status of your current shift is displayed if you attempt to run this command when already clocked in.

## Clocking Out

![Clock-Out Demo][Clock-Out Demo]

```
shift out

    [-m <"OPTIONAL MESSAGE">]
```

Use this command to clock-out. The clock-in time is then written to the timesheet or database.

You can record a message corresponding to your clock in by including the `-m` flag.

Your clock-in time and message as well as your shift duration is displayed.

A warning is displayed if you attempt to run this command when already clocked out.

## Check Current Status

```
shift status
```

Use this command to display the current shift status.

### If `shift` Is Active

![Active Status Demo][Active Status Demo]

### If `shift` Is Inactive

![Inactive Status Demo][Inactive Status Demo]

> ***NOTE:*** If you have never run `shift` prior to running the `status` command, an error message will inform you to track a shift before attempting to run the command.

## Amend Shift Message

![Amend Demo][Amend Demo]

```
shift amend (in|out) "YOUR NEW MESSAGE"

    [-d <DATE_or_DAY_OF_THE_WEEK>]
    [-m <MONTH>]
    [-y <YEAR>]
```

Use this command to amend the most recent shift's clock-in or clock-out message.

If used without any optional flags, `shift` will target shifts recorded on the current day.

You can search for recorded shifts on a different day, month, and/or year by including the `-d`, `-m`, and/or `-y` flags. Combine these flags to narrow your search.

> ***NOTE:*** The `-d` flag accepts either a day of the week or a date. Type the entire day of the week, ie. Monday, or provide the date using the MM-DD-YYYY or MM/DD/YYYY format.

> ***NOTE:*** Type the entire month name when using the `-m` flag, ie. January.

> ***NOTE:*** Type the entire year in YYYY format when using the `-y` flag, ie. 2021.

## List Tracked Shifts

![List Demo][List Demo]

```
shift list [all]

    [-d <DATE_or_DAY_OF_THE_WEEK>]
    [-m <MONTH>]
    [-y <YEAR>]
```

Use this command to display recorded shifts.

If used without the optional positional argument `all`, `shift` will display all shifts recorded on the current day.

You can list all recorded shifts within the current month by including the `all` argument.

You can display the recorded shifts for a different day, month, and/or year by including the `-d`, `-m`, and/or `-y` flags. Combine these flags in addition to `all` to narrow your search.

> ***NOTE:*** The `-d` flag accepts either a day of the week or a date. Type the entire day of the week, ie. Monday, or provide the date using the MM-DD-YYYY or MM/DD/YYYY format.

> ***NOTE:*** Type the entire month name when using the `-m` flag, ie. January.

> ***NOTE:*** Type the entire year in YYYY format when using the `-y` flag, ie. 2021.

## Delete a Shift

![Delete Demo][Delete Demo]

```
shift delete

    [-d <DATE_or_DAY_OF_THE_WEEK>]
    [-m <MONTH>]
    [-y <YEAR>]
```

Use this command to delete a recorded shift.

If used without any optional flags, `shift` will target shifts recorded on the current day.

You can search for recorded shifts on a different day, month, and/or year by including the `-d`, `-m`, and/or `-y` flags. Combine these flags to narrow your search.

> ***NOTE:*** The `-d` flag accepts either a day of the week or a date. Type the entire day of the week, ie. Monday, or provide the date using the MM-DD-YYYY or MM/DD/YYYY format.

> ***NOTE:*** Type the entire month name when using the `-m` flag, ie. January.

> ***NOTE:*** Type the entire year in YYYY format when using the `-y` flag, ie. 2021.

# Storage/How to Set the Storage Method

```
shift storage 

    [set (timesheet|database)]
```

## Display Current Storage Method

![Storage Demo][Storage Demo]

You can check the current storage method by using the `storage` command without additional sub-commands.

## Set New Storage Method

![Set Storage Demo][Set Storage Demo]

You can change the storage method by including the `set` sub-command and providing `timesheet` or `database` as its value to switch from one to another.

> ***NOTE***: You cannot change storage options while clocked in. `shift` will throw an error if you attempt to do so.

# Multiple "Instances" of `shift`

You may want to use multiple "instances" of `shift`. For example, if you are working two different jobs, or want to track time spent on personal projects in addition to your day job.

It is quite a simple solution - just copy the `shift` executable into different directories like so:

```
day_job/
????????? shift <--- executable

URS/
????????? shift <--- executable
```

You are now able to track the time you spent doing different things :thumbsup:

<!-- LINKS -->
[Go]: https://golang.org/
[clck]: https://github.com/LukeDSchenk/clck
[Releases]: https://github.com/JosephLai241/shift/releases

<!-- DEMO LINKS -->
[Active Status Demo]: https://github.com/JosephLai241/shift/blob/demo/screenshots/active_status.png
[Amend Demo]: https://github.com/JosephLai241/shift/blob/demo/screenshots/amend.png
[Clock-In Demo]: https://github.com/JosephLai241/shift/blob/demo/screenshots/in.png
[Clock-Out Demo]: https://github.com/JosephLai241/shift/blob/demo/screenshots/out.png
[Delete Demo]: https://github.com/JosephLai241/shift/blob/demo/screenshots/delete.png
[Inactive Status Demo]: https://github.com/JosephLai241/shift/blob/demo/screenshots/inactive_status.png
[List Demo]: https://github.com/JosephLai241/shift/blob/demo/screenshots/list.png
[Storage Demo]: https://github.com/JosephLai241/shift/blob/demo/screenshots/storage.png
[Set Storage Demo]: https://github.com/JosephLai241/shift/blob/demo/screenshots/set_storage.png
