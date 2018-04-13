#!/bin/bash
go build && mv aetest aegis_app
zip hello.zip -r templates aegis_app
mv hello.zip ./aegis_configs/dev/
cp ./*.go ./aegis_configs/dev/
cd aegis_configs/dev/
pwd
aegis deploy
rm -f ./*.go
rm -f ../../aegis_app