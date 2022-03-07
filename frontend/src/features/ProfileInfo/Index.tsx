import * as React from 'react';
import { KeyValue } from '@components/KeyValue';
import { Title } from '@components/Title';
import { ProfileInfoProps } from './ProfileInfo.typings';
import {
    profileInfoCn,
    profileInfoNameCn,
    profileInfoRoleCn,
    profileInfoAvatarCn,
    profileInfoPersonalCn,
} from './ProfileInfo.const';
import { ContentCard } from '@src/components/ContentCard';

import './Index.scss';
import { Icon } from '@yandex/ui/Icon/bundle';

export const ProfileInfo: React.FC<ProfileInfoProps> = (props) => {
    const {
        name,
        login,
        surname,
        role,
        avatar,
        patronymic,
        birthday,
        region,
    } = props;

    return (
        <ContentCard className={profileInfoCn}>
            <Icon className={profileInfoAvatarCn} url={avatar ?? '/img/avatar-com.svg'} />
            <div>
                <Title className={profileInfoNameCn}>{surname}</Title>
                <Title className={profileInfoNameCn}>{name}</Title>
                <Title className={profileInfoNameCn}>{patronymic}</Title>
                <div className={profileInfoRoleCn}>{role}</div>
            </div>
            <div className={profileInfoPersonalCn}>
                <KeyValue keyName="Почта" value={login} />
                <KeyValue keyName="День рождения" value={birthday} />
                <KeyValue keyName="Регион" value={region} />
            </div>
        </ContentCard>
    );
};
