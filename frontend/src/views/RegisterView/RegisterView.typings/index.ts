import { UserContextConsumerProps } from '@views/BaseView/BaseView.typings';
import { RoleType } from '@views/ProfileView/ProfileView.typings';

export type RegisterViewProps = UserContextConsumerProps;
export type RegisterViewState = RegisterFields & {
    progress: boolean;
    showAlert: boolean;
    isSuccess: boolean;
}

export type RegisterFields = {
    name: string;
    surname: string;
    login: string;
    password: string;
}

export type RegisterRequest = {
    userName: string;
    userSurname: string;
    userMail: string;
    userRole: RoleType;
    password: string;
}
