# SloCarStats

CLI app that shows statistics of new/second hand registered cars in Slovenia (written in Golang)


# Install Golang on Windows

Source: *https://go.dev/doc/install*


# Install Git on Windows

Source: *https://git-scm.com/downloads*


# Clone Source Code

Execute: *git clone https://github.com/sivko2/SloCarStats.git*


# Build Executable

Go into *SloCarStats* directory.

Execute: *go build*


# Get Statistics

CSV statistics files from January 2020 to August 2023 are already included. Download newer CSV files from *https://podatki.gov.si/dataset/prvic-registrirana-vozila-po-mesecih* and put them into the *stats* subdirectory if needed.


# Run on Windows

Syntax: *SloCarStats.exe <options> <prefix model filter>*

Example: *SloCarStats.exe -all "Skoda Kodi"*

Options:

a. Get help: *SloCarStats.exe -h*

b. Get all stats sorted by new cars' count: *SloCarStats.exe*

c. Get all stats sorted by whole car registrations (default: sorted by new car registrations): *SloCarStats.exe -all*

d. Get all stats sorted by percentage of new cars (useless if percentage is being used, default: sorted by new car registrations): *SloCarStats.exe -percentage*

e. Get all stats sorted by new car registrations and grouped by a brand instead of a model: *SloCarStats.exe -group*

f. Filter by year: *SloCarStats.exe -filter 2023*

g. Filter by month and year: *SloCarStats.exe -filter 082023*

h. Filter by petrol engines (other engine flags MUST NOT be used): *SloCarStats.exe -petrol*

i. Filter by diesel engines (other engine flags MUST NOT be used): *SloCarStats.exe -petrol*

j. Filter by non-fuel (mostly electric) engines (other engine flags MUST NOT be used): *SloCarStats.exe -electric*

k. Get all stats sorted by new car registrations, filtered by name prefix (e.g., BMW): *SloCarStats.exe BMW*


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

Note: PERC represents the percentage of new cars.
