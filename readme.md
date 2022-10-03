# Exercise 1 - Quiz Game via CSV

_Credit to [Gophercises](https://gophercises.com/)!_


> _Gophercises is an excellent collection of exercises for learning 
the Go programming language. They are provided FREE of charge, and you only have to provide
your email address. I think that's a pretty fair trade. I am publishing these exercise 
solutions here as part of my own learning, and **also** in case they can help someone else. 
Please don't read ahead without trying the exercises first!_

[https://gophercises.com/](https://gophercises.com/)

---

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

---

## How to Run

- `go mod init exercises/ex1`
- `go run .`



## Questions, Research

A list of some of the reference items I encountered while writing this. In no
particular order:

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

### Read Console Input

- Ref: [https://tutorialedge.net/golang/reading-console-input-golang/](https://tutorialedge.net/golang/reading-console-input-golang/)
- Ref: [https://freshman.tech/snippets/go/read-console-input/](https://freshman.tech/snippets/go/read-console-input/)

```go
reader := bufio.NewReader(os.Stdin)
text, _ := reader.ReadString('\n')
// Removing line ending char(s). Windows, CRLF (\r\n). Unix-based, CR (\r).
text = strings.Replace(text, "\r\n", "", -1)
```


### CSV Resources

- csv package: [https://pkg.go.dev/encoding/csv](https://pkg.go.dev/encoding/csv)
- CSV tutorial: [https://medium.com/@barunthapa/working-with-csv-in-go-50a4f540e623](https://medium.com/@barunthapa/working-with-csv-in-go-50a4f540e623)


### Print int in string

Ref: [https://golang.cafe/blog/golang-int-to-string-conversion-example.html](https://golang.cafe/blog/golang-int-to-string-conversion-example.html)

```go
input = "test"
fmt.Printf("length: %v", len(input))
```

### Time Duration

Ref: [https://pkg.go.dev/time#Duration](https://pkg.go.dev/time#Duration)

### Exiting a Program

Ref: [https://gobyexample.com/exit](https://gobyexample.com/exit)