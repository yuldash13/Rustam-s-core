export const baseUrl = __ENV.K6_BASE_URL || 'http://localhost:8080';

export const jsonHeaders = {
    headers: {
        'Content-Type': 'application/json',
    },
};

export const defaultOptions = {
    vus: Number(__ENV.K6_VUS || 10),
    duration: __ENV.K6_DURATION || '30s',
};

export function okStatus(res) {
    return res.status === 200 || res.status === 201;
}
