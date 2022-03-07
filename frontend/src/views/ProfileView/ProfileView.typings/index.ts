import { UserContextConsumerProps } from '@views/BaseView/BaseView.typings';
import { ListProps } from '@features/List/List.typings';

export type RoleType = 'Администратор' | 'Студент' | 'Тьютор';
export type UserType = {
    userID: number;
    userMail: string;
    userRole: RoleType;
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
