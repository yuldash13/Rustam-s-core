package library

import (
	"context"
	"encoding/json"
	"strings"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"github.com/project/library/internal/entity"
	"github.com/project/library/internal/usecase/repository"
)

func (l *LibraryImpl) CreateBook(ctx context.Context, book entity.Book) (entity.Book, error) {
	span := trace.SpanFromContext(ctx)
	l.logger.Info("creating book",
		zap.String("trace_id", span.SpanContext().TraceID().String()),
	)

	if book.Name == "" {
		return entity.Book{}, entity.EmptyName
	}

	if !l.outboxEnabled {
		created, err := l.booksRepo.CreateBook(ctx, book)
		if err != nil {
			l.logger.Error("failed to create book",
				zap.String("trace_id", span.SpanContext().TraceID().String()),
				zap.String("handler", "AddBook"),
				zap.String("operation", "create_book"),
				zap.Error(err),
				zap.String("name", book.Name),
			)
			return entity.Book{}, err
		}
		l.logger.Info("book created",
			zap.String("trace_id", span.SpanContext().TraceID().String()),
			zap.String("handler", "AddBook"),
			zap.String("operation", "create_book"),
			zap.String("book_id", created.ID),
			zap.String("name", book.Name),
		)
		return created, nil
	}

	var created entity.Book
	err := l.transactor.WithTx(ctx, func(txCtx context.Context) error {
		var err error
		created, err = l.booksRepo.CreateBook(txCtx, book)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(map[string]string{"id": created.ID})
		if err != nil {
			return err
		}
		return l.outboxRepo.SendMessage(txCtx, created.ID, repository.OutboxKindBook, payload)
	})
	if err != nil {
		span.RecordError(err)
		return entity.Book{}, err
	}

	span.SetAttributes(attribute.String("book.id", created.ID))
	l.logger.Info("book created",
		zap.String("trace_id", span.SpanContext().TraceID().String()),
		zap.String("handler", "AddBook"),
		zap.String("operation", "create_book"),
		zap.String("book_id", created.ID),
		zap.String("author_id", strings.Join(created.AuthorID, ",")),
	)
	return created, nil
}

func (l *LibraryImpl) GetBook(ctx context.Context, bookID string) (entity.Book, error) {
	span := trace.SpanFromContext(ctx)
	l.logger.Info("getting book",
		zap.String("trace_id", span.SpanContext().TraceID().String()),
		zap.String("book_id", bookID),
	)

	if bookID == "" {
		return entity.Book{}, entity.EmptyID
	}

	book, err := l.booksRepo.GetBook(ctx, bookID)
	if err != nil {
		span.RecordError(err)
		l.logger.Error("failed to get book",
			zap.Error(err),
			zap.String("book_id", bookID),
		)
		return entity.Book{}, err
	}

	span.SetAttributes(attribute.String("book.id", bookID))
	l.logger.Info("book retrieved",
		zap.String("trace_id", span.SpanContext().TraceID().String()),
		zap.String("handler", "GetBookInfo"),
		zap.String("operation", "get_book"),
		zap.String("book_id", book.ID), zap.String("name", book.Name))
	return book, nil
}

func (l *LibraryImpl) UpdateBook(ctx context.Context, book entity.Book) error {
	span := trace.SpanFromContext(ctx)
	l.logger.Info("updating book",
		zap.String("trace_id", span.SpanContext().TraceID().String()),
		zap.String("book_id", book.ID),
		zap.String("name", book.Name),
	)

	if book.ID == "" {
		return entity.EmptyID
	}
	if book.Name == "" {
		return entity.EmptyName
	}

	err := l.booksRepo.UpdateBook(ctx, book)
	if err != nil {
		span.RecordError(err)
		l.logger.Error("failed to update book",
			zap.Error(err),
			zap.String("book_id", book.ID),
			zap.String("name", book.Name),
		)
		return err
	}

	span.SetAttributes(attribute.String("book.id", book.ID))
	l.logger.Info("book updated",
		zap.String("trace_id", span.SpanContext().TraceID().String()),
		zap.String("handler", "UpdateBook"),
		zap.String("operation", "update_book"),
		zap.String("book_id", book.ID),
		zap.String("name", book.Name),
	)
	return nil
}

func (l *LibraryImpl) GetBookByAuthorID(ctx context.Context, authorID string) ([]entity.Book, error) {
	span := trace.SpanFromContext(ctx)
	l.logger.Info("getting author books",
		zap.String("trace_id", span.SpanContext().TraceID().String()),
		zap.String("author_id", authorID),
	)

	if authorID == "" {
		return []entity.Book{}, entity.EmptyID
	}

	books, err := l.booksRepo.GetBookByAuthorID(ctx, authorID)
	if err != nil {
		span.RecordError(err)
		l.logger.Error("failed to get author books",
			zap.Error(err),
			zap.String("author_id", authorID),
		)
		return []entity.Book{}, err
	}

	span.SetAttributes(attribute.String("author.id", authorID))
	l.logger.Info("author books retrieved",
		zap.String("trace_id", span.SpanContext().TraceID().String()),
		zap.String("handler", "GetAuthorBooks"),
		zap.String("operation", "get_book_by_author_id"),
		zap.String("author_id", authorID),
	)
	return books, nil
}
