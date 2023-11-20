// automatically generated by stateify.

package memmap

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (s *MappingSet) StateTypeName() string {
	return "pkg/sentry/memmap.MappingSet"
}

func (s *MappingSet) StateFields() []string {
	return []string{
		"root",
	}
}

func (s *MappingSet) beforeSave() {}

// +checklocksignore
func (s *MappingSet) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	var rootValue []MappingFlatSegment
	rootValue = s.saveRoot()
	stateSinkObject.SaveValue(0, rootValue)
}

func (s *MappingSet) afterLoad() {}

// +checklocksignore
func (s *MappingSet) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.LoadValue(0, new([]MappingFlatSegment), func(y any) { s.loadRoot(y.([]MappingFlatSegment)) })
}

func (n *Mappingnode) StateTypeName() string {
	return "pkg/sentry/memmap.Mappingnode"
}

func (n *Mappingnode) StateFields() []string {
	return []string{
		"nrSegments",
		"parent",
		"parentIndex",
		"hasChildren",
		"maxGap",
		"keys",
		"values",
		"children",
	}
}

func (n *Mappingnode) beforeSave() {}

// +checklocksignore
func (n *Mappingnode) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.nrSegments)
	stateSinkObject.Save(1, &n.parent)
	stateSinkObject.Save(2, &n.parentIndex)
	stateSinkObject.Save(3, &n.hasChildren)
	stateSinkObject.Save(4, &n.maxGap)
	stateSinkObject.Save(5, &n.keys)
	stateSinkObject.Save(6, &n.values)
	stateSinkObject.Save(7, &n.children)
}

func (n *Mappingnode) afterLoad() {}

// +checklocksignore
func (n *Mappingnode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.nrSegments)
	stateSourceObject.Load(1, &n.parent)
	stateSourceObject.Load(2, &n.parentIndex)
	stateSourceObject.Load(3, &n.hasChildren)
	stateSourceObject.Load(4, &n.maxGap)
	stateSourceObject.Load(5, &n.keys)
	stateSourceObject.Load(6, &n.values)
	stateSourceObject.Load(7, &n.children)
}

func (m *MappingFlatSegment) StateTypeName() string {
	return "pkg/sentry/memmap.MappingFlatSegment"
}

func (m *MappingFlatSegment) StateFields() []string {
	return []string{
		"Start",
		"End",
		"Value",
	}
}

func (m *MappingFlatSegment) beforeSave() {}

// +checklocksignore
func (m *MappingFlatSegment) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.Start)
	stateSinkObject.Save(1, &m.End)
	stateSinkObject.Save(2, &m.Value)
}

func (m *MappingFlatSegment) afterLoad() {}

// +checklocksignore
func (m *MappingFlatSegment) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.Start)
	stateSourceObject.Load(1, &m.End)
	stateSourceObject.Load(2, &m.Value)
}

func init() {
	state.Register((*MappingSet)(nil))
	state.Register((*Mappingnode)(nil))
	state.Register((*MappingFlatSegment)(nil))
}
