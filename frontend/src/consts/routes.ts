const localhostAPI = 'http://localhost:8080';

const HostApi = localhostAPI;

export const ClientRoutes = {
    registerPage: '/register',
    conversationPage: '/conversation',
    profilePage: '/profile',
    loginPage: '/login',
    notFoundPage: '*',
    homePage: '/',
};

export const RoutesServerApi = {
    Auth: `${HostApi}/login`,
    Login: `${HostApi}/login`,
    Logout: `${HostApi}/logout`,
    Register: `${HostApi}/register`,
    Profile: `${HostApi}/profile`,
    Profiles: `${HostApi}/profiles`,
    CSRF: `${HostApi}/csrf`,
};
