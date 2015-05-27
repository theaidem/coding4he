# Coding4he

## Installation

```
git clone
cd  coding4he && go build
```

## Usage

Start server:

```
./coding4he
```

You can also provide connection flags:

```
./coding4he -tcp_port=1313 -http_port=8008
```

For send some arbitrary text data, you can use ``nc`` utility:

```
nc localhost 5555 < test.txt
```

OR

```
echo -n "test out the server" | nc localhost 5555
```

Then you would see [stats](http://localhost:8080/stats)

## TODO

* Refactor

* Make tests