import * as React from 'react';
import { compose } from '@bem-react/core';
import {
    Button as ButtonDesktop,
    withSizeM,
    withViewDefault,
} from '@yandex/ui/Button/desktop';
import { Link } from 'react-router-dom';
import { HeaderPropsType } from './Header.typings';
import { cnHeader, headerCn, headerContainerCn } from './Header.consts';

import './index.scss';

const Button = compose(withSizeM, withViewDefault)(ButtonDesktop);

export const Header: React.FC<HeaderPropsType> = () => {
    return (
        <div className={headerCn}>
            <div className={headerContainerCn}>
                <Link to="/" className={cnHeader('Link')}>
                    <div className={cnHeader('Title')}>Шаг в будущее</div>
                </Link>
                <Link to="/login" className={cnHeader('Link')}>
                    <Button view="default" size="m">Войти</Button>
                </Link>
            </div>
        </div>
    );
};
