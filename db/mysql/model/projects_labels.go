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

// ProjectsLabel is an object representing the database table.
type ProjectsLabel struct {
	ID        int `boil:"id" json:"id" toml:"id" yaml:"id"`
	ProjectID int `boil:"project_id" json:"project_id" toml:"project_id" yaml:"project_id"`
	LabelID   int `boil:"label_id" json:"label_id" toml:"label_id" yaml:"label_id"`

	R *projectsLabelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L projectsLabelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProjectsLabelColumns = struct {
	ID        string
	ProjectID string
	LabelID   string
}{
	ID:        "id",
	ProjectID: "project_id",
	LabelID:   "label_id",
}

// Generated where

var ProjectsLabelWhere = struct {
	ID        whereHelperint
	ProjectID whereHelperint
	LabelID   whereHelperint
}{
	ID:        whereHelperint{field: "`projects_labels`.`id`"},
	ProjectID: whereHelperint{field: "`projects_labels`.`project_id`"},
	LabelID:   whereHelperint{field: "`projects_labels`.`label_id`"},
}

// ProjectsLabelRels is where relationship names are stored.
var ProjectsLabelRels = struct {
	Project string
	Label   string
}{
	Project: "Project",
	Label:   "Label",
}

// projectsLabelR is where relationships are stored.
type projectsLabelR struct {
	Project *Project
	Label   *Label
}

// NewStruct creates a new relationship struct
func (*projectsLabelR) NewStruct() *projectsLabelR {
	return &projectsLabelR{}
}

// projectsLabelL is where Load methods for each relationship are stored.
type projectsLabelL struct{}

var (
	projectsLabelAllColumns            = []string{"id", "project_id", "label_id"}
	projectsLabelColumnsWithoutDefault = []string{"project_id", "label_id"}
	projectsLabelColumnsWithDefault    = []string{"id"}
	projectsLabelPrimaryKeyColumns     = []string{"id"}
)

type (
	// ProjectsLabelSlice is an alias for a slice of pointers to ProjectsLabel.
	// This should generally be used opposed to []ProjectsLabel.
	ProjectsLabelSlice []*ProjectsLabel
	// ProjectsLabelHook is the signature for custom ProjectsLabel hook methods
	ProjectsLabelHook func(context.Context, boil.ContextExecutor, *ProjectsLabel) error

	projectsLabelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	projectsLabelType                 = reflect.TypeOf(&ProjectsLabel{})
	projectsLabelMapping              = queries.MakeStructMapping(projectsLabelType)
	projectsLabelPrimaryKeyMapping, _ = queries.BindMapping(projectsLabelType, projectsLabelMapping, projectsLabelPrimaryKeyColumns)
	projectsLabelInsertCacheMut       sync.RWMutex
	projectsLabelInsertCache          = make(map[string]insertCache)
	projectsLabelUpdateCacheMut       sync.RWMutex
	projectsLabelUpdateCache          = make(map[string]updateCache)
	projectsLabelUpsertCacheMut       sync.RWMutex
	projectsLabelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var projectsLabelBeforeInsertHooks []ProjectsLabelHook
var projectsLabelBeforeUpdateHooks []ProjectsLabelHook
var projectsLabelBeforeDeleteHooks []ProjectsLabelHook
var projectsLabelBeforeUpsertHooks []ProjectsLabelHook

var projectsLabelAfterInsertHooks []ProjectsLabelHook
var projectsLabelAfterSelectHooks []ProjectsLabelHook
var projectsLabelAfterUpdateHooks []ProjectsLabelHook
var projectsLabelAfterDeleteHooks []ProjectsLabelHook
var projectsLabelAfterUpsertHooks []ProjectsLabelHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ProjectsLabel) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectsLabelBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ProjectsLabel) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectsLabelBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ProjectsLabel) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectsLabelBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ProjectsLabel) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectsLabelBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ProjectsLabel) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectsLabelAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ProjectsLabel) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectsLabelAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ProjectsLabel) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectsLabelAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ProjectsLabel) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectsLabelAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ProjectsLabel) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectsLabelAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddProjectsLabelHook registers your hook function for all future operations.
func AddProjectsLabelHook(hookPoint boil.HookPoint, projectsLabelHook ProjectsLabelHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		projectsLabelBeforeInsertHooks = append(projectsLabelBeforeInsertHooks, projectsLabelHook)
	case boil.BeforeUpdateHook:
		projectsLabelBeforeUpdateHooks = append(projectsLabelBeforeUpdateHooks, projectsLabelHook)
	case boil.BeforeDeleteHook:
		projectsLabelBeforeDeleteHooks = append(projectsLabelBeforeDeleteHooks, projectsLabelHook)
	case boil.BeforeUpsertHook:
		projectsLabelBeforeUpsertHooks = append(projectsLabelBeforeUpsertHooks, projectsLabelHook)
	case boil.AfterInsertHook:
		projectsLabelAfterInsertHooks = append(projectsLabelAfterInsertHooks, projectsLabelHook)
	case boil.AfterSelectHook:
		projectsLabelAfterSelectHooks = append(projectsLabelAfterSelectHooks, projectsLabelHook)
	case boil.AfterUpdateHook:
		projectsLabelAfterUpdateHooks = append(projectsLabelAfterUpdateHooks, projectsLabelHook)
	case boil.AfterDeleteHook:
		projectsLabelAfterDeleteHooks = append(projectsLabelAfterDeleteHooks, projectsLabelHook)
	case boil.AfterUpsertHook:
		projectsLabelAfterUpsertHooks = append(projectsLabelAfterUpsertHooks, projectsLabelHook)
	}
}

// One returns a single projectsLabel record from the query.
func (q projectsLabelQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ProjectsLabel, error) {
	o := &ProjectsLabel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for projects_labels")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ProjectsLabel records from the query.
func (q projectsLabelQuery) All(ctx context.Context, exec boil.ContextExecutor) (ProjectsLabelSlice, error) {
	var o []*ProjectsLabel

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to ProjectsLabel slice")
	}

	if len(projectsLabelAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ProjectsLabel records in the query.
func (q projectsLabelQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count projects_labels rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q projectsLabelQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if projects_labels exists")
	}

	return count > 0, nil
}

// Project pointed to by the foreign key.
func (o *ProjectsLabel) Project(mods ...qm.QueryMod) projectQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.ProjectID),
	}

	queryMods = append(queryMods, mods...)

	query := Projects(queryMods...)
	queries.SetFrom(query.Query, "`projects`")

	return query
}

// Label pointed to by the foreign key.
func (o *ProjectsLabel) Label(mods ...qm.QueryMod) labelQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.LabelID),
	}

	queryMods = append(queryMods, mods...)

	query := Labels(queryMods...)
	queries.SetFrom(query.Query, "`labels`")

	return query
}

// LoadProject allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (projectsLabelL) LoadProject(ctx context.Context, e boil.ContextExecutor, singular bool, maybeProjectsLabel interface{}, mods queries.Applicator) error {
	var slice []*ProjectsLabel
	var object *ProjectsLabel

	if singular {
		object = maybeProjectsLabel.(*ProjectsLabel)
	} else {
		slice = *maybeProjectsLabel.(*[]*ProjectsLabel)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &projectsLabelR{}
		}
		args = append(args, object.ProjectID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &projectsLabelR{}
			}

			for _, a := range args {
				if a == obj.ProjectID {
					continue Outer
				}
			}

			args = append(args, obj.ProjectID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`projects`), qm.WhereIn(`projects.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Project")
	}

	var resultSlice []*Project
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Project")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for projects")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for projects")
	}

	if len(projectsLabelAfterSelectHooks) != 0 {
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
		object.R.Project = foreign
		if foreign.R == nil {
			foreign.R = &projectR{}
		}
		foreign.R.ProjectsLabels = append(foreign.R.ProjectsLabels, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ProjectID == foreign.ID {
				local.R.Project = foreign
				if foreign.R == nil {
					foreign.R = &projectR{}
				}
				foreign.R.ProjectsLabels = append(foreign.R.ProjectsLabels, local)
				break
			}
		}
	}

	return nil
}

// LoadLabel allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (projectsLabelL) LoadLabel(ctx context.Context, e boil.ContextExecutor, singular bool, maybeProjectsLabel interface{}, mods queries.Applicator) error {
	var slice []*ProjectsLabel
	var object *ProjectsLabel

	if singular {
		object = maybeProjectsLabel.(*ProjectsLabel)
	} else {
		slice = *maybeProjectsLabel.(*[]*ProjectsLabel)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &projectsLabelR{}
		}
		args = append(args, object.LabelID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &projectsLabelR{}
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

	if len(projectsLabelAfterSelectHooks) != 0 {
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
		foreign.R.ProjectsLabels = append(foreign.R.ProjectsLabels, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.LabelID == foreign.ID {
				local.R.Label = foreign
				if foreign.R == nil {
					foreign.R = &labelR{}
				}
				foreign.R.ProjectsLabels = append(foreign.R.ProjectsLabels, local)
				break
			}
		}
	}

	return nil
}

// SetProject of the projectsLabel to the related item.
// Sets o.R.Project to related.
// Adds o to related.R.ProjectsLabels.
func (o *ProjectsLabel) SetProject(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Project) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `projects_labels` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"project_id"}),
		strmangle.WhereClause("`", "`", 0, projectsLabelPrimaryKeyColumns),
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

	o.ProjectID = related.ID
	if o.R == nil {
		o.R = &projectsLabelR{
			Project: related,
		}
	} else {
		o.R.Project = related
	}

	if related.R == nil {
		related.R = &projectR{
			ProjectsLabels: ProjectsLabelSlice{o},
		}
	} else {
		related.R.ProjectsLabels = append(related.R.ProjectsLabels, o)
	}

	return nil
}

// SetLabel of the projectsLabel to the related item.
// Sets o.R.Label to related.
// Adds o to related.R.ProjectsLabels.
func (o *ProjectsLabel) SetLabel(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Label) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `projects_labels` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"label_id"}),
		strmangle.WhereClause("`", "`", 0, projectsLabelPrimaryKeyColumns),
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
		o.R = &projectsLabelR{
			Label: related,
		}
	} else {
		o.R.Label = related
	}

	if related.R == nil {
		related.R = &labelR{
			ProjectsLabels: ProjectsLabelSlice{o},
		}
	} else {
		related.R.ProjectsLabels = append(related.R.ProjectsLabels, o)
	}

	return nil
}

// ProjectsLabels retrieves all the records using an executor.
func ProjectsLabels(mods ...qm.QueryMod) projectsLabelQuery {
	mods = append(mods, qm.From("`projects_labels`"))
	return projectsLabelQuery{NewQuery(mods...)}
}

// FindProjectsLabel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProjectsLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ProjectsLabel, error) {
	projectsLabelObj := &ProjectsLabel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `projects_labels` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, projectsLabelObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from projects_labels")
	}

	return projectsLabelObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ProjectsLabel) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no projects_labels provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(projectsLabelColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	projectsLabelInsertCacheMut.RLock()
	cache, cached := projectsLabelInsertCache[key]
	projectsLabelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			projectsLabelAllColumns,
			projectsLabelColumnsWithDefault,
			projectsLabelColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(projectsLabelType, projectsLabelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(projectsLabelType, projectsLabelMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `projects_labels` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `projects_labels` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `projects_labels` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, projectsLabelPrimaryKeyColumns))
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
		return errors.Wrap(err, "model: unable to insert into projects_labels")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == projectsLabelMapping["id"] {
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
		return errors.Wrap(err, "model: unable to populate default values for projects_labels")
	}

CacheNoHooks:
	if !cached {
		projectsLabelInsertCacheMut.Lock()
		projectsLabelInsertCache[key] = cache
		projectsLabelInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ProjectsLabel.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ProjectsLabel) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	projectsLabelUpdateCacheMut.RLock()
	cache, cached := projectsLabelUpdateCache[key]
	projectsLabelUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			projectsLabelAllColumns,
			projectsLabelPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update projects_labels, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `projects_labels` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, projectsLabelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(projectsLabelType, projectsLabelMapping, append(wl, projectsLabelPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "model: unable to update projects_labels row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for projects_labels")
	}

	if !cached {
		projectsLabelUpdateCacheMut.Lock()
		projectsLabelUpdateCache[key] = cache
		projectsLabelUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q projectsLabelQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for projects_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for projects_labels")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProjectsLabelSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectsLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `projects_labels` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, projectsLabelPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in projectsLabel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all projectsLabel")
	}
	return rowsAff, nil
}

var mySQLProjectsLabelUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ProjectsLabel) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no projects_labels provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(projectsLabelColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLProjectsLabelUniqueColumns, o)

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

	projectsLabelUpsertCacheMut.RLock()
	cache, cached := projectsLabelUpsertCache[key]
	projectsLabelUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			projectsLabelAllColumns,
			projectsLabelColumnsWithDefault,
			projectsLabelColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			projectsLabelAllColumns,
			projectsLabelPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("model: unable to upsert projects_labels, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "projects_labels", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `projects_labels` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(projectsLabelType, projectsLabelMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(projectsLabelType, projectsLabelMapping, ret)
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
		return errors.Wrap(err, "model: unable to upsert for projects_labels")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == projectsLabelMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(projectsLabelType, projectsLabelMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for projects_labels")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for projects_labels")
	}

CacheNoHooks:
	if !cached {
		projectsLabelUpsertCacheMut.Lock()
		projectsLabelUpsertCache[key] = cache
		projectsLabelUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ProjectsLabel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ProjectsLabel) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no ProjectsLabel provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), projectsLabelPrimaryKeyMapping)
	sql := "DELETE FROM `projects_labels` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from projects_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for projects_labels")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q projectsLabelQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no projectsLabelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from projects_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for projects_labels")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProjectsLabelSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(projectsLabelBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectsLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `projects_labels` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, projectsLabelPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from projectsLabel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for projects_labels")
	}

	if len(projectsLabelAfterDeleteHooks) != 0 {
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
func (o *ProjectsLabel) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindProjectsLabel(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProjectsLabelSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProjectsLabelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectsLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `projects_labels`.* FROM `projects_labels` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, projectsLabelPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in ProjectsLabelSlice")
	}

	*o = slice

	return nil
}

// ProjectsLabelExists checks if the ProjectsLabel row exists.
func ProjectsLabelExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `projects_labels` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if projects_labels exists")
	}

	return exists, nil
}
