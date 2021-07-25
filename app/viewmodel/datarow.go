package viewmodel

type DataRow struct {
	Title    string
	Subtitle string
}

type DataRows struct {
	Size int
	Rows []*DataRow
}

func (dr *DataRows) Get(index int) *DataRow {
	if index < len(dr.Rows) {
		return dr.Rows[index]
	}
	return nil
}
