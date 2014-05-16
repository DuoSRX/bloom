# Bloom

Bloom is a Bloom Filter implemented in Go. While it is usable, it's meant to be a learning tool.

## Usage

```go
bloom := bloom.New(1000)
bloom.add("foo")
bloom.IsPresent("foo") // => true
bloom.IsPresent("bar") // => false
```
