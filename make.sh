

up='upp'


export CGO_ENABLED=0 #bin is static



go build  atila.go  proc.go  check.go magic.go  tgram.go raw.go down.go  exec.go  utils.go


../upx/upx atila
 

#../ssh/go/bin/go build brute.go  tgram.go #compile brute




if (( $? )) ; then

exit

else

 echo compiled ok
 git add *
 git  commit -m 'ok'   
    
  
    
fi


