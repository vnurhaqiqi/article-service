package filter

import "strings"

func Where(field, operator string, value interface{}) FilterFunc {
	return func(f *Filter) {
		f.conditions = append(f.conditions, condition{
			Field:    field,
			Operator: operator,
			Value:    value,
		})
	}
}

func Equal(field string, value interface{}) FilterFunc {
	return func(f *Filter) {
		f.conditions = append(f.conditions, condition{
			Field:    field,
			Operator: "=",
			Value:    value,
		})
	}
}

func In(field string, value interface{}) FilterFunc {
	return func(f *Filter) {
		f.conditions = append(f.conditions, condition{
			Field:    field,
			Operator: "IN",
			Value:    value,
		})
	}
}

func NotIn(field string, value interface{}) FilterFunc {
	return func(f *Filter) {
		f.conditions = append(f.conditions, condition{
			Field:    field,
			Operator: "NOT IN",
			Value:    value,
		})
	}
}

func IsNull(field string) FilterFunc {
	return func(f *Filter) {
		f.conditions = append(f.conditions, condition{
			Field:    field,
			Operator: "IS NULL",
		})
	}
}

func IsNotNull(field string) FilterFunc {
	return func(f *Filter) {
		f.conditions = append(f.conditions, condition{
			Field:    field,
			Operator: "IS NOT NULL",
		})
	}
}

func Like(field string, value interface{}) FilterFunc {
	return func(f *Filter) {
		f.conditions = append(f.conditions, condition{
			Field:    field,
			Operator: "LIKE",
			Value:    value,
		})
	}
}

func Match(fields string, value interface{}) FilterFunc {
	return func(f *Filter) {
		f.conditions = append(f.conditions, condition{
			Field:    fields,
			Operator: "MATCH",
			Value:    value,
		})
	}
}

func OrderBy(direction string, columns ...string) FilterFunc {
	return func(f *Filter) {
		f.order = order{
			Direction: strings.ToUpper(direction),
			Columns:   columns,
		}
	}
}

func GroupBy(columns []string) FilterFunc {
	return func(f *Filter) {
		f.group.Columns = append(f.group.Columns, columns...)
	}
}

func Limit(size int) FilterFunc {
	return func(f *Filter) {
		f.limit.Size = size
	}
}

func Paginate(size int, page int) FilterFunc {
	return func(f *Filter) {
		f.limit.Size = size
		f.limit.Page = page
	}
}

func InnerJoin(tableName, foreignKey, referenceKey string) FilterFunc {
	return func(f *Filter) {
		f.joins = append(f.joins, join{
			Args:         "INNER JOIN",
			Table:        tableName,
			ForeignKey:   foreignKey,
			ReferenceKey: referenceKey,
		})
	}
}

// Named Exec
func NamedEqual(field string) NamedFilterFunc {
	return func(f *NamedFilter) {
		f.namedConditions = append(f.namedConditions, namedCondition{
			Field:    field,
			Operator: "=",
		})
	}
}
