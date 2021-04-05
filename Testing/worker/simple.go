package main // Tell go runtime your main lives here

import ("fmt") // We need the fmt package to print to the stdout


var done = make(chan bool)
var msgs = make(chan int)


func main () {
   go produce()
   go consume()
   <- done
}


func produce() {
	i := 0
    for{
		i
        msgs <- i
    }
    done <- true
}


func consume() {
    for {
      msg := <-msgs
      fmt.Println(msg)
   }
}