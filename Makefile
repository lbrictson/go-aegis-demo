.PHONY: init test

test:
	. env/bin/activate	
	./env/bin/nosetests --with-xunit

deploydev:
	aegis_configs/dev/deploy.sh
	

deployprod:
	aegis_configs/prod/deploy.sh
