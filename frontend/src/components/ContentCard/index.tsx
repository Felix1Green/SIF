import * as React from 'react';
import { classnames } from '@bem-react/classnames';
import { ContentCardPropsType } from './ContentCard.typings';
import { contentCardCn, contentCardTitle } from './ContentCard.consts';

import './index.scss';

export const ContentCard: React.FC<ContentCardPropsType> = (props) => {
    const { children, className, title } = props;
    return (
        <div className={classnames(className, contentCardCn)}>
            { title && <div className={contentCardTitle}>{title}</div> }
            {children}
        </div>
    );
};
