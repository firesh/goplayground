# Go Play Ground Code

## Go Life Circle

* Check go func() life circle when main() quit. See [go_life/main.go](go_life/main.go)
* Check go func() life circle when main() is running and parent func() quit. See [go_life_in_func/main.go](go_life_in_func/main.go)

### Result
```
MacBook-Pro-Wang:go_life wang$ go run main.go
main start
go func start
count: 0
count: 1
count: 2
count: 3
main quit

MacBook-Pro-Wang:go_life_in_func wang$ go run main.go
main start
go func1 start
go func2 start
count: 0
count: 1
count: 2
count: 3
go func1 quit
count: 4
count: 5
count: 6
count: 7
count: 8
count: 9
count: 10
go func2 quit
main quit
```

### Conclution
* When func main() quits, go func() will be termiated and deffer will not be involved.
* go func() will continue when parent go func() quits.

## SHA

### SHA encode `""` message
There are several SHA functions. Codes compare the `""` encoding message. We can identify those functions by the encoded message.
> SHA-1 is not secure enough for application any more.
 
### Result
* HA-1: 0xda39a3ee5e6b4b0d3255bfef95601890afd80709
* SHA-2 SHA-256: 0xe3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
* SHA-3 SHA3-256: 0xa7ffc6f8bf1ed76651c14756a061d662f580ff4de43b49fa82d80a4b80f8434a
* SHA-3 Keccak256: 0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470