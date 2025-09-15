package parser

func parse(fileName string) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid file upload"})
		return
	}

	openedFile, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to open uploaded file"})
		return
	}
	defer openedFile.Close()

	// Check the file extension
	fileExt := strings.ToLower(filepath.Ext(file.Filename))

	switch fileExt {
	case ".zip":
		// If it's a .zip file, read it as a zip and extract files inside
		r, err := zip.NewReader(openedFile, file.Size)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to read zip file"})
			return
		}

		demo, err := parseZip(r.File)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, demo)
		return
	case ".dem":
		demo, err := parseDemo(openedFile)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		
		db.InsertDemo(demo)
		c.JSON(200, demo)
		return
	default:
		c.JSON(400, gin.H{"error": "Invalid file type. Only .zip or .dem are allowed"})
		return
	}
}

func parseZip(files []*zip.File) (*structs.DemoData, error) {
	series_id := uuid.New().String()

	var firstDemo *structs.DemoData

	for _, f := range files {
		if filepath.Ext(f.Name) == ".dem" {
			fmt.Println("parse", f.Name)

			// Open the file to get an io.ReadCloser
			rc, err := f.Open()
			if err != nil {
				fmt.Printf("failed to open file %s: %v\n", f.Name, err)
				continue
			}
			defer rc.Close()

			go_parser := demoinfocs.NewParser(rc)
			defer go_parser.Close()

			demo_id := uuid.New().String()

			context := &handlers.HandlerContext{
				Parser: go_parser,
				DemoData: &structs.DemoData {
					DemoID: demo_id,
					SeriesID: series_id,  
					NumRounds: 0,
					Map: "",
					UploadDate: time.Now().Format(time.RFC3339), 
				},
			}

			go_parser.RegisterNetMessageHandler(func(msg *msg.CSVCMsg_ServerInfo) {
				context.DemoData.Map = msg.GetMapName()
			})

			parser.CreateHandlers(context)

			if err := go_parser.ParseToEnd(); err != nil {
				return nil, err
			}

			if context.FirstRound != nil {
				err := utils.WriteRounds(context.AllRounds, demo_id)
				if err != nil {
					return nil, fmt.Errorf("failed writing all rounds file for %s", f.Name)
				}
				db.InsertDemo(context.DemoData)

				if firstDemo == nil {
					firstDemo = context.DemoData
				}
			} else {
				return nil, fmt.Errorf("no round parsed for %s", f.Name)
			}
		}
	}

	return firstDemo, nil
}

func parseDemo(uploadedFile multipart.File) (*structs.DemoData, error) {
	go_parser := demoinfocs.NewParser(uploadedFile)
	defer go_parser.Close()

	demo_id := uuid.New().String()

	context := &handlers.HandlerContext{
		Parser: go_parser,
		DemoData: &structs.DemoData {
			DemoID: demo_id,
			SeriesID: uuid.New().String(),  
			NumRounds: 0,
			Map: "",
			UploadDate: time.Now().Format(time.RFC3339), 
		},
	}

	go_parser.RegisterNetMessageHandler(func(msg *msg.CSVCMsg_ServerInfo) {
		context.DemoData.Map = msg.GetMapName()
	})

	parser.CreateHandlers(context)

	if err := go_parser.ParseToEnd(); err != nil {
		return nil, err
	}

	if context.FirstRound != nil {
		err := utils.WriteRounds(context.AllRounds, demo_id)
		if err != nil {
			return nil, fmt.Errorf("failed writing all rounds file")
		}
		return context.DemoData, nil
	} else {
		return nil, fmt.Errorf("no round parsed")
	}
}