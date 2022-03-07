import * as React from 'react';
import { classnames } from '@bem-react/classnames';
import { DisclaimerProps } from './Disclaimer.typings';
import { cnDisclaimer, disclaimerCn } from './Disclaimer.const';

import './index.scss';

export const Disclaimer: React.FC<DisclaimerProps> = (props) => {
    const {
        show,
        type,
        children,
    } = props;

    if (!show) {
        return null;
    }

    return (
        <div className={classnames(disclaimerCn, cnDisclaimer({ type }))}>
            {children}
        </div>
    );
};
