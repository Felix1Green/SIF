const localhostAPI = 'http://localhost:8080';

const HostApi = localhostAPI;

export const ClientRoutes = {
    registerPage: '/register',
    profilePage: '/profile',
    loginPage: '/login',
    notFoundPage: '*',
    homePage: '/',
};

export const RoutesServerApi = {
    Login: `${HostApi}/login`,
    Logout: `${HostApi}/logout`,
    Register: `${HostApi}/register`,
    CSRF: `${HostApi}/csrf`,
};
