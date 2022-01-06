import * as React from 'react';
import { KeyValueProps } from '@components/KeyValue/KeyValue.typyings';
import { keyValueCn, keyValueKeyCn, keyValueValueCn } from '@components/KeyValue/KeyValue.consts';

import './index.scss';

export const KeyValue: React.FC<KeyValueProps> = (props) => {
    const {
        keyName,
        value,
    } = props;

    return (
        <div className={keyValueCn}>
            <div className={keyValueKeyCn}>{keyName}</div>
            <div className={keyValueValueCn}>{value}</div>
        </div>
    );
};
