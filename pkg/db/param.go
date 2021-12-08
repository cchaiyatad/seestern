package db

type PSParam struct {
	CntStr string
	Vendor string
	DBName string
}

type InitParam struct {
	CntStr   string
	TargetDB []string
	Verbose  bool
}
