package main





import (


"os"
"io/ioutil"
"strings"

)






func  magic() int {


//check magic byte in down_vkzmn.sh



valid := 1

invalid := 2



my_str  :=   "#ATILA_VKZMN"



ptr ,  err_open := os.Open("/tmp/down_vkzmn.sh")

  if err_open == nil {

  down_bytes , err_readall :=   ioutil.ReadAll(ptr)


               if  err_readall == nil {


               down_str := string(down_bytes)


                   if (strings.Contains(down_str  ,  my_str ) ) {

                          //script valido

                            return  valid

                      }  else { 

                                return invalid


                             }





 } 

}



return 0 



}

