# Potemkin

This is the simplest possible service for leaving one-time notes.
This can be compiled into a single binary, all batteries included:

1. Embedded DB (boltdb)
2. All static files can be compiled to binary

## Installation

You need to have go and make tools installed:

1. Clone this repository
2. Run ```make```


## Configuration

- ```-config``` — path to configuration file

Possible contents of a configuration file is listed below:

```
{
  "port": "which port should be used to start server",
  "domain": "domain that this server is binded to, need this to generate correct link",
  "dbLocation": "location of the db file on disk",
  "clientKey": "ReCAPTCHA site key",
  "secretKey": "ReCAPTCHA secret key"
}
```

Every config parameter can be redefined with setting environment variable thanks to gonfig.
