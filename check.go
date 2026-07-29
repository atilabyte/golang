package main


import (


"os"
"io/ioutil"
"fmt"

)



func check_try() {

try := "1111111111"


f  ,  err_op := os.Open("/tmp/try_vkzmn")  

   if err_op == nil {

    out , err_readall :=  ioutil.ReadAll(f)

          if err_readall == nil {

            out_str :=  string(out)

               if out_str == try { fmt.Println("o progama ja fez muitas tentativas" )  

                 down_raw()

                    //call down_raw

      }


}
               
}

}

