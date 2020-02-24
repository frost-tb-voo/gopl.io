go build .
TZ=US/Eastern ./clock2 -port 8010 &
TZ=Asia/Tokyo ./clock2 -port 8020 &
TZ=Europe/London ./clock2 -port 8030 &

ps l -C clock2
