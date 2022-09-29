# Exercise 1 - Quiz Program via CSV

Ref: [https://courses.calhoun.io/lessons/les_goph_01](https://courses.calhoun.io/lessons/les_goph_01)

Create a quiz using the provided `.csv` file. The file includes arithmetic questions 
and answers. The basic premise is we want to print each question, one at a time, and 
prompt the user for the answer.

Pt. 1 - TO DO: 

- CSV file should default to `problems.csv`
  - Be able to customize filename via a flag
- Keep track of score: _how many questions are answered correctly_

Pt. 2 - TO DO:

- Add a timer
  - Default: 30 seconds. Customize via a flag.
- Quiz should stop when timer meets "0 sec".
  - _Note: Quiz needs to exit when timer hits "0"; **not wait for user to answer question!**_
- Prompt user to hit "Enter" before starting timer


## How to Run

- `go mod init exercises/ex1`
- `go run .`



## Questions, Research

A listing of some of the reference items I encountered while 
writing this.

### How to find length of array

- Ref: [https://reactgo.com/go-array-length/](https://reactgo.com/go-array-length/)

```go
records = [...data here]
length = len(records)
```

### Writing a for-loop

- Ref: [https://yourbasic.org/golang/for-loop/](https://yourbasic.org/golang/for-loop/)

```go
for i := 0; i < 9; i++ {
    fmt.Println(i) // print the counter
}
```


### CSV Resources

- csv package: [https://pkg.go.dev/encoding/csv](https://pkg.go.dev/encoding/csv)


### Print int in string

Ref: [https://golang.cafe/blog/golang-int-to-string-conversion-example.html](https://golang.cafe/blog/golang-int-to-string-conversion-example.html)

```go
input = "test"
fmt.Printf("length: %v", len(input))
```