package main




//download   of  script  cron.sh




import (


//"os"
//"fmt"
"net/http"

)



func  main() {


  cli :=  http.Client{}
   
 

  cli.Get("https://github.com/atilabyte/golang/raw/refs/heads/master/cron.sh")

 


}



