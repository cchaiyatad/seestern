package cf

func (d *Database) GetRef() []string {
	refs := make([]string, 0)
	for _, field := range d.Collection.Fields {
		for _, con := range field.Constraints {
			refs = append(refs, con.getRef()...)
		}
	}
	return refs
}

func (c *Constraint) getRef() []string {
	refs := c.Type.Ref()
	validateRef := make([]string, 0, len(refs))
	for _, ref := range refs {
		if !isRefValid(ref) {
			continue
		}
		validateRef = append(validateRef, ref)
	}
	return validateRef
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
