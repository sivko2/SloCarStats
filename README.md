# SloCarStats - Slovenian Car Statistics 2020-2023

This is a command-line application that shows statistics of new and second hand registered cars in Slovenia. It is written in Golang.


# Installing Golang on Windows

Source: *https://go.dev/doc/install*


# Installing Git on Windows

Source: *https://git-scm.com/downloads*


# Cloning Source Code

Open a command prompt, create a project directory, go into it and execute: *git clone https://github.com/sivko2/SloCarStats.git*


# Building Executable

Go into *SloCarStats* directory.

Execute: *go build*


# Getting Statistics Files

CSV statistics files from January 2020 to now are already included. Download newer CSV files from *https://podatki.gov.si/dataset/prvic-registrirana-vozila-po-mesecih* and put them into the *stats* subdirectory when needed.


# Running on Windows

A syntax: *SloCarStats.exe <options> <prefix for brand/model filtering>*

An example to show statistics for all Skoda Kodiaqs since 1st Jan 2020: *SloCarStats.exe -all "Skoda Kodiaq"*

Possible options:

a. Get help: *SloCarStats.exe -h*

b. Get all stats sorted by new car registrations count: *SloCarStats.exe*

c. Get all stats sorted by whole car registrations (new + second-hand together, ignored if percentage is being used, default is sorted by new car registrations): *SloCarStats.exe -all*

d. Get all stats sorted by percentage of new car registrations (default: sorted by new car registrations): *SloCarStats.exe -percentage*

e. Get all stats sorted by new car registrations and grouped by brands: *SloCarStats.exe -group*

f. Filter by year (e.g., 2023): *SloCarStats.exe -filter 2023*

g. Filter by month and year (e.g., Aug 2023): *SloCarStats.exe -filter 082023*

h. Filter by petrol engines (other engine flags MUST NOT be used): *SloCarStats.exe -petrol*

i. Filter by diesel engines (other engine flags MUST NOT be used): *SloCarStats.exe -diesel*

j. Filter by non-fuel (electric mostly) engines (other engine flags MUST NOT be used): *SloCarStats.exe -electric*

k. Get all stats sorted by new car registrations, filtered by name prefix - case insensitive (e.g., BMW X): *SloCarStats.exe BMW X*

l. Get first N (e.g., 10) models/brands registrations: *SloCarStats.exe -count 10*

Full syntax description: *SloCarStats.exe [-all | -percentage] [-group] [-filter <MMYYYY> | -filter <YYYY>] [-petrol | -diesel | -electric] [-count <N>]*


# Example of Output

```
+------+----------------------------------------------------+-------+-------+-------+------+
| #    | BRAND AND MODEL                                    | NEW   | OLD   | SUM   | PERC |
+------+----------------------------------------------------+-------+-------+-------+------+
|    1 | RENAULT CLIO                                       |  1055 |   275 |  1330 |  79% |
|    2 | SKODA OCTAVIA                                      |   998 |   841 |  1839 |  54% |
|    3 | VOLKSWAGEN T-ROC                                   |   859 |   117 |   976 |  88% |
|    4 | RENAULT CAPTUR                                     |   857 |   120 |   977 |  87% |
...
| 1590 | VOLKSWAGEN LT 31                                   |     0 |     1 |     1 |   0% |
| 1591 | VOLVO V60 T6 TWIN ENGINE                           |     0 |     1 |     1 |   0% |
+------+----------------------------------------------------+-------+-------+-------+------+
|      | SUM                                                | 34474 | 25884 | 60358 |      |
+------+----------------------------------------------------+-------+-------+-------+------+
```

Note: PERC represents the percentage of new car registrations compared to all car registrations.
