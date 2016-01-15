//************************************************************************//
// API "congo": Application User Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/bketelsen/gorma/example
// --design=github.com/bketelsen/gorma/example/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/raphael/goa"

// ProposalModel type
type ProposalModel struct {
	UserID    *int
	Abstract  *string
	CreatedAt *interface{}
	DeletedAt *interface{}
	Detail    *string
	Firstname *string
	ID        string
	M2reviews []*ReviewModel
	Reviews   []*ReviewModel
	Title     *string
	UpdatedAt *interface{}
	Withdrawn *bool
}

// Validate validates the type instance.
func (ut *ProposalModel) Validate() (err error) {
	if ut.ID == "" {
		err = goa.MissingAttributeError(`response`, "id", err)
	}

	if ut.Abstract != nil {
		if len(*ut.Abstract) < 50 {
			err = goa.InvalidLengthError(`response.abstract`, *ut.Abstract, len(*ut.Abstract), 50, true, err)
		}
	}
	if ut.Abstract != nil {
		if len(*ut.Abstract) > 500 {
			err = goa.InvalidLengthError(`response.abstract`, *ut.Abstract, len(*ut.Abstract), 500, false, err)
		}
	}
	if ut.Detail != nil {
		if len(*ut.Detail) < 100 {
			err = goa.InvalidLengthError(`response.detail`, *ut.Detail, len(*ut.Detail), 100, true, err)
		}
	}
	if ut.Detail != nil {
		if len(*ut.Detail) > 2000 {
			err = goa.InvalidLengthError(`response.detail`, *ut.Detail, len(*ut.Detail), 2000, false, err)
		}
	}
	if ut.Firstname != nil {
		if len(*ut.Firstname) < 2 {
			err = goa.InvalidLengthError(`response.firstname`, *ut.Firstname, len(*ut.Firstname), 2, true, err)
		}
	}
	for _, e := range ut.M2reviews {
		if e.Comment != nil {
			if len(*e.Comment) < 10 {
				err = goa.InvalidLengthError(`response.m2reviews[*].comment`, *e.Comment, len(*e.Comment), 10, true, err)
			}
		}
		if e.Comment != nil {
			if len(*e.Comment) > 200 {
				err = goa.InvalidLengthError(`response.m2reviews[*].comment`, *e.Comment, len(*e.Comment), 200, false, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating < 1 {
				err = goa.InvalidRangeError(`response.m2reviews[*].rating`, *e.Rating, 1, true, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating > 5 {
				err = goa.InvalidRangeError(`response.m2reviews[*].rating`, *e.Rating, 5, false, err)
			}
		}
		for _, e := range e.Reviewers {
			if e.Bio != nil {
				if len(*e.Bio) > 500 {
					err = goa.InvalidLengthError(`response.m2reviews[*].reviewers[*].bio`, *e.Bio, len(*e.Bio), 500, false, err)
				}
			}
			if e.Email != nil {
				if err2 := goa.ValidateFormat(goa.FormatEmail, *e.Email); err2 != nil {
					err = goa.InvalidFormatError(`response.m2reviews[*].reviewers[*].email`, *e.Email, goa.FormatEmail, err2, err)
				}
			}
		}
	}
	for _, e := range ut.Reviews {
		if e.Comment != nil {
			if len(*e.Comment) < 10 {
				err = goa.InvalidLengthError(`response.reviews[*].comment`, *e.Comment, len(*e.Comment), 10, true, err)
			}
		}
		if e.Comment != nil {
			if len(*e.Comment) > 200 {
				err = goa.InvalidLengthError(`response.reviews[*].comment`, *e.Comment, len(*e.Comment), 200, false, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating < 1 {
				err = goa.InvalidRangeError(`response.reviews[*].rating`, *e.Rating, 1, true, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating > 5 {
				err = goa.InvalidRangeError(`response.reviews[*].rating`, *e.Rating, 5, false, err)
			}
		}
		for _, e := range e.Reviewers {
			if e.Bio != nil {
				if len(*e.Bio) > 500 {
					err = goa.InvalidLengthError(`response.reviews[*].reviewers[*].bio`, *e.Bio, len(*e.Bio), 500, false, err)
				}
			}
			if e.Email != nil {
				if err2 := goa.ValidateFormat(goa.FormatEmail, *e.Email); err2 != nil {
					err = goa.InvalidFormatError(`response.reviews[*].reviewers[*].email`, *e.Email, goa.FormatEmail, err2, err)
				}
			}
		}
	}
	if ut.Title != nil {
		if len(*ut.Title) < 10 {
			err = goa.InvalidLengthError(`response.title`, *ut.Title, len(*ut.Title), 10, true, err)
		}
	}
	if ut.Title != nil {
		if len(*ut.Title) > 200 {
			err = goa.InvalidLengthError(`response.title`, *ut.Title, len(*ut.Title), 200, false, err)
		}
	}
	return
}

// MarshalProposalModel validates and renders an instance of ProposalModel into a interface{}
func MarshalProposalModel(source *ProposalModel, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp15 := map[string]interface{}{
		"UserID":     source.UserID,
		"abstract":   source.Abstract,
		"created_at": source.CreatedAt,
		"deleted_at": source.DeletedAt,
		"detail":     source.Detail,
		"firstname":  source.Firstname,
		"id":         source.ID,
		"title":      source.Title,
		"updated_at": source.UpdatedAt,
		"withdrawn":  source.Withdrawn,
	}
	if source.M2reviews != nil {
		tmp16 := make([]map[string]interface{}, len(source.M2reviews))
		for tmp17, tmp18 := range source.M2reviews {
			tmp16[tmp17], err = MarshalReviewModel(tmp18, err)
		}
		tmp15["m2reviews"] = tmp16
	}
	if source.Reviews != nil {
		tmp19 := make([]map[string]interface{}, len(source.Reviews))
		for tmp20, tmp21 := range source.Reviews {
			tmp19[tmp20], err = MarshalReviewModel(tmp21, err)
		}
		tmp15["reviews"] = tmp19
	}
	target = tmp15
	return
}

// UnmarshalProposalModel unmarshals and validates a raw interface{} into an instance of ProposalModel
func UnmarshalProposalModel(source interface{}, inErr error) (target *ProposalModel, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(ProposalModel)
		if v, ok := val["UserID"]; ok {
			var tmp22 int
			if f, ok := v.(float64); ok {
				tmp22 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.UserID`, v, "int", err)
			}
			target.UserID = &tmp22
		}
		if v, ok := val["abstract"]; ok {
			var tmp23 string
			if val, ok := v.(string); ok {
				tmp23 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Abstract`, v, "string", err)
			}
			target.Abstract = &tmp23
		}
		if v, ok := val["created_at"]; ok {
			var tmp24 interface{}
			tmp24 = v
			target.CreatedAt = &tmp24
		}
		if v, ok := val["deleted_at"]; ok {
			var tmp25 interface{}
			tmp25 = v
			target.DeletedAt = &tmp25
		}
		if v, ok := val["detail"]; ok {
			var tmp26 string
			if val, ok := v.(string); ok {
				tmp26 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Detail`, v, "string", err)
			}
			target.Detail = &tmp26
		}
		if v, ok := val["firstname"]; ok {
			var tmp27 string
			if val, ok := v.(string); ok {
				tmp27 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Firstname`, v, "string", err)
			}
			target.Firstname = &tmp27
		}
		if v, ok := val["id"]; ok {
			var tmp28 string
			if val, ok := v.(string); ok {
				tmp28 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "string", err)
			}
			target.ID = tmp28
		} else {
			err = goa.MissingAttributeError(`load`, "id", err)
		}
		if v, ok := val["m2reviews"]; ok {
			var tmp29 []*ReviewModel
			if val, ok := v.([]interface{}); ok {
				tmp29 = make([]*ReviewModel, len(val))
				for tmp30, v := range val {
					tmp29[tmp30], err = UnmarshalReviewModel(v, err)
				}
			} else {
				err = goa.InvalidAttributeTypeError(`load.M2reviews`, v, "array", err)
			}
			target.M2reviews = tmp29
		}
		if v, ok := val["reviews"]; ok {
			var tmp31 []*ReviewModel
			if val, ok := v.([]interface{}); ok {
				tmp31 = make([]*ReviewModel, len(val))
				for tmp32, v := range val {
					tmp31[tmp32], err = UnmarshalReviewModel(v, err)
				}
			} else {
				err = goa.InvalidAttributeTypeError(`load.Reviews`, v, "array", err)
			}
			target.Reviews = tmp31
		}
		if v, ok := val["title"]; ok {
			var tmp33 string
			if val, ok := v.(string); ok {
				tmp33 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Title`, v, "string", err)
			}
			target.Title = &tmp33
		}
		if v, ok := val["updated_at"]; ok {
			var tmp34 interface{}
			tmp34 = v
			target.UpdatedAt = &tmp34
		}
		if v, ok := val["withdrawn"]; ok {
			var tmp35 bool
			if val, ok := v.(bool); ok {
				tmp35 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Withdrawn`, v, "bool", err)
			}
			target.Withdrawn = &tmp35
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}

// ReviewModel type
type ReviewModel struct {
	ProposalID *int
	UserID     *int
	Comment    *string
	ID         string
	Rating     *int
	Reviewers  []*UserModel
}

// Validate validates the type instance.
func (ut *ReviewModel) Validate() (err error) {
	if ut.ID == "" {
		err = goa.MissingAttributeError(`response`, "id", err)
	}

	if ut.Comment != nil {
		if len(*ut.Comment) < 10 {
			err = goa.InvalidLengthError(`response.comment`, *ut.Comment, len(*ut.Comment), 10, true, err)
		}
	}
	if ut.Comment != nil {
		if len(*ut.Comment) > 200 {
			err = goa.InvalidLengthError(`response.comment`, *ut.Comment, len(*ut.Comment), 200, false, err)
		}
	}
	if ut.Rating != nil {
		if *ut.Rating < 1 {
			err = goa.InvalidRangeError(`response.rating`, *ut.Rating, 1, true, err)
		}
	}
	if ut.Rating != nil {
		if *ut.Rating > 5 {
			err = goa.InvalidRangeError(`response.rating`, *ut.Rating, 5, false, err)
		}
	}
	for _, e := range ut.Reviewers {
		if e.Bio != nil {
			if len(*e.Bio) > 500 {
				err = goa.InvalidLengthError(`response.reviewers[*].bio`, *e.Bio, len(*e.Bio), 500, false, err)
			}
		}
		if e.Email != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *e.Email); err2 != nil {
				err = goa.InvalidFormatError(`response.reviewers[*].email`, *e.Email, goa.FormatEmail, err2, err)
			}
		}
	}
	return
}

// MarshalReviewModel validates and renders an instance of ReviewModel into a interface{}
func MarshalReviewModel(source *ReviewModel, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp36 := map[string]interface{}{
		"ProposalID": source.ProposalID,
		"UserID":     source.UserID,
		"comment":    source.Comment,
		"id":         source.ID,
		"rating":     source.Rating,
	}
	if source.Reviewers != nil {
		tmp37 := make([]map[string]interface{}, len(source.Reviewers))
		for tmp38, tmp39 := range source.Reviewers {
			tmp37[tmp38], err = MarshalUserModel(tmp39, err)
		}
		tmp36["reviewers"] = tmp37
	}
	target = tmp36
	return
}

// UnmarshalReviewModel unmarshals and validates a raw interface{} into an instance of ReviewModel
func UnmarshalReviewModel(source interface{}, inErr error) (target *ReviewModel, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(ReviewModel)
		if v, ok := val["ProposalID"]; ok {
			var tmp40 int
			if f, ok := v.(float64); ok {
				tmp40 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ProposalID`, v, "int", err)
			}
			target.ProposalID = &tmp40
		}
		if v, ok := val["UserID"]; ok {
			var tmp41 int
			if f, ok := v.(float64); ok {
				tmp41 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.UserID`, v, "int", err)
			}
			target.UserID = &tmp41
		}
		if v, ok := val["comment"]; ok {
			var tmp42 string
			if val, ok := v.(string); ok {
				tmp42 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Comment`, v, "string", err)
			}
			target.Comment = &tmp42
		}
		if v, ok := val["id"]; ok {
			var tmp43 string
			if val, ok := v.(string); ok {
				tmp43 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "string", err)
			}
			target.ID = tmp43
		} else {
			err = goa.MissingAttributeError(`load`, "id", err)
		}
		if v, ok := val["rating"]; ok {
			var tmp44 int
			if f, ok := v.(float64); ok {
				tmp44 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.Rating`, v, "int", err)
			}
			target.Rating = &tmp44
		}
		if v, ok := val["reviewers"]; ok {
			var tmp45 []*UserModel
			if val, ok := v.([]interface{}); ok {
				tmp45 = make([]*UserModel, len(val))
				for tmp46, v := range val {
					tmp45[tmp46], err = UnmarshalUserModel(v, err)
				}
			} else {
				err = goa.InvalidAttributeTypeError(`load.Reviewers`, v, "array", err)
			}
			target.Reviewers = tmp45
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}

// UserModel type
type UserModel struct {
	Bio       *string
	City      *string
	Country   *string
	Email     *string
	Firstname *string
	ID        string
	Lastname  *string
	Role      *string
	State     *string
}

// Validate validates the type instance.
func (ut *UserModel) Validate() (err error) {
	if ut.ID == "" {
		err = goa.MissingAttributeError(`response`, "id", err)
	}

	if ut.Bio != nil {
		if len(*ut.Bio) > 500 {
			err = goa.InvalidLengthError(`response.bio`, *ut.Bio, len(*ut.Bio), 500, false, err)
		}
	}
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2, err)
		}
	}
	return
}

// MarshalUserModel validates and renders an instance of UserModel into a interface{}
func MarshalUserModel(source *UserModel, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp47 := map[string]interface{}{
		"bio":       source.Bio,
		"city":      source.City,
		"country":   source.Country,
		"email":     source.Email,
		"firstname": source.Firstname,
		"id":        source.ID,
		"lastname":  source.Lastname,
		"role":      source.Role,
		"state":     source.State,
	}
	target = tmp47
	return
}

// UnmarshalUserModel unmarshals and validates a raw interface{} into an instance of UserModel
func UnmarshalUserModel(source interface{}, inErr error) (target *UserModel, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(UserModel)
		if v, ok := val["bio"]; ok {
			var tmp48 string
			if val, ok := v.(string); ok {
				tmp48 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Bio`, v, "string", err)
			}
			target.Bio = &tmp48
		}
		if v, ok := val["city"]; ok {
			var tmp49 string
			if val, ok := v.(string); ok {
				tmp49 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.City`, v, "string", err)
			}
			target.City = &tmp49
		}
		if v, ok := val["country"]; ok {
			var tmp50 string
			if val, ok := v.(string); ok {
				tmp50 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Country`, v, "string", err)
			}
			target.Country = &tmp50
		}
		if v, ok := val["email"]; ok {
			var tmp51 string
			if val, ok := v.(string); ok {
				tmp51 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Email`, v, "string", err)
			}
			target.Email = &tmp51
		}
		if v, ok := val["firstname"]; ok {
			var tmp52 string
			if val, ok := v.(string); ok {
				tmp52 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Firstname`, v, "string", err)
			}
			target.Firstname = &tmp52
		}
		if v, ok := val["id"]; ok {
			var tmp53 string
			if val, ok := v.(string); ok {
				tmp53 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "string", err)
			}
			target.ID = tmp53
		} else {
			err = goa.MissingAttributeError(`load`, "id", err)
		}
		if v, ok := val["lastname"]; ok {
			var tmp54 string
			if val, ok := v.(string); ok {
				tmp54 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Lastname`, v, "string", err)
			}
			target.Lastname = &tmp54
		}
		if v, ok := val["role"]; ok {
			var tmp55 string
			if val, ok := v.(string); ok {
				tmp55 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Role`, v, "string", err)
			}
			target.Role = &tmp55
		}
		if v, ok := val["state"]; ok {
			var tmp56 string
			if val, ok := v.(string); ok {
				tmp56 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.State`, v, "string", err)
			}
			target.State = &tmp56
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}
