package libwtf

type Attribute struct {
	Key string
	Value string
}

type Record struct {
	Severity uint8
	Attributes []Attribute
}

type DagNode struct {
	Children []*DagNode
}

type Dag *DagNode

func MakeRecord(severity uint8, attributes []Attribute) Record {
	record := Record{
		Severity: severity,
		Attributes: attributes,
	}
	fmt.Println("Made record=%+v", record)
	return record
}

func Emit(record Record) int {
	// TODO
	return 0
}

func Collect(attributes []Attribute) []Record {
	// TODO
	return make([]Record, 0)
}

func Analyze(records []Record) Dag {
	if len(records) == 0 {
		fmt.Println("No records")
		return nil
	}
	// TODO
	return nil
}

// combines Collect() and Analyze()
func WheresTheFault(attributes []Attribute) Dag {
	records := Collect(attributes)
	dag := Analyze(records)
	return dag
}