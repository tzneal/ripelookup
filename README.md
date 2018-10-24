Overview
========

ripelookup is used to lookup IP ownership information for an IP address. It
uses IANA registration information to determine which server to query for more
information.

Generating Tables
=================

The v4/v6 tables files are generated from IANA data and allow looking up the
appropriate server for a particular IP address. To regenerate the table, run:

```sh
go run gen/gen.go
```

Usage
=====
```go
        ip := "8.8.8.8"
        server, _ := ripelookup.DetermineServer(net.ParseIP(ip))
        records, _ := ripelookup.WhoisIP(ip, server)
        fmt.Println(records[0])
```
