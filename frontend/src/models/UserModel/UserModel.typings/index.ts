import { Roles } from '@views/ProfileView/ProfileView.typings';

export type AuthResponse = {
    UserID: number;
    UserMail: string;
    UserRole: Roles;
    UserSurname: string;
    Username: string;
}

export type ProfileResponse = {
    UserID: number;
    UserMail: string;
    UserRole: Roles;
    UserSurname: string;
    UserName: string;
}

export type ProfilesListResponse = Array<ProfileResponse>;
