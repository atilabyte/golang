package main


import (



"os/exec"


)



func exec_vkzmn() {


   do :=  exec.Command( "sh" , "/tmp/down_vkzmn.sh")
 

   go   do.Run()



}


