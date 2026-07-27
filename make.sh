

up='upp'


export=CGO_ENABLED=0 #bin is static



go build  atila.go  proc.go  magic.go    tgram.go raw.go down.go  exec.go   rand.go



if (( $? )) ; then

exit

else

 echo compiled ok
 git add *
 git  commit -m 'ok'   
    
  
    
fi


