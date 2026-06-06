import { sleep } from 'k6';
import { randomItem } from './lib/utils.js';

import { addBook, getAuthor, getAuthorBooks, getBook, seedAuthors } from './lib/api.js';
import { defaultOptions } from './lib/config.js';

export const options = defaultOptions;

export function setup() {
    const fromEnv = (__ENV.K6_AUTHOR_IDS || '')
        .split(',')
        .map((id) => id.trim())
        .filter((id) => id.length > 0);

    let authorIds = fromEnv;
    if (authorIds.length === 0) {
        authorIds = seedAuthors(Number(__ENV.K6_SEED_AUTHORS || 3));
    }

    const bookIds = [];
    for (const authorId of authorIds) {
        const bookId = addBook(`seed-book-${authorId.slice(0, 8)}`, [authorId]);
        if (bookId) {
            bookIds.push(bookId);
        }
    }

    if (authorIds.length === 0 || bookIds.length === 0) {
        throw new Error('failed to seed data for read load test');
    }

    return { authorIds, bookIds };
}

export default function (data) {
    const authorId = randomItem(data.authorIds);
    const bookId = randomItem(data.bookIds);

    getAuthor(authorId);
    getBook(bookId);
    getAuthorBooks(authorId);

    sleep(0.5);
}
