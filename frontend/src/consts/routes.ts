const localhostAPI = 'http://localhost:3001/api';

const HostApi = localhostAPI;

export const ClientRoutes = {
    registerPage: '/register',
    profilePage: '/profile',
    loginPage: '/login',
    notFoundPage: '*',
    homePage: '/',
};

export const RoutesServerApi = {
    LoginApi: `${HostApi}/login`,
    CSRF: `${HostApi}/csrf`,
};
