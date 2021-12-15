package db

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

type GenParam struct {
	CntStr  string
	Vendor  string
	File    string
	Outpath string

	Verbose  bool
	IsDrop   bool
	IsInsert bool
}
