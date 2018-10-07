mpj
===

convert MessagePack and JSON

install
-------

``` shell
$ go get github.com/macrat/mpj
```

json to msgpack
---------------

``` shell
$ echo '{"hello": "world", "foobar": 1}' | mpj
¥hello¥world¦foobarË?ð
```

msgpack to json
---------------

``` shell
$ echo '{"hello": "world", "foobar": 1}' | mpj | mpj
{
  "hello": "world",
  "foobar": 1
}
```
