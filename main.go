package main 


import (


"time"

)







func main() {


for {


go proc()

time.Sleep (10 * time.Millisecond)

}




}
