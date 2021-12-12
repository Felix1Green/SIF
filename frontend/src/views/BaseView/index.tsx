import * as React from 'react';
import {
    Routes,
    Route
} from 'react-router-dom';
import { Header } from '@components/Header';
import { Footer } from '@components/Footer';
import LoginView from '@views/LoginView';
import NotFoundView from '@views/NotFoundView';
import { BaseViewPropsType } from './BaseView.typings';

import './index.scss';
import { ClientRoutes } from '@consts/routes';

export const BaseView: React.FC<BaseViewPropsType> = () => {
    return (
        <>
            <Header />
            <div className={'Main'}>
                <Routes>
                    <Route path={ClientRoutes.homePage} element={<span>Hello</span>} />
                    <Route path={ClientRoutes.loginPage} element={<LoginView />} />
                    <Route path={ClientRoutes.registerPage} element={<span>Register</span>} />
                    <Route path={ClientRoutes.profilePage} element={<span>Profile</span>} />
                    <Route path={ClientRoutes.notFoundPage} element={<NotFoundView />} />
                </Routes>
            </div>
            <Footer />
        </>
    );
};
