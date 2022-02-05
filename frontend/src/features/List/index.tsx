import * as React from 'react';
import { ListProps } from './List.typings';
import { ContentCard } from '@components/ContentCard';
import { Link } from 'react-router-dom';
import { ClientRoutes } from '@consts/routes';
import { Button } from '@yandex/ui/Button/desktop/bundle';
import { UsersListItem } from '@components/UsersListItem';
import { Spin } from '@yandex/ui/Spin/desktop/bundle';
import {
    listAppendCn,
    listCn,
    listContainerCn, listLoadingCn
} from '@features/List/List.consts';

import './index.scss';
import { Info } from '@components/Info';

export const List: React.FC<ListProps> = props => {
    const {
        usersList,
        title,
    } = props;

    return (
        <ContentCard className={listCn} title={title}>
            {
                (usersList === undefined) ?
                    <Spin className={listLoadingCn} progress view="default" size="l" />
                    :
                    (usersList && usersList.length !== 0) ?
                        <div className={listContainerCn}>
                            {usersList.map((value, index) =>
                                <UsersListItem key={index} name={value.UserMail} surname={value.surname} />)}
                        </div>
                        :
                        <Info show={true} type={'info'}>
                            Нет зарегистрированных пользователей
                        </Info>
            }
            <Link className={listAppendCn} to={ClientRoutes.registerPage}>
                <Button view="action" size="m">Добавить</Button>
            </Link>
        </ContentCard>
    );
};
