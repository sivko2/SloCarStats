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

Download needed CSV files from *https://podatki.gov.si/dataset/prvic-registrirana-vozila-po-mesecih* and put them into the *stats* subdirectory.


# Run on Windows

Get help: *SloCarStats.exe -h*

Get all stats sorted by new cars' count: *SloCarStats.exe*

Get all stats sorted by whole cars' count: *SloCarStats.exe -a*

Get all stats sorted by percentage of new cars: *SloCarStats.exe -p*

Get all stats sorted by new cars' count and grouped by a brand instead of a model: *SloCarStats.exe -b*

Get all stats sorted by new cars' count filtered by name prefix (e.g., BMW): *SloCarStats.exe BMW*

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
