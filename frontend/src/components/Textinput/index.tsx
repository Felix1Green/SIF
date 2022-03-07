import * as React from 'react';
import { Text } from '@yandex/ui/Text/desktop/bundle';
import {
    ITextinputProps,
    Textinput as TextinputBase
} from '@yandex/ui/Textinput/desktop/bundle';
import { inputCn, labelCn } from './Textinput.const';

import './index.scss';

export const Textinput: React.FC<ITextinputProps> = (props: ITextinputProps) => {
    const {
        required = true,
        value,
        onChange,
        label,
        type,
    } = props;

    return (
        <>
            { label && <Text className={labelCn}>{label}</Text>}
            <TextinputBase
                size="m"
                view="default"
                type={type}
                value={value}
                className={inputCn}
                onChange={onChange}
                required={required}
                hasClear={true}
            />
        </>
    );
};

