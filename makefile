build64:
	GOOS=linux GOARCH=amd64 go build -o partner_system github.com/ubertrip/partner-system

staging_deploy: build64
	scp -C ./partner_system oz@95.85.12.25:/home/oz/staging-partner-system/partner_system.tmp
	ssh oz@95.85.12.25 killall partner_system; exit 0
	sleep 1;
	rsync -czavP env/ oz@95.85.12.25:/home/oz/staging-partner-system/env
	ssh oz@95.85.12.25 "mv ~/staging-partner-system/partner_system.tmp ~/staging-partner-system/partner_system; cd ~/staging-partner-system; (setsid nohup ./partner_system --env=staging > ./partner_system.out 2>&1 &)"