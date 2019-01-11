docker build -t sln/hlc2018 .
docker tag sln/hlc2018 stor.highloadcup.ru/accounts/quiet_ant
docker push stor.highloadcup.ru/accounts/quiet_ant