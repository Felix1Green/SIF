import * as React from 'react';
import { KeyValue } from '@components/KeyValue';
import { ProfileInfoProps } from './ProfileInfo.typings';
import { profileInfoCn } from './ProfileInfo.consts';
import { ContentCard } from '@src/components/ContentCard';
import { Button } from '@yandex/ui/Button/desktop/bundle';

import './Index.scss';

export const ProfileInfo: React.FC<ProfileInfoProps> = (props) => {
    const {
        name,
        login,
        surname,
        role,
        onLogout,
    } = props;

    return (
        <ContentCard className={profileInfoCn}>
            <KeyValue keyName="Имя" value={name} />
            <KeyValue keyName="Фамилия" value={surname} />
            <KeyValue keyName="Логин" value={login} />
            <KeyValue keyName="Роль" value={role} />
            <Button view="default" size="m" onClick={onLogout}>Выйти</Button>
        </ContentCard>
    );
};
