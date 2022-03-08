import { ProfilesListType, Roles, UserWrapper } from '@views/ProfileView/ProfileView.typings';
import { fetchPOST, fetchGET } from '@src/helpers/fetcher';
import { RoutesServerApi } from '@consts/routes';
import { LoginRequest } from '@views/LoginView/LoginView.typings';
import { mapProfilesListFromGetProfiles, mapUserFromAuth } from '@models/UserModel/UserModel.helpers';
import { AuthResponse, ProfilesListResponse } from '@models/UserModel/UserModel.typings';
import { RegisterFields, RegisterRequest } from '@views/RegisterView/RegisterView.typings';

export default class UserModel {
    async login(login: string, password: string): Promise<boolean> {
        try {
            const response = await fetchPOST<LoginRequest>(RoutesServerApi.Login, {
                'username': login,
                'password': password,
            });

            // TODO: Замедление срабатывания метода для демонстрации загрузки
            await new Promise((resolve) => setTimeout(async () => {resolve('');}, 1000));

            if (!response || !response.ok) {
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

    async register(registerFields: RegisterFields) {
        const { name, surname, role, login, password } = registerFields;
        try {
            const response = await fetchPOST<RegisterRequest>(RoutesServerApi.Register, {
                'password': password,
                'userMail': login,
                'userName': name,
                'userRole': role,
                'userSurname': surname,
            });

            if (!response || !response.ok) {
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

            if (!response || !response.ok) {
                return false;
            }

            return response.ok;
        } catch (err) {
            return false;
        }
    }

    async getProfileUsersList(): Promise<ProfilesListType> {
        try {
            // TODO: Замедление срабатывания метода для демонстрации загрузки
            await new Promise((resolve) => setTimeout(async () => {resolve('hello');}, 500));

            const response = await fetchGET(RoutesServerApi.Profiles);

            if (!response || !response.ok) {
                return null;
            }

            const profilesResponse = await response.json() as ProfilesListResponse;
            return mapProfilesListFromGetProfiles(profilesResponse);
        } catch (err) {
            return null;
        }
    }

    async getUserProfile(): Promise<UserWrapper> {
        try {
            const response = await fetchGET(RoutesServerApi.Auth);

            if (!response || !response.ok) {
                return null;
            }

            const authResponse = await response.json() as AuthResponse;
            return mapUserFromAuth(authResponse);
        } catch (err) {
            return null;
        }
    }
}
