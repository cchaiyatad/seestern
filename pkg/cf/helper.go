package cf

type databaseCollectionInfo map[string]map[string]struct{}

func (ssconfig *SSConfig) GetDatabaseCollectionInfo() databaseCollectionInfo {
	info := make(databaseCollectionInfo)

	for _, db := range ssconfig.Databases {
		dbName := db.D_name
		collName := db.Collection.C_name
		if _, ok := info[dbName]; !ok {
			info[dbName] = make(map[string]struct{})
		}
		info[dbName][collName] = struct{}{}
	}

	return info
}
