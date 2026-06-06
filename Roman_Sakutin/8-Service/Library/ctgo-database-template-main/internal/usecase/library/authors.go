package library

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"encoding/json"

	"github.com/project/library/internal/entity"
	"github.com/project/library/internal/usecase/repository"
)

func (l *LibraryImpl) CreateAuthor(ctx context.Context, author entity.Author) (entity.Author, error) {
	span := trace.SpanFromContext(ctx)
	l.logger.Info("creating author",
		zap.String("trace_id", span.SpanContext().TraceID().String()),
		zap.String("name", author.Name),
	)

	if author.Name == "" {
		return entity.Author{}, entity.InvalidArgument
	}

	if !l.outboxEnabled {
		created, err := l.authorsRepo.CreateAuthor(ctx, author)
		if err != nil {
			l.logger.Error("failed to create author",
				zap.String("trace_id", span.SpanContext().TraceID().String()),
				zap.String("handler", "RegisterAuthor"),
				zap.String("operation", "create_author"),
				zap.Error(err),
				zap.String("name", author.Name),
			)
			return entity.Author{}, err
		}
		l.logger.Info("author created",
			zap.String("trace_id", span.SpanContext().TraceID().String()),
			zap.String("handler", "RegisterAuthor"),
			zap.String("operation", "create_author"),
			zap.String("author_id", created.ID),
			zap.String("name", author.Name),
		)
		return created, nil
	}

	var created entity.Author
	err := l.transactor.WithTx(ctx, func(txCtx context.Context) error {
		var err error
		created, err = l.authorsRepo.CreateAuthor(txCtx, author)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(map[string]string{"id": created.ID})
		if err != nil {
			return err
		}
		return l.outboxRepo.SendMessage(txCtx, created.ID, repository.OutboxKindAuthor, payload)
	})
	if err != nil {
		span.RecordError(err)
		l.logger.Error("failed to create author",
			zap.Error(err),
			zap.String("name", author.Name),
		)
		return entity.Author{}, err
	}

	span.SetAttributes(attribute.String("author.id", created.ID))
	l.logger.Info("author created",
		zap.String("trace_id", span.SpanContext().TraceID().String()),
		zap.String("handler", "RegisterAuthor"),
		zap.String("operation", "create_author"),
		zap.String("author_id", created.ID),
		zap.String("name", author.Name))
	return created, nil
}

func (l *LibraryImpl) GetAuthor(ctx context.Context, authorID string) (entity.Author, error) {
	span := trace.SpanFromContext(ctx)
	l.logger.Info("getting author",
		zap.String("trace_id", span.SpanContext().TraceID().String()),
		zap.String("author_id", authorID),
	)

	if authorID == "" {
		return entity.Author{}, entity.InvalidArgument
	}

	author, err := l.authorsRepo.GetAuthor(ctx, authorID)
	if err != nil {
		span.RecordError(err)
		l.logger.Error("failed to get author",
			zap.Error(err),
			zap.String("author_id", authorID),
		)
		return entity.Author{}, err
	}

	span.SetAttributes(attribute.String("author.id", authorID))
	l.logger.Info("author retrieved",
		zap.String("trace_id", span.SpanContext().TraceID().String()),
		zap.String("handler", "GetAuthorInfo"),
		zap.String("operation", "get_author"),
		zap.String("author_id", authorID),
	)
	return author, nil
}

func (l *LibraryImpl) UpdateAuthor(ctx context.Context, author entity.Author) error {
	span := trace.SpanFromContext(ctx)
	l.logger.Info("updating author",
		zap.String("trace_id", span.SpanContext().TraceID().String()),
		zap.String("author_id", author.ID),
		zap.String("name", author.Name),
	)

	if author.ID == "" {
		return entity.EmptyID
	}
	if author.Name == "" {
		return entity.EmptyName
	}

	err := l.authorsRepo.UpdateAuthor(ctx, author)
	if err != nil {
		span.RecordError(err)
		l.logger.Error("failed to update author",
			zap.Error(err),
			zap.String("author_id", author.ID),
			zap.String("name", author.Name),
		)
		return err
	}

	span.SetAttributes(attribute.String("author.id", author.ID))
	l.logger.Info("author updated",
		zap.String("trace_id", span.SpanContext().TraceID().String()),
		zap.String("handler", "ChangeAuthorInfo"),
		zap.String("operation", "update_author"),
		zap.String("author_id", author.ID),
		zap.String("name", author.Name),
	)
	return nil
}
