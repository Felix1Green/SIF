// import CSRF from './CSRF';

export const fetchPOST = async function (url: string, body?: object, headersOptions?: HeadersInit): Promise<Response> {
    return await fetch(url, {
        method: 'POST',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
            // 'X-CSRF-TOKEN': localStorage['X-CSRF-Token'],
            ... headersOptions,
        },
        body: JSON.stringify(body),
    });
};

export const fetchGET = async function (url: string): Promise<Response | undefined> {
    return await fetch(url, {
        method: 'GET',
        credentials: 'include',
    });
};
