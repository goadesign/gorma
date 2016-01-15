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
	Abstract  *string
	CreatedAt *string
	DeletedAt *string
	Detail    *string
	FirstName *string
	ID        int
	M2reviews *string
	Reviews   *string
	Title     *string
	UpdatedAt *string
	UserId    *int
	Withdrawn *bool
}

// Validate validates the type instance.
func (ut *ProposalModel) Validate() (err error) {

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
	if ut.FirstName != nil {
		if len(*ut.FirstName) < 2 {
			err = goa.InvalidLengthError(`response.first_name`, *ut.FirstName, len(*ut.FirstName), 2, true, err)
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
		"abstract":   source.Abstract,
		"created_at": source.CreatedAt,
		"deleted_at": source.DeletedAt,
		"detail":     source.Detail,
		"first_name": source.FirstName,
		"id":         source.ID,
		"m2reviews":  source.M2reviews,
		"reviews":    source.Reviews,
		"title":      source.Title,
		"updated_at": source.UpdatedAt,
		"user_id":    source.UserId,
		"withdrawn":  source.Withdrawn,
	}
	target = tmp15
	return
}

// UnmarshalProposalModel unmarshals and validates a raw interface{} into an instance of ProposalModel
func UnmarshalProposalModel(source interface{}, inErr error) (target *ProposalModel, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(ProposalModel)
		if v, ok := val["abstract"]; ok {
			var tmp16 string
			if val, ok := v.(string); ok {
				tmp16 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Abstract`, v, "string", err)
			}
			target.Abstract = &tmp16
		}
		if v, ok := val["created_at"]; ok {
			var tmp17 string
			if val, ok := v.(string); ok {
				tmp17 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.CreatedAt`, v, "string", err)
			}
			target.CreatedAt = &tmp17
		}
		if v, ok := val["deleted_at"]; ok {
			var tmp18 string
			if val, ok := v.(string); ok {
				tmp18 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.DeletedAt`, v, "string", err)
			}
			target.DeletedAt = &tmp18
		}
		if v, ok := val["detail"]; ok {
			var tmp19 string
			if val, ok := v.(string); ok {
				tmp19 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Detail`, v, "string", err)
			}
			target.Detail = &tmp19
		}
		if v, ok := val["first_name"]; ok {
			var tmp20 string
			if val, ok := v.(string); ok {
				tmp20 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.FirstName`, v, "string", err)
			}
			target.FirstName = &tmp20
		}
		if v, ok := val["id"]; ok {
			var tmp21 int
			if f, ok := v.(float64); ok {
				tmp21 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = tmp21
		} else {
			err = goa.MissingAttributeError(`load`, "id", err)
		}
		if v, ok := val["m2reviews"]; ok {
			var tmp22 string
			if val, ok := v.(string); ok {
				tmp22 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.M2reviews`, v, "string", err)
			}
			target.M2reviews = &tmp22
		}
		if v, ok := val["reviews"]; ok {
			var tmp23 string
			if val, ok := v.(string); ok {
				tmp23 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Reviews`, v, "string", err)
			}
			target.Reviews = &tmp23
		}
		if v, ok := val["title"]; ok {
			var tmp24 string
			if val, ok := v.(string); ok {
				tmp24 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Title`, v, "string", err)
			}
			target.Title = &tmp24
		}
		if v, ok := val["updated_at"]; ok {
			var tmp25 string
			if val, ok := v.(string); ok {
				tmp25 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.UpdatedAt`, v, "string", err)
			}
			target.UpdatedAt = &tmp25
		}
		if v, ok := val["user_id"]; ok {
			var tmp26 int
			if f, ok := v.(float64); ok {
				tmp26 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.UserId`, v, "int", err)
			}
			target.UserId = &tmp26
		}
		if v, ok := val["withdrawn"]; ok {
			var tmp27 bool
			if val, ok := v.(bool); ok {
				tmp27 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Withdrawn`, v, "bool", err)
			}
			target.Withdrawn = &tmp27
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
	User       *string
	Comment    *string
	ID         int
	ProposalId *int
	Rating     *int
	Reviewers  *string
}

// Validate validates the type instance.
func (ut *ReviewModel) Validate() (err error) {

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
	return
}

// MarshalReviewModel validates and renders an instance of ReviewModel into a interface{}
func MarshalReviewModel(source *ReviewModel, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp28 := map[string]interface{}{
		"User":        source.User,
		"comment":     source.Comment,
		"id":          source.ID,
		"proposal_id": source.ProposalId,
		"rating":      source.Rating,
		"reviewers":   source.Reviewers,
	}
	target = tmp28
	return
}

// UnmarshalReviewModel unmarshals and validates a raw interface{} into an instance of ReviewModel
func UnmarshalReviewModel(source interface{}, inErr error) (target *ReviewModel, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(ReviewModel)
		if v, ok := val["User"]; ok {
			var tmp29 string
			if val, ok := v.(string); ok {
				tmp29 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.User`, v, "string", err)
			}
			target.User = &tmp29
		}
		if v, ok := val["comment"]; ok {
			var tmp30 string
			if val, ok := v.(string); ok {
				tmp30 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Comment`, v, "string", err)
			}
			target.Comment = &tmp30
		}
		if v, ok := val["id"]; ok {
			var tmp31 int
			if f, ok := v.(float64); ok {
				tmp31 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = tmp31
		} else {
			err = goa.MissingAttributeError(`load`, "id", err)
		}
		if v, ok := val["proposal_id"]; ok {
			var tmp32 int
			if f, ok := v.(float64); ok {
				tmp32 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ProposalId`, v, "int", err)
			}
			target.ProposalId = &tmp32
		}
		if v, ok := val["rating"]; ok {
			var tmp33 int
			if f, ok := v.(float64); ok {
				tmp33 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.Rating`, v, "int", err)
			}
			target.Rating = &tmp33
		}
		if v, ok := val["reviewers"]; ok {
			var tmp34 string
			if val, ok := v.(string); ok {
				tmp34 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Reviewers`, v, "string", err)
			}
			target.Reviewers = &tmp34
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
	CreatedAt *string
	Email     *string
	// First name Description
	FirstName *string
	ID        int
	LastName  *string
	Role      string
	State     *string
	UpdatedAt *string
}

// Validate validates the type instance.
func (ut *UserModel) Validate() (err error) {

	if ut.Role == "" {
		err = goa.MissingAttributeError(`response`, "role", err)
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
	tmp35 := map[string]interface{}{
		"bio":        source.Bio,
		"city":       source.City,
		"country":    source.Country,
		"created_at": source.CreatedAt,
		"email":      source.Email,
		"first_name": source.FirstName,
		"id":         source.ID,
		"last_name":  source.LastName,
		"role":       source.Role,
		"state":      source.State,
		"updated_at": source.UpdatedAt,
	}
	target = tmp35
	return
}

// UnmarshalUserModel unmarshals and validates a raw interface{} into an instance of UserModel
func UnmarshalUserModel(source interface{}, inErr error) (target *UserModel, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(UserModel)
		if v, ok := val["bio"]; ok {
			var tmp36 string
			if val, ok := v.(string); ok {
				tmp36 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Bio`, v, "string", err)
			}
			target.Bio = &tmp36
		}
		if v, ok := val["city"]; ok {
			var tmp37 string
			if val, ok := v.(string); ok {
				tmp37 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.City`, v, "string", err)
			}
			target.City = &tmp37
		}
		if v, ok := val["country"]; ok {
			var tmp38 string
			if val, ok := v.(string); ok {
				tmp38 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Country`, v, "string", err)
			}
			target.Country = &tmp38
		}
		if v, ok := val["created_at"]; ok {
			var tmp39 string
			if val, ok := v.(string); ok {
				tmp39 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.CreatedAt`, v, "string", err)
			}
			target.CreatedAt = &tmp39
		}
		if v, ok := val["email"]; ok {
			var tmp40 string
			if val, ok := v.(string); ok {
				tmp40 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Email`, v, "string", err)
			}
			target.Email = &tmp40
		}
		if v, ok := val["first_name"]; ok {
			var tmp41 string
			if val, ok := v.(string); ok {
				tmp41 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.FirstName`, v, "string", err)
			}
			target.FirstName = &tmp41
		}
		if v, ok := val["id"]; ok {
			var tmp42 int
			if f, ok := v.(float64); ok {
				tmp42 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = tmp42
		} else {
			err = goa.MissingAttributeError(`load`, "id", err)
		}
		if v, ok := val["last_name"]; ok {
			var tmp43 string
			if val, ok := v.(string); ok {
				tmp43 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.LastName`, v, "string", err)
			}
			target.LastName = &tmp43
		}
		if v, ok := val["role"]; ok {
			var tmp44 string
			if val, ok := v.(string); ok {
				tmp44 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Role`, v, "string", err)
			}
			target.Role = tmp44
		} else {
			err = goa.MissingAttributeError(`load`, "role", err)
		}
		if v, ok := val["state"]; ok {
			var tmp45 string
			if val, ok := v.(string); ok {
				tmp45 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.State`, v, "string", err)
			}
			target.State = &tmp45
		}
		if v, ok := val["updated_at"]; ok {
			var tmp46 string
			if val, ok := v.(string); ok {
				tmp46 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.UpdatedAt`, v, "string", err)
			}
			target.UpdatedAt = &tmp46
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	if target != nil {
		err = goa.ReportError(err, target.Validate())
	}
	return
}
