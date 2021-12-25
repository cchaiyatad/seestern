package cf

func (d *Database) GetRef() []string {
	var ref []string
	return ref
}

func (d *Database) getNym() (string, bool) {
	if d == nil {
		return "", false
	}
	return CreateNym(d.D_name, d.Collection.C_name), true
}

func (d *Database) isEqualToNym(nym string) bool {
	if thisNym, ok := d.getNym(); ok {
		return thisNym == nym
	}
	return false
}
