import { sleep } from 'k6';
import { randomItem } from './lib/utils.js';

import {
    addBook,
    getAuthor,
    getAuthorBooks,
    getBook,
    registerAuthor,
    seedAuthors,
} from './lib/api.js';
import { defaultOptions } from './lib/config.js';

export const options = defaultOptions;

export function setup() {
    const authorIds = seedAuthors(Number(__ENV.K6_SEED_AUTHORS || 3));
    const bookIds = [];

    for (const authorId of authorIds) {
        const bookId = addBook(`mixed-seed-${authorId.slice(0, 8)}`, [authorId]);
        if (bookId) {
            bookIds.push(bookId);
        }
    }

    if (authorIds.length === 0) {
        throw new Error('failed to seed authors for mixed load test');
    }

    return { authorIds, bookIds };
}

export default function (data) {
    const roll = Math.random();

    if (roll < 0.4) {
        registerAuthor(`author-${__VU}-${__ITER}`);
    } else if (roll < 0.7) {
        addBook(`book-${__VU}-${__ITER}`, [randomItem(data.authorIds)]);
    } else {
        const authorId = randomItem(data.authorIds);
        getAuthor(authorId);
        if (data.bookIds.length > 0) {
            getBook(randomItem(data.bookIds));
        }
        getAuthorBooks(authorId);
    }

    sleep(0.5);
}
