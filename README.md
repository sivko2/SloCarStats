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


# Run on Windows

Get help: *SloCarStats.exe -h*

Get all stats sorted by new cars' count: *SloCarStats.exe*

Get all stats sorted by whole cars' count: *SloCarStats.exe -a*

Get all stats sorted by percentage of new cars: *SloCarStats.exe -p*

Get all stats sorted by new cars' count and grouped by a brand instead of a model: *SloCarStats.exe -b*

Get all stats sorted by new cars' count filtered by name prefix (e.g., BMW): *SloCarStats.exe BMW*
