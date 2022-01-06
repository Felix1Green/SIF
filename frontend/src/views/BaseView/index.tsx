import * as React from 'react';
import {
    Routes,
    Route
} from 'react-router-dom';
import { Header } from '@components/Header';
import { Footer } from '@components/Footer';
import LoginView from '@views/LoginView';
import ProfileView from '@views/ProfileView';
import NotFoundView from '@views/NotFoundView';
import { BaseViewPropsType } from './BaseView.typings';

import './index.scss';
import { ClientRoutes } from '@consts/routes';
import { UserContext } from '../../index';

export const BaseView: React.FC<BaseViewPropsType> = () => {
    return (
        <UserContext.Consumer>
            {({ user, setUser }) => (
                <>
                    <Header user={user}/>
                    <div className={'Main'}>
                        <Routes>
                            <Route path={ClientRoutes.homePage} element={<span>Hello</span>}/>
                            <Route path={ClientRoutes.loginPage} element={
                                <LoginView
                                    user={user}
                                    setUser={setUser}
                                />
                            }/>
                            <Route path={ClientRoutes.registerPage} element={<span>Register</span>}/>
                            <Route path={ClientRoutes.profilePage} element={
                                <ProfileView
                                    user={user}
                                    setUser={setUser}
                                />
                            }/>
                            <Route path={ClientRoutes.notFoundPage} element={<NotFoundView/>}/>
                        </Routes>
                    </div>
                    <Footer/>
                </>
            )}
        </UserContext.Consumer>
    );
};
