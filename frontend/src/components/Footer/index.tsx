import * as React from 'react';
import { FooterPropsType } from './Footer.typings';
import { footerCn } from './Footer.consts';

import './index.scss';

export const Footer: React.FC<FooterPropsType> = () => {
    return (
        <div className={footerCn}>
            <div>© Jigglypuff 2021</div>
        </div>
    );
};
