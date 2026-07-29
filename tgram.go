package main

import (
	"bytes"
	"encoding/json"
	"net/http"
        "fmt"
       "time"
      


)



var token string = "8788643517:AAH8qGbylRhwyXLRwQW6pXx8ufFExli6TuE"



func bot(msg string ) {






      time.Sleep  ( 60   *  time.Second)  //evite    spam 



	cli := http.Client{}


       



	str := map[string]string{

		"chat_id": "7127446120",

		"text":    "teste" , //msg 


	}




	jm, _ := json.Marshal(str)

	data := bytes.NewBufferString(string(jm))

	req, _ := http.NewRequest("POST", "https://api.telegram.org/bot"+token+"/sendMessage", data)

	req.Header.Set("Content-Type", "application/json")

	r, _ := cli.Do(req)

       

  
       if  r.StatusCode  ==    200 { fmt.Println( " send ok"  ) }




}
