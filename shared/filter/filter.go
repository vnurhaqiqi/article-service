package filter

import (
	"fmt"
	"strings"
)

type Filter struct {
	conditions
	order
	limit
	group
	joins
}

type FilterFunc func(f *Filter)

func New(filterFunc ...FilterFunc) *Filter {
	f := &Filter{}
	for _, filter := range filterFunc {
		filter(f)
	}

	return f
}

// QueryClause that generate clause and its arguments
// the operator has three conditions "AND", "OR", and "NOT".
func (f *Filter) QueryClause(operator string) ([]interface{}, string) {
	var (
		args    = make([]interface{}, 0)
		clauses = make([]string, 0)
	)

	if len(f.conditions) == 0 {
		return args, ""
	}

	for _, c := range f.conditions {
		clauses = append(clauses, c.buildQuery())
		if c.Value != nil {
			args = append(args, c.Value)
		}
	}

	return args, strings.Join(clauses, fmt.Sprintf(" %s ", operator))
}

func (f *Filter) Joins() string {
	if len(f.joins) == 0 {
		return ""
	}

	states := make([]string, 0)
	for _, join := range f.joins {
		states = append(states, fmt.Sprintf(" %s %s ON %s = %s", join.Args, join.Table, join.ForeignKey, join.ReferenceKey))
	}

	return strings.Join(states, "\n")
}

// Paginate function that generate limit query and its offset.
func (f Filter) Paginate() string {
	if f.limit.Page == 0 && f.limit.Size == 0 {
		return ""
	}

	if f.limit.Size > 0 && f.limit.Page == 0 {
		return fmt.Sprintf(" LIMIT %d", f.limit.Size)
	}

	offset := (f.limit.Page - 1) * f.limit.Size

	return fmt.Sprintf(" LIMIT %d OFFSET %d", f.limit.Size, offset)
}

// SortBy function that generate Order By query.
func (o Filter) SortBy() string {
	if len(o.order.Columns) == 0 {
		return ""
	}

	return fmt.Sprintf(" ORDER BY %s %s", strings.Join(o.order.Columns, ","), o.Direction)
}

// Group function that generate Group By column.
func (o Filter) Group() string {
	if len(o.group.Columns) == 0 {
		return ""
	}

	return fmt.Sprintf(" GROUP BY %s", strings.Join(o.group.Columns, ","))
}

type condition struct {
	Field    string
	Operator string
	Value    interface{}
}

type conditions []condition

func (filter *condition) buildQuery() string {
	var clause string

	switch strings.ToLower(filter.Operator) {
	case "in":
		clause = fmt.Sprintf("%s IN(?)", filter.Field)
	case "not in":
		clause = fmt.Sprintf("%s NOT IN(?)", filter.Field)
	case "is null":
		clause = fmt.Sprintf("%s IS NULL", filter.Field)
	case "like":
		clause = fmt.Sprintf("%s LIKE '%%%v%%'", filter.Field, filter.Value)
	case "match":
		clause = fmt.Sprintf("MATCH%s AGAINST('+%%%v%%*' IN BOOLEAN MODE)", filter.Field, filter.Value)
	default:
		clause = fmt.Sprintf("%s %s ?", filter.Field, filter.Operator)
	}

	return clause
}

type order struct {
	Direction string
	Columns   []string
}

type limit struct {
	Size int
	Page int
}

type group struct {
	Columns []string
}

type join struct {
	Args         string
	Table        string
	ForeignKey   string
	ReferenceKey string
}

type joins []join

// Named Query

type NamedFilter struct {
	namedConditions
}

type NamedFilterFunc func(f *NamedFilter)

func NewNamedStmt(namedFilterFunc ...NamedFilterFunc) *NamedFilter {
	f := &NamedFilter{}
	for _, filter := range namedFilterFunc {
		filter(f)
	}

	return f
}

type namedCondition struct {
	Field    string
	Operator string
	Value    interface{}
}

type namedConditions []namedCondition

func (f *NamedFilter) NamedQueryClause(operator string) string {
	var (
		clauses = make([]string, 0)
	)

	if len(f.namedConditions) == 0 {
		return ""
	}

	for _, c := range f.namedConditions {
		clauses = append(clauses, c.namedBuildQuery())
	}

	return strings.Join(clauses, fmt.Sprintf(" %s ", operator))
}

func (filter namedCondition) namedBuildQuery() string {
	var clause string

	switch strings.ToLower(filter.Operator) {
	case "in":
		clause = fmt.Sprintf("%s IN(:%s)", filter.Field, filter.Field)
	case "not in":
		clause = fmt.Sprintf("%s NOT IN(:%s)", filter.Field, filter.Field)
	default:
		clause = fmt.Sprintf("%s %s :%s", filter.Field, filter.Operator, filter.Field)
	}

	return clause
}
