import Statuses from '../consts/statuses';
// import CSRF from '../helpers/CSRF';

export const fetchDataPOST = async function (url: string, body?: object, headersOptions?: HeadersInit): Promise<Response> {
    const response = await fetch(url, {
        method: 'POST',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
            // 'X-CSRF-TOKEN': localStorage['X-CSRF-Token'],
            ... headersOptions,
        },
        body: JSON.stringify(body),
    });

    if (!response.ok) {
        if (response.status === Statuses.Forbidden) {
            // await CSRF.getCSRF();
            await fetchDataPOST(url, body, headersOptions);
        }
    }

    return response;
};

export const fetchGET = async function (url: string): Promise<Response | undefined> {
    const response = await fetch(url, {
        method: 'GET',
        credentials: 'include',
    });

    if (response.ok) {
        return response;
    }
};
