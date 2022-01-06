import * as React from 'react';
import { classnames } from '@bem-react/classnames';
import { InfoPropsType } from './Info.typings';
import { cnInfo, infoCn } from './Info.consts';

import './index.scss';

export const Info: React.FC<InfoPropsType> = (props) => {
    const {
        show,
        type,
        children,
    } = props;

    if (!show) {
        return null;
    }

    return (
        <div className={classnames(infoCn, cnInfo({ type }))}>
            {children}
        </div>
    );
};
