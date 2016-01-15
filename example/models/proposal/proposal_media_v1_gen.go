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

// Proposal views
type ProposalViewEnum string

const (
	// Proposal default view
	ProposalDefaultView ProposalViewEnum = "default"

	// Proposal link view
	ProposalLinkView ProposalViewEnum = "link"
)

// LoadProposal loads raw data into an instance of Proposal
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadProposal(raw interface{}) (res *Proposal, err error) {
	res, err = UnmarshalProposal(raw, err)
	return
}

// Dump produces raw data from an instance of Proposal running all the
// validations. See LoadProposal for the definition of raw data.

func (mt *Proposal) Dump(view ProposalViewEnum) (res map[string]interface{}, err error) {
	if view == ProposalDefaultView {
		res, err = MarshalProposal(mt, err)
	}
	if view == ProposalLinkView {
		res, err = MarshalProposalLink(mt, err)
	}
	return
}

// MarshalProposal validates and renders an instance of Proposal into a interface{}
// using view "default".
func MarshalProposal(source *Proposal, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp15 := map[string]interface{}{
		"abstract": source.Abstract,
		"detail":   source.Detail,
		"href":     source.Href,
		"id":       source.ID,
		"title":    source.Title,
	}
	target = tmp15
	return
}

// MarshalProposalLink validates and renders an instance of Proposal into a interface{}
// using view "link".
func MarshalProposalLink(source *Proposal, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp16 := map[string]interface{}{
		"href":  source.Href,
		"id":    source.ID,
		"title": source.Title,
	}
	target = tmp16
	return
}

// UnmarshalProposal unmarshals and validates a raw interface{} into an instance of Proposal
func UnmarshalProposal(source interface{}, inErr error) (target *Proposal, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Proposal)
		if v, ok := val["abstract"]; ok {
			var tmp17 string
			if val, ok := v.(string); ok {
				tmp17 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Abstract`, v, "string", err)
			}
			target.Abstract = &tmp17
		}
		if v, ok := val["detail"]; ok {
			var tmp18 string
			if val, ok := v.(string); ok {
				tmp18 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Detail`, v, "string", err)
			}
			target.Detail = &tmp18
		}
		if v, ok := val["href"]; ok {
			var tmp19 string
			if val, ok := v.(string); ok {
				tmp19 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Href`, v, "string", err)
			}
			target.Href = &tmp19
		}
		if v, ok := val["id"]; ok {
			var tmp20 int
			if f, ok := v.(float64); ok {
				tmp20 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = &tmp20
		}
		if v, ok := val["title"]; ok {
			var tmp21 string
			if val, ok := v.(string); ok {
				tmp21 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Title`, v, "string", err)
			}
			target.Title = &tmp21
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}
