import * as React from 'react';
import { Icon } from '@yandex/ui/Icon/desktop/bundle';
import { ListItemProps } from './ListItem.typings';
import {
    listItemCn,
    listItemTextCn,
    listItemTitleCn,
    listItemActionCn,
    listItemDescriptionCn,
} from './ListItem.const';

import './index.scss';

export const ListItem: React.FC<ListItemProps> = props => {
    const {
        title,
        description,
        actions,
    } = props;

    return (
        <div className={listItemCn}>
            <div className={listItemTextCn}>
                <span className={listItemTitleCn}>{title}</span>
                {description && <span className={listItemDescriptionCn}> · {description}</span>}
            </div>
            {actions && actions.map(((value, i) => (
                <Icon key={i} className={listItemActionCn} onClick={value.onClick} url={value.icon} />
            )))}
        </div>
    );
};
