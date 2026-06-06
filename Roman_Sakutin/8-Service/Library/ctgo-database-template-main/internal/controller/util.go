package controller

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/project/library/generated/api/library"
	"github.com/project/library/internal/entity"
)

// TimeToTimestamp конвертирует time.Time в timestamppb.Timestamp
func TimeToTimestamp(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

// TimestampToTime конвертирует timestamppb.Timestamp в time.Time
func TimestampToTime(ts *timestamppb.Timestamp) time.Time {
	if ts == nil {
		return time.Time{}
	}
	return ts.AsTime()
}

// EntityAuthorToProto конвертирует entity.Author в generated.Author (для GetAuthorInfo)
// Но в protobuf нет отдельного типа Author, только в ответах используется
// Для GetAuthorInfoResponse используется структура с id и name
func EntityAuthorToGetAuthorInfoResponse(author entity.Author) *library.GetAuthorInfoResponse {
	return &library.GetAuthorInfoResponse{
		Id:   author.ID,
		Name: author.Name,
	}
}

// EntityBookToProto конвертирует entity.Book в generated.Book
func EntityBookToProto(book entity.Book) *library.Book {
	protoBook := &library.Book{
		Id:       book.ID,
		Name:     book.Name,
		AuthorId: book.AuthorID,
	}
	if !book.CreatedAt.IsZero() {
		protoBook.CreatedAt = timestamppb.New(book.CreatedAt)
	}
	if !book.UpdatedAt.IsZero() {
		protoBook.UpdatedAt = timestamppb.New(book.UpdatedAt)
	}
	return protoBook
}

// EntityBooksToProto конвертирует массив entity.Book в массив generated.Book
func EntityBooksToProto(books []entity.Book) []*library.Book {
	result := make([]*library.Book, 0, len(books))
	for _, book := range books {
		result = append(result, EntityBookToProto(book))
	}
	return result
}
