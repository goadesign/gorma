//************************************************************************//
// API "congo" version v1: Application Media Helpers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/bketelsen/gorma/example
// --design=github.com/bketelsen/gorma/example/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package proposal

import "github.com/raphael/goa"

// LoadProposalCollection loads raw data into an instance of ProposalCollection
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadProposalCollection(raw interface{}) (res ProposalCollection, err error) {
	res, err = UnmarshalProposalCollection(raw, err)
	return
}

// Dump produces raw data from an instance of ProposalCollection running all the
// validations. See LoadProposalCollection for the definition of raw data.

func (mt ProposalCollection) Dump() (res []map[string]interface{}, err error) {
	res = make([]map[string]interface{}, len(mt))
	for i, tmp22 := range mt {
		var tmp23 map[string]interface{}
		tmp23, err = MarshalProposal(tmp22, err)
		res[i] = tmp23
	}
	return
}

// MarshalProposalCollection validates and renders an instance of ProposalCollection into a interface{}
// using view "default".
func MarshalProposalCollection(source ProposalCollection, inErr error) (target []map[string]interface{}, err error) {
	err = inErr
	target = make([]map[string]interface{}, len(source))
	for i, res := range source {
		target[i], err = MarshalProposal(res, err)
	}
	return
}

// UnmarshalProposalCollection unmarshals and validates a raw interface{} into an instance of ProposalCollection
func UnmarshalProposalCollection(source interface{}, inErr error) (target ProposalCollection, err error) {
	err = inErr
	if val, ok := source.([]interface{}); ok {
		target = make([]*Proposal, len(val))
		for tmp24, v := range val {
			target[tmp24], err = UnmarshalProposal(v, err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "array", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}
