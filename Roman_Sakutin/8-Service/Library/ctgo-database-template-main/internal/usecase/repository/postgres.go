package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/project/library/internal/entity"
)

var _ AuthorsRepository = (*postgresRepository)(nil)
var _ BooksRepository = (*postgresRepository)(nil)

type postgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *postgresRepository {
	return &postgresRepository{
		db: db,
	}
}

func (p *postgresRepository) CreateAuthor(ctx context.Context, author entity.Author) (entity.Author, error) {
	if _, ok := TxFromContext(ctx); ok {
		return p.insertAuthor(ctx, author)
	}

	tx, err := p.db.Begin(ctx)
	if err != nil {
		return entity.Author{}, err
	}
	defer func() { _ = tx.Rollback(ctx) }()

	inner := InjectTx(ctx, tx)
	res, err := p.insertAuthor(inner, author)
	if err != nil {
		return entity.Author{}, err
	}
	if err := tx.Commit(ctx); err != nil {
		return entity.Author{}, err
	}
	return res, nil
}

func (p *postgresRepository) insertAuthor(ctx context.Context, author entity.Author) (entity.Author, error) {
	const queryAuthor = `
INSERT INTO author (name)
VALUES ($1)
RETURNING id, created_at, updated_at
`
	result := entity.Author{ID: author.ID, Name: author.Name}
	err := p.q(ctx).QueryRow(ctx, queryAuthor, author.Name).Scan(&result.ID, &result.CreatedAt, &result.UpdatedAt)
	return result, mapPostgresErr(err)
}

func (p *postgresRepository) CreateBook(ctx context.Context, book entity.Book) (entity.Book, error) {
	if _, ok := TxFromContext(ctx); ok {
		return p.insertBook(ctx, book)
	}

	tx, err := p.db.Begin(ctx)
	if err != nil {
		return entity.Book{}, err
	}
	defer func() { _ = tx.Rollback(ctx) }()

	inner := InjectTx(ctx, tx)
	res, err := p.insertBook(inner, book)
	if err != nil {
		return entity.Book{}, err
	}
	if err := tx.Commit(ctx); err != nil {
		return entity.Book{}, err
	}
	return res, nil
}

func (p *postgresRepository) insertBook(ctx context.Context, book entity.Book) (entity.Book, error) {
	const createBook = `
INSERT INTO book (name)
VALUES ($1)
RETURNING id, created_at, updated_at
`
	result := entity.Book{Name: book.Name}
	err := p.q(ctx).QueryRow(ctx, createBook, book.Name).Scan(&result.ID, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return entity.Book{}, mapPostgresErr(err)
	}

	const createAuthorBook = `
INSERT INTO author_book (author_id, book_id)
VALUES ($1, $2)
`
	for _, authorID := range book.AuthorID {
		if _, err := p.q(ctx).Exec(ctx, createAuthorBook, authorID, result.ID); err != nil {
			return entity.Book{}, mapPostgresErr(err)
		}
	}

	result.AuthorID = book.AuthorID
	return result, nil
}

func (p *postgresRepository) GetAuthor(ctx context.Context, authorID string) (entity.Author, error) {
	const getAuthor = `
SELECT id, name, created_at, updated_at 
FROM author 
WHERE id = $1
`
	var result entity.Author
	err := p.db.QueryRow(ctx, getAuthor, authorID).Scan(
		&result.ID,
		&result.Name,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Author{}, entity.NotFound
		}
		return entity.Author{}, mapPostgresErr(err)
	}

	return result, nil
}

func (p *postgresRepository) GetBook(ctx context.Context, bookID string) (entity.Book, error) {
	tx, err := p.db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.RepeatableRead})
	if err != nil {
		return entity.Book{}, err
	}
	defer tx.Rollback(ctx)

	const getBook = `
SELECT b.id, b.name, b.created_at, b.updated_at,
       COALESCE((
           SELECT array_agg(ab.author_id::text ORDER BY ab.author_id)
           FROM author_book ab
           WHERE ab.book_id = b.id
       ), '{}')
FROM book b
WHERE b.id = $1
FOR SHARE OF b
`
	var result entity.Book
	err = tx.QueryRow(ctx, getBook, bookID).Scan(
		&result.ID,
		&result.Name,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.AuthorID,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Book{}, entity.NotFound
		}
		return entity.Book{}, mapPostgresErr(err)
	}

	if err = tx.Commit(ctx); err != nil {
		return entity.Book{}, err
	}

	return result, nil
}

func (p *postgresRepository) UpdateAuthor(ctx context.Context, author entity.Author) error {
	tx, err := p.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	const updateAuthor = `
UPDATE author 
SET name = $1
WHERE id = $2
`
	result, err := tx.Exec(ctx, updateAuthor, author.Name, author.ID)
	if err != nil {
		return mapPostgresErr(err)
	}
	if n := result.RowsAffected(); n == 0 {
		return entity.NotFound
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

func (p *postgresRepository) UpdateBook(ctx context.Context, book entity.Book) error {
	tx, err := p.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	authorIDs := append([]string(nil), book.AuthorID...)

	const lockBook = `
SELECT id
FROM book
WHERE id = $1
FOR UPDATE
`
	var lockedID string
	err = tx.QueryRow(ctx, lockBook, book.ID).Scan(&lockedID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.NotFound
		}
		return mapPostgresErr(err)
	}

	const lockAuthorBook = `
SELECT 1
FROM author_book
WHERE book_id = $1
FOR UPDATE
`
	rows, err := tx.Query(ctx, lockAuthorBook, book.ID)
	if err != nil {
		return mapPostgresErr(err)
	}
	rows.Close()

	const updateBook = `
UPDATE book 
SET name = $1
WHERE id = $2
`
	if _, err := tx.Exec(ctx, updateBook, book.Name, book.ID); err != nil {
		return mapPostgresErr(err)
	}

	const deleteAuthorBook = `
DELETE FROM author_book
WHERE book_id = $1
`
	if _, err := tx.Exec(ctx, deleteAuthorBook, book.ID); err != nil {
		return err
	}

	const insertAuthorBook = `
INSERT INTO author_book 
(author_id, book_id)
VALUES ($1, $2)
`
	for _, authorID := range authorIDs {
		if _, err := tx.Exec(ctx, insertAuthorBook, authorID, book.ID); err != nil {
			return mapPostgresErr(err)
		}
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

func (p *postgresRepository) GetBookByAuthorID(ctx context.Context, authorID string) ([]entity.Book, error) {
	const getAuthorBook = `
SELECT book_id
FROM author_book
WHERE author_id = $1
`
	rows, err := p.db.Query(ctx, getAuthorBook, authorID)
	if err != nil {
		return []entity.Book{}, err
	}
	defer rows.Close()

	slice := make([]string, 0)
	for rows.Next() {
		var bookID string
		if err := rows.Scan(&bookID); err != nil {
			return []entity.Book{}, err
		}
		slice = append(slice, bookID)
	}

	if err := rows.Err(); err != nil {
		return []entity.Book{}, err
	}

	var books = make([]entity.Book, 0)
	for _, bookID := range slice {
		result, err := p.GetBook(ctx, bookID)
		if err != nil {
			return []entity.Book{}, err
		}
		books = append(books, result)
	}

	return books, nil
}

func (p *postgresRepository) q(ctx context.Context) PoolTx {
	if tx, ok := TxFromContext(ctx); ok {
		return tx
	}
	return p.db
}
