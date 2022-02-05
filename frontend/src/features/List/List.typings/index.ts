import { UsersListType } from '@views/ProfileView/ProfileView.typings';

export type ListProps = {
    usersList: UsersListType;
    type: 'projects' | 'users';
    title: string;
};
