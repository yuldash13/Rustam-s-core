import http from 'k6/http';
import { check } from 'k6';

import { baseUrl, jsonHeaders, okStatus } from './config.js';

export function registerAuthor(name) {
    const res = http.post(
        `${baseUrl}/v1/library/author`,
        JSON.stringify({ name }),
        jsonHeaders,
    );

    check(res, {
        'register author status ok': okStatus,
    });

    if (!okStatus(res)) {
        return null;
    }

    const body = res.json();
    return body.id || null;
}

export function addBook(name, authorIds) {
    const res = http.post(
        `${baseUrl}/v1/library/book`,
        JSON.stringify({ name, authorIds }),
        jsonHeaders,
    );

    check(res, {
        'add book status ok': okStatus,
    });

    if (!okStatus(res)) {
        return null;
    }

    const body = res.json();
    return body.book ? body.book.id : null;
}

export function getAuthor(id) {
    const res = http.get(`${baseUrl}/v1/library/author/${id}`);

    check(res, {
        'get author status ok': (r) => r.status === 200,
    });

    return res;
}

export function getBook(id) {
    const res = http.get(`${baseUrl}/v1/library/book/${id}`);

    check(res, {
        'get book status ok': (r) => r.status === 200,
    });

    return res;
}

export function getAuthorBooks(authorId) {
    const res = http.get(`${baseUrl}/v1/library/author_books/${authorId}`);

    check(res, {
        'get author books status ok': (r) => r.status === 200,
    });

    return res;
}

export function seedAuthors(count) {
    const authorIds = [];

    for (let i = 0; i < count; i++) {
        const id = registerAuthor(`SeedAuthor${i}`);
        if (id) {
            authorIds.push(id);
        }
    }

    return authorIds;
}
