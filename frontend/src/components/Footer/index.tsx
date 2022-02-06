import * as React from 'react';
import { FooterPropsType } from './Footer.typings';
import {
    footerCn, footerContactCn,
    footerContactsCn,
    footerCopyrightCn, footerIconCn,
    footerSocialCn,
} from './Footer.consts';

import './index.scss';
import { headerBurgerMenuCn } from '@components/Header/Header.consts';
import { Icon } from '@yandex/ui/Icon/bundle';

const contacts = [
    { name: 'apfn@step-into-the-future.ru', url: 'mailto:apfn@step-into-the-future.ru' },
    { name: 'Программа "Шаг в будущее"', url: 'http://www.step-into-the-future.ru' },
    { name: 'Форум "Шаг в будущее"', url: 'https://шагвбудущее.рф/' },
    { name: 'МГТУ им. Н.Э.Баумана', url: 'https://bmstu.ru/' },
];
const social = [
    { iconUrl: '/icons/social/vk.svg', url: 'https://vk.com/officestep' },
    { iconUrl: '/icons/social/youtube.svg', url: 'https://youtube.com/channel/UCSch2bdI2PdE9pAN_QN3vvQ' },
];

export const Footer: React.FC<FooterPropsType> = () => {
    return (
        <div className={footerCn}>
            <div className={footerContactsCn}>
                {contacts.map(value => {
                    return <a target="_blank" className={footerContactCn} href={value.url}>{value.name}</a>;
                })}
                <div className={footerSocialCn}>
                    {social.map(value => {
                        return <a target="_blank" href={value.url}><Icon className={footerIconCn} url={value.iconUrl}/></a>;
                    })}
                </div>
            </div>
            <div className={footerCopyrightCn}>© Стартапы будущего {new Date().getFullYear()}</div>
        </div>
    );
};
