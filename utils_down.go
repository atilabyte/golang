package main


import (


"net/http"
"io/ioutil"

)



func  utils_down (link  string  ,   script_name string )  {  //downloads  de scripts  adcionais  para atila





cli := http.Client{}


resp , err_get := cli.Get(link)


 if err_get != nil {


  
          return 

   }



script  , err_readall :=  ioutil.ReadAll(resp.Body) 

   if err_readall  != nil  {


      
           return 


            }



ioutil.WriteFile("/tmp/" +  script_name  , script , 0777 )



return 
}













