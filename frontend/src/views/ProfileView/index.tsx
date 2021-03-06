import * as React from 'react';
import UserModel from '@models/UserModel';
import { ProfileViewProps, ProfileViewState, Roles } from './ProfileView.typings';
import { profileCn, profileLogoutCn, profileManageCn } from './ProfileView.const';
import { Navigate } from 'react-router-dom';
import { Link } from '@yandex/ui/Link/desktop/bundle';
import { ContentCard } from '@components/ContentCard';
import { ClientRoutes } from '@consts/routes';
import { ProfileInfo } from '@features/ProfileInfo';
import { List } from '@features/List';

import './index.scss';

export default class ProfileView extends React.Component<ProfileViewProps, ProfileViewState> {
    private userModel: UserModel;

    constructor(props: ProfileViewProps) {
        super(props);

        this.userModel = new UserModel();

        this.state = {
            isLogout: false,
            usersList: undefined,
        };
    }

    async componentDidMount() {
        await this.getUsersList();
    }

    onLogout = async () => {
        await this.userModel.logout();
        this.setState({ isLogout: true }, () => {
            this.props.setUser && this.props.setUser(null);
        });
    }

    getUsersList = async () => {
        const usersList = (await this.userModel.getProfileUsersList())?.
            filter(value => value.userMail !== this.props.user?.userMail).
            map(value => {
                return {
                    id: `${value.userID}`,
                    url: `profile/${value.userID}`,
                    title: `${value.userSurname} ${value.userName}`,
                    description: value.userMail,
                    actions: [
                        {
                            icon: '/icons/chat.svg',
                        }, {
                            icon: '/icons/trash.svg',
                            onClick: () => this.onDeleteUserFromList(`${value.userID}`),
                        },
                    ],
                };
        });
        this.setState({ usersList });
    }

    onDeleteUserFromList = (id: string) => {
        this.setState({ usersList: this.state.usersList?.filter(value => value.id !== id) });
    }

    render() {
        if (!this.props.user) {
            return <Navigate to={ClientRoutes.loginPage} />;
        }
        if (this.state.isLogout) {
            return <Navigate to="/" replace={true} />;
        }

        const {
            usersList,
        } = this.state;

        const {
            userName,
            userRole,
            userMail,
            userSurname,
        } = this.props.user;

        return (
            <div className={profileCn}>
                <ProfileInfo
                    name={userName}
                    surname={userSurname}
                    patronymic="??????????????????????"
                    login={userMail}
                    role={userRole}
                    avatar="/img/avatar-com.svg"
                    birthday="20.07.2007"
                    region="???????????????????? ??????????????, ????????????"
                />
                {userRole !== Roles.Student && <List
                    list={usersList}
                    icon='/icons/add-user.svg'
                    title="???????????????????????????????????? ????????????????????????"
                />}
                <ContentCard title="???????????????????? ??????????????????" icon="/icons/settings.svg" collapsed={false}>
                    <div className={profileManageCn}><Link className={profileLogoutCn} view="default" onClick={this.onLogout}>??????????</Link></div>
                </ContentCard>
            </div>
        );
    }
}

