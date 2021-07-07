package dal

type Meta struct {
	Key   string `db:"key"`
	Value string `db:"value"`
}

type TMArea struct {
	UUID    string `db:"uuid"`
	Title   string `db:"title"`
	Visible bool   `db:"title"`
	Index   int32  `db:"index"`
	// CachedTags []byte `db:"cachedTags"`
}

type TMSettings struct {
}

type TMTag struct {
}

type TMContact struct{}

type TMChecklistItem struct{
	UUID string `db:"uuid"`

}



type TMTask struct{

}

type TMMetaItem struct{}

type TMCommand struct{}

type TMTombstone struct{}
