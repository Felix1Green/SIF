import { UserType, UserWrapper } from '@views/ProfileView/ProfileView.typings';
import { fetchDataPOST } from '@models/index';
import { RoutesServerApi } from '@consts/routes';

export default class UserModel {
    async login(login: string, password: string) {
        return fetchDataPOST(RoutesServerApi.Login, {
            'Email': login,
            'Password': password,
        });
    }

    async register(name: string, surname: string, login: string, password: string) {
        // TODO: Замедление срабатывания метода для демонстрации загрузки
        // await new Promise((resolve) => setTimeout(async () => {resolve('hello');}, 5000));

        return fetchDataPOST(RoutesServerApi.Register, {
            'Email': login,
            'Password': password,
        });
    }

    async logout() {
        return fetchDataPOST(RoutesServerApi.Logout);
    }

    async getUserProfile() {
        // TODO: Замедление срабатывания метода для демонстрации загрузки
        // await new Promise((resolve) => setTimeout(async () => {resolve('hello');}, 500));

        return new Promise<UserType>((resolve) => {
            resolve({
                login: 'antontintul@gmail.com',
                name: 'Антон',
                surname: 'Тинтул',
                role: 'Администратор',
            });
        });
    }

    async getProfileUsersList() {
        // TODO: Замедление срабатывания метода для демонстрации загрузки
        await new Promise((resolve) => setTimeout(async () => {resolve('hello');}, 500));

        return new Promise<Array<UserType>>((resolve) => {
            resolve([
                { name: 'Григорий', surname: 'Горбачев', login: 'griggor@yandex.ru', role: 'Administrator' },
                { name: 'Георгий', surname: 'Шашурин', login: 'georgshash@yandex.ru', role: 'Tutor' },
                { name: 'Василий', surname: 'Петров', login: 'vaspetr@yandex.ru', role: 'Student' },
                { name: 'Петр', surname: 'Попов', login: 'popovpetr@yandex.ru', role: 'Student' },
            ]);
        });
    }

    async isAuthorized() {
        // TODO: Замедление срабатывания метода для демонстрации загрузки
        // await new Promise((resolve) => setTimeout(async () => {resolve('hello');}, 500));

        return new Promise<UserWrapper>((resolve) => {
            resolve(null);
        });
    }
}
