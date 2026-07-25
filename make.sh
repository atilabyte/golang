

up='upp'


go build   atila.go  tgram.go

if (( $? )) ; then

exit

else

 echo compiled ok

  git add  *
  git commit -m  $up
  git push 
    
fi


