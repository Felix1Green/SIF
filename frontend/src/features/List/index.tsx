import * as React from 'react';
import { Link } from 'react-router-dom';
import { Button } from '@yandex/ui/Button/desktop/bundle';
import { Spin } from '@yandex/ui/Spin/desktop/bundle';
import { Disclaimer } from '@components/Disclaimer';
import { ClientRoutes } from '@consts/routes';
import { ContentCard } from '@components/ContentCard';
import { ListProps } from './List.typings';
import { ListItem } from './List.components/ListItem';
import {
    listCn,
    listAppendCn,
    listLoadingCn,
    listContainerCn,
} from './List.const';

import './index.scss';

export const List: React.FC<ListProps> = props => {
    const {
        icon,
        list,
        title,
    } = props;

    return (
        <ContentCard className={listCn} title={title} collapsed={false} icon={icon}>
            {!list ?
                <Spin className={listLoadingCn} progress view="default" size="l" />
                :
                list.length ?
                    <div className={listContainerCn}>
                        {list.map(value =>
                            <ListItem
                                key={value.id}
                                id={value.id}
                                url={value.url}
                                title={value.title}
                                description={value.description}
                                actions={value.actions}
                            />
                        )}
                    </div>
                    :
                    <Disclaimer show={true} type={'info'}>Нет данных</Disclaimer>
            }
            <Link className={listAppendCn} to={ClientRoutes.registerPage}>
                <Button view="action" size="m">Добавить</Button>
            </Link>
        </ContentCard>
    );
};
