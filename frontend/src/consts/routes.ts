const localhostAPI = 'http://localhost:3001/api';

const HostApi = localhostAPI;

export const RoutesClientApi = {
    profilePage: '/profile',
    loginPage: '/login',
    mainPage: '/',
    notFoundPage: '',
};

export const RoutesServerApi = {
    LoginApi: `${HostApi}/login`,
    CSRF: `${HostApi}/csrf`,
};
