# go-bitflyer

## Install

```bash
go get -u github.com/uphy/go-bitflyer/bitflyer
```

## Commands

### markets

```bash
$ bitflyer markets
+-----------------+---------------+
|  PRODUCT CODE   |     ALIAS     |
+-----------------+---------------+
| BTC_JPY         |               |
| FX_BTC_JPY      |               |
| ETH_BTC         |               |
| BCH_BTC         |               |
| BTCJPY01DEC2017 | BTCJPY_MAT1WK |
| BTCJPY08DEC2017 | BTCJPY_MAT2WK |
+-----------------+---------------+
```

### ticker

```bash
$ bitflyer ticker --productcode FX_BTC_JPY
Product Code: FX_BTC_JPY
LTP         : 1390000.000
```
