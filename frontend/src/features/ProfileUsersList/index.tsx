import * as React from 'react';
import { ProfileUsersListProps } from './ProfileUsersList.typings';
import { ContentCard } from '@components/ContentCard';
import { Link } from 'react-router-dom';
import { ClientRoutes } from '@consts/routes';
import { Button } from '@yandex/ui/Button/desktop/bundle';
import { UsersList } from '@components/UsersListItem';
import { Spin } from '@yandex/ui/Spin/desktop/bundle';
import {
    profileUsersListAppendCn,
    profileUsersListCn,
    profileUsersListContainerCn, profileUsersListLoadingCn
} from '@features/ProfileUsersList/ProfileUsersList.consts';

import './index.scss';
import { Info } from '@components/Info';

export const ProfileUsersList: React.FC<ProfileUsersListProps> = props => {
    const {
        usersList,
    } = props;

    return (
        <ContentCard className={profileUsersListCn} title="Пользователи">
            {
                (usersList === null) ?
                    <Spin className={profileUsersListLoadingCn} progress view="default" size="l" />
                    :
                    (usersList.length !== 0) ?
                        <div className={profileUsersListContainerCn}>
                            {usersList.map((value, index) =>
                                <UsersList key={index} name={value.name} surname={value.surname} />)}
                        </div>
                        :
                        <Info show={true} type={'info'}>
                            Нет зарегистрированных пользователей
                        </Info>
            }
            <Link className={profileUsersListAppendCn} to={ClientRoutes.registerPage}>
                <Button view="action" size="m">Добавить</Button>
            </Link>
        </ContentCard>
    );
};
