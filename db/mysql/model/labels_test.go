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

func testLabels(t *testing.T) {
	t.Parallel()

	query := Labels()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testLabelsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Label{}
	if err = randomize.Struct(seed, o, labelDBTypes, true, labelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
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

	count, err := Labels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLabelsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Label{}
	if err = randomize.Struct(seed, o, labelDBTypes, true, labelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Labels().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Labels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLabelsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Label{}
	if err = randomize.Struct(seed, o, labelDBTypes, true, labelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LabelSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Labels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLabelsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Label{}
	if err = randomize.Struct(seed, o, labelDBTypes, true, labelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := LabelExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Label exists: %s", err)
	}
	if !e {
		t.Errorf("Expected LabelExists to return true, but got false.")
	}
}

func testLabelsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Label{}
	if err = randomize.Struct(seed, o, labelDBTypes, true, labelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	labelFound, err := FindLabel(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if labelFound == nil {
		t.Error("want a record, got nil")
	}
}

func testLabelsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Label{}
	if err = randomize.Struct(seed, o, labelDBTypes, true, labelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Labels().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testLabelsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Label{}
	if err = randomize.Struct(seed, o, labelDBTypes, true, labelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Labels().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testLabelsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	labelOne := &Label{}
	labelTwo := &Label{}
	if err = randomize.Struct(seed, labelOne, labelDBTypes, false, labelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
	}
	if err = randomize.Struct(seed, labelTwo, labelDBTypes, false, labelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = labelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = labelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Labels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testLabelsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	labelOne := &Label{}
	labelTwo := &Label{}
	if err = randomize.Struct(seed, labelOne, labelDBTypes, false, labelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
	}
	if err = randomize.Struct(seed, labelTwo, labelDBTypes, false, labelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = labelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = labelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Labels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func labelBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Label) error {
	*o = Label{}
	return nil
}

func labelAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Label) error {
	*o = Label{}
	return nil
}

func labelAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Label) error {
	*o = Label{}
	return nil
}

func labelBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Label) error {
	*o = Label{}
	return nil
}

func labelAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Label) error {
	*o = Label{}
	return nil
}

func labelBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Label) error {
	*o = Label{}
	return nil
}

func labelAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Label) error {
	*o = Label{}
	return nil
}

func labelBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Label) error {
	*o = Label{}
	return nil
}

func labelAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Label) error {
	*o = Label{}
	return nil
}

func testLabelsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Label{}
	o := &Label{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, labelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Label object: %s", err)
	}

	AddLabelHook(boil.BeforeInsertHook, labelBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	labelBeforeInsertHooks = []LabelHook{}

	AddLabelHook(boil.AfterInsertHook, labelAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	labelAfterInsertHooks = []LabelHook{}

	AddLabelHook(boil.AfterSelectHook, labelAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	labelAfterSelectHooks = []LabelHook{}

	AddLabelHook(boil.BeforeUpdateHook, labelBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	labelBeforeUpdateHooks = []LabelHook{}

	AddLabelHook(boil.AfterUpdateHook, labelAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	labelAfterUpdateHooks = []LabelHook{}

	AddLabelHook(boil.BeforeDeleteHook, labelBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	labelBeforeDeleteHooks = []LabelHook{}

	AddLabelHook(boil.AfterDeleteHook, labelAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	labelAfterDeleteHooks = []LabelHook{}

	AddLabelHook(boil.BeforeUpsertHook, labelBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	labelBeforeUpsertHooks = []LabelHook{}

	AddLabelHook(boil.AfterUpsertHook, labelAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	labelAfterUpsertHooks = []LabelHook{}
}

func testLabelsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Label{}
	if err = randomize.Struct(seed, o, labelDBTypes, true, labelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Labels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLabelsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Label{}
	if err = randomize.Struct(seed, o, labelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(labelColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Labels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLabelToManyCards(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Label
	var b, c Card

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, labelDBTypes, true, labelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, cardDBTypes, false, cardColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, cardDBTypes, false, cardColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec("insert into `cards_labels` (`label_id`, `card_id`) values (?, ?)", a.ID, b.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into `cards_labels` (`label_id`, `card_id`) values (?, ?)", a.ID, c.ID)
	if err != nil {
		t.Fatal(err)
	}

	check, err := a.Cards().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ID == b.ID {
			bFound = true
		}
		if v.ID == c.ID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := LabelSlice{&a}
	if err = a.L.LoadCards(ctx, tx, false, (*[]*Label)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Cards); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Cards = nil
	if err = a.L.LoadCards(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Cards); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testLabelToManyAddOpCards(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Label
	var b, c, d, e Card

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, labelDBTypes, false, strmangle.SetComplement(labelPrimaryKeyColumns, labelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Card{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, cardDBTypes, false, strmangle.SetComplement(cardPrimaryKeyColumns, cardColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Card{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddCards(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if first.R.Labels[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.Labels[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}

		if a.R.Cards[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Cards[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Cards().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testLabelToManySetOpCards(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Label
	var b, c, d, e Card

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, labelDBTypes, false, strmangle.SetComplement(labelPrimaryKeyColumns, labelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Card{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, cardDBTypes, false, strmangle.SetComplement(cardPrimaryKeyColumns, cardColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetCards(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Cards().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetCards(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Cards().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	// The following checks cannot be implemented since we have no handle
	// to these when we call Set(). Leaving them here as wishful thinking
	// and to let people know there's dragons.
	//
	// if len(b.R.Labels) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.Labels) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.Labels[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.Labels[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.Cards[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Cards[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testLabelToManyRemoveOpCards(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Label
	var b, c, d, e Card

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, labelDBTypes, false, strmangle.SetComplement(labelPrimaryKeyColumns, labelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Card{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, cardDBTypes, false, strmangle.SetComplement(cardPrimaryKeyColumns, cardColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddCards(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Cards().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveCards(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Cards().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if len(b.R.Labels) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.Labels) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.Labels[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Labels[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if len(a.R.Cards) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Cards[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Cards[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testLabelToOneBoardUsingBoard(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Label
	var foreign Board

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, labelDBTypes, false, labelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, boardDBTypes, false, boardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Board struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.BoardID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Board().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := LabelSlice{&local}
	if err = local.L.LoadBoard(ctx, tx, false, (*[]*Label)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Board == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Board = nil
	if err = local.L.LoadBoard(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Board == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testLabelToOneSetOpBoardUsingBoard(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Label
	var b, c Board

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, labelDBTypes, false, strmangle.SetComplement(labelPrimaryKeyColumns, labelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, boardDBTypes, false, strmangle.SetComplement(boardPrimaryKeyColumns, boardColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, boardDBTypes, false, strmangle.SetComplement(boardPrimaryKeyColumns, boardColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Board{&b, &c} {
		err = a.SetBoard(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Board != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Labels[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.BoardID != x.ID {
			t.Error("foreign key was wrong value", a.BoardID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.BoardID))
		reflect.Indirect(reflect.ValueOf(&a.BoardID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.BoardID != x.ID {
			t.Error("foreign key was wrong value", a.BoardID, x.ID)
		}
	}
}

func testLabelsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Label{}
	if err = randomize.Struct(seed, o, labelDBTypes, true, labelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
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

func testLabelsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Label{}
	if err = randomize.Struct(seed, o, labelDBTypes, true, labelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LabelSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testLabelsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Label{}
	if err = randomize.Struct(seed, o, labelDBTypes, true, labelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Labels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	labelDBTypes = map[string]string{`ID`: `int`, `BoardID`: `int`, `Name`: `varchar`, `Color`: `int`, `CreatedAt`: `datetime`, `UpdatedAt`: `datetime`}
	_            = bytes.MinRead
)

func testLabelsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(labelPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(labelAllColumns) == len(labelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Label{}
	if err = randomize.Struct(seed, o, labelDBTypes, true, labelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Labels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, labelDBTypes, true, labelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testLabelsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(labelAllColumns) == len(labelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Label{}
	if err = randomize.Struct(seed, o, labelDBTypes, true, labelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Labels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, labelDBTypes, true, labelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(labelAllColumns, labelPrimaryKeyColumns) {
		fields = labelAllColumns
	} else {
		fields = strmangle.SetComplement(
			labelAllColumns,
			labelPrimaryKeyColumns,
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

	slice := LabelSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testLabelsUpsert(t *testing.T) {
	t.Parallel()

	if len(labelAllColumns) == len(labelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLLabelUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Label{}
	if err = randomize.Struct(seed, &o, labelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Label: %s", err)
	}

	count, err := Labels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, labelDBTypes, false, labelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Label struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Label: %s", err)
	}

	count, err = Labels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
