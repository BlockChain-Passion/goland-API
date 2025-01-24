package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Item struct {
	Task   string
	Status string
}

type DB struct {
	pool *pgxpool.Pool
}

func New(user, password, host, dbname string, port int) (*DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, dbname)

	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("failed while connecting Database connection: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed while pinging Database connection: %w", err)
	}
	return &DB{pool: pool}, nil
}

func (db *DB) InsertItem(ctx context.Context, item Item) error {
	query := `INSERT INTO todo_items (task,status) VALUES ($1, $2)`
	_, err := db.pool.Exec(ctx, query, item.Task, item.Status)
	if err != nil {
		fmt.Printf("failed while inserting item into database: %v\n", err)
		return err
	}
	return nil
}

func (db *DB) GetAllItems(ctx context.Context) ([]Item, error) {
	query := `select task,status from todo_items`
	rows, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error while fetching rows from Db :%w", err)
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.Task, &item.Status); err != nil {
			return nil, fmt.Errorf("failed while fetching rows from Db :%w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed while fetching rows from Db :%w", err)
	}
	return items, nil
}

func (db *DB) Close() {
	db.pool.Close()
}
