import { UserContextConsumerProps } from '@views/BaseView/BaseView.typings';

export type UserType = {
    name: string;
    surname: string;
    login: string;
    role: string;
};
export type UserWrapper = UserType | null;

export type ProfileViewProps = UserContextConsumerProps;
export type ProfileUsersListType = Array<UserType> | null;

export type ProfileViewState = {
    isLogout?: boolean;
    usersList: ProfileUsersListType;
};
