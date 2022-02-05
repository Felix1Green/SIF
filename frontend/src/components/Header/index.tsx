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
import { ClientRoutes } from '@consts/routes';
import './index.scss';

const Button = compose(withSizeM, withViewDefault)(ButtonDesktop);

export const Header: React.FC<HeaderPropsType> = (props) => {
    return (
        <div className={headerCn}>
            <div className={headerContainerCn}>
                <Link to="/" className={cnHeader('Link')}>
                    <div className={cnHeader('Title')}>Шаг в будущее</div>
                </Link>
                { (!props.user) ? (
                    <Link to="/login" className={cnHeader('Link')}>
                        <Button view="default" size="m">Войти</Button>
                    </Link>
                ) : (
                    <div>
                        <Link to={ClientRoutes.profilePage} className={cnHeader('Link')}>
                            <Button view="default" size="m">Профиль</Button>
                        </Link>
                        <Link to={ClientRoutes.conversationPage} className={cnHeader('Link')}>
                            <Button view="default" size="m" className={cnHeader('Messanger')}>Сообщения</Button>
                        </Link>
                    </div>
                )}
            </div>
        </div>
    );
};
