# Signer service
Signer service should have restricted access. We suggest to don't use CD for it in production basing on security reason.
## Sign generation:
Signer service stores encrypted private keys from eth validation address and from pokt bridge wallet adress to sign transactions.
One part of encryption passphrase should be store in signer service env variables, other part should be stored at chain listiners.
For generating encoded private key, you shoud have golang installed, or create build of `/cmd/generator/main.go`
if you have golang installed just run `/cmd/generator/main.go` with decrypted pk as an argument.
For example let encrypt eth pk `0000000000000000000000000000000000000000000000000000000000000000`
```
go run /cmd/generator/main.go 0000000000000000000000000000000000000000000000000000000000000000
```
if you have compiled app run 
```
go run ./generator 0000000000000000000000000000000000000000000000000000000000000000
```
you should see the next result
```
Encoded pk: cPBiOu1g/zycSauVMQ3xcILElZarRvCZ7oVI+JupsEbdjIzvcpHN8waUsThtDUaeZyH1Xay5PogjA80OdaJjM7U9lj9O51KPkdaqmBEs1rM5k/ny1TlYqZamV6c=
Signer passcode: zX*KZX672yUG2&GavTSGHU7AXiSE3RZ&
Bridge passcode: E3fvkPtMCtA$1etxA9&w4v$6F5v7B_NU
```

use _Encoded pk_ as _SIGNER_ETH_PK_ at signer service
use _Signer passcode_ as _SIGNER_ETH_PASS_ at signer service
use _Bridge passcode_ as signer pass env for pokt and eth chain services

## Building process:
Use `docker-compose.yml` or `Dockerfile` to build application.

## Dependecies updates:
Evm chain service uses proto lib from protocol repository. To update proto version use `make update_chain_client` and `make update_sign_client` command

## App Env description:
Examples are in `.env.example` file
- _SIGNER_POKT_PK_ - encrypted pokt wallet private key
- _SIGNER_POKT_PASS_ - pokt key passphrase
- _SIGNER_ETH_PK_ - encrypted eth validation address private key
- _SIGNER_ETH_PASS_ - eth key passphrase
- _SIGNER_PORT_ - signer service grpc address
