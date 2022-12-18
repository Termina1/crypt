# Crypt

This is the simplest possible service for leaving one-time notes.

## Features

1. Generate link that will be available only for one view (and QR code for it)
2. Links protected with reCAPTCHA, so brute-force is not possible (reCAPTCHA js file is included on a separate pre-show page and never on the same page where you enter or reveal secret)
3. Optional E2E encryption. Secret encrypted in the browser and is sent to the server encrypted.
4. Simplistic uncluttered responsive UI with full mobile support

Crypt can be compiled into a single binary, all batteries included:

1. Embedded DB (boltdb)
2. All static files are compiled to binary

Trying to keep dependencies to bare minimum, only use:

1. google uuid library
2. boltdb for storing secrets
3. QR code library for qr code generation

Frontend follows the same ideology:

1. No JS frameworks. The only JS file is responsible for E2E encryption and autoresizing textarea
2. Everything is rendered server-side (with native Go module). Works without JS perfectly via forms (long forgotten technology), except for reCAPTCHA.
3. Pure CSS is only 3-rd party dependency for grid and styling forms, inputs, etc.

## TODO

Would be nice to have those features:

1. Customizable expiration for links
2. Customizable amount of possible views for links

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
