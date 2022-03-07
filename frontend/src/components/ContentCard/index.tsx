import * as React from 'react';
import { Icon } from '@yandex/ui/Icon/desktop/bundle';
import { classnames } from '@bem-react/classnames';
import { ContentCardProps } from './ContentCard.typings';
import {
    cnContentCard,
    contentCardCn, contentCardCollapse,
    contentCardHeader,
    contentCardIcon,
    contentCardTitle,
} from './ContentCard.const';

import { Title } from '@components/Title';
import { useCollapse } from '@components/ContentCard/ContentCard.hooks';

import './index.scss';

export const ContentCard: React.FC<ContentCardProps> = (props) => {
    const { children, className, title, icon, collapsed } = props;
    const [ collapseState, setCollapseState ] = useCollapse(collapsed);

    const headerClassName = collapsed !== undefined ?
        cnContentCard('Header', { collapsing: true }) : contentCardHeader;

    return (
        <div className={classnames(className, contentCardCn)}>
            {title && <div className={headerClassName} onClick={setCollapseState}>
                {icon && <Icon className={contentCardIcon} url={icon} />}
                <Title className={contentCardTitle}>{title}</Title>
                {collapseState !== undefined && <Icon className={contentCardCollapse} url={'/icons/caret-down.svg'} />}
            </div>}
            {collapseState !== undefined ? collapseState && children : children}
        </div>
    );
};
