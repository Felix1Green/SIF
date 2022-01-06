import { UserWrapper } from '@views/ProfileView/ProfileView.typings';

export type BaseViewPropsType = {
};

export type UserContextConsumerProps = {
    user: UserWrapper;
    setUser?: (receivedUser: UserWrapper) => void;
};
