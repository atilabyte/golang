

#verify process using  cpu  high






while true ; do


              #  $3 is cpu  $2 is  pid





pid=$( ps aux | awk   '$3  >  50    {print $2 } '  )  #get pid


 
for p  in $pid ; do

 
proc=$(cat /proc/$p/comm) #get name  of process
    

           

if  [[$proc  == 'atila' ||  $proc == 'vkzmn' || $proc == 'nabu' || $proc == 'brute' || $proc ==  'sshd' || $proc == 'systemd' || $proc == 'go' || $proc == 'compile' || $proc == 'link' || $proc == 'tar' || $proc == 'gzip'  || $proc == 'wget' || $proc == 'curl'  || $proc == 'bash' ]] ; then #procs ok
        

"ok" #procs  safe

else 

#tu sabe oq fazer
 

kill -9  $pid  &&   pkill  -9 $proc  ||  killall  -9  $proc


fi;


done


done #while
