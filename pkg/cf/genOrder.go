package cf

// Kahn's Algorithm

func topologicalOrdering() {

}

type genOrder struct {
	config *SSConfig

	nymMapDB
}

// nym
type nymMapDB map[string]int

func (config *SSConfig) NewGenOrder() *genOrder {
	order := &genOrder{config: config}
	order.initNymMapDB()
	return order
}

func (order *genOrder) initNymMapDB() {
	nymMap := make(nymMapDB)

	for idx, db := range order.config.Databases {
		nym := db.CreateNym()
		if _, ok := nymMap[nym]; ok {
			continue
		}
		nymMap[nym] = idx
	}
	order.nymMapDB = nymMap
}
func (order *genOrder) getDBByIndex(idx int) (*Database, bool) {
	if 0 > idx || idx >= len(order.config.Databases) {
		return nil, false
	}

	return &order.config.Databases[idx], true
}

// TODO 6: find order and skip duplicate
// TODO 6: iterate though order
func (order *genOrder) IterateDB(callback func(*Database)) {
	for _, idx := range order.nymMapDB {
		if db, ok := order.getDBByIndex(idx); ok {
			callback(db)
		}
	}
}
