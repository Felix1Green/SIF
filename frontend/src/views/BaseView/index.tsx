import * as React from 'react';
import { BaseViewPropsType } from './BaseView.typings';
import { Header } from '../../components/Header';
import { Footer } from '../../components/Footer';

export const BaseView: React.FC<BaseViewPropsType> = () => {
    return (
        <>
            <Header />
            <div className={'Main'}></div>
            <Footer />
        </>
    );
};
