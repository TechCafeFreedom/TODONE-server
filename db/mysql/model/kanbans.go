// Code generated by SQLBoiler 4.1.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Kanban is an object representing the database table.
type Kanban struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID    int       `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	BoardID   int       `boil:"board_id" json:"board_id" toml:"board_id" yaml:"board_id"`
	Title     string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	IsArchive bool      `boil:"is_archive" json:"is_archive" toml:"is_archive" yaml:"is_archive"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *kanbanR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L kanbanL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var KanbanColumns = struct {
	ID        string
	UserID    string
	BoardID   string
	Title     string
	IsArchive string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	BoardID:   "board_id",
	Title:     "title",
	IsArchive: "is_archive",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// Generated where

var KanbanWhere = struct {
	ID        whereHelperint
	UserID    whereHelperint
	BoardID   whereHelperint
	Title     whereHelperstring
	IsArchive whereHelperbool
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint{field: "`kanbans`.`id`"},
	UserID:    whereHelperint{field: "`kanbans`.`user_id`"},
	BoardID:   whereHelperint{field: "`kanbans`.`board_id`"},
	Title:     whereHelperstring{field: "`kanbans`.`title`"},
	IsArchive: whereHelperbool{field: "`kanbans`.`is_archive`"},
	CreatedAt: whereHelpertime_Time{field: "`kanbans`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`kanbans`.`updated_at`"},
}

// KanbanRels is where relationship names are stored.
var KanbanRels = struct {
	User  string
	Board string
	Cards string
}{
	User:  "User",
	Board: "Board",
	Cards: "Cards",
}

// kanbanR is where relationships are stored.
type kanbanR struct {
	User  *User     `boil:"User" json:"User" toml:"User" yaml:"User"`
	Board *Board    `boil:"Board" json:"Board" toml:"Board" yaml:"Board"`
	Cards CardSlice `boil:"Cards" json:"Cards" toml:"Cards" yaml:"Cards"`
}

// NewStruct creates a new relationship struct
func (*kanbanR) NewStruct() *kanbanR {
	return &kanbanR{}
}

// kanbanL is where Load methods for each relationship are stored.
type kanbanL struct{}

var (
	kanbanAllColumns            = []string{"id", "user_id", "board_id", "title", "is_archive", "created_at", "updated_at"}
	kanbanColumnsWithoutDefault = []string{"user_id", "board_id", "title"}
	kanbanColumnsWithDefault    = []string{"id", "is_archive", "created_at", "updated_at"}
	kanbanPrimaryKeyColumns     = []string{"id"}
)

type (
	// KanbanSlice is an alias for a slice of pointers to Kanban.
	// This should generally be used opposed to []Kanban.
	KanbanSlice []*Kanban
	// KanbanHook is the signature for custom Kanban hook methods
	KanbanHook func(context.Context, boil.ContextExecutor, *Kanban) error

	kanbanQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	kanbanType                 = reflect.TypeOf(&Kanban{})
	kanbanMapping              = queries.MakeStructMapping(kanbanType)
	kanbanPrimaryKeyMapping, _ = queries.BindMapping(kanbanType, kanbanMapping, kanbanPrimaryKeyColumns)
	kanbanInsertCacheMut       sync.RWMutex
	kanbanInsertCache          = make(map[string]insertCache)
	kanbanUpdateCacheMut       sync.RWMutex
	kanbanUpdateCache          = make(map[string]updateCache)
	kanbanUpsertCacheMut       sync.RWMutex
	kanbanUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var kanbanBeforeInsertHooks []KanbanHook
var kanbanBeforeUpdateHooks []KanbanHook
var kanbanBeforeDeleteHooks []KanbanHook
var kanbanBeforeUpsertHooks []KanbanHook

var kanbanAfterInsertHooks []KanbanHook
var kanbanAfterSelectHooks []KanbanHook
var kanbanAfterUpdateHooks []KanbanHook
var kanbanAfterDeleteHooks []KanbanHook
var kanbanAfterUpsertHooks []KanbanHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Kanban) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range kanbanBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Kanban) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range kanbanBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Kanban) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range kanbanBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Kanban) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range kanbanBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Kanban) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range kanbanAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Kanban) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range kanbanAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Kanban) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range kanbanAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Kanban) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range kanbanAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Kanban) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range kanbanAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddKanbanHook registers your hook function for all future operations.
func AddKanbanHook(hookPoint boil.HookPoint, kanbanHook KanbanHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		kanbanBeforeInsertHooks = append(kanbanBeforeInsertHooks, kanbanHook)
	case boil.BeforeUpdateHook:
		kanbanBeforeUpdateHooks = append(kanbanBeforeUpdateHooks, kanbanHook)
	case boil.BeforeDeleteHook:
		kanbanBeforeDeleteHooks = append(kanbanBeforeDeleteHooks, kanbanHook)
	case boil.BeforeUpsertHook:
		kanbanBeforeUpsertHooks = append(kanbanBeforeUpsertHooks, kanbanHook)
	case boil.AfterInsertHook:
		kanbanAfterInsertHooks = append(kanbanAfterInsertHooks, kanbanHook)
	case boil.AfterSelectHook:
		kanbanAfterSelectHooks = append(kanbanAfterSelectHooks, kanbanHook)
	case boil.AfterUpdateHook:
		kanbanAfterUpdateHooks = append(kanbanAfterUpdateHooks, kanbanHook)
	case boil.AfterDeleteHook:
		kanbanAfterDeleteHooks = append(kanbanAfterDeleteHooks, kanbanHook)
	case boil.AfterUpsertHook:
		kanbanAfterUpsertHooks = append(kanbanAfterUpsertHooks, kanbanHook)
	}
}

// One returns a single kanban record from the query.
func (q kanbanQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Kanban, error) {
	o := &Kanban{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for kanbans")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Kanban records from the query.
func (q kanbanQuery) All(ctx context.Context, exec boil.ContextExecutor) (KanbanSlice, error) {
	var o []*Kanban

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to Kanban slice")
	}

	if len(kanbanAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Kanban records in the query.
func (q kanbanQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count kanbans rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q kanbanQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if kanbans exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *Kanban) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "`users`")

	return query
}

// Board pointed to by the foreign key.
func (o *Kanban) Board(mods ...qm.QueryMod) boardQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.BoardID),
	}

	queryMods = append(queryMods, mods...)

	query := Boards(queryMods...)
	queries.SetFrom(query.Query, "`boards`")

	return query
}

// Cards retrieves all the card's Cards with an executor.
func (o *Kanban) Cards(mods ...qm.QueryMod) cardQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`cards`.`kanban_id`=?", o.ID),
	)

	query := Cards(queryMods...)
	queries.SetFrom(query.Query, "`cards`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`cards`.*"})
	}

	return query
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (kanbanL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeKanban interface{}, mods queries.Applicator) error {
	var slice []*Kanban
	var object *Kanban

	if singular {
		object = maybeKanban.(*Kanban)
	} else {
		slice = *maybeKanban.(*[]*Kanban)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &kanbanR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &kanbanR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(kanbanAfterSelectHooks) != 0 {
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
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.Kanbans = append(foreign.R.Kanbans, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.Kanbans = append(foreign.R.Kanbans, local)
				break
			}
		}
	}

	return nil
}

// LoadBoard allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (kanbanL) LoadBoard(ctx context.Context, e boil.ContextExecutor, singular bool, maybeKanban interface{}, mods queries.Applicator) error {
	var slice []*Kanban
	var object *Kanban

	if singular {
		object = maybeKanban.(*Kanban)
	} else {
		slice = *maybeKanban.(*[]*Kanban)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &kanbanR{}
		}
		args = append(args, object.BoardID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &kanbanR{}
			}

			for _, a := range args {
				if a == obj.BoardID {
					continue Outer
				}
			}

			args = append(args, obj.BoardID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`boards`),
		qm.WhereIn(`boards.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Board")
	}

	var resultSlice []*Board
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Board")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for boards")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for boards")
	}

	if len(kanbanAfterSelectHooks) != 0 {
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
		object.R.Board = foreign
		if foreign.R == nil {
			foreign.R = &boardR{}
		}
		foreign.R.Kanbans = append(foreign.R.Kanbans, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.BoardID == foreign.ID {
				local.R.Board = foreign
				if foreign.R == nil {
					foreign.R = &boardR{}
				}
				foreign.R.Kanbans = append(foreign.R.Kanbans, local)
				break
			}
		}
	}

	return nil
}

// LoadCards allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (kanbanL) LoadCards(ctx context.Context, e boil.ContextExecutor, singular bool, maybeKanban interface{}, mods queries.Applicator) error {
	var slice []*Kanban
	var object *Kanban

	if singular {
		object = maybeKanban.(*Kanban)
	} else {
		slice = *maybeKanban.(*[]*Kanban)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &kanbanR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &kanbanR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`cards`),
		qm.WhereIn(`cards.kanban_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load cards")
	}

	var resultSlice []*Card
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice cards")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on cards")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for cards")
	}

	if len(cardAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Cards = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &cardR{}
			}
			foreign.R.Kanban = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.KanbanID {
				local.R.Cards = append(local.R.Cards, foreign)
				if foreign.R == nil {
					foreign.R = &cardR{}
				}
				foreign.R.Kanban = local
				break
			}
		}
	}

	return nil
}

// SetUser of the kanban to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Kanbans.
func (o *Kanban) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `kanbans` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, kanbanPrimaryKeyColumns),
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

	o.UserID = related.ID
	if o.R == nil {
		o.R = &kanbanR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			Kanbans: KanbanSlice{o},
		}
	} else {
		related.R.Kanbans = append(related.R.Kanbans, o)
	}

	return nil
}

// SetBoard of the kanban to the related item.
// Sets o.R.Board to related.
// Adds o to related.R.Kanbans.
func (o *Kanban) SetBoard(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Board) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `kanbans` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"board_id"}),
		strmangle.WhereClause("`", "`", 0, kanbanPrimaryKeyColumns),
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

	o.BoardID = related.ID
	if o.R == nil {
		o.R = &kanbanR{
			Board: related,
		}
	} else {
		o.R.Board = related
	}

	if related.R == nil {
		related.R = &boardR{
			Kanbans: KanbanSlice{o},
		}
	} else {
		related.R.Kanbans = append(related.R.Kanbans, o)
	}

	return nil
}

// AddCards adds the given related objects to the existing relationships
// of the kanban, optionally inserting them as new records.
// Appends related to o.R.Cards.
// Sets related.R.Kanban appropriately.
func (o *Kanban) AddCards(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Card) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.KanbanID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `cards` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"kanban_id"}),
				strmangle.WhereClause("`", "`", 0, cardPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.KanbanID = o.ID
		}
	}

	if o.R == nil {
		o.R = &kanbanR{
			Cards: related,
		}
	} else {
		o.R.Cards = append(o.R.Cards, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &cardR{
				Kanban: o,
			}
		} else {
			rel.R.Kanban = o
		}
	}
	return nil
}

// Kanbans retrieves all the records using an executor.
func Kanbans(mods ...qm.QueryMod) kanbanQuery {
	mods = append(mods, qm.From("`kanbans`"))
	return kanbanQuery{NewQuery(mods...)}
}

// FindKanban retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindKanban(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Kanban, error) {
	kanbanObj := &Kanban{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `kanbans` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, kanbanObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from kanbans")
	}

	return kanbanObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Kanban) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no kanbans provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(kanbanColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	kanbanInsertCacheMut.RLock()
	cache, cached := kanbanInsertCache[key]
	kanbanInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			kanbanAllColumns,
			kanbanColumnsWithDefault,
			kanbanColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(kanbanType, kanbanMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(kanbanType, kanbanMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `kanbans` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `kanbans` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `kanbans` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, kanbanPrimaryKeyColumns))
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
		return errors.Wrap(err, "model: unable to insert into kanbans")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == kanbanMapping["id"] {
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
		return errors.Wrap(err, "model: unable to populate default values for kanbans")
	}

CacheNoHooks:
	if !cached {
		kanbanInsertCacheMut.Lock()
		kanbanInsertCache[key] = cache
		kanbanInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Kanban.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Kanban) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	kanbanUpdateCacheMut.RLock()
	cache, cached := kanbanUpdateCache[key]
	kanbanUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			kanbanAllColumns,
			kanbanPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update kanbans, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `kanbans` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, kanbanPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(kanbanType, kanbanMapping, append(wl, kanbanPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "model: unable to update kanbans row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for kanbans")
	}

	if !cached {
		kanbanUpdateCacheMut.Lock()
		kanbanUpdateCache[key] = cache
		kanbanUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q kanbanQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for kanbans")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for kanbans")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o KanbanSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), kanbanPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `kanbans` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, kanbanPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in kanban slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all kanban")
	}
	return rowsAff, nil
}

var mySQLKanbanUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Kanban) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no kanbans provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(kanbanColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLKanbanUniqueColumns, o)

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

	kanbanUpsertCacheMut.RLock()
	cache, cached := kanbanUpsertCache[key]
	kanbanUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			kanbanAllColumns,
			kanbanColumnsWithDefault,
			kanbanColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			kanbanAllColumns,
			kanbanPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("model: unable to upsert kanbans, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "kanbans", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `kanbans` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(kanbanType, kanbanMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(kanbanType, kanbanMapping, ret)
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
		return errors.Wrap(err, "model: unable to upsert for kanbans")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == kanbanMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(kanbanType, kanbanMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for kanbans")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for kanbans")
	}

CacheNoHooks:
	if !cached {
		kanbanUpsertCacheMut.Lock()
		kanbanUpsertCache[key] = cache
		kanbanUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Kanban record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Kanban) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no Kanban provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), kanbanPrimaryKeyMapping)
	sql := "DELETE FROM `kanbans` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from kanbans")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for kanbans")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q kanbanQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no kanbanQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from kanbans")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for kanbans")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o KanbanSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(kanbanBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), kanbanPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `kanbans` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, kanbanPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from kanban slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for kanbans")
	}

	if len(kanbanAfterDeleteHooks) != 0 {
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
func (o *Kanban) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindKanban(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *KanbanSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := KanbanSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), kanbanPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `kanbans`.* FROM `kanbans` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, kanbanPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in KanbanSlice")
	}

	*o = slice

	return nil
}

// KanbanExists checks if the Kanban row exists.
func KanbanExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `kanbans` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if kanbans exists")
	}

	return exists, nil
}
