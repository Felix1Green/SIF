import * as React from 'react';
import Button from '@mui/material/Button';
import { HeaderPropsType } from './Header.typings';
import { cnHeader, headerCn } from './Header.consts';

import './index.scss';

export const Header: React.FC<HeaderPropsType> = () => {
    return (
        <div className={headerCn}>
            <div className={cnHeader('Title')}>Шаг в будущее</div>
            <Button variant="contained">Войти</Button>
        </div>
    );
};
