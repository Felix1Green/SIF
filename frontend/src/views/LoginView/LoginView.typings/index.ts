import { UserContextConsumerProps } from '@views/BaseView/BaseView.typings';

export type LoginViewState = {
    login: string;
    password: string;
    showAlert: boolean;
    progress: boolean;
    isAuth: boolean;
};

export type LoginViewProps = UserContextConsumerProps;
