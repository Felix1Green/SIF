import * as React from 'react';
import UserModel from '@models/UserModel';
import { ProfileViewProps, ProfileViewState } from './ProfileView.typings';
import {
    leftColumnCn,
    profileCn,
    rightColumnCn,
} from './ProfileView.consts';

import { Navigate } from 'react-router-dom';

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
        let usersList = await this.userModel.getProfileUsersList();
        if (usersList) {
            usersList = usersList.filter(value => {
                return value.UserMail !== this.props.user?.UserMail;
            });
        }
        this.setState({ usersList });
    }

    onLogout = async () => {
        await this.userModel.logout();
        this.setState({ isLogout: true }, () => {
            this.props.setUser && this.props.setUser(null);
        });
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
            name,
            surname,
            role,
            UserMail,
        } = this.props.user;

        return (
            <div className={profileCn}>
                <div className={leftColumnCn}>
                    <ProfileInfo
                        name={name}
                        surname={surname}
                        login={UserMail}
                        role={role}
                        onLogout={this.onLogout}
                    />
                </div>
                <div className={rightColumnCn}>
                    <List
                        usersList={usersList}
                        type="users"
                        title="Пользователь"
                    />
                </div>
            </div>
        );
    }
}

