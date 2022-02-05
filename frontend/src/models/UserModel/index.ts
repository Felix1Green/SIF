import { UsersListType, UserType, UserWrapper } from '@views/ProfileView/ProfileView.typings';
import { fetchPOST, fetchGET } from '@src/helpers/fetcher';
import { RoutesServerApi } from '@consts/routes';

export default class UserModel {
    async login(login: string, password: string): Promise<boolean> {
        try {
            const response = await fetchPOST(RoutesServerApi.Login, {
                'Username': login,
                'Password': password,
            });

            // TODO: Замедление срабатывания метода для демонстрации загрузки
            await new Promise((resolve) => setTimeout(async () => {resolve('');}, 1000));

            if (!response || (response && !response.ok)) {
                return false;
            }

            // TODO: Временное решение пока с бэка не будет приходить результат в status
            try {
                const answer = await response.json();
                if (answer && answer.ErrorCode) {
                    return false;
                }
            } catch (err) {
                return response.ok;
            }

            return response.ok;
        } catch (err) {
            return false;
        }
    }

    async register(name: string, surname: string, login: string, password: string) {
        try {
            const response = await fetchPOST(RoutesServerApi.Register, {
                'UserMail': login,
                'Password': password,
            });

            if (!response || (response && !response.ok)) {
                return false;
            }

            return response.ok;
        } catch (err) {
            return false;
        }
    }

    async logout() {
        try {
            const response = await fetchPOST(RoutesServerApi.Logout);

            if (!response || (response && !response.ok)) {
                return false;
            }

            return response.ok;
        } catch (err) {
            return false;
        }
    }

    async getProfileUsersList(): Promise<UsersListType> {
        try {
            // TODO: Замедление срабатывания метода для демонстрации загрузки
            await new Promise((resolve) => setTimeout(async () => {resolve('hello');}, 500));

            const response = await fetchGET(RoutesServerApi.Profiles);

            if (!response || (response && !response.ok)) {
                return null;
            }

            return await response.json();
        } catch (err) {
            return null;
        }
    }

    async getUserProfile(): Promise<UserWrapper> {
        try {
            const response = await fetchGET(RoutesServerApi.Auth);

            if (!response || (response && !response.ok)) {
                return null;
            }

            return response.json();
        } catch (err) {
            return null;
        }
    }
}
