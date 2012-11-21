#!/bin/bash

wget -c \
	http://www.census.gov/genealogy/www/data/1990surnames/dist.all.last \
	http://www.census.gov/genealogy/www/data/1990surnames/dist.female.first \
	http://www.census.gov/genealogy/www/data/1990surnames/dist.male.first

cat dist.all.last | cut -d' ' -f1 > data/names/dist.all.last.txt
cat dist.female.first | cut -d' ' -f1 > data/names/dist.female.first.txt
cat dist.male.first | cut -d' ' -f1 > data/names/dist.male.first.txt

tools/txt2json/txt2json data/names/*.txt

wget -c http://sourceforge.net/projects/zips/files/zips/zips.csv.zip/zips.csv.zip
unzip zips.csv.zip
mv zips.csv data/postal/ -v

rm -vf \
	dist.*.* \
	data/names/*.txt \
	zips.csv.zip

