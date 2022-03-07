import { AuthResponse, ProfileResponse, ProfilesListResponse } from '@models/UserModel/UserModel.typings';
import { ProfilesListType, UserWrapper } from '@views/ProfileView/ProfileView.typings';

export const mapUserFromAuth = (body: AuthResponse): UserWrapper => {
    return {
        userID: body.UserID,
        userName: body.Username,
        userSurname: body.UserSurname,
        userRole: body.UserRole,
        userMail: body.UserMail,
    };
};

export const mapProfilesListFromGetProfiles = (body: ProfilesListResponse): ProfilesListType => {
    return body.map((value: ProfileResponse) => {
        return {
            userID: value.UserID,
            userName: value.UserName,
            userSurname: value.UserSurname,
            userRole: value.UserRole,
            userMail: value.UserMail,
        };
    });
};
