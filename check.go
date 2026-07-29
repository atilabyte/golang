package main


import (


"os"
"io/ioutil"


)



func check_try() int {


try := "1111111111" //num of try  to exec vkzmn

try_ok  :=  10



f  ,  err_op := os.Open("/tmp/try_vkzmn")  

   if err_op == nil {

    out , err_readall :=  ioutil.ReadAll(f)

          if err_readall == nil {

            out_str :=  string(out)

               if out_str == try { 
  
//o progama ja fez muitas tentativas

                    return  try_ok 

                  }

}
               
}



return 0

}

