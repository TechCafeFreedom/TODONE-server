// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// TodosLabel is an object representing the database table.
type TodosLabel struct {
	ID      int `boil:"id" json:"id" toml:"id" yaml:"id"`
	TodoID  int `boil:"todo_id" json:"todo_id" toml:"todo_id" yaml:"todo_id"`
	LabelID int `boil:"label_id" json:"label_id" toml:"label_id" yaml:"label_id"`

	R *todosLabelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L todosLabelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TodosLabelColumns = struct {
	ID      string
	TodoID  string
	LabelID string
}{
	ID:      "id",
	TodoID:  "todo_id",
	LabelID: "label_id",
}

// Generated where

var TodosLabelWhere = struct {
	ID      whereHelperint
	TodoID  whereHelperint
	LabelID whereHelperint
}{
	ID:      whereHelperint{field: "`todos_labels`.`id`"},
	TodoID:  whereHelperint{field: "`todos_labels`.`todo_id`"},
	LabelID: whereHelperint{field: "`todos_labels`.`label_id`"},
}

// TodosLabelRels is where relationship names are stored.
var TodosLabelRels = struct {
	Todo  string
	Label string
}{
	Todo:  "Todo",
	Label: "Label",
}

// todosLabelR is where relationships are stored.
type todosLabelR struct {
	Todo  *Todo
	Label *Label
}

// NewStruct creates a new relationship struct
func (*todosLabelR) NewStruct() *todosLabelR {
	return &todosLabelR{}
}

// todosLabelL is where Load methods for each relationship are stored.
type todosLabelL struct{}

var (
	todosLabelAllColumns            = []string{"id", "todo_id", "label_id"}
	todosLabelColumnsWithoutDefault = []string{"todo_id", "label_id"}
	todosLabelColumnsWithDefault    = []string{"id"}
	todosLabelPrimaryKeyColumns     = []string{"id"}
)

type (
	// TodosLabelSlice is an alias for a slice of pointers to TodosLabel.
	// This should generally be used opposed to []TodosLabel.
	TodosLabelSlice []*TodosLabel
	// TodosLabelHook is the signature for custom TodosLabel hook methods
	TodosLabelHook func(context.Context, boil.ContextExecutor, *TodosLabel) error

	todosLabelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	todosLabelType                 = reflect.TypeOf(&TodosLabel{})
	todosLabelMapping              = queries.MakeStructMapping(todosLabelType)
	todosLabelPrimaryKeyMapping, _ = queries.BindMapping(todosLabelType, todosLabelMapping, todosLabelPrimaryKeyColumns)
	todosLabelInsertCacheMut       sync.RWMutex
	todosLabelInsertCache          = make(map[string]insertCache)
	todosLabelUpdateCacheMut       sync.RWMutex
	todosLabelUpdateCache          = make(map[string]updateCache)
	todosLabelUpsertCacheMut       sync.RWMutex
	todosLabelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var todosLabelBeforeInsertHooks []TodosLabelHook
var todosLabelBeforeUpdateHooks []TodosLabelHook
var todosLabelBeforeDeleteHooks []TodosLabelHook
var todosLabelBeforeUpsertHooks []TodosLabelHook

var todosLabelAfterInsertHooks []TodosLabelHook
var todosLabelAfterSelectHooks []TodosLabelHook
var todosLabelAfterUpdateHooks []TodosLabelHook
var todosLabelAfterDeleteHooks []TodosLabelHook
var todosLabelAfterUpsertHooks []TodosLabelHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TodosLabel) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range todosLabelBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TodosLabel) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range todosLabelBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TodosLabel) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range todosLabelBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TodosLabel) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range todosLabelBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TodosLabel) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range todosLabelAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TodosLabel) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range todosLabelAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TodosLabel) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range todosLabelAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TodosLabel) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range todosLabelAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TodosLabel) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range todosLabelAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTodosLabelHook registers your hook function for all future operations.
func AddTodosLabelHook(hookPoint boil.HookPoint, todosLabelHook TodosLabelHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		todosLabelBeforeInsertHooks = append(todosLabelBeforeInsertHooks, todosLabelHook)
	case boil.BeforeUpdateHook:
		todosLabelBeforeUpdateHooks = append(todosLabelBeforeUpdateHooks, todosLabelHook)
	case boil.BeforeDeleteHook:
		todosLabelBeforeDeleteHooks = append(todosLabelBeforeDeleteHooks, todosLabelHook)
	case boil.BeforeUpsertHook:
		todosLabelBeforeUpsertHooks = append(todosLabelBeforeUpsertHooks, todosLabelHook)
	case boil.AfterInsertHook:
		todosLabelAfterInsertHooks = append(todosLabelAfterInsertHooks, todosLabelHook)
	case boil.AfterSelectHook:
		todosLabelAfterSelectHooks = append(todosLabelAfterSelectHooks, todosLabelHook)
	case boil.AfterUpdateHook:
		todosLabelAfterUpdateHooks = append(todosLabelAfterUpdateHooks, todosLabelHook)
	case boil.AfterDeleteHook:
		todosLabelAfterDeleteHooks = append(todosLabelAfterDeleteHooks, todosLabelHook)
	case boil.AfterUpsertHook:
		todosLabelAfterUpsertHooks = append(todosLabelAfterUpsertHooks, todosLabelHook)
	}
}

// One returns a single todosLabel record from the query.
func (q todosLabelQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TodosLabel, error) {
	o := &TodosLabel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for todos_labels")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TodosLabel records from the query.
func (q todosLabelQuery) All(ctx context.Context, exec boil.ContextExecutor) (TodosLabelSlice, error) {
	var o []*TodosLabel

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to TodosLabel slice")
	}

	if len(todosLabelAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TodosLabel records in the query.
func (q todosLabelQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count todos_labels rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q todosLabelQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if todos_labels exists")
	}

	return count > 0, nil
}

// Todo pointed to by the foreign key.
func (o *TodosLabel) Todo(mods ...qm.QueryMod) todoQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.TodoID),
	}

	queryMods = append(queryMods, mods...)

	query := Todos(queryMods...)
	queries.SetFrom(query.Query, "`todos`")

	return query
}

// Label pointed to by the foreign key.
func (o *TodosLabel) Label(mods ...qm.QueryMod) labelQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.LabelID),
	}

	queryMods = append(queryMods, mods...)

	query := Labels(queryMods...)
	queries.SetFrom(query.Query, "`labels`")

	return query
}

// LoadTodo allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (todosLabelL) LoadTodo(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTodosLabel interface{}, mods queries.Applicator) error {
	var slice []*TodosLabel
	var object *TodosLabel

	if singular {
		object = maybeTodosLabel.(*TodosLabel)
	} else {
		slice = *maybeTodosLabel.(*[]*TodosLabel)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &todosLabelR{}
		}
		args = append(args, object.TodoID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &todosLabelR{}
			}

			for _, a := range args {
				if a == obj.TodoID {
					continue Outer
				}
			}

			args = append(args, obj.TodoID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`todos`), qm.WhereIn(`todos.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Todo")
	}

	var resultSlice []*Todo
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Todo")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for todos")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for todos")
	}

	if len(todosLabelAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Todo = foreign
		if foreign.R == nil {
			foreign.R = &todoR{}
		}
		foreign.R.TodosLabels = append(foreign.R.TodosLabels, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TodoID == foreign.ID {
				local.R.Todo = foreign
				if foreign.R == nil {
					foreign.R = &todoR{}
				}
				foreign.R.TodosLabels = append(foreign.R.TodosLabels, local)
				break
			}
		}
	}

	return nil
}

// LoadLabel allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (todosLabelL) LoadLabel(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTodosLabel interface{}, mods queries.Applicator) error {
	var slice []*TodosLabel
	var object *TodosLabel

	if singular {
		object = maybeTodosLabel.(*TodosLabel)
	} else {
		slice = *maybeTodosLabel.(*[]*TodosLabel)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &todosLabelR{}
		}
		args = append(args, object.LabelID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &todosLabelR{}
			}

			for _, a := range args {
				if a == obj.LabelID {
					continue Outer
				}
			}

			args = append(args, obj.LabelID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`labels`), qm.WhereIn(`labels.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Label")
	}

	var resultSlice []*Label
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Label")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for labels")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for labels")
	}

	if len(todosLabelAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Label = foreign
		if foreign.R == nil {
			foreign.R = &labelR{}
		}
		foreign.R.TodosLabels = append(foreign.R.TodosLabels, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.LabelID == foreign.ID {
				local.R.Label = foreign
				if foreign.R == nil {
					foreign.R = &labelR{}
				}
				foreign.R.TodosLabels = append(foreign.R.TodosLabels, local)
				break
			}
		}
	}

	return nil
}

// SetTodo of the todosLabel to the related item.
// Sets o.R.Todo to related.
// Adds o to related.R.TodosLabels.
func (o *TodosLabel) SetTodo(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Todo) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `todos_labels` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"todo_id"}),
		strmangle.WhereClause("`", "`", 0, todosLabelPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TodoID = related.ID
	if o.R == nil {
		o.R = &todosLabelR{
			Todo: related,
		}
	} else {
		o.R.Todo = related
	}

	if related.R == nil {
		related.R = &todoR{
			TodosLabels: TodosLabelSlice{o},
		}
	} else {
		related.R.TodosLabels = append(related.R.TodosLabels, o)
	}

	return nil
}

// SetLabel of the todosLabel to the related item.
// Sets o.R.Label to related.
// Adds o to related.R.TodosLabels.
func (o *TodosLabel) SetLabel(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Label) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `todos_labels` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"label_id"}),
		strmangle.WhereClause("`", "`", 0, todosLabelPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.LabelID = related.ID
	if o.R == nil {
		o.R = &todosLabelR{
			Label: related,
		}
	} else {
		o.R.Label = related
	}

	if related.R == nil {
		related.R = &labelR{
			TodosLabels: TodosLabelSlice{o},
		}
	} else {
		related.R.TodosLabels = append(related.R.TodosLabels, o)
	}

	return nil
}

// TodosLabels retrieves all the records using an executor.
func TodosLabels(mods ...qm.QueryMod) todosLabelQuery {
	mods = append(mods, qm.From("`todos_labels`"))
	return todosLabelQuery{NewQuery(mods...)}
}

// FindTodosLabel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTodosLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*TodosLabel, error) {
	todosLabelObj := &TodosLabel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `todos_labels` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, todosLabelObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from todos_labels")
	}

	return todosLabelObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TodosLabel) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no todos_labels provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(todosLabelColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	todosLabelInsertCacheMut.RLock()
	cache, cached := todosLabelInsertCache[key]
	todosLabelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			todosLabelAllColumns,
			todosLabelColumnsWithDefault,
			todosLabelColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(todosLabelType, todosLabelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(todosLabelType, todosLabelMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `todos_labels` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `todos_labels` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `todos_labels` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, todosLabelPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to insert into todos_labels")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == todosLabelMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for todos_labels")
	}

CacheNoHooks:
	if !cached {
		todosLabelInsertCacheMut.Lock()
		todosLabelInsertCache[key] = cache
		todosLabelInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TodosLabel.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TodosLabel) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	todosLabelUpdateCacheMut.RLock()
	cache, cached := todosLabelUpdateCache[key]
	todosLabelUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			todosLabelAllColumns,
			todosLabelPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update todos_labels, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `todos_labels` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, todosLabelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(todosLabelType, todosLabelMapping, append(wl, todosLabelPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update todos_labels row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for todos_labels")
	}

	if !cached {
		todosLabelUpdateCacheMut.Lock()
		todosLabelUpdateCache[key] = cache
		todosLabelUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q todosLabelQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for todos_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for todos_labels")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TodosLabelSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("model: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), todosLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `todos_labels` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, todosLabelPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in todosLabel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all todosLabel")
	}
	return rowsAff, nil
}

var mySQLTodosLabelUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TodosLabel) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no todos_labels provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(todosLabelColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTodosLabelUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	todosLabelUpsertCacheMut.RLock()
	cache, cached := todosLabelUpsertCache[key]
	todosLabelUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			todosLabelAllColumns,
			todosLabelColumnsWithDefault,
			todosLabelColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			todosLabelAllColumns,
			todosLabelPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("model: unable to upsert todos_labels, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "todos_labels", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `todos_labels` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(todosLabelType, todosLabelMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(todosLabelType, todosLabelMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to upsert for todos_labels")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == todosLabelMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(todosLabelType, todosLabelMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for todos_labels")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for todos_labels")
	}

CacheNoHooks:
	if !cached {
		todosLabelUpsertCacheMut.Lock()
		todosLabelUpsertCache[key] = cache
		todosLabelUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TodosLabel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TodosLabel) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no TodosLabel provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), todosLabelPrimaryKeyMapping)
	sql := "DELETE FROM `todos_labels` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from todos_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for todos_labels")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q todosLabelQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no todosLabelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from todos_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for todos_labels")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TodosLabelSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(todosLabelBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), todosLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `todos_labels` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, todosLabelPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from todosLabel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for todos_labels")
	}

	if len(todosLabelAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TodosLabel) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTodosLabel(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TodosLabelSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TodosLabelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), todosLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `todos_labels`.* FROM `todos_labels` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, todosLabelPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in TodosLabelSlice")
	}

	*o = slice

	return nil
}

// TodosLabelExists checks if the TodosLabel row exists.
func TodosLabelExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `todos_labels` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if todos_labels exists")
	}

	return exists, nil
}