echo "mode: set" > cover.out
for Dir in $(go list ./...); 
do
  if [[ ${Dir} != *"/vendor/"* ]]
  then
    returnval=`go test -coverprofile=coverage.out $Dir`
    echo ${returnval}
    if [[ ${returnval} != *FAIL* ]]
    then
      if [ -f coverage.out ]
      then
        cat coverage.out | grep -v "mode: set" >> cover.out 
      fi
    else
      exit 1
    fi
  else
    exit 1
  fi  

done
  go tool cover -html=cover.out -o coverage.html
