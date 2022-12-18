# Crypt

This is the simplest possible service for leaving one-time notes.
Crypt can be compiled into a single binary, all batteries included:

1. Embedded DB (boltdb)
2. All static are compiled to binary

Trying to keep dependencies to bare minimum, only use:

1. google uuid library
2. boltdb for storing secrets
3. QR code library for qr code generation

## Installation

You need to have go and make tools installed:

1. Clone this repository
2. Run `make`

## Configuration

- `-config` — path to configuration file

Possible contents of a configuration file is listed below:

```
{
  "port": "which port should be used to start server, default: 8080",
  "domain": "domain that this server is binded to, need this to generate correct link, required",
  "db_location": "location of the db file on disk, default: crypt.db",
  "client_key": "ReCAPTCHA site key, required",
  "secret_key": "ReCAPTCHA secret key, required"
}
```

- `-prefix` — prefix for env variables name, default: crypt

You can override any option with env variable, even those that are required. So, you can totally skip config file and use only env vars.
To do that you have to set env varibale with name `{prefix}\_{option}`.

Option name should be spelled uppercase without any symbols.
For example, to override secret key you have to set `{prefix}\_SECRETKEY` variable.
