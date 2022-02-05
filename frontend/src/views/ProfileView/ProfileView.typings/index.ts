import { UserContextConsumerProps } from '@views/BaseView/BaseView.typings';

export type RoleType = 'administrator' | 'student' | 'tutor';
export type UserType = {
    name: string;
    surname: string;
    login: string;
    role: RoleType;
    UserMail: string;
};
export type UserWrapper = UserType | null | undefined;

export type ProfileViewProps = UserContextConsumerProps;
export type UsersListType = Array<UserType> | null | undefined;

export type ProfileViewState = {
    isLogout?: boolean;
    usersList: UsersListType;
};
