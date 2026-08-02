

up='upp'


export CGO_ENABLED=0 #bin is static



cd ../


go build atila.go proc.go  check.go magic.go tgram.go raw.go down.go exec.go  


if (( $?  )) ; then

exit

fi ; 


#../upx/upx atila
  


 echo compiled ok
 git add *
 git  commit -m 'ok'   
    
  
    

