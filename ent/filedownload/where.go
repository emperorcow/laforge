// Code generated by entc, DO NOT EDIT.

package filedownload

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/gen0cide/laforge/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// SourceType applies equality check predicate on the "source_type" field. It's identical to SourceTypeEQ.
func SourceType(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSourceType), v))
	})
}

// Source applies equality check predicate on the "source" field. It's identical to SourceEQ.
func Source(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSource), v))
	})
}

// Destination applies equality check predicate on the "destination" field. It's identical to DestinationEQ.
func Destination(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDestination), v))
	})
}

// Template applies equality check predicate on the "template" field. It's identical to TemplateEQ.
func Template(v bool) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTemplate), v))
	})
}

// Mode applies equality check predicate on the "mode" field. It's identical to ModeEQ.
func Mode(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMode), v))
	})
}

// Disabled applies equality check predicate on the "disabled" field. It's identical to DisabledEQ.
func Disabled(v bool) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDisabled), v))
	})
}

// Md5 applies equality check predicate on the "md5" field. It's identical to Md5EQ.
func Md5(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMd5), v))
	})
}

// AbsPath applies equality check predicate on the "abs_path" field. It's identical to AbsPathEQ.
func AbsPath(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAbsPath), v))
	})
}

// SourceTypeEQ applies the EQ predicate on the "source_type" field.
func SourceTypeEQ(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSourceType), v))
	})
}

// SourceTypeNEQ applies the NEQ predicate on the "source_type" field.
func SourceTypeNEQ(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSourceType), v))
	})
}

// SourceTypeIn applies the In predicate on the "source_type" field.
func SourceTypeIn(vs ...string) predicate.FileDownload {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FileDownload(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSourceType), v...))
	})
}

// SourceTypeNotIn applies the NotIn predicate on the "source_type" field.
func SourceTypeNotIn(vs ...string) predicate.FileDownload {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FileDownload(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSourceType), v...))
	})
}

// SourceTypeGT applies the GT predicate on the "source_type" field.
func SourceTypeGT(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSourceType), v))
	})
}

// SourceTypeGTE applies the GTE predicate on the "source_type" field.
func SourceTypeGTE(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSourceType), v))
	})
}

// SourceTypeLT applies the LT predicate on the "source_type" field.
func SourceTypeLT(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSourceType), v))
	})
}

// SourceTypeLTE applies the LTE predicate on the "source_type" field.
func SourceTypeLTE(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSourceType), v))
	})
}

// SourceTypeContains applies the Contains predicate on the "source_type" field.
func SourceTypeContains(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSourceType), v))
	})
}

// SourceTypeHasPrefix applies the HasPrefix predicate on the "source_type" field.
func SourceTypeHasPrefix(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSourceType), v))
	})
}

// SourceTypeHasSuffix applies the HasSuffix predicate on the "source_type" field.
func SourceTypeHasSuffix(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSourceType), v))
	})
}

// SourceTypeEqualFold applies the EqualFold predicate on the "source_type" field.
func SourceTypeEqualFold(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSourceType), v))
	})
}

// SourceTypeContainsFold applies the ContainsFold predicate on the "source_type" field.
func SourceTypeContainsFold(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSourceType), v))
	})
}

// SourceEQ applies the EQ predicate on the "source" field.
func SourceEQ(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSource), v))
	})
}

// SourceNEQ applies the NEQ predicate on the "source" field.
func SourceNEQ(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSource), v))
	})
}

// SourceIn applies the In predicate on the "source" field.
func SourceIn(vs ...string) predicate.FileDownload {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FileDownload(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSource), v...))
	})
}

// SourceNotIn applies the NotIn predicate on the "source" field.
func SourceNotIn(vs ...string) predicate.FileDownload {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FileDownload(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSource), v...))
	})
}

// SourceGT applies the GT predicate on the "source" field.
func SourceGT(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSource), v))
	})
}

// SourceGTE applies the GTE predicate on the "source" field.
func SourceGTE(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSource), v))
	})
}

// SourceLT applies the LT predicate on the "source" field.
func SourceLT(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSource), v))
	})
}

// SourceLTE applies the LTE predicate on the "source" field.
func SourceLTE(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSource), v))
	})
}

// SourceContains applies the Contains predicate on the "source" field.
func SourceContains(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSource), v))
	})
}

// SourceHasPrefix applies the HasPrefix predicate on the "source" field.
func SourceHasPrefix(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSource), v))
	})
}

// SourceHasSuffix applies the HasSuffix predicate on the "source" field.
func SourceHasSuffix(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSource), v))
	})
}

// SourceEqualFold applies the EqualFold predicate on the "source" field.
func SourceEqualFold(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSource), v))
	})
}

// SourceContainsFold applies the ContainsFold predicate on the "source" field.
func SourceContainsFold(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSource), v))
	})
}

// DestinationEQ applies the EQ predicate on the "destination" field.
func DestinationEQ(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDestination), v))
	})
}

// DestinationNEQ applies the NEQ predicate on the "destination" field.
func DestinationNEQ(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDestination), v))
	})
}

// DestinationIn applies the In predicate on the "destination" field.
func DestinationIn(vs ...string) predicate.FileDownload {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FileDownload(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDestination), v...))
	})
}

// DestinationNotIn applies the NotIn predicate on the "destination" field.
func DestinationNotIn(vs ...string) predicate.FileDownload {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FileDownload(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDestination), v...))
	})
}

// DestinationGT applies the GT predicate on the "destination" field.
func DestinationGT(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDestination), v))
	})
}

// DestinationGTE applies the GTE predicate on the "destination" field.
func DestinationGTE(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDestination), v))
	})
}

// DestinationLT applies the LT predicate on the "destination" field.
func DestinationLT(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDestination), v))
	})
}

// DestinationLTE applies the LTE predicate on the "destination" field.
func DestinationLTE(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDestination), v))
	})
}

// DestinationContains applies the Contains predicate on the "destination" field.
func DestinationContains(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDestination), v))
	})
}

// DestinationHasPrefix applies the HasPrefix predicate on the "destination" field.
func DestinationHasPrefix(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDestination), v))
	})
}

// DestinationHasSuffix applies the HasSuffix predicate on the "destination" field.
func DestinationHasSuffix(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDestination), v))
	})
}

// DestinationEqualFold applies the EqualFold predicate on the "destination" field.
func DestinationEqualFold(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDestination), v))
	})
}

// DestinationContainsFold applies the ContainsFold predicate on the "destination" field.
func DestinationContainsFold(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDestination), v))
	})
}

// TemplateEQ applies the EQ predicate on the "template" field.
func TemplateEQ(v bool) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTemplate), v))
	})
}

// TemplateNEQ applies the NEQ predicate on the "template" field.
func TemplateNEQ(v bool) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTemplate), v))
	})
}

// ModeEQ applies the EQ predicate on the "mode" field.
func ModeEQ(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMode), v))
	})
}

// ModeNEQ applies the NEQ predicate on the "mode" field.
func ModeNEQ(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMode), v))
	})
}

// ModeIn applies the In predicate on the "mode" field.
func ModeIn(vs ...string) predicate.FileDownload {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FileDownload(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMode), v...))
	})
}

// ModeNotIn applies the NotIn predicate on the "mode" field.
func ModeNotIn(vs ...string) predicate.FileDownload {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FileDownload(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMode), v...))
	})
}

// ModeGT applies the GT predicate on the "mode" field.
func ModeGT(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMode), v))
	})
}

// ModeGTE applies the GTE predicate on the "mode" field.
func ModeGTE(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMode), v))
	})
}

// ModeLT applies the LT predicate on the "mode" field.
func ModeLT(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMode), v))
	})
}

// ModeLTE applies the LTE predicate on the "mode" field.
func ModeLTE(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMode), v))
	})
}

// ModeContains applies the Contains predicate on the "mode" field.
func ModeContains(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMode), v))
	})
}

// ModeHasPrefix applies the HasPrefix predicate on the "mode" field.
func ModeHasPrefix(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMode), v))
	})
}

// ModeHasSuffix applies the HasSuffix predicate on the "mode" field.
func ModeHasSuffix(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMode), v))
	})
}

// ModeEqualFold applies the EqualFold predicate on the "mode" field.
func ModeEqualFold(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMode), v))
	})
}

// ModeContainsFold applies the ContainsFold predicate on the "mode" field.
func ModeContainsFold(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMode), v))
	})
}

// DisabledEQ applies the EQ predicate on the "disabled" field.
func DisabledEQ(v bool) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDisabled), v))
	})
}

// DisabledNEQ applies the NEQ predicate on the "disabled" field.
func DisabledNEQ(v bool) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDisabled), v))
	})
}

// Md5EQ applies the EQ predicate on the "md5" field.
func Md5EQ(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMd5), v))
	})
}

// Md5NEQ applies the NEQ predicate on the "md5" field.
func Md5NEQ(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMd5), v))
	})
}

// Md5In applies the In predicate on the "md5" field.
func Md5In(vs ...string) predicate.FileDownload {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FileDownload(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMd5), v...))
	})
}

// Md5NotIn applies the NotIn predicate on the "md5" field.
func Md5NotIn(vs ...string) predicate.FileDownload {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FileDownload(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMd5), v...))
	})
}

// Md5GT applies the GT predicate on the "md5" field.
func Md5GT(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMd5), v))
	})
}

// Md5GTE applies the GTE predicate on the "md5" field.
func Md5GTE(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMd5), v))
	})
}

// Md5LT applies the LT predicate on the "md5" field.
func Md5LT(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMd5), v))
	})
}

// Md5LTE applies the LTE predicate on the "md5" field.
func Md5LTE(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMd5), v))
	})
}

// Md5Contains applies the Contains predicate on the "md5" field.
func Md5Contains(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMd5), v))
	})
}

// Md5HasPrefix applies the HasPrefix predicate on the "md5" field.
func Md5HasPrefix(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMd5), v))
	})
}

// Md5HasSuffix applies the HasSuffix predicate on the "md5" field.
func Md5HasSuffix(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMd5), v))
	})
}

// Md5EqualFold applies the EqualFold predicate on the "md5" field.
func Md5EqualFold(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMd5), v))
	})
}

// Md5ContainsFold applies the ContainsFold predicate on the "md5" field.
func Md5ContainsFold(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMd5), v))
	})
}

// AbsPathEQ applies the EQ predicate on the "abs_path" field.
func AbsPathEQ(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAbsPath), v))
	})
}

// AbsPathNEQ applies the NEQ predicate on the "abs_path" field.
func AbsPathNEQ(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAbsPath), v))
	})
}

// AbsPathIn applies the In predicate on the "abs_path" field.
func AbsPathIn(vs ...string) predicate.FileDownload {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FileDownload(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAbsPath), v...))
	})
}

// AbsPathNotIn applies the NotIn predicate on the "abs_path" field.
func AbsPathNotIn(vs ...string) predicate.FileDownload {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FileDownload(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAbsPath), v...))
	})
}

// AbsPathGT applies the GT predicate on the "abs_path" field.
func AbsPathGT(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAbsPath), v))
	})
}

// AbsPathGTE applies the GTE predicate on the "abs_path" field.
func AbsPathGTE(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAbsPath), v))
	})
}

// AbsPathLT applies the LT predicate on the "abs_path" field.
func AbsPathLT(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAbsPath), v))
	})
}

// AbsPathLTE applies the LTE predicate on the "abs_path" field.
func AbsPathLTE(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAbsPath), v))
	})
}

// AbsPathContains applies the Contains predicate on the "abs_path" field.
func AbsPathContains(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAbsPath), v))
	})
}

// AbsPathHasPrefix applies the HasPrefix predicate on the "abs_path" field.
func AbsPathHasPrefix(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAbsPath), v))
	})
}

// AbsPathHasSuffix applies the HasSuffix predicate on the "abs_path" field.
func AbsPathHasSuffix(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAbsPath), v))
	})
}

// AbsPathEqualFold applies the EqualFold predicate on the "abs_path" field.
func AbsPathEqualFold(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAbsPath), v))
	})
}

// AbsPathContainsFold applies the ContainsFold predicate on the "abs_path" field.
func AbsPathContainsFold(v string) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAbsPath), v))
	})
}

// HasTag applies the HasEdge predicate on the "tag" edge.
func HasTag() predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TagTable, TagColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagWith applies the HasEdge predicate on the "tag" edge with a given conditions (other predicates).
func HasTagWith(preds ...predicate.Tag) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TagTable, TagColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.FileDownload) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.FileDownload) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.FileDownload) predicate.FileDownload {
	return predicate.FileDownload(func(s *sql.Selector) {
		p(s.Not())
	})
}
