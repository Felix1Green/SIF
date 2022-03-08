import { UserContextConsumerProps } from '@views/BaseView/BaseView.typings';
import { ListProps } from '@features/List/List.typings';

export enum Roles {
    Default = '',
    Administrator = 'Администратор',
    Student = 'Студент',
    Tutor = 'Тьютор',
}

export type UserType = {
    userID: number;
    userMail: string;
    userRole: Roles;
    userSurname: string;
    userName: string;
};
export type UserWrapper = UserType | null | undefined;

export type ProfileViewProps = UserContextConsumerProps;
export type ProfilesListType = Array<UserType> | null | undefined;

export type ProfileViewState = {
    isLogout?: boolean;
    usersList?: ListProps['list'];
};
