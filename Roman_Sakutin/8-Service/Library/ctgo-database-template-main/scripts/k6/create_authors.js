import http from 'k6/http';
import { check, sleep } from 'k6';

import { baseUrl, defaultOptions, jsonHeaders, okStatus } from './lib/config.js';

export const options = defaultOptions;

export default function () {
    const name = `LoadAuthor${__VU}${__ITER}`;

    const res = http.post(
        `${baseUrl}/v1/library/author`,
        JSON.stringify({ name }),
        jsonHeaders,
    );

    check(res, {
        'status is 200 or 201': okStatus,
        'author id returned': (r) => {
            if (!okStatus(r)) {
                return false;
            }
            const body = r.json();
            return typeof body.id === 'string' && body.id.length > 0;
        },
    });

    sleep(1);
}
