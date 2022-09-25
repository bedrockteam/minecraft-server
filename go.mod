module github.com/bedrockteam/minecraft-server

go 1.19

//replace github.com/bedrockteam/skin-bot => ../../../bedrock-lol/skin-bot

require (
	github.com/bedrockteam/skin-bot v0.1.2-3
	github.com/df-mc/dragonfly v0.8.5
	github.com/pelletier/go-toml v1.9.5
	github.com/sandertv/gophertunnel v1.24.9
	github.com/sirupsen/logrus v1.9.0
)

require (
	github.com/brentp/intintmap v0.0.0-20190211203843-30dc0ade9af9 // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/df-mc/atomic v1.10.0 // indirect
	github.com/df-mc/goleveldb v1.1.9 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/go-gl/mathgl v1.0.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/klauspost/compress v1.15.10 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/muhammadmuzzammil1998/jsonc v1.0.0 // indirect
	github.com/rabbitmq/amqp091-go v1.5.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/sandertv/go-raknet v1.12.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	golang.org/x/crypto v0.0.0-20220924013350-4ba4fb4dd9e7 // indirect
	golang.org/x/exp v0.0.0-20220921164117-439092de6870 // indirect
	golang.org/x/image v0.0.0-20220902085622-e7cb96979f69 // indirect
	golang.org/x/net v0.0.0-20220923203811-8be639271d50 // indirect
	golang.org/x/oauth2 v0.0.0-20220909003341-f21342109be1 // indirect
	golang.org/x/sys v0.0.0-20220919091848-fb04ddd9f9c8 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

//replace github.com/df-mc/dragonfly => ../../dragonfly
replace github.com/df-mc/dragonfly => github.com/olebeck/dragonfly v0.8.5-1

//replace github.com/sandertv/gophertunnel => ./gophertunnel
replace github.com/sandertv/gophertunnel => github.com/olebeck/gophertunnel v1.24.8-7
