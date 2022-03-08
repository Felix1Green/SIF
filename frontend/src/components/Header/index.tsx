import * as React from 'react';
import { Link, NavLink } from 'react-router-dom';
import { HeaderPropsType } from './Header.typings';
import {
    headerCn,
    headerTitleCn,
    headerBurgerMenuCn,
    headerNavigationCn,
    headerNavigationLinkCn, headerUserIconsCn, headerUserContainerCn, headerAuthCn
} from './Header.const';
import { ClientRoutes } from '@consts/routes';
import { Icon } from '@yandex/ui/Icon/bundle';

import './index.scss';
import { Title } from '@components/Title';

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
    const userAvatar = <Icon className={headerUserIconsCn} url="/img/avatar-com.svg" />;

    return (
        <div className={headerCn}>
            {burgerMenu}
            <Link to="/">
                <Title className={headerTitleCn}>{serviceTitle}</Title>
            </Link>
            <div className={headerNavigationCn}>
                {navigationTabs.map((value, i) => {
                    return <NavLink key={i} to={value.url}><div className={headerNavigationLinkCn}>{value.name}</div></NavLink>;
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
