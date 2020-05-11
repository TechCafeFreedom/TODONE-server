// Code generated by SQLBoiler 4.1.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testBoards(t *testing.T) {
	t.Parallel()

	query := Boards()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBoardsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Board{}
	if err = randomize.Struct(seed, o, boardDBTypes, true, boardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Boards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBoardsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Board{}
	if err = randomize.Struct(seed, o, boardDBTypes, true, boardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Boards().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Boards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBoardsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Board{}
	if err = randomize.Struct(seed, o, boardDBTypes, true, boardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BoardSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Boards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBoardsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Board{}
	if err = randomize.Struct(seed, o, boardDBTypes, true, boardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BoardExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Board exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BoardExists to return true, but got false.")
	}
}

func testBoardsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Board{}
	if err = randomize.Struct(seed, o, boardDBTypes, true, boardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	boardFound, err := FindBoard(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if boardFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBoardsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Board{}
	if err = randomize.Struct(seed, o, boardDBTypes, true, boardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Boards().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBoardsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Board{}
	if err = randomize.Struct(seed, o, boardDBTypes, true, boardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Boards().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBoardsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boardOne := &Board{}
	boardTwo := &Board{}
	if err = randomize.Struct(seed, boardOne, boardDBTypes, false, boardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}
	if err = randomize.Struct(seed, boardTwo, boardDBTypes, false, boardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = boardOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = boardTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Boards().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBoardsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	boardOne := &Board{}
	boardTwo := &Board{}
	if err = randomize.Struct(seed, boardOne, boardDBTypes, false, boardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}
	if err = randomize.Struct(seed, boardTwo, boardDBTypes, false, boardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = boardOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = boardTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Boards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func boardBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Board) error {
	*o = Board{}
	return nil
}

func boardAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Board) error {
	*o = Board{}
	return nil
}

func boardAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Board) error {
	*o = Board{}
	return nil
}

func boardBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Board) error {
	*o = Board{}
	return nil
}

func boardAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Board) error {
	*o = Board{}
	return nil
}

func boardBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Board) error {
	*o = Board{}
	return nil
}

func boardAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Board) error {
	*o = Board{}
	return nil
}

func boardBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Board) error {
	*o = Board{}
	return nil
}

func boardAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Board) error {
	*o = Board{}
	return nil
}

func testBoardsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Board{}
	o := &Board{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, boardDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Board object: %s", err)
	}

	AddBoardHook(boil.BeforeInsertHook, boardBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	boardBeforeInsertHooks = []BoardHook{}

	AddBoardHook(boil.AfterInsertHook, boardAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	boardAfterInsertHooks = []BoardHook{}

	AddBoardHook(boil.AfterSelectHook, boardAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	boardAfterSelectHooks = []BoardHook{}

	AddBoardHook(boil.BeforeUpdateHook, boardBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	boardBeforeUpdateHooks = []BoardHook{}

	AddBoardHook(boil.AfterUpdateHook, boardAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	boardAfterUpdateHooks = []BoardHook{}

	AddBoardHook(boil.BeforeDeleteHook, boardBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	boardBeforeDeleteHooks = []BoardHook{}

	AddBoardHook(boil.AfterDeleteHook, boardAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	boardAfterDeleteHooks = []BoardHook{}

	AddBoardHook(boil.BeforeUpsertHook, boardBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	boardBeforeUpsertHooks = []BoardHook{}

	AddBoardHook(boil.AfterUpsertHook, boardAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	boardAfterUpsertHooks = []BoardHook{}
}

func testBoardsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Board{}
	if err = randomize.Struct(seed, o, boardDBTypes, true, boardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Boards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBoardsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Board{}
	if err = randomize.Struct(seed, o, boardDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(boardColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Boards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBoardToManyKanbans(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Board
	var b, c Kanban

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, boardDBTypes, true, boardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, kanbanDBTypes, false, kanbanColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, kanbanDBTypes, false, kanbanColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.BoardID = a.ID
	c.BoardID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Kanbans().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.BoardID == b.BoardID {
			bFound = true
		}
		if v.BoardID == c.BoardID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BoardSlice{&a}
	if err = a.L.LoadKanbans(ctx, tx, false, (*[]*Board)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Kanbans); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Kanbans = nil
	if err = a.L.LoadKanbans(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Kanbans); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBoardToManyLabels(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Board
	var b, c Label

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, boardDBTypes, true, boardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, labelDBTypes, false, labelColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, labelDBTypes, false, labelColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.BoardID = a.ID
	c.BoardID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Labels().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.BoardID == b.BoardID {
			bFound = true
		}
		if v.BoardID == c.BoardID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BoardSlice{&a}
	if err = a.L.LoadLabels(ctx, tx, false, (*[]*Board)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Labels); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Labels = nil
	if err = a.L.LoadLabels(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Labels); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBoardToManyUsersBoards(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Board
	var b, c UsersBoard

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, boardDBTypes, true, boardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, usersBoardDBTypes, false, usersBoardColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, usersBoardDBTypes, false, usersBoardColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.BoardID = a.ID
	c.BoardID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.UsersBoards().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.BoardID == b.BoardID {
			bFound = true
		}
		if v.BoardID == c.BoardID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BoardSlice{&a}
	if err = a.L.LoadUsersBoards(ctx, tx, false, (*[]*Board)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UsersBoards); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UsersBoards = nil
	if err = a.L.LoadUsersBoards(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UsersBoards); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBoardToManyAddOpKanbans(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Board
	var b, c, d, e Kanban

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, boardDBTypes, false, strmangle.SetComplement(boardPrimaryKeyColumns, boardColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Kanban{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, kanbanDBTypes, false, strmangle.SetComplement(kanbanPrimaryKeyColumns, kanbanColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Kanban{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddKanbans(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.BoardID {
			t.Error("foreign key was wrong value", a.ID, first.BoardID)
		}
		if a.ID != second.BoardID {
			t.Error("foreign key was wrong value", a.ID, second.BoardID)
		}

		if first.R.Board != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Board != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Kanbans[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Kanbans[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Kanbans().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testBoardToManyAddOpLabels(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Board
	var b, c, d, e Label

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, boardDBTypes, false, strmangle.SetComplement(boardPrimaryKeyColumns, boardColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Label{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, labelDBTypes, false, strmangle.SetComplement(labelPrimaryKeyColumns, labelColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Label{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddLabels(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.BoardID {
			t.Error("foreign key was wrong value", a.ID, first.BoardID)
		}
		if a.ID != second.BoardID {
			t.Error("foreign key was wrong value", a.ID, second.BoardID)
		}

		if first.R.Board != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Board != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Labels[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Labels[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Labels().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testBoardToManyAddOpUsersBoards(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Board
	var b, c, d, e UsersBoard

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, boardDBTypes, false, strmangle.SetComplement(boardPrimaryKeyColumns, boardColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*UsersBoard{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, usersBoardDBTypes, false, strmangle.SetComplement(usersBoardPrimaryKeyColumns, usersBoardColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*UsersBoard{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUsersBoards(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.BoardID {
			t.Error("foreign key was wrong value", a.ID, first.BoardID)
		}
		if a.ID != second.BoardID {
			t.Error("foreign key was wrong value", a.ID, second.BoardID)
		}

		if first.R.Board != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Board != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UsersBoards[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UsersBoards[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UsersBoards().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testBoardToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Board
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, boardDBTypes, false, boardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := BoardSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*Board)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBoardToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Board
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, boardDBTypes, false, strmangle.SetComplement(boardPrimaryKeyColumns, boardColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Boards[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testBoardsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Board{}
	if err = randomize.Struct(seed, o, boardDBTypes, true, boardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBoardsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Board{}
	if err = randomize.Struct(seed, o, boardDBTypes, true, boardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BoardSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBoardsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Board{}
	if err = randomize.Struct(seed, o, boardDBTypes, true, boardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Boards().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	boardDBTypes = map[string]string{`ID`: `int`, `UserID`: `int`, `Title`: `varchar`, `Description`: `varchar`, `CreatedAt`: `datetime`, `UpdatedAt`: `datetime`}
	_            = bytes.MinRead
)

func testBoardsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(boardPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(boardAllColumns) == len(boardPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Board{}
	if err = randomize.Struct(seed, o, boardDBTypes, true, boardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Boards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, boardDBTypes, true, boardPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBoardsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(boardAllColumns) == len(boardPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Board{}
	if err = randomize.Struct(seed, o, boardDBTypes, true, boardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Boards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, boardDBTypes, true, boardPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(boardAllColumns, boardPrimaryKeyColumns) {
		fields = boardAllColumns
	} else {
		fields = strmangle.SetComplement(
			boardAllColumns,
			boardPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := BoardSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBoardsUpsert(t *testing.T) {
	t.Parallel()

	if len(boardAllColumns) == len(boardPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLBoardUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Board{}
	if err = randomize.Struct(seed, &o, boardDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Board: %s", err)
	}

	count, err := Boards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, boardDBTypes, false, boardPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Board: %s", err)
	}

	count, err = Boards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
