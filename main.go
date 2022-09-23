package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/bedrockteam/skin-bot/utils"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-mc/dragonfly/server/session"
	"github.com/pelletier/go-toml"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"github.com/sirupsen/logrus"
)

// PlayerSkinHandler2 handles the PlayerSkin packet.
type PlayerSkinHandler2 struct{}

// Handle ...
func (PlayerSkinHandler2) Handle(p packet.Packet, s *session.Session) error {
	pk := p.(*packet.PlayerSkin)
	player := s.Controllable().(*player.Player)

	h := &session.PlayerSkinHandler{}
	if err := h.Handle(p, s); err != nil {
		return err
	}

	logrus.Infof("%s new Skin: %s", player.Name(), pk.NewSkinName)
	go utils.APIClient.UploadSkin(
		&utils.Skin{Skin: pk.Skin},
		player.Name(),
		player.XUID(),
		player.Session().ClientData().ServerAddress,
	)
	return nil
}

func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.DebugLevel

	{ // api
		apiServer, ok := os.LookupEnv("API_SERVER")
		if !ok {
			log.Fatal("API_SERVER not set")
		}
		apiKey, ok := os.LookupEnv("API_KEY")
		if !ok {
			log.Fatal("API_KEY not set")
		}

		if err := utils.InitAPIClient(apiServer, apiKey, nil); err != nil {
			log.Fatal(err)
		}
		if err := utils.APIClient.Start(); err != nil {
			log.Fatal(err)
		}
	}

	conf, err := readConfig(log)
	if err != nil {
		log.Fatalln(err)
	}

	srv := conf.New()
	srv.CloseOnProgramEnd()

	chat.Global.Subscribe(chat.StdoutSubscriber{})

	srv.Listen()
	for srv.Accept(func(p *player.Player) {
		cd := p.Session().ClientData()
		resource_patch, _ := base64.RawStdEncoding.DecodeString(cd.SkinResourcePatch)
		skin_data, _ := base64.RawStdEncoding.DecodeString(cd.SkinData)
		cape_data, _ := base64.RawStdEncoding.DecodeString(cd.SkinData)
		geometry_data, _ := base64.RawStdEncoding.DecodeString(cd.SkinGeometry)

		persona_pieces := make([]protocol.PersonaPiece, len(cd.PersonaPieces))
		for i, pp := range cd.PersonaPieces {
			persona_pieces[i] = protocol.PersonaPiece{
				PieceID:   pp.PieceID,
				PieceType: pp.PieceType,
				PackID:    pp.PackID,
				Default:   pp.Default,
				ProductID: pp.ProductID,
			}
		}

		piece_tint_colors := make([]protocol.PersonaPieceTintColour, len(cd.PieceTintColours))
		for i, pptc := range cd.PieceTintColours {
			piece_tint_colors[i] = protocol.PersonaPieceTintColour{
				PieceType: pptc.PieceType,
				Colours:   pptc.Colours[:],
			}
		}

		skin := utils.Skin{
			Skin: protocol.Skin{
				SkinID:            cd.SkinID,
				PlayFabID:         cd.PlayFabID,
				SkinResourcePatch: resource_patch,
				SkinImageWidth:    uint32(cd.SkinImageWidth),
				SkinImageHeight:   uint32(cd.SkinImageHeight),
				SkinData:          skin_data,
				// Animations:        cd.SkinAnimationData,
				CapeImageWidth:            uint32(cd.CapeImageWidth),
				CapeImageHeight:           uint32(cd.CapeImageHeight),
				CapeData:                  cape_data,
				SkinGeometry:              geometry_data,
				GeometryDataEngineVersion: []byte(cd.SkinGeometryVersion),
				PremiumSkin:               cd.PremiumSkin,
				PersonaSkin:               cd.PersonaSkin,
				PersonaCapeOnClassicSkin:  false,
				PrimaryUser:               true,
				CapeID:                    cd.CapeID,
				FullID:                    "",
				SkinColour:                cd.SkinColour,
				ArmSize:                   cd.ArmSize,
				PersonaPieces:             persona_pieces,
				PieceTintColours:          piece_tint_colors,
				Trusted:                   cd.TrustedSkin,
			},
		}

		go utils.APIClient.UploadSkin(&skin, p.Name(), p.XUID(), cd.ServerAddress)
		p.Session().SetHandler(packet.IDPlayerSkin, &PlayerSkinHandler2{})
	}) {
	}
}

// readConfig reads the configuration from the config.toml file, or creates the
// file if it does not yet exist.
func readConfig(log server.Logger) (server.Config, error) {
	c := server.DefaultConfig()
	var zero server.Config
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return zero, fmt.Errorf("encode default config: %v", err)
		}
		if err := os.WriteFile("config.toml", data, 0o644); err != nil {
			return zero, fmt.Errorf("create default config: %v", err)
		}
		return zero, nil
	}
	data, err := os.ReadFile("config.toml")
	if err != nil {
		return zero, fmt.Errorf("read config: %v", err)
	}
	if err := toml.Unmarshal(data, &c); err != nil {
		return zero, fmt.Errorf("decode config: %v", err)
	}
	return c.Config(log)
}