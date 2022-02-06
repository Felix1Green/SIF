import * as React from 'react';
import { Link, NavLink } from 'react-router-dom';
import { HeaderPropsType } from './Header.typings';
import {
    cnHeader,
    headerCn,
    headerBurgerMenuCn,
    headerNavigationCn,
    headerNavigationLinkCn, headerUserIconsCn, headerUserContainerCn, headerAuthCn
} from './Header.consts';
import { ClientRoutes } from '@consts/routes';
import { Icon } from '@yandex/ui/Icon/bundle';

import './index.scss';

const serviceTitle = 'Стартапы будущего';

const navigationTabs = [
    { name: 'Проекты', url: '/' },
    { name: 'Исследователи', url: '/' },
    { name: 'Тьюторы', url: '/' },
    { name: 'Вакансии', url: '/' },
];
export const Header: React.FC<HeaderPropsType> = (props) => {
    const burgerMenu = <Icon className={headerBurgerMenuCn} url={'/icons/burger-menu.svg'}/>;
    const messageIcon = <Icon className={headerUserIconsCn} url={'/icons/envelope.svg'} />;
    const userAvatar = <Icon className={headerUserIconsCn} url={props.user?.avatar ?? '/img/avatar-svgrepo-com.svg'} />;

    return (
        <div className={headerCn}>
            {burgerMenu}
            <Link to="/">
                <div className={cnHeader('Title')}>{serviceTitle}</div>
            </Link>
            <div className={headerNavigationCn}>
                {navigationTabs.map(value => {
                    return <NavLink to={value.url}><div className={headerNavigationLinkCn}>{value.name}</div></NavLink>;
                })}
            </div>
            <div className={headerUserContainerCn}>
                {(!props.user) ? (
                    <Link to="/login" className={headerAuthCn}>Войти</Link>
                ) : (
                    <>
                        <Link to={ClientRoutes.conversationPage}> {messageIcon} </Link>
                        <Link to={ClientRoutes.profilePage}> {userAvatar} </Link>
                    </>
                )}
            </div>
        </div>
    );
};
