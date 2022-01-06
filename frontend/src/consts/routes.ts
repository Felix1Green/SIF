const localhostAPI = 'http://localhost:8080';

const HostApi = localhostAPI;

export const ClientRoutes = {
    registerPage: '/profile/register',
    profilePage: '/profile',
    loginPage: '/login',
    notFoundPage: '*',
    homePage: '/',
};

export const RoutesServerApi = {
    Login: `${HostApi}/login`,
    Logout: `${HostApi}/logout`,
    Register: `${HostApi}/logout`,
    CSRF: `${HostApi}/csrf`,
};
