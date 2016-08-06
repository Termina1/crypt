# Potemkin

This is the simplest possible service for leaving one-time notes.
This can be compiled in single binary, all batteries included:

1. Embedded DB (boltdb)
2. All static files can be compiled to binary

## Installation

You need to have go and make tools installed:

1. Clone this repository
2. run ```make```


## Configuration

- ```-port``` — which port should be used to start server
- ```-domain``` — domain that this server is binded to, need this to generate correct link
