import { UserContextConsumerProps } from '@views/BaseView/BaseView.typings';

export type LoginViewState = LoginFields & {
    showAlert: boolean;
    progress: boolean;
    isAuth: boolean;
};

export type LoginViewProps = UserContextConsumerProps;

export type LoginFields = {
    login: string;
    password: string;
}

export type LoginRequest = {
    username: string;
    password: string;
}
