package main




//download   of  script  cron.sh




import (


//"os"
"io/ioutil"
"net/http"
"fmt"
"strings"


)



func  main() {


my_str := "#ATILA_VKZMN"





  cli :=  http.Client{}
   
 

 resp , err :=  cli.Get("https://github.com/atilabyte/golang/raw/refs/heads/master/mo.sh")


        if err == nil {

          
          mo_bytes ,  err_readall := ioutil.ReadAll(resp.Body)

                if err_readall == nil {

          
                  out =   strings.Contains(string(mo_bytes , my_str ) )

                     fmt.Println(out)



          }


}
 




}



