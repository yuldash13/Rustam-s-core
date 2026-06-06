import { check, sleep } from 'k6';
import { randomItem } from './lib/utils.js';

import { addBook, seedAuthors } from './lib/api.js';
import { defaultOptions } from './lib/config.js';

export const options = defaultOptions;

export function setup() {
    const fromEnv = (__ENV.K6_AUTHOR_IDS || '')
        .split(',')
        .map((id) => id.trim())
        .filter((id) => id.length > 0);

    if (fromEnv.length > 0) {
        return { authorIds: fromEnv };
    }

    const count = Number(__ENV.K6_SEED_AUTHORS || 5);
    const authorIds = seedAuthors(count);

    if (authorIds.length === 0) {
        throw new Error('failed to seed authors for book load test');
    }

    return { authorIds };
}

export default function (data) {
    const authorId = randomItem(data.authorIds);
    const bookId = addBook(`book-vu${__VU}-iter${__ITER}`, [authorId]);

    check(bookId, {
        'book id returned': (id) => typeof id === 'string' && id.length > 0,
    });

    sleep(1);
}
