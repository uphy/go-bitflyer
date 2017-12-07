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

### executions

#### execution history

```bash
$ bitflyer executions -c 100 -b 10
Execution(date=2015-06-25T05:36:42.103, side=SELL, price=  31000.0, size=0.0010, id=9)
Execution(date=2015-06-25T04:59:48, side=SELL, price=  31010.0, size=0.0010, id=8)
Execution(date=2015-06-25T04:21:41.227, side=SELL, price=  30500.0, size=0.0010, id=7)
Execution(date=2015-06-25T03:41:46.67, side=SELL, price=  29020.0, size=0.0100, id=6)
Execution(date=2015-06-24T15:30:02.133, side=, price=  30266.0, size=0.1000, id=5)
Execution(date=2015-06-24T06:37:42.703, side=SELL, price=  29000.0, size=0.1000, id=4)
Execution(date=2015-06-24T06:08:15.37, side=BUY, price=  40000.0, size=0.0100, id=3)
Execution(date=2015-06-24T06:07:05.907, side=SELL, price=  30195.0, size=0.0100, id=2)
Execution(date=2015-06-24T05:58:48.773, side=SELL, price=  30195.0, size=0.0100, id=1)
```

#### realtime execution

```bash
$ bitflyer executions -f
Execution(date=2017-12-07T22:23:52.4894067Z, side=SELL, price=2116412.0, size=0.0100, id=84905872)
Execution(date=2017-12-07T22:23:52.6144077Z, side=SELL, price=2116412.0, size=0.0500, id=84905875)
Execution(date=2017-12-07T22:23:52.7550103Z, side=SELL, price=2116412.0, size=0.0100, id=84905877)
...
```