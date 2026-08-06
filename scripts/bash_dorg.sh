


install_url='https://github.com/atilabyte/golang/raw/refs/heads/master/scripts/install.sh'




#script usado para  fica monitorando  o atila
 
#dorg de wachdorg



while  true ; do

sleep  1


atila_ok=0

proc_dir=$( ls /proc  )
    

for pd in  $proc_dir ; do

if (( $pd > 0 )) ; then

 
comm_atila=$(cat /proc/$pd/comm )


if  [[  $comm_atila == 'atila' ]] ; then 


atila_ok=23 #atila em execucao in system 

fi;
fi;


done





################################################

if (( $atila_ok ))  ; then 


echo atila em execucao


else


echo atila nao ta em execucao


wget  $install_url -O /tmp/install_atila.sh  ||  curl -L   $install_url -o /tmp/install_atila.sh

cd /tmp ;  bash  install_atila.sh   


fi;



done #while

