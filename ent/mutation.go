// Code generated by ent, DO NOT EDIT.

package ent

import (
	"calling-bill/ent/call"
	"calling-bill/ent/predicate"
	"calling-bill/ent/user"
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCall = "Call"
	TypeUser = "User"
)

// CallMutation represents an operation that mutates the Call nodes in the graph.
type CallMutation struct {
	config
	op             Op
	typ            string
	id             *int
	duration       *int
	addduration    *int
	block_count    *int
	addblock_count *int
	clearedFields  map[string]struct{}
	user           map[int]struct{}
	removeduser    map[int]struct{}
	cleareduser    bool
	done           bool
	oldValue       func(context.Context) (*Call, error)
	predicates     []predicate.Call
}

var _ ent.Mutation = (*CallMutation)(nil)

// callOption allows management of the mutation configuration using functional options.
type callOption func(*CallMutation)

// newCallMutation creates new mutation for the Call entity.
func newCallMutation(c config, op Op, opts ...callOption) *CallMutation {
	m := &CallMutation{
		config:        c,
		op:            op,
		typ:           TypeCall,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCallID sets the ID field of the mutation.
func withCallID(id int) callOption {
	return func(m *CallMutation) {
		var (
			err   error
			once  sync.Once
			value *Call
		)
		m.oldValue = func(ctx context.Context) (*Call, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Call.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCall sets the old Call of the mutation.
func withCall(node *Call) callOption {
	return func(m *CallMutation) {
		m.oldValue = func(context.Context) (*Call, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CallMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CallMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CallMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CallMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Call.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetDuration sets the "duration" field.
func (m *CallMutation) SetDuration(i int) {
	m.duration = &i
	m.addduration = nil
}

// Duration returns the value of the "duration" field in the mutation.
func (m *CallMutation) Duration() (r int, exists bool) {
	v := m.duration
	if v == nil {
		return
	}
	return *v, true
}

// OldDuration returns the old "duration" field's value of the Call entity.
// If the Call object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CallMutation) OldDuration(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDuration is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDuration requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDuration: %w", err)
	}
	return oldValue.Duration, nil
}

// AddDuration adds i to the "duration" field.
func (m *CallMutation) AddDuration(i int) {
	if m.addduration != nil {
		*m.addduration += i
	} else {
		m.addduration = &i
	}
}

// AddedDuration returns the value that was added to the "duration" field in this mutation.
func (m *CallMutation) AddedDuration() (r int, exists bool) {
	v := m.addduration
	if v == nil {
		return
	}
	return *v, true
}

// ResetDuration resets all changes to the "duration" field.
func (m *CallMutation) ResetDuration() {
	m.duration = nil
	m.addduration = nil
}

// SetBlockCount sets the "block_count" field.
func (m *CallMutation) SetBlockCount(i int) {
	m.block_count = &i
	m.addblock_count = nil
}

// BlockCount returns the value of the "block_count" field in the mutation.
func (m *CallMutation) BlockCount() (r int, exists bool) {
	v := m.block_count
	if v == nil {
		return
	}
	return *v, true
}

// OldBlockCount returns the old "block_count" field's value of the Call entity.
// If the Call object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CallMutation) OldBlockCount(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBlockCount is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBlockCount requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBlockCount: %w", err)
	}
	return oldValue.BlockCount, nil
}

// AddBlockCount adds i to the "block_count" field.
func (m *CallMutation) AddBlockCount(i int) {
	if m.addblock_count != nil {
		*m.addblock_count += i
	} else {
		m.addblock_count = &i
	}
}

// AddedBlockCount returns the value that was added to the "block_count" field in this mutation.
func (m *CallMutation) AddedBlockCount() (r int, exists bool) {
	v := m.addblock_count
	if v == nil {
		return
	}
	return *v, true
}

// ResetBlockCount resets all changes to the "block_count" field.
func (m *CallMutation) ResetBlockCount() {
	m.block_count = nil
	m.addblock_count = nil
}

// AddUserIDs adds the "user" edge to the User entity by ids.
func (m *CallMutation) AddUserIDs(ids ...int) {
	if m.user == nil {
		m.user = make(map[int]struct{})
	}
	for i := range ids {
		m.user[ids[i]] = struct{}{}
	}
}

// ClearUser clears the "user" edge to the User entity.
func (m *CallMutation) ClearUser() {
	m.cleareduser = true
}

// UserCleared reports if the "user" edge to the User entity was cleared.
func (m *CallMutation) UserCleared() bool {
	return m.cleareduser
}

// RemoveUserIDs removes the "user" edge to the User entity by IDs.
func (m *CallMutation) RemoveUserIDs(ids ...int) {
	if m.removeduser == nil {
		m.removeduser = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.user, ids[i])
		m.removeduser[ids[i]] = struct{}{}
	}
}

// RemovedUser returns the removed IDs of the "user" edge to the User entity.
func (m *CallMutation) RemovedUserIDs() (ids []int) {
	for id := range m.removeduser {
		ids = append(ids, id)
	}
	return
}

// UserIDs returns the "user" edge IDs in the mutation.
func (m *CallMutation) UserIDs() (ids []int) {
	for id := range m.user {
		ids = append(ids, id)
	}
	return
}

// ResetUser resets all changes to the "user" edge.
func (m *CallMutation) ResetUser() {
	m.user = nil
	m.cleareduser = false
	m.removeduser = nil
}

// Where appends a list predicates to the CallMutation builder.
func (m *CallMutation) Where(ps ...predicate.Call) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *CallMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Call).
func (m *CallMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CallMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.duration != nil {
		fields = append(fields, call.FieldDuration)
	}
	if m.block_count != nil {
		fields = append(fields, call.FieldBlockCount)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CallMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case call.FieldDuration:
		return m.Duration()
	case call.FieldBlockCount:
		return m.BlockCount()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CallMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case call.FieldDuration:
		return m.OldDuration(ctx)
	case call.FieldBlockCount:
		return m.OldBlockCount(ctx)
	}
	return nil, fmt.Errorf("unknown Call field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CallMutation) SetField(name string, value ent.Value) error {
	switch name {
	case call.FieldDuration:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDuration(v)
		return nil
	case call.FieldBlockCount:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBlockCount(v)
		return nil
	}
	return fmt.Errorf("unknown Call field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CallMutation) AddedFields() []string {
	var fields []string
	if m.addduration != nil {
		fields = append(fields, call.FieldDuration)
	}
	if m.addblock_count != nil {
		fields = append(fields, call.FieldBlockCount)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CallMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case call.FieldDuration:
		return m.AddedDuration()
	case call.FieldBlockCount:
		return m.AddedBlockCount()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CallMutation) AddField(name string, value ent.Value) error {
	switch name {
	case call.FieldDuration:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddDuration(v)
		return nil
	case call.FieldBlockCount:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddBlockCount(v)
		return nil
	}
	return fmt.Errorf("unknown Call numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CallMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CallMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CallMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Call nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CallMutation) ResetField(name string) error {
	switch name {
	case call.FieldDuration:
		m.ResetDuration()
		return nil
	case call.FieldBlockCount:
		m.ResetBlockCount()
		return nil
	}
	return fmt.Errorf("unknown Call field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CallMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.user != nil {
		edges = append(edges, call.EdgeUser)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CallMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case call.EdgeUser:
		ids := make([]ent.Value, 0, len(m.user))
		for id := range m.user {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CallMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removeduser != nil {
		edges = append(edges, call.EdgeUser)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CallMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case call.EdgeUser:
		ids := make([]ent.Value, 0, len(m.removeduser))
		for id := range m.removeduser {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CallMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cleareduser {
		edges = append(edges, call.EdgeUser)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CallMutation) EdgeCleared(name string) bool {
	switch name {
	case call.EdgeUser:
		return m.cleareduser
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CallMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Call unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CallMutation) ResetEdge(name string) error {
	switch name {
	case call.EdgeUser:
		m.ResetUser()
		return nil
	}
	return fmt.Errorf("unknown Call edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	username      *string
	clearedFields map[string]struct{}
	calls         map[int]struct{}
	removedcalls  map[int]struct{}
	clearedcalls  bool
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUsername sets the "username" field.
func (m *UserMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *UserMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername resets all changes to the "username" field.
func (m *UserMutation) ResetUsername() {
	m.username = nil
}

// AddCallIDs adds the "calls" edge to the Call entity by ids.
func (m *UserMutation) AddCallIDs(ids ...int) {
	if m.calls == nil {
		m.calls = make(map[int]struct{})
	}
	for i := range ids {
		m.calls[ids[i]] = struct{}{}
	}
}

// ClearCalls clears the "calls" edge to the Call entity.
func (m *UserMutation) ClearCalls() {
	m.clearedcalls = true
}

// CallsCleared reports if the "calls" edge to the Call entity was cleared.
func (m *UserMutation) CallsCleared() bool {
	return m.clearedcalls
}

// RemoveCallIDs removes the "calls" edge to the Call entity by IDs.
func (m *UserMutation) RemoveCallIDs(ids ...int) {
	if m.removedcalls == nil {
		m.removedcalls = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.calls, ids[i])
		m.removedcalls[ids[i]] = struct{}{}
	}
}

// RemovedCalls returns the removed IDs of the "calls" edge to the Call entity.
func (m *UserMutation) RemovedCallsIDs() (ids []int) {
	for id := range m.removedcalls {
		ids = append(ids, id)
	}
	return
}

// CallsIDs returns the "calls" edge IDs in the mutation.
func (m *UserMutation) CallsIDs() (ids []int) {
	for id := range m.calls {
		ids = append(ids, id)
	}
	return
}

// ResetCalls resets all changes to the "calls" edge.
func (m *UserMutation) ResetCalls() {
	m.calls = nil
	m.clearedcalls = false
	m.removedcalls = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.username != nil {
		fields = append(fields, user.FieldUsername)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldUsername:
		return m.Username()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldUsername:
		return m.OldUsername(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldUsername:
		m.ResetUsername()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.calls != nil {
		edges = append(edges, user.EdgeCalls)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeCalls:
		ids := make([]ent.Value, 0, len(m.calls))
		for id := range m.calls {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedcalls != nil {
		edges = append(edges, user.EdgeCalls)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeCalls:
		ids := make([]ent.Value, 0, len(m.removedcalls))
		for id := range m.removedcalls {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedcalls {
		edges = append(edges, user.EdgeCalls)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeCalls:
		return m.clearedcalls
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeCalls:
		m.ResetCalls()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
