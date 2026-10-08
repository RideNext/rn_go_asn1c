#!/bin/bash 
mkdir -p BinaryTemplates
#mkdir -p protocol
mak_list=`find asn1 -name "*.asn"`
rm -rf BynaryTemplates/*
rm -rf protocols/*
#rm -rf protocol/*
for each in $mak_list;
do
	asnname=`basename  $each`
	withoutext=`echo $asnname | cut -f1 -d'.'`
	asname=`echo $withoutext | cut -f1 -d'-'`
	version=`echo $withoutext | cut -f2 -d'-'`
	echo "Creating code for $asname for version $version"
	mkdir -p protocols/${asname}/${version}/
	if [ "$asname" == "rrc" ] ; then
            python2 genasnpy.py -i $asnname -t uaper
	    rm -rf protocol/${asname}/${version}/__pycache__
	    #continue
	elif [ "$asname" == "nas" ] ; then
	    continue
        elif [ "$asname" == "pyint" ] ; then
            python2 genasnpy.py -i $asnname -t per
	    rm protocol/${asname}/${version}/*.py
	    rm -rf protocol/${asname}/${version}/__pycache__
	    continue
	elif [ "$asname" == "e2sm*" ] ; then
            python2 genasnpy.py -i $asnname -t per
        else
            python2 genasnpy.py -i $asnname -t per -p $asnname
        fi
	echo "creating template"
	cd protocols/${asname}/${version} 
	   go mod init ridenext.co.in/goasn1/${asname}
	cd -
	mkdir -p protocols/${asname}/cmd/
	cp libpy/${asname}.go protocols/${asname}/cmd/
	cd protocols/${asname}/cmd/
	    rm -f go.mod
	    go mod init ridenext.co.in/goasn1/test
	    echo "replace ridenext.co.in/goasn1/${asname} => ../${version}/" >> go.mod
	    go mod tidy
	cd -

done

mkdir -p BlueFox/messages
rm -rf BlueFox/messages/*
cp -arf  BinaryTemplates BlueFox/messages/
rm -rf BlueFox/protocols
cp -arf protocols BlueFox/
#sh build.sh

