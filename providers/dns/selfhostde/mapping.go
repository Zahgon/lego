package selfhostde

const (
	lineSep   = ","
	recordSep = ":"
)

type Seq struct {
	cursor int
	ids    []string
}

func NewSeq(ids ...string) *Seq { _ = "STUB: not implemented"; return nil }

func (s *Seq) Next() string { _ = "STUB: not implemented"; return "" }

func parseRecordsMapping(raw string) (map[string]*Seq, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseLine(line string) (string, *Seq, error) { _ = "STUB: not implemented"; return "", nil, nil }

func safeIndex(v, sep string) (int, error) { _ = "STUB: not implemented"; return 0, nil }
