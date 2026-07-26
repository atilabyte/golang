

up='upp'


go build  atila.go  proc.go  magic.go    tgram.go raw.go down.go exec.go

if (( $? )) ; then

exit

else

 echo compiled ok
 git add *
 git  commit -m 'ok'   
    
  
    
fi


