// Package postgres contains the types for schema 'public'.
package postgres

// GENERATED BY XO. DO NOT EDIT.

import "errors"

// AuthGroup represents a row from 'public.auth_group'.
type AuthGroup struct {
	ID   int    `json:"id"`   // id
	Name string `json:"name"` // name

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the AuthGroup exists in the database.
func (ag *AuthGroup) Exists() bool {
	return ag._exists
}

// Deleted provides information if the AuthGroup has been deleted from the database.
func (ag *AuthGroup) Deleted() bool {
	return ag._deleted
}

// Insert inserts the AuthGroup to the database.
func (ag *AuthGroup) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if ag._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.auth_group (` +
		`name` +
		`) VALUES (` +
		`$1` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, ag.Name)
	err = db.QueryRow(sqlstr, ag.Name).Scan(&ag.ID)
	if err != nil {
		return err
	}

	// set existence
	ag._exists = true

	return nil
}

// Update updates the AuthGroup in the database.
func (ag *AuthGroup) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ag._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if ag._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.auth_group SET (` +
		`name` +
		`) = ( ` +
		`$1` +
		`) WHERE id = $2`

	// run query
	XOLog(sqlstr, ag.Name, ag.ID)
	_, err = db.Exec(sqlstr, ag.Name, ag.ID)
	return err
}

// Save saves the AuthGroup to the database.
func (ag *AuthGroup) Save(db XODB) error {
	if ag.Exists() {
		return ag.Update(db)
	}

	return ag.Insert(db)
}

// Upsert performs an upsert for AuthGroup.
//
// NOTE: PostgreSQL 9.5+ only
func (ag *AuthGroup) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if ag._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.auth_group (` +
		`id, name` +
		`) VALUES (` +
		`$1, $2` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, name` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.name` +
		`)`

	// run query
	XOLog(sqlstr, ag.ID, ag.Name)
	_, err = db.Exec(sqlstr, ag.ID, ag.Name)
	if err != nil {
		return err
	}

	// set existence
	ag._exists = true

	return nil
}

// Delete deletes the AuthGroup from the database.
func (ag *AuthGroup) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ag._exists {
		return nil
	}

	// if deleted, bail
	if ag._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.auth_group WHERE id = $1`

	// run query
	XOLog(sqlstr, ag.ID)
	_, err = db.Exec(sqlstr, ag.ID)
	if err != nil {
		return err
	}

	// set deleted
	ag._deleted = true

	return nil
}

// AuthGroupsByName retrieves a row from 'public.auth_group' as a AuthGroup.
//
// Generated from index 'auth_group_name_a6ea08ec_like'.
func AuthGroupsByName(db XODB, name string) ([]*AuthGroup, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name ` +
		`FROM public.auth_group ` +
		`WHERE name = $1`

	// run query
	XOLog(sqlstr, name)
	q, err := db.Query(sqlstr, name)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AuthGroup{}
	for q.Next() {
		ag := AuthGroup{
			_exists: true,
		}

		// scan
		err = q.Scan(&ag.ID, &ag.Name)
		if err != nil {
			return nil, err
		}

		res = append(res, &ag)
	}

	return res, nil
}

// AuthGroupByName retrieves a row from 'public.auth_group' as a AuthGroup.
//
// Generated from index 'auth_group_name_key'.
func AuthGroupByName(db XODB, name string) (*AuthGroup, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name ` +
		`FROM public.auth_group ` +
		`WHERE name = $1`

	// run query
	XOLog(sqlstr, name)
	ag := AuthGroup{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, name).Scan(&ag.ID, &ag.Name)
	if err != nil {
		return nil, err
	}

	return &ag, nil
}

// AuthGroupByID retrieves a row from 'public.auth_group' as a AuthGroup.
//
// Generated from index 'auth_group_pkey'.
func AuthGroupByID(db XODB, id int) (*AuthGroup, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name ` +
		`FROM public.auth_group ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	ag := AuthGroup{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&ag.ID, &ag.Name)
	if err != nil {
		return nil, err
	}

	return &ag, nil
}
