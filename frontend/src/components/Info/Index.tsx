import * as React from 'react';
import { InfoPropsType } from './Info.typings';
import { cnInfo, infoCn, infoTextCn } from './Info.consts';

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
        <div className={`${infoCn} ${cnInfo({ type })}`}>
            <span className={infoTextCn}>
                {children}
            </span>
        </div>
    );
};
