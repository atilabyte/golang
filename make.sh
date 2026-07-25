

up='upp'


go build   main.go  atila.go  tgram.go 

if (( $? )) ; then

exit

else

 echo compiled ok
 git add *
 git  commit -m 'ok'   
    gi push
  
    
fi


