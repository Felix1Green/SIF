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

export const BaseView: React.FC<BaseViewPropsType> = () => {
    return (
        <>
            <Header />
            <div className={'Main'}>
                <Routes>
                    <Route path="/" element={<span>Hello</span>} />
                    <Route path="/login" element={<LoginView />} />
                    <Route path="/register" element={<span>Register</span>} />
                    <Route path='*' element={<NotFoundView />} />
                </Routes>
            </div>
            <Footer />
        </>
    );
};
