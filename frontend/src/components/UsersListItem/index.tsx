import * as React from 'react';
import { UsersListItemProps } from './UsersListItem.typings';
import { usersListItemCn } from './UsersListItem.consts';

import './index.scss';

export const UsersListItem: React.FC<UsersListItemProps> = props => {
    const {
        name,
        surname,
    } = props;

    return (
        <div className={usersListItemCn}>
            {name} {surname}
        </div>
    );
};
