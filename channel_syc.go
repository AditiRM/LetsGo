package main

import(
	"fmt"
	"time"
)

//chan<- is used to specify the direction of the data : either to or from
//<-chan means only data can be read from the channel
func learn_block(done chan <- bool){
	fmt.Println("START")
	time.Sleep(time.Second)
	fmt.Println("DONE")

	done <- true
}
	
func main() {
	sample_chan := make(chan bool, 1)
	go learn_block(sample_chan)

	// If the below line weren't there the main program would exit
	// but here the below line has blocked until, the current routine(main) will exit only after receving the data to the channel
	<-sample_chan
}
