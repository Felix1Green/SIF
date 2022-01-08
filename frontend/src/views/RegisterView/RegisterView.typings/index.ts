import { UserContextConsumerProps } from '@views/BaseView/BaseView.typings';

export type RegisterViewProps = UserContextConsumerProps;
export type RegisterViewState = {
    name: string;
    surname: string;
    login: string;
    password: string;
    progress: boolean;
    showAlert: boolean;
    isSuccess: boolean;
}
