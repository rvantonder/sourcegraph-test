package lsif

import "github.com/sourcegraph/sourcegraph/enterprise/cmd/precise-code-intel-worker/internal/correlation/datastructures"

type Element struct { /* all structs must go */ }

type Edge struct { /* all structs must go */ }

type MetaData struct { /* all structs must go */ }

type Document struct { /* all structs must go */ }

type Range struct { /* all structs must go */ }

func (d Range) SetDefinitionResultID(id int) Range {
	return Range{
		StartLine:          d.StartLine,
		StartCharacter:     d.StartCharacter,
		EndLine:            d.EndLine,
		EndCharacter:       d.EndCharacter,
		DefinitionResultID: id,
		ReferenceResultID:  d.ReferenceResultID,
		HoverResultID:      d.HoverResultID,
		MonikerIDs:         d.MonikerIDs,
	}
}

func (d Range) SetReferenceResultID(id int) Range {
	return Range{
		StartLine:          d.StartLine,
		StartCharacter:     d.StartCharacter,
		EndLine:            d.EndLine,
		EndCharacter:       d.EndCharacter,
		DefinitionResultID: d.DefinitionResultID,
		ReferenceResultID:  id,
		HoverResultID:      d.HoverResultID,
		MonikerIDs:         d.MonikerIDs,
	}
}

func (d Range) SetHoverResultID(id int) Range {
	return Range{
		StartLine:          d.StartLine,
		StartCharacter:     d.StartCharacter,
		EndLine:            d.EndLine,
		EndCharacter:       d.EndCharacter,
		DefinitionResultID: d.DefinitionResultID,
		ReferenceResultID:  d.ReferenceResultID,
		HoverResultID:      id,
		MonikerIDs:         d.MonikerIDs,
	}
}

func (d Range) SetMonikerIDs(ids *datastructures.IDSet) Range {
	return Range{
		StartLine:          d.StartLine,
		StartCharacter:     d.StartCharacter,
		EndLine:            d.EndLine,
		EndCharacter:       d.EndCharacter,
		DefinitionResultID: d.DefinitionResultID,
		ReferenceResultID:  d.ReferenceResultID,
		HoverResultID:      d.HoverResultID,
		MonikerIDs:         ids,
	}
}

type ResultSet struct { /* all structs must go */ }

func (d ResultSet) SetDefinitionResultID(id int) ResultSet {
	return ResultSet{
		DefinitionResultID: id,
		ReferenceResultID:  d.ReferenceResultID,
		HoverResultID:      d.HoverResultID,
		MonikerIDs:         d.MonikerIDs,
	}
}

func (d ResultSet) SetReferenceResultID(id int) ResultSet {
	return ResultSet{
		DefinitionResultID: d.DefinitionResultID,
		ReferenceResultID:  id,
		HoverResultID:      d.HoverResultID,
		MonikerIDs:         d.MonikerIDs,
	}
}

func (d ResultSet) SetHoverResultID(id int) ResultSet {
	return ResultSet{
		DefinitionResultID: d.DefinitionResultID,
		ReferenceResultID:  d.ReferenceResultID,
		HoverResultID:      id,
		MonikerIDs:         d.MonikerIDs,
	}
}

func (d ResultSet) SetMonikerIDs(ids *datastructures.IDSet) ResultSet {
	return ResultSet{
		DefinitionResultID: d.DefinitionResultID,
		ReferenceResultID:  d.ReferenceResultID,
		HoverResultID:      d.HoverResultID,
		MonikerIDs:         ids,
	}
}

type Moniker struct { /* all structs must go */ }

func (d Moniker) SetPackageInformationID(id int) Moniker {
	return Moniker{
		Kind:                 d.Kind,
		Scheme:               d.Scheme,
		Identifier:           d.Identifier,
		PackageInformationID: id,
	}
}

type PackageInformation struct { /* all structs must go */ }

type DiagnosticResult struct { /* all structs must go */ }

type Diagnostic struct { /* all structs must go */ }
