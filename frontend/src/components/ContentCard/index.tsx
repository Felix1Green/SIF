import * as React from 'react';
import { ContentCardPropsType } from './ContentCard.typings';
import { contentCardCn } from './ContentCard.consts';

import './index.scss';

export const ContentCard: React.FC<ContentCardPropsType> = (props) => {
    const { children, className } = props;
    return (
        <div className={`${className} ${contentCardCn}`}>
            {children}
        </div>
    );
};
