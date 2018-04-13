#!/bin/bash
go build && mv aetest aegis_app
zip hello.zip -r templates aegis_app
mv hello.zip ./aegis_configs/prod/
cp ./*.go ./aegis_configs/prod/
cd aegis_configs/prod/
pwd
aegis deploy
rm -f ./*.go
rm -f ../../aegis_app