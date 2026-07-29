package main


import (


"os"
"io/ioutil"
"fmt"
"strings"
)



func check_try() string {



try := "11111" //num of try  to exec vkzmn
 
 
try_ok  :=  "try_ok"



f  ,  err_op := os.Open("/tmp/try_vkzmn")  

   if err_op == nil {

    out , err_readall :=  ioutil.ReadAll(f)

          if err_readall == nil {

            out_str :=  string(out)

          

               if strings.Contains(out_str , try ) { 

  
                    fmt.Println("numero de tentativa ok" )
                       
                        return try_ok


                  }

}
               
}

return ""

}

