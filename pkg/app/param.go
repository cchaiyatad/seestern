package app

type PSParam struct {
	CntStr string
	Vendor string
	DBName string
}

type InitParam struct {
	CntStr      string
	Vendor      string
	TargetColls []string
	Outpath     string
	Verbose     bool
	FileType    string
}

func (param InitParam) isWriteFile() bool {
	return param.Outpath != ""
}

type GenParam struct {
	CntStr  string
	Vendor  string
	File    string
	Outpath string

	Verbose  bool
	IsDrop   bool
	IsInsert bool
}

func (param GenParam) shouldConnectDB() bool {
	return param.IsDrop || param.IsInsert
}

func (param GenParam) shouldGenJson() bool {
	return param.Verbose || param.Outpath != ""
}

func (param GenParam) isWriteFile() bool {
	return param.Outpath != ""
}
