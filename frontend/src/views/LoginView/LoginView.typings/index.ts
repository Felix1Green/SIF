export type LoginViewState = {
    login: string;
    password: string;
    showAlert: boolean;
    progress: boolean;
    user: boolean;
};

export type LoginViewProps = Record<string, never>;
