// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"hmdp/ent/migrate"

	"hmdp/ent/blog"
	"hmdp/ent/seckillvoucher"
	"hmdp/ent/shop"
	"hmdp/ent/shoptype"
	"hmdp/ent/user"
	"hmdp/ent/voucher"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Blog is the client for interacting with the Blog builders.
	Blog *BlogClient
	// SeckillVoucher is the client for interacting with the SeckillVoucher builders.
	SeckillVoucher *SeckillVoucherClient
	// Shop is the client for interacting with the Shop builders.
	Shop *ShopClient
	// ShopType is the client for interacting with the ShopType builders.
	ShopType *ShopTypeClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// Voucher is the client for interacting with the Voucher builders.
	Voucher *VoucherClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Blog = NewBlogClient(c.config)
	c.SeckillVoucher = NewSeckillVoucherClient(c.config)
	c.Shop = NewShopClient(c.config)
	c.ShopType = NewShopTypeClient(c.config)
	c.User = NewUserClient(c.config)
	c.Voucher = NewVoucherClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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
		ctx:            ctx,
		config:         cfg,
		Blog:           NewBlogClient(cfg),
		SeckillVoucher: NewSeckillVoucherClient(cfg),
		Shop:           NewShopClient(cfg),
		ShopType:       NewShopTypeClient(cfg),
		User:           NewUserClient(cfg),
		Voucher:        NewVoucherClient(cfg),
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
		ctx:            ctx,
		config:         cfg,
		Blog:           NewBlogClient(cfg),
		SeckillVoucher: NewSeckillVoucherClient(cfg),
		Shop:           NewShopClient(cfg),
		ShopType:       NewShopTypeClient(cfg),
		User:           NewUserClient(cfg),
		Voucher:        NewVoucherClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Blog.
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
	for _, n := range []interface{ Use(...Hook) }{
		c.Blog, c.SeckillVoucher, c.Shop, c.ShopType, c.User, c.Voucher,
	} {
		n.Use(hooks...)
	}
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	for _, n := range []interface{ Intercept(...Interceptor) }{
		c.Blog, c.SeckillVoucher, c.Shop, c.ShopType, c.User, c.Voucher,
	} {
		n.Intercept(interceptors...)
	}
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *BlogMutation:
		return c.Blog.mutate(ctx, m)
	case *SeckillVoucherMutation:
		return c.SeckillVoucher.mutate(ctx, m)
	case *ShopMutation:
		return c.Shop.mutate(ctx, m)
	case *ShopTypeMutation:
		return c.ShopType.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	case *VoucherMutation:
		return c.Voucher.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// BlogClient is a client for the Blog schema.
type BlogClient struct {
	config
}

// NewBlogClient returns a client for the Blog from the given config.
func NewBlogClient(c config) *BlogClient {
	return &BlogClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `blog.Hooks(f(g(h())))`.
func (c *BlogClient) Use(hooks ...Hook) {
	c.hooks.Blog = append(c.hooks.Blog, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `blog.Intercept(f(g(h())))`.
func (c *BlogClient) Intercept(interceptors ...Interceptor) {
	c.inters.Blog = append(c.inters.Blog, interceptors...)
}

// Create returns a builder for creating a Blog entity.
func (c *BlogClient) Create() *BlogCreate {
	mutation := newBlogMutation(c.config, OpCreate)
	return &BlogCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Blog entities.
func (c *BlogClient) CreateBulk(builders ...*BlogCreate) *BlogCreateBulk {
	return &BlogCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Blog.
func (c *BlogClient) Update() *BlogUpdate {
	mutation := newBlogMutation(c.config, OpUpdate)
	return &BlogUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BlogClient) UpdateOne(b *Blog) *BlogUpdateOne {
	mutation := newBlogMutation(c.config, OpUpdateOne, withBlog(b))
	return &BlogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BlogClient) UpdateOneID(id int64) *BlogUpdateOne {
	mutation := newBlogMutation(c.config, OpUpdateOne, withBlogID(id))
	return &BlogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Blog.
func (c *BlogClient) Delete() *BlogDelete {
	mutation := newBlogMutation(c.config, OpDelete)
	return &BlogDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BlogClient) DeleteOne(b *Blog) *BlogDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *BlogClient) DeleteOneID(id int64) *BlogDeleteOne {
	builder := c.Delete().Where(blog.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BlogDeleteOne{builder}
}

// Query returns a query builder for Blog.
func (c *BlogClient) Query() *BlogQuery {
	return &BlogQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeBlog},
		inters: c.Interceptors(),
	}
}

// Get returns a Blog entity by its id.
func (c *BlogClient) Get(ctx context.Context, id int64) (*Blog, error) {
	return c.Query().Where(blog.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BlogClient) GetX(ctx context.Context, id int64) *Blog {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *BlogClient) Hooks() []Hook {
	return c.hooks.Blog
}

// Interceptors returns the client interceptors.
func (c *BlogClient) Interceptors() []Interceptor {
	return c.inters.Blog
}

func (c *BlogClient) mutate(ctx context.Context, m *BlogMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&BlogCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&BlogUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&BlogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&BlogDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Blog mutation op: %q", m.Op())
	}
}

// SeckillVoucherClient is a client for the SeckillVoucher schema.
type SeckillVoucherClient struct {
	config
}

// NewSeckillVoucherClient returns a client for the SeckillVoucher from the given config.
func NewSeckillVoucherClient(c config) *SeckillVoucherClient {
	return &SeckillVoucherClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `seckillvoucher.Hooks(f(g(h())))`.
func (c *SeckillVoucherClient) Use(hooks ...Hook) {
	c.hooks.SeckillVoucher = append(c.hooks.SeckillVoucher, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `seckillvoucher.Intercept(f(g(h())))`.
func (c *SeckillVoucherClient) Intercept(interceptors ...Interceptor) {
	c.inters.SeckillVoucher = append(c.inters.SeckillVoucher, interceptors...)
}

// Create returns a builder for creating a SeckillVoucher entity.
func (c *SeckillVoucherClient) Create() *SeckillVoucherCreate {
	mutation := newSeckillVoucherMutation(c.config, OpCreate)
	return &SeckillVoucherCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SeckillVoucher entities.
func (c *SeckillVoucherClient) CreateBulk(builders ...*SeckillVoucherCreate) *SeckillVoucherCreateBulk {
	return &SeckillVoucherCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SeckillVoucher.
func (c *SeckillVoucherClient) Update() *SeckillVoucherUpdate {
	mutation := newSeckillVoucherMutation(c.config, OpUpdate)
	return &SeckillVoucherUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SeckillVoucherClient) UpdateOne(sv *SeckillVoucher) *SeckillVoucherUpdateOne {
	mutation := newSeckillVoucherMutation(c.config, OpUpdateOne, withSeckillVoucher(sv))
	return &SeckillVoucherUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SeckillVoucherClient) UpdateOneID(id uint64) *SeckillVoucherUpdateOne {
	mutation := newSeckillVoucherMutation(c.config, OpUpdateOne, withSeckillVoucherID(id))
	return &SeckillVoucherUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SeckillVoucher.
func (c *SeckillVoucherClient) Delete() *SeckillVoucherDelete {
	mutation := newSeckillVoucherMutation(c.config, OpDelete)
	return &SeckillVoucherDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SeckillVoucherClient) DeleteOne(sv *SeckillVoucher) *SeckillVoucherDeleteOne {
	return c.DeleteOneID(sv.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SeckillVoucherClient) DeleteOneID(id uint64) *SeckillVoucherDeleteOne {
	builder := c.Delete().Where(seckillvoucher.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SeckillVoucherDeleteOne{builder}
}

// Query returns a query builder for SeckillVoucher.
func (c *SeckillVoucherClient) Query() *SeckillVoucherQuery {
	return &SeckillVoucherQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSeckillVoucher},
		inters: c.Interceptors(),
	}
}

// Get returns a SeckillVoucher entity by its id.
func (c *SeckillVoucherClient) Get(ctx context.Context, id uint64) (*SeckillVoucher, error) {
	return c.Query().Where(seckillvoucher.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SeckillVoucherClient) GetX(ctx context.Context, id uint64) *SeckillVoucher {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryGetForm queries the getForm edge of a SeckillVoucher.
func (c *SeckillVoucherClient) QueryGetForm(sv *SeckillVoucher) *VoucherQuery {
	query := (&VoucherClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := sv.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(seckillvoucher.Table, seckillvoucher.FieldID, id),
			sqlgraph.To(voucher.Table, voucher.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, seckillvoucher.GetFormTable, seckillvoucher.GetFormColumn),
		)
		fromV = sqlgraph.Neighbors(sv.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SeckillVoucherClient) Hooks() []Hook {
	return c.hooks.SeckillVoucher
}

// Interceptors returns the client interceptors.
func (c *SeckillVoucherClient) Interceptors() []Interceptor {
	return c.inters.SeckillVoucher
}

func (c *SeckillVoucherClient) mutate(ctx context.Context, m *SeckillVoucherMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SeckillVoucherCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SeckillVoucherUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SeckillVoucherUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SeckillVoucherDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown SeckillVoucher mutation op: %q", m.Op())
	}
}

// ShopClient is a client for the Shop schema.
type ShopClient struct {
	config
}

// NewShopClient returns a client for the Shop from the given config.
func NewShopClient(c config) *ShopClient {
	return &ShopClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `shop.Hooks(f(g(h())))`.
func (c *ShopClient) Use(hooks ...Hook) {
	c.hooks.Shop = append(c.hooks.Shop, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `shop.Intercept(f(g(h())))`.
func (c *ShopClient) Intercept(interceptors ...Interceptor) {
	c.inters.Shop = append(c.inters.Shop, interceptors...)
}

// Create returns a builder for creating a Shop entity.
func (c *ShopClient) Create() *ShopCreate {
	mutation := newShopMutation(c.config, OpCreate)
	return &ShopCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Shop entities.
func (c *ShopClient) CreateBulk(builders ...*ShopCreate) *ShopCreateBulk {
	return &ShopCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Shop.
func (c *ShopClient) Update() *ShopUpdate {
	mutation := newShopMutation(c.config, OpUpdate)
	return &ShopUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ShopClient) UpdateOne(s *Shop) *ShopUpdateOne {
	mutation := newShopMutation(c.config, OpUpdateOne, withShop(s))
	return &ShopUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ShopClient) UpdateOneID(id int64) *ShopUpdateOne {
	mutation := newShopMutation(c.config, OpUpdateOne, withShopID(id))
	return &ShopUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Shop.
func (c *ShopClient) Delete() *ShopDelete {
	mutation := newShopMutation(c.config, OpDelete)
	return &ShopDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ShopClient) DeleteOne(s *Shop) *ShopDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ShopClient) DeleteOneID(id int64) *ShopDeleteOne {
	builder := c.Delete().Where(shop.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ShopDeleteOne{builder}
}

// Query returns a query builder for Shop.
func (c *ShopClient) Query() *ShopQuery {
	return &ShopQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeShop},
		inters: c.Interceptors(),
	}
}

// Get returns a Shop entity by its id.
func (c *ShopClient) Get(ctx context.Context, id int64) (*Shop, error) {
	return c.Query().Where(shop.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ShopClient) GetX(ctx context.Context, id int64) *Shop {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ShopClient) Hooks() []Hook {
	return c.hooks.Shop
}

// Interceptors returns the client interceptors.
func (c *ShopClient) Interceptors() []Interceptor {
	return c.inters.Shop
}

func (c *ShopClient) mutate(ctx context.Context, m *ShopMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ShopCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ShopUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ShopUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ShopDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Shop mutation op: %q", m.Op())
	}
}

// ShopTypeClient is a client for the ShopType schema.
type ShopTypeClient struct {
	config
}

// NewShopTypeClient returns a client for the ShopType from the given config.
func NewShopTypeClient(c config) *ShopTypeClient {
	return &ShopTypeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `shoptype.Hooks(f(g(h())))`.
func (c *ShopTypeClient) Use(hooks ...Hook) {
	c.hooks.ShopType = append(c.hooks.ShopType, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `shoptype.Intercept(f(g(h())))`.
func (c *ShopTypeClient) Intercept(interceptors ...Interceptor) {
	c.inters.ShopType = append(c.inters.ShopType, interceptors...)
}

// Create returns a builder for creating a ShopType entity.
func (c *ShopTypeClient) Create() *ShopTypeCreate {
	mutation := newShopTypeMutation(c.config, OpCreate)
	return &ShopTypeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ShopType entities.
func (c *ShopTypeClient) CreateBulk(builders ...*ShopTypeCreate) *ShopTypeCreateBulk {
	return &ShopTypeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ShopType.
func (c *ShopTypeClient) Update() *ShopTypeUpdate {
	mutation := newShopTypeMutation(c.config, OpUpdate)
	return &ShopTypeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ShopTypeClient) UpdateOne(st *ShopType) *ShopTypeUpdateOne {
	mutation := newShopTypeMutation(c.config, OpUpdateOne, withShopType(st))
	return &ShopTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ShopTypeClient) UpdateOneID(id int64) *ShopTypeUpdateOne {
	mutation := newShopTypeMutation(c.config, OpUpdateOne, withShopTypeID(id))
	return &ShopTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ShopType.
func (c *ShopTypeClient) Delete() *ShopTypeDelete {
	mutation := newShopTypeMutation(c.config, OpDelete)
	return &ShopTypeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ShopTypeClient) DeleteOne(st *ShopType) *ShopTypeDeleteOne {
	return c.DeleteOneID(st.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ShopTypeClient) DeleteOneID(id int64) *ShopTypeDeleteOne {
	builder := c.Delete().Where(shoptype.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ShopTypeDeleteOne{builder}
}

// Query returns a query builder for ShopType.
func (c *ShopTypeClient) Query() *ShopTypeQuery {
	return &ShopTypeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeShopType},
		inters: c.Interceptors(),
	}
}

// Get returns a ShopType entity by its id.
func (c *ShopTypeClient) Get(ctx context.Context, id int64) (*ShopType, error) {
	return c.Query().Where(shoptype.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ShopTypeClient) GetX(ctx context.Context, id int64) *ShopType {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ShopTypeClient) Hooks() []Hook {
	return c.hooks.ShopType
}

// Interceptors returns the client interceptors.
func (c *ShopTypeClient) Interceptors() []Interceptor {
	return c.inters.ShopType
}

func (c *ShopTypeClient) mutate(ctx context.Context, m *ShopTypeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ShopTypeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ShopTypeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ShopTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ShopTypeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ShopType mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uint64) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uint64) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uint64) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uint64) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// VoucherClient is a client for the Voucher schema.
type VoucherClient struct {
	config
}

// NewVoucherClient returns a client for the Voucher from the given config.
func NewVoucherClient(c config) *VoucherClient {
	return &VoucherClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `voucher.Hooks(f(g(h())))`.
func (c *VoucherClient) Use(hooks ...Hook) {
	c.hooks.Voucher = append(c.hooks.Voucher, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `voucher.Intercept(f(g(h())))`.
func (c *VoucherClient) Intercept(interceptors ...Interceptor) {
	c.inters.Voucher = append(c.inters.Voucher, interceptors...)
}

// Create returns a builder for creating a Voucher entity.
func (c *VoucherClient) Create() *VoucherCreate {
	mutation := newVoucherMutation(c.config, OpCreate)
	return &VoucherCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Voucher entities.
func (c *VoucherClient) CreateBulk(builders ...*VoucherCreate) *VoucherCreateBulk {
	return &VoucherCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Voucher.
func (c *VoucherClient) Update() *VoucherUpdate {
	mutation := newVoucherMutation(c.config, OpUpdate)
	return &VoucherUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VoucherClient) UpdateOne(v *Voucher) *VoucherUpdateOne {
	mutation := newVoucherMutation(c.config, OpUpdateOne, withVoucher(v))
	return &VoucherUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VoucherClient) UpdateOneID(id uint64) *VoucherUpdateOne {
	mutation := newVoucherMutation(c.config, OpUpdateOne, withVoucherID(id))
	return &VoucherUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Voucher.
func (c *VoucherClient) Delete() *VoucherDelete {
	mutation := newVoucherMutation(c.config, OpDelete)
	return &VoucherDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *VoucherClient) DeleteOne(v *Voucher) *VoucherDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *VoucherClient) DeleteOneID(id uint64) *VoucherDeleteOne {
	builder := c.Delete().Where(voucher.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VoucherDeleteOne{builder}
}

// Query returns a query builder for Voucher.
func (c *VoucherClient) Query() *VoucherQuery {
	return &VoucherQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeVoucher},
		inters: c.Interceptors(),
	}
}

// Get returns a Voucher entity by its id.
func (c *VoucherClient) Get(ctx context.Context, id uint64) (*Voucher, error) {
	return c.Query().Where(voucher.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VoucherClient) GetX(ctx context.Context, id uint64) *Voucher {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryGetMore queries the getMore edge of a Voucher.
func (c *VoucherClient) QueryGetMore(v *Voucher) *SeckillVoucherQuery {
	query := (&SeckillVoucherClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(voucher.Table, voucher.FieldID, id),
			sqlgraph.To(seckillvoucher.Table, seckillvoucher.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, voucher.GetMoreTable, voucher.GetMoreColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *VoucherClient) Hooks() []Hook {
	return c.hooks.Voucher
}

// Interceptors returns the client interceptors.
func (c *VoucherClient) Interceptors() []Interceptor {
	return c.inters.Voucher
}

func (c *VoucherClient) mutate(ctx context.Context, m *VoucherMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&VoucherCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&VoucherUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&VoucherUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&VoucherDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Voucher mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Blog, SeckillVoucher, Shop, ShopType, User, Voucher []ent.Hook
	}
	inters struct {
		Blog, SeckillVoucher, Shop, ShopType, User, Voucher []ent.Interceptor
	}
)
