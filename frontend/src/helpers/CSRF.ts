import { RoutesServerApi } from '@consts/routes';
import { fetchGET } from './fetcher';

export default class CSRF {
    static async getCSRF() {
        const response = await fetchGET(RoutesServerApi.CSRF);

        if (response && response.ok) {
            const token: string = (await response.json())['Token'];
            localStorage.setItem('X-CSRF-Token', token);
        }

        return response;
    }
}
