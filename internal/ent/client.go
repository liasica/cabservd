// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/auroraride/cabservd/internal/ent/migrate"

	"github.com/auroraride/cabservd/internal/ent/bin"
	"github.com/auroraride/cabservd/internal/ent/cabinet"
	"github.com/auroraride/cabservd/internal/ent/console"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Bin is the client for interacting with the Bin builders.
	Bin *BinClient
	// Cabinet is the client for interacting with the Cabinet builders.
	Cabinet *CabinetClient
	// Console is the client for interacting with the Console builders.
	Console *ConsoleClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Bin = NewBinClient(c.config)
	c.Cabinet = NewCabinetClient(c.config)
	c.Console = NewConsoleClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Bin:     NewBinClient(cfg),
		Cabinet: NewCabinetClient(cfg),
		Console: NewConsoleClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Bin:     NewBinClient(cfg),
		Cabinet: NewCabinetClient(cfg),
		Console: NewConsoleClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Bin.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Bin.Use(hooks...)
	c.Cabinet.Use(hooks...)
	c.Console.Use(hooks...)
}

// BinClient is a client for the Bin schema.
type BinClient struct {
	config
}

// NewBinClient returns a client for the Bin from the given config.
func NewBinClient(c config) *BinClient {
	return &BinClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `bin.Hooks(f(g(h())))`.
func (c *BinClient) Use(hooks ...Hook) {
	c.hooks.Bin = append(c.hooks.Bin, hooks...)
}

// Create returns a builder for creating a Bin entity.
func (c *BinClient) Create() *BinCreate {
	mutation := newBinMutation(c.config, OpCreate)
	return &BinCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Bin entities.
func (c *BinClient) CreateBulk(builders ...*BinCreate) *BinCreateBulk {
	return &BinCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Bin.
func (c *BinClient) Update() *BinUpdate {
	mutation := newBinMutation(c.config, OpUpdate)
	return &BinUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BinClient) UpdateOne(b *Bin) *BinUpdateOne {
	mutation := newBinMutation(c.config, OpUpdateOne, withBin(b))
	return &BinUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BinClient) UpdateOneID(id uint64) *BinUpdateOne {
	mutation := newBinMutation(c.config, OpUpdateOne, withBinID(id))
	return &BinUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Bin.
func (c *BinClient) Delete() *BinDelete {
	mutation := newBinMutation(c.config, OpDelete)
	return &BinDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BinClient) DeleteOne(b *Bin) *BinDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *BinClient) DeleteOneID(id uint64) *BinDeleteOne {
	builder := c.Delete().Where(bin.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BinDeleteOne{builder}
}

// Query returns a query builder for Bin.
func (c *BinClient) Query() *BinQuery {
	return &BinQuery{
		config: c.config,
	}
}

// Get returns a Bin entity by its id.
func (c *BinClient) Get(ctx context.Context, id uint64) (*Bin, error) {
	return c.Query().Where(bin.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BinClient) GetX(ctx context.Context, id uint64) *Bin {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *BinClient) Hooks() []Hook {
	return c.hooks.Bin
}

// CabinetClient is a client for the Cabinet schema.
type CabinetClient struct {
	config
}

// NewCabinetClient returns a client for the Cabinet from the given config.
func NewCabinetClient(c config) *CabinetClient {
	return &CabinetClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `cabinet.Hooks(f(g(h())))`.
func (c *CabinetClient) Use(hooks ...Hook) {
	c.hooks.Cabinet = append(c.hooks.Cabinet, hooks...)
}

// Create returns a builder for creating a Cabinet entity.
func (c *CabinetClient) Create() *CabinetCreate {
	mutation := newCabinetMutation(c.config, OpCreate)
	return &CabinetCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Cabinet entities.
func (c *CabinetClient) CreateBulk(builders ...*CabinetCreate) *CabinetCreateBulk {
	return &CabinetCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Cabinet.
func (c *CabinetClient) Update() *CabinetUpdate {
	mutation := newCabinetMutation(c.config, OpUpdate)
	return &CabinetUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CabinetClient) UpdateOne(ca *Cabinet) *CabinetUpdateOne {
	mutation := newCabinetMutation(c.config, OpUpdateOne, withCabinet(ca))
	return &CabinetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CabinetClient) UpdateOneID(id uint64) *CabinetUpdateOne {
	mutation := newCabinetMutation(c.config, OpUpdateOne, withCabinetID(id))
	return &CabinetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Cabinet.
func (c *CabinetClient) Delete() *CabinetDelete {
	mutation := newCabinetMutation(c.config, OpDelete)
	return &CabinetDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CabinetClient) DeleteOne(ca *Cabinet) *CabinetDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CabinetClient) DeleteOneID(id uint64) *CabinetDeleteOne {
	builder := c.Delete().Where(cabinet.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CabinetDeleteOne{builder}
}

// Query returns a query builder for Cabinet.
func (c *CabinetClient) Query() *CabinetQuery {
	return &CabinetQuery{
		config: c.config,
	}
}

// Get returns a Cabinet entity by its id.
func (c *CabinetClient) Get(ctx context.Context, id uint64) (*Cabinet, error) {
	return c.Query().Where(cabinet.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CabinetClient) GetX(ctx context.Context, id uint64) *Cabinet {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CabinetClient) Hooks() []Hook {
	return c.hooks.Cabinet
}

// ConsoleClient is a client for the Console schema.
type ConsoleClient struct {
	config
}

// NewConsoleClient returns a client for the Console from the given config.
func NewConsoleClient(c config) *ConsoleClient {
	return &ConsoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `console.Hooks(f(g(h())))`.
func (c *ConsoleClient) Use(hooks ...Hook) {
	c.hooks.Console = append(c.hooks.Console, hooks...)
}

// Create returns a builder for creating a Console entity.
func (c *ConsoleClient) Create() *ConsoleCreate {
	mutation := newConsoleMutation(c.config, OpCreate)
	return &ConsoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Console entities.
func (c *ConsoleClient) CreateBulk(builders ...*ConsoleCreate) *ConsoleCreateBulk {
	return &ConsoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Console.
func (c *ConsoleClient) Update() *ConsoleUpdate {
	mutation := newConsoleMutation(c.config, OpUpdate)
	return &ConsoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ConsoleClient) UpdateOne(co *Console) *ConsoleUpdateOne {
	mutation := newConsoleMutation(c.config, OpUpdateOne, withConsole(co))
	return &ConsoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ConsoleClient) UpdateOneID(id uint64) *ConsoleUpdateOne {
	mutation := newConsoleMutation(c.config, OpUpdateOne, withConsoleID(id))
	return &ConsoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Console.
func (c *ConsoleClient) Delete() *ConsoleDelete {
	mutation := newConsoleMutation(c.config, OpDelete)
	return &ConsoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ConsoleClient) DeleteOne(co *Console) *ConsoleDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ConsoleClient) DeleteOneID(id uint64) *ConsoleDeleteOne {
	builder := c.Delete().Where(console.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ConsoleDeleteOne{builder}
}

// Query returns a query builder for Console.
func (c *ConsoleClient) Query() *ConsoleQuery {
	return &ConsoleQuery{
		config: c.config,
	}
}

// Get returns a Console entity by its id.
func (c *ConsoleClient) Get(ctx context.Context, id uint64) (*Console, error) {
	return c.Query().Where(console.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ConsoleClient) GetX(ctx context.Context, id uint64) *Console {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCabinet queries the cabinet edge of a Console.
func (c *ConsoleClient) QueryCabinet(co *Console) *CabinetQuery {
	query := &CabinetQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(console.Table, console.FieldID, id),
			sqlgraph.To(cabinet.Table, cabinet.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, console.CabinetTable, console.CabinetColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBin queries the bin edge of a Console.
func (c *ConsoleClient) QueryBin(co *Console) *BinQuery {
	query := &BinQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(console.Table, console.FieldID, id),
			sqlgraph.To(bin.Table, bin.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, console.BinTable, console.BinColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ConsoleClient) Hooks() []Hook {
	return c.hooks.Console
}
