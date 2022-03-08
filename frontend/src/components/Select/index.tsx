import * as React from 'react';
import { Text } from '@yandex/ui/Text/desktop/bundle';
import { Select as SelectBase } from '@yandex/ui/Select/desktop/bundle';
import { ISelectProps } from './Select.typings';
import { selectComponentCn, selectLabelCn } from './Select.const';

import './index.scss';

export const Select: React.FC<ISelectProps> = (props: ISelectProps) => {
    const { required = true, label, ...selectProps } = props;

    return (
        <>
            { label && <Text className={selectLabelCn}>{label}</Text>}
            <SelectBase
                size="m"
                view="default"
                className={selectComponentCn}
                required={required}
                {...selectProps}
            />
        </>
    );
};

