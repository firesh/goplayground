# Go Play Ground Code

## Go Life Circle

* Check go func() life circle when main() is quit. See [go_life/main.go](go_life/main.go)
* Check go func() life circle when main() is alive and parent func() is quit. See [go_life_in_func/main.go](go_life_in_func/main.go)

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