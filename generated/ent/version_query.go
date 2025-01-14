// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/satisfactorymodding/smr-api/generated/ent/mod"
	"github.com/satisfactorymodding/smr-api/generated/ent/predicate"
	"github.com/satisfactorymodding/smr-api/generated/ent/version"
	"github.com/satisfactorymodding/smr-api/generated/ent/versiondependency"
	"github.com/satisfactorymodding/smr-api/generated/ent/versiontarget"
)

// VersionQuery is the builder for querying Version entities.
type VersionQuery struct {
	config
	ctx                     *QueryContext
	order                   []version.OrderOption
	inters                  []Interceptor
	predicates              []predicate.Version
	withMod                 *ModQuery
	withDependencies        *ModQuery
	withTargets             *VersionTargetQuery
	withVersionDependencies *VersionDependencyQuery
	modifiers               []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VersionQuery builder.
func (vq *VersionQuery) Where(ps ...predicate.Version) *VersionQuery {
	vq.predicates = append(vq.predicates, ps...)
	return vq
}

// Limit the number of records to be returned by this query.
func (vq *VersionQuery) Limit(limit int) *VersionQuery {
	vq.ctx.Limit = &limit
	return vq
}

// Offset to start from.
func (vq *VersionQuery) Offset(offset int) *VersionQuery {
	vq.ctx.Offset = &offset
	return vq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vq *VersionQuery) Unique(unique bool) *VersionQuery {
	vq.ctx.Unique = &unique
	return vq
}

// Order specifies how the records should be ordered.
func (vq *VersionQuery) Order(o ...version.OrderOption) *VersionQuery {
	vq.order = append(vq.order, o...)
	return vq
}

// QueryMod chains the current query on the "mod" edge.
func (vq *VersionQuery) QueryMod() *ModQuery {
	query := (&ModClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(version.Table, version.FieldID, selector),
			sqlgraph.To(mod.Table, mod.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, version.ModTable, version.ModColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDependencies chains the current query on the "dependencies" edge.
func (vq *VersionQuery) QueryDependencies() *ModQuery {
	query := (&ModClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(version.Table, version.FieldID, selector),
			sqlgraph.To(mod.Table, mod.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, version.DependenciesTable, version.DependenciesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTargets chains the current query on the "targets" edge.
func (vq *VersionQuery) QueryTargets() *VersionTargetQuery {
	query := (&VersionTargetClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(version.Table, version.FieldID, selector),
			sqlgraph.To(versiontarget.Table, versiontarget.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, version.TargetsTable, version.TargetsColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVersionDependencies chains the current query on the "version_dependencies" edge.
func (vq *VersionQuery) QueryVersionDependencies() *VersionDependencyQuery {
	query := (&VersionDependencyClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(version.Table, version.FieldID, selector),
			sqlgraph.To(versiondependency.Table, versiondependency.VersionColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, version.VersionDependenciesTable, version.VersionDependenciesColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Version entity from the query.
// Returns a *NotFoundError when no Version was found.
func (vq *VersionQuery) First(ctx context.Context) (*Version, error) {
	nodes, err := vq.Limit(1).All(setContextOp(ctx, vq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{version.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vq *VersionQuery) FirstX(ctx context.Context) *Version {
	node, err := vq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Version ID from the query.
// Returns a *NotFoundError when no Version ID was found.
func (vq *VersionQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = vq.Limit(1).IDs(setContextOp(ctx, vq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{version.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vq *VersionQuery) FirstIDX(ctx context.Context) string {
	id, err := vq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Version entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Version entity is found.
// Returns a *NotFoundError when no Version entities are found.
func (vq *VersionQuery) Only(ctx context.Context) (*Version, error) {
	nodes, err := vq.Limit(2).All(setContextOp(ctx, vq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{version.Label}
	default:
		return nil, &NotSingularError{version.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vq *VersionQuery) OnlyX(ctx context.Context) *Version {
	node, err := vq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Version ID in the query.
// Returns a *NotSingularError when more than one Version ID is found.
// Returns a *NotFoundError when no entities are found.
func (vq *VersionQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = vq.Limit(2).IDs(setContextOp(ctx, vq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{version.Label}
	default:
		err = &NotSingularError{version.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vq *VersionQuery) OnlyIDX(ctx context.Context) string {
	id, err := vq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Versions.
func (vq *VersionQuery) All(ctx context.Context) ([]*Version, error) {
	ctx = setContextOp(ctx, vq.ctx, ent.OpQueryAll)
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Version, *VersionQuery]()
	return withInterceptors[[]*Version](ctx, vq, qr, vq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vq *VersionQuery) AllX(ctx context.Context) []*Version {
	nodes, err := vq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Version IDs.
func (vq *VersionQuery) IDs(ctx context.Context) (ids []string, err error) {
	if vq.ctx.Unique == nil && vq.path != nil {
		vq.Unique(true)
	}
	ctx = setContextOp(ctx, vq.ctx, ent.OpQueryIDs)
	if err = vq.Select(version.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vq *VersionQuery) IDsX(ctx context.Context) []string {
	ids, err := vq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vq *VersionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vq.ctx, ent.OpQueryCount)
	if err := vq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vq, querierCount[*VersionQuery](), vq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vq *VersionQuery) CountX(ctx context.Context) int {
	count, err := vq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vq *VersionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vq.ctx, ent.OpQueryExist)
	switch _, err := vq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vq *VersionQuery) ExistX(ctx context.Context) bool {
	exist, err := vq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VersionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vq *VersionQuery) Clone() *VersionQuery {
	if vq == nil {
		return nil
	}
	return &VersionQuery{
		config:                  vq.config,
		ctx:                     vq.ctx.Clone(),
		order:                   append([]version.OrderOption{}, vq.order...),
		inters:                  append([]Interceptor{}, vq.inters...),
		predicates:              append([]predicate.Version{}, vq.predicates...),
		withMod:                 vq.withMod.Clone(),
		withDependencies:        vq.withDependencies.Clone(),
		withTargets:             vq.withTargets.Clone(),
		withVersionDependencies: vq.withVersionDependencies.Clone(),
		// clone intermediate query.
		sql:  vq.sql.Clone(),
		path: vq.path,
	}
}

// WithMod tells the query-builder to eager-load the nodes that are connected to
// the "mod" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VersionQuery) WithMod(opts ...func(*ModQuery)) *VersionQuery {
	query := (&ModClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withMod = query
	return vq
}

// WithDependencies tells the query-builder to eager-load the nodes that are connected to
// the "dependencies" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VersionQuery) WithDependencies(opts ...func(*ModQuery)) *VersionQuery {
	query := (&ModClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withDependencies = query
	return vq
}

// WithTargets tells the query-builder to eager-load the nodes that are connected to
// the "targets" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VersionQuery) WithTargets(opts ...func(*VersionTargetQuery)) *VersionQuery {
	query := (&VersionTargetClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withTargets = query
	return vq
}

// WithVersionDependencies tells the query-builder to eager-load the nodes that are connected to
// the "version_dependencies" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VersionQuery) WithVersionDependencies(opts ...func(*VersionDependencyQuery)) *VersionQuery {
	query := (&VersionDependencyClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withVersionDependencies = query
	return vq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Version.Query().
//		GroupBy(version.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vq *VersionQuery) GroupBy(field string, fields ...string) *VersionGroupBy {
	vq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VersionGroupBy{build: vq}
	grbuild.flds = &vq.ctx.Fields
	grbuild.label = version.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Version.Query().
//		Select(version.FieldCreatedAt).
//		Scan(ctx, &v)
func (vq *VersionQuery) Select(fields ...string) *VersionSelect {
	vq.ctx.Fields = append(vq.ctx.Fields, fields...)
	sbuild := &VersionSelect{VersionQuery: vq}
	sbuild.label = version.Label
	sbuild.flds, sbuild.scan = &vq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VersionSelect configured with the given aggregations.
func (vq *VersionQuery) Aggregate(fns ...AggregateFunc) *VersionSelect {
	return vq.Select().Aggregate(fns...)
}

func (vq *VersionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vq); err != nil {
				return err
			}
		}
	}
	for _, f := range vq.ctx.Fields {
		if !version.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vq.path != nil {
		prev, err := vq.path(ctx)
		if err != nil {
			return err
		}
		vq.sql = prev
	}
	return nil
}

func (vq *VersionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Version, error) {
	var (
		nodes       = []*Version{}
		_spec       = vq.querySpec()
		loadedTypes = [4]bool{
			vq.withMod != nil,
			vq.withDependencies != nil,
			vq.withTargets != nil,
			vq.withVersionDependencies != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Version).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Version{config: vq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(vq.modifiers) > 0 {
		_spec.Modifiers = vq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := vq.withMod; query != nil {
		if err := vq.loadMod(ctx, query, nodes, nil,
			func(n *Version, e *Mod) { n.Edges.Mod = e }); err != nil {
			return nil, err
		}
	}
	if query := vq.withDependencies; query != nil {
		if err := vq.loadDependencies(ctx, query, nodes,
			func(n *Version) { n.Edges.Dependencies = []*Mod{} },
			func(n *Version, e *Mod) { n.Edges.Dependencies = append(n.Edges.Dependencies, e) }); err != nil {
			return nil, err
		}
	}
	if query := vq.withTargets; query != nil {
		if err := vq.loadTargets(ctx, query, nodes,
			func(n *Version) { n.Edges.Targets = []*VersionTarget{} },
			func(n *Version, e *VersionTarget) { n.Edges.Targets = append(n.Edges.Targets, e) }); err != nil {
			return nil, err
		}
	}
	if query := vq.withVersionDependencies; query != nil {
		if err := vq.loadVersionDependencies(ctx, query, nodes,
			func(n *Version) { n.Edges.VersionDependencies = []*VersionDependency{} },
			func(n *Version, e *VersionDependency) {
				n.Edges.VersionDependencies = append(n.Edges.VersionDependencies, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vq *VersionQuery) loadMod(ctx context.Context, query *ModQuery, nodes []*Version, init func(*Version), assign func(*Version, *Mod)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Version)
	for i := range nodes {
		fk := nodes[i].ModID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(mod.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "mod_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (vq *VersionQuery) loadDependencies(ctx context.Context, query *ModQuery, nodes []*Version, init func(*Version), assign func(*Version, *Mod)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*Version)
	nids := make(map[string]map[*Version]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(version.DependenciesTable)
		s.Join(joinT).On(s.C(mod.FieldID), joinT.C(version.DependenciesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(version.DependenciesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(version.DependenciesPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullString)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullString).String
				inValue := values[1].(*sql.NullString).String
				if nids[inValue] == nil {
					nids[inValue] = map[*Version]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Mod](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "dependencies" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (vq *VersionQuery) loadTargets(ctx context.Context, query *VersionTargetQuery, nodes []*Version, init func(*Version), assign func(*Version, *VersionTarget)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Version)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(versiontarget.FieldVersionID)
	}
	query.Where(predicate.VersionTarget(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(version.TargetsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.VersionID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "version_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (vq *VersionQuery) loadVersionDependencies(ctx context.Context, query *VersionDependencyQuery, nodes []*Version, init func(*Version), assign func(*Version, *VersionDependency)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Version)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(versiondependency.FieldVersionID)
	}
	query.Where(predicate.VersionDependency(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(version.VersionDependenciesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.VersionID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "version_id" returned %v for node %v`, fk, n)
		}
		assign(node, n)
	}
	return nil
}

func (vq *VersionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vq.querySpec()
	if len(vq.modifiers) > 0 {
		_spec.Modifiers = vq.modifiers
	}
	_spec.Node.Columns = vq.ctx.Fields
	if len(vq.ctx.Fields) > 0 {
		_spec.Unique = vq.ctx.Unique != nil && *vq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vq.driver, _spec)
}

func (vq *VersionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(version.Table, version.Columns, sqlgraph.NewFieldSpec(version.FieldID, field.TypeString))
	_spec.From = vq.sql
	if unique := vq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vq.path != nil {
		_spec.Unique = true
	}
	if fields := vq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, version.FieldID)
		for i := range fields {
			if fields[i] != version.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if vq.withMod != nil {
			_spec.Node.AddColumnOnce(version.FieldModID)
		}
	}
	if ps := vq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vq *VersionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vq.driver.Dialect())
	t1 := builder.Table(version.Table)
	columns := vq.ctx.Fields
	if len(columns) == 0 {
		columns = version.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vq.sql != nil {
		selector = vq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vq.ctx.Unique != nil && *vq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range vq.modifiers {
		m(selector)
	}
	for _, p := range vq.predicates {
		p(selector)
	}
	for _, p := range vq.order {
		p(selector)
	}
	if offset := vq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (vq *VersionQuery) Modify(modifiers ...func(s *sql.Selector)) *VersionSelect {
	vq.modifiers = append(vq.modifiers, modifiers...)
	return vq.Select()
}

// VersionGroupBy is the group-by builder for Version entities.
type VersionGroupBy struct {
	selector
	build *VersionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vgb *VersionGroupBy) Aggregate(fns ...AggregateFunc) *VersionGroupBy {
	vgb.fns = append(vgb.fns, fns...)
	return vgb
}

// Scan applies the selector query and scans the result into the given value.
func (vgb *VersionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vgb.build.ctx, ent.OpQueryGroupBy)
	if err := vgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VersionQuery, *VersionGroupBy](ctx, vgb.build, vgb, vgb.build.inters, v)
}

func (vgb *VersionGroupBy) sqlScan(ctx context.Context, root *VersionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vgb.fns))
	for _, fn := range vgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vgb.flds)+len(vgb.fns))
		for _, f := range *vgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VersionSelect is the builder for selecting fields of Version entities.
type VersionSelect struct {
	*VersionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vs *VersionSelect) Aggregate(fns ...AggregateFunc) *VersionSelect {
	vs.fns = append(vs.fns, fns...)
	return vs
}

// Scan applies the selector query and scans the result into the given value.
func (vs *VersionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vs.ctx, ent.OpQuerySelect)
	if err := vs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VersionQuery, *VersionSelect](ctx, vs.VersionQuery, vs, vs.inters, v)
}

func (vs *VersionSelect) sqlScan(ctx context.Context, root *VersionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vs.fns))
	for _, fn := range vs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (vs *VersionSelect) Modify(modifiers ...func(s *sql.Selector)) *VersionSelect {
	vs.modifiers = append(vs.modifiers, modifiers...)
	return vs
}
