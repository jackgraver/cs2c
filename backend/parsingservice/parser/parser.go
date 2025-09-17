package parser

import (
	"fmt"
	"os"
	"parsingservice/parser/handlers"
	"parsingservice/parser/structs"
	"parsingservice/parser/utils"
	"time"

	"github.com/google/uuid"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/msg"
) 

func CreateHandlers(context *handlers.HandlerContext) {
	context.Parser.RegisterNetMessageHandler(func(msg *msg.CSVCMsg_ServerInfo) {
		context.DemoData.Map = msg.GetMapName()
	})
	handlers.RegisterGrenadeHandlers(context)
	handlers.RegisterPlayerHandler(context)
	handlers.RegisterRoundHandler(context)
	handlers.RegisterBombHandler(context)
	handlers.RegisterKillHandlers(context)
}

func Parse(fileName string) (bool, error){
	filePath := utils.JoinToStoreDir(fileName )//filepath.Join(storeDir, fileName)
	demFile, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return false, err
	}
	demoinfocs_parser := demoinfocs.NewParser(demFile)
	defer demoinfocs_parser.Close()

	demo_id := uuid.New().String()
	context := &handlers.HandlerContext{
		Parser: demoinfocs_parser,
		DemoData: &structs.DemoData {
			DemoID: demo_id,
			SeriesID: uuid.New().String(),  
			NumRounds: 0,
			Map: "",
			UploadDate: time.Now().Format(time.RFC3339), 
		},
		CurrentRound: EmptyRound(),
		FirstRound: EmptyRound(),
		ActiveSmokes: make(map[int]structs.SmokeMolly),
		ActiveMollies: make(map[int]structs.SmokeMolly),
		InAirGrenades: make(map[int]structs.InAirGrenade),
	}
	CreateHandlers(context)

	if err := demoinfocs_parser.ParseToEnd(); err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
		return false, err
	}

	if context.FirstRound != nil {
		err := utils.WriteRounds(context.AllRounds, demo_id)
		if err != nil {
			return false, fmt.Errorf("failed writing all rounds file")
		}

		demFile.Close()
		err = os.Remove(filePath)
		if err != nil {
			return false, fmt.Errorf("failed to remove file: %w", err)
		}

		return true, nil
	} else {
		return false, fmt.Errorf("no round parsed")
	}

}


func EmptyRound() *structs.RoundData {
	return &structs.RoundData{
		RoundNum: 0,
		Ticks: []structs.TickData{},
	}
}